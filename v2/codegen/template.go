package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

const (
	schemaLoc = "./apischema/auto-merge/"
)

type providerRes struct {
	Fpath     string
	Path      string
	Name      string
	ChildPath string
}

func main() {
	schemas, _ := ioutil.ReadDir(schemaLoc)
	cmdbResources := []providerRes{}
	for _, schema := range schemas {
		fileName := schema.Name()
		if ignoreSchemas(fileName) {
			log.Printf("[DEBUG] schema: %s marked for skipping", fileName)
			continue
		}
		log.Printf("[DEBUG] processing: %s", fileName)
		funcMap := template.FuncMap{
			"replace":          replace,
			"mixedCase":        mixedCase,
			"typeLookup":       typeLookup,
			"schemaTypeLookup": schemaTypeLookup,
			"flatten":          flatten,
			"resFlatten":       resFlatten,
			"valiLookup":       valiLookup,
			"title":            title,
			"subcategory":      subcategory,
			"diffLookup":       diffLookup,
			"readExample":      readExample,
			"debug":            debugx,
		}
		t := template.Must(template.New("main").Funcs(funcMap).ParseGlob("./templates/*.gotmpl"))

		m := map[string]interface{}{}
		jsondata := getSchema(fileName)
		if err := json.Unmarshal([]byte(jsondata), &m); err != nil {
			panic(err)
		}
		// Add relative path to json struct for use in templates.
		n := addPaths(m)
		// Add extra values to use in templates.
		r := addResourceInfo(n)

		cp, ok := r["child_path"].(string)
		if !ok {
			cp = ""
		}
		tmp := providerRes{
			Fpath:     r["fpath"].(string),
			Path:      r["path"].(string),
			Name:      r["name"].(string),
			ChildPath: cp,
		}
		cmdbResources = append(cmdbResources, tmp)

		fileName = strings.TrimSuffix(fileName, ".json")

		render("resource", fileName, t, r)
		fileName += "_test"
		render("resource_test", fileName, t, r)
	}
	providerResourceRender(cmdbResources)
}

// ignore these cause need work
func ignoreSchemas(name string) bool {
	ign := []string{
		"router.access_list_rule.json",
		"router.access_list6_rule.json",
		"router.prefix_list_rule.json",
		"router.prefix_list6_rule.json",
		"router.route_map_rule.json",
		// duplicate keys
		"extender_controller_dataplan.json",
		// empty ?
		"firewall_vendor_mac_summary.json",
	}
	for _, v := range ign {
		if name == v {
			return true
		}
	}
	return false
}

func isComplex(category string) bool {
	return category == "complex"
}

func providerResourceRender(resources []providerRes) {
	funcMap := template.FuncMap{
		"replace":    replace,
		"resFlatten": resFlatten,
	}
	t := template.Must(template.New("main").Funcs(funcMap).ParseGlob("./templates/provider_resource.gotmpl"))
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "provider_resource", resources); err != nil {
		panic(err)
	}
	f, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	perm := int(0755)
	os.WriteFile("../internal/provider/provider_resources.go", f, os.FileMode(perm))
}

func debugx(v interface{}) string {
	fmt.Println(v)
	return ""
}

func render(templateName, fileName string, t *template.Template, m map[string]interface{}) {
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, templateName, m); err != nil {
		panic(err)
	}

	// lint
	// f := buf.Bytes()
	log.Printf("[DEBUG] linting: %s", fileName)
	f, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	log.Printf("[DEBUG] removed unused imports: %s", fileName)
	f, err = imports.Process("", f, nil)
	if err != nil {
		panic(err)
	}
	// Debug
	// f := buf.Bytes()
	// fmt.Println(f)

	perm := int(0755)
	os.WriteFile(fmt.Sprintf("../internal/provider/resource_%s.go", fileName), f, os.FileMode(perm))
}

func addPaths(m map[string]interface{}) map[string]interface{} {
	path := mixedCase(m["path"].(string))
	if _, ok := m["child_path"].(string); ok {
		name, _ := m["name"].(string)
		path += mixedCase(name)
	}
	resource := mixedCase(m["results"].(map[string]interface{})["name"].(string))
	tmp := m["results"]
	m["fpath"] = mixedCase(path + " " + resource)
	m["results"].(map[string]interface{})["fpath"] = mixedCase(path + " " + resource)
	recursePaths(tmp, path+resource)
	return m
}

func recursePaths(v interface{}, pre string) interface{} {
	child, ok := v.(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			if v2, ok := v.(map[string]interface{}); ok {
				name := mixedCase(v2["name"].(string))
				fpath := mixedCase(pre + " " + name)
				v2["fpath"] = fpath
				recursePaths(v, fpath)
			}
		}
	} else {
		return v
	}
	return v
}

func addSensitive(m map[string]interface{}) map[string]interface{} {
	tmp := m["results"].(map[string]interface{})
	recurseSensitive(tmp)
	return m
}

func recurseSensitive(m map[string]interface{}) map[string]interface{} {
	if child, ok := m["children"].(map[string]interface{}); ok {
		for _, v := range child {
			if v2, ok := v.(map[string]interface{}); ok {
				if s, ok := v2["type"].(string); ok {
					if strings.Contains(s, "password") {
						v2["sensitive"] = true
					}
				}
				recurseSensitive(v2)
			}
		}
	} else {
		return m
	}
	return m
}

// mark if name begins with number cause Go struct no likey and need to prepend
func addNumberName(m map[string]interface{}) map[string]interface{} {
	tmp := m["results"].(map[string]interface{})
	recurseNumberName(tmp)
	return m
}

func recurseNumberName(m map[string]interface{}) map[string]interface{} {
	if child, ok := m["children"].(map[string]interface{}); ok {
		for _, v := range child {
			if v2, ok := v.(map[string]interface{}); ok {
				if s, ok := v2["name"].(string); ok {
					if s[0] >= '0' && s[0] <= '9' {
						v2["name_begin_with_int"] = true
					}
				}
				recurseNumberName(v2)
			}
		}
	} else {
		return m
	}
	return m
}

func addDataSourceInfo(m map[string]interface{}) map[string]interface{} {
	m = replaceTopLevelId(m)
	m = addDataSourceSchemaRequired(m)
	m = addSensitive(m)
	m = addNumberName(m)
	return m
}

func addDataSourceSchemaRequired(m map[string]interface{}) map[string]interface{} {
	mkey, exists := m["results"].(map[string]interface{})["mkey"].(string)
	if !exists {
		return m
	}
	if child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{}); ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			if name == mkey {
				v.(map[string]interface{})["schema_required"] = true
			}

		}
	}
	return m
}

func addResourceSchemaRequired(m map[string]interface{}) map[string]interface{} {
	mkey, ok := m["results"].(map[string]interface{})["mkey"].(string)
	if !ok {
		return m
	}
	if child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{}); ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			if name == "seq-num" || name == "id" {
				v.(map[string]interface{})["schema_opt_force"] = true
			} else if name == mkey {
				v.(map[string]interface{})["schema_required"] = true
			}

		}
	}
	return m
}

func addResourceInfo(m map[string]interface{}) map[string]interface{} {
	m = addDynSortTable(m)
	m = replaceTopLevelId(m)
	m = addResourceSchemaRequired(m)
	m = addSensitive(m)
	m = complexEmptyOnDestroy(m)
	m = addNumberName(m)
	return m
}

func replaceTopLevelId(m map[string]interface{}) map[string]interface{} {
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			// id and count are reserved in TF
			if name == "id" || name == "count" {
				v.(map[string]interface{})["NameReserved"] = true
			}
		}
	}
	return m
}

func addDynSortTable(m map[string]interface{}) map[string]interface{} {
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		dyn := false
		for _, v := range child {
			if category, ok := v.(map[string]interface{})["category"].(string); ok {
				if category == "table" {
					dyn = true
				}
			}
			m["results"].(map[string]interface{})["dynamic_sort_table"] = dyn
		}
	}
	return m
}

func complexEmptyOnDestroy(m map[string]interface{}) map[string]interface{} {
	category := m["results"].(map[string]interface{})["category"].(string)
	if category == "complex" {
		if child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{}); ok {
			complex_parent := false
			for _, v := range child {
				if category, ok := v.(map[string]interface{})["category"].(string); ok {
					if category == "table" {
						complex_parent = true
					}
				}
				m["results"].(map[string]interface{})["complex_parent"] = complex_parent
			}
		}
	}
	return m
}

func getSchema(name string) (content []byte) {
	file := schemaLoc
	file += name
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func flatten(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "+", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	return s
}

func resFlatten(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "+", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ToLower(s)
	return s
}

func schemaTypeLookup(s string) string {
	m := map[string]string{
		"string":                 "TypeString",
		"option":                 "TypeString",
		"ipv4-address":           "TypeString",
		"ipv4-address-any":       "TypeString",
		"ipv4-classnet":          "TypeString",
		"ipv4-classnet-host":     "TypeString",
		"ipv4-classnet-any":      "TypeString",
		"ipv4-netmask-any":       "TypeString",
		"ipv4-address-multicast": "TypeString",
		"ipv4-netmask":           "TypeString",
		"ipv6-address":           "TypeString",
		"ipv6-prefix":            "TypeString",
		"ipv6-network":           "TypeString",
		"var-string":             "TypeString",
		"password":               "TypeString",
		"integer":                "TypeInt",
		"user":                   "TypeString",
		"password-2":             "TypeString",
		"password-3":             "TypeString",
		"varlen_password":        "TypeString",
		"mac-address":            "TypeString",
		"password_aes256":        "TypeString",
		"uuid":                   "TypeString",
		"ether-type":             "TypeString",
	}
	s, ok := m[s]
	if !ok {
		log.Fatalf("SchemaType not found %s", s)
	}
	return fmt.Sprintf("Type:         schema.%s,", s)
}

func typeLookup(s string) string {
	m := map[string]string{
		"string":                 "string",
		"option":                 "string",
		"ipv4-address":           "string",
		"ipv4-address-any":       "string",
		"ipv4-classnet":          "string",
		"ipv4-classnet-host":     "string",
		"ipv4-classnet-any":      "string",
		"ipv4-netmask":           "string",
		"ipv4-netmask-any":       "string",
		"ipv4-address-multicast": "string",
		"ipv6-address":           "string",
		"ipv6-prefix":            "string",
		"ipv6-network":           "string",
		"var-string":             "string",
		"password":               "string",
		"integer":                "int64",
		"user":                   "string",
		"password-2":             "string",
		"password-3":             "string",
		"varlen_password":        "string",
		"mac-address":            "string",
		"password_aes256":        "string",
		"uuid":                   "string",
		"ether-type":             "string",
	}
	s, ok := m[s]
	if !ok {
		log.Fatalf("Type not found %s", s)
	}
	return s
}

func valiLookup(values map[string]interface{}) string {
	vtype := values["type"].(string)
	size, _ := values["size"].(float64)
	multi_val, _ := values["multiple_values"].(bool)
	o, _ := values["options"].([]interface{})
	min, _ := values["min-value"].(float64)
	max, _ := values["max-value"].(float64)
	m := map[string]string{
		"string":                 valiStringLen(size),
		"option":                 valiOptions(o, multi_val),
		"ipv4-address":           "validation.IsIPv4Address",
		"ipv4-address-any":       "validation.IsIPv4Address",
		"ipv4-classnet":          "validators.FortiValidateIPv4Classnet",
		"ipv4-classnet-host":     "validators.FortiValidateIPv4ClassnetHost",
		"ipv4-classnet-any":      "validators.FortiValidateIPv4ClassnetAny",
		"ipv4-netmask":           "validators.FortiValidateIPv4Netmask",
		"ipv4-netmask-any":       "validators.FortiValidateIPv4Netmask",
		"ipv4-address-multicast": "validation.IsIPv4Address",
		"ipv6-address":           "validation.IsIPv6Address",
		"ipv6-prefix":            "validators.FortiValidateIPv6Prefix",
		"ipv6-network":           "validators.FortiValidateIPv6Network",
		"var-string":             valiStringLen(size),
		"password":               "",
		"integer":                valiInt(int(min), int(max)),
		"user":                   "",
		"password-2":             "",
		"password-3":             "",
		"varlen_password":        "",
		"mac-address":            "",
		"password_aes256":        "",
		"uuid":                   "",
		"ether-type":             "",
	}
	s, ok := m[vtype]
	if !ok {
		s = "VALI-ERROR"
	}
	if s == "" {
		return ""
	} else {
		return fmt.Sprintf("ValidateFunc: %s,", s)
	}
}

func diffLookup(values map[string]interface{}) string {
	vtype := values["type"].(string)
	multi_val, _ := values["multiple_values"].(bool)
	m := map[string]string{
		"string":                 "",
		"option":                 diffOptions(multi_val),
		"ipv4-address":           "",
		"ipv4-address-any":       "",
		"ipv4-classnet":          "",
		"ipv4-classnet-host":     "",
		"ipv4-classnet-any":      "",
		"ipv4-netmask":           "",
		"ipv4-netmask-any":       "",
		"ipv4-address-multicast": "",
		"ipv6-address":           "suppressors.DiffIPEqual",
		"ipv6-prefix":            "suppressors.DiffCidrEqual",
		"ipv6-network":           "suppressors.DiffCidrEqual",
		"var-string":             "",
		"password":               "",
		"integer":                "",
		"user":                   "",
		"password-2":             "",
		"password-3":             "",
		"varlen_password":        "",
		"mac-address":            "",
		"password_aes256":        "",
		"uuid":                   "",
		"ether-type":             "",
	}
	s, ok := m[vtype]
	if !ok {
		s = "DIFF-LOOKUP-ERROR"
	}
	if s == "" {
		return ""
	} else {
		return fmt.Sprintf("DiffSuppressFunc: %s,", s)
	}
}

func diffOptions(multi_val bool) string {
	if multi_val {
		return "suppressors.DiffFakeListEqual"
	} else {
		return ""
	}
}

func valiStringLen(v float64) string {
	i := int(v)
	return fmt.Sprintf("validation.StringLenBetween(0, %d)", i)
}

func valiInt(min, max int) string {
	if max > 2147483647 {
		return ""
	}
	return fmt.Sprintf("validation.IntBetween(%d, %d)", min, max)
}

func valiOptions(opts []interface{}, multi_val bool) string {
	if multi_val {
		return ""
	}
	l := make([]string, 0)
	for _, v := range opts {
		l = append(l, fmt.Sprintf("%q", v.(map[string]interface{})["name"].(string)))
	}
	s := strings.Join(l, ", ")
	return fmt.Sprintf("validation.StringInSlice([]string{%s}, false)", s)
}

func mixedCase(v string) string {
	v = strings.ReplaceAll(v, ".", " ")
	v = strings.ReplaceAll(v, "-", " ")
	v = strings.ReplaceAll(v, "+", "")
	v = strings.ReplaceAll(v, "(", "")
	v = strings.ReplaceAll(v, ")", "")
	return strings.ReplaceAll(strings.Title(v), " ", "")
}

func title(v string) string {
	return strings.Title(v)
}

func subcategory(input string) string {
	s := strings.Split(input, ".")
	return s[0]
}

func readExample(name, typ string) string {
	file, err := os.Open(fmt.Sprintf("./examples/%s.%s.txt", name, typ))
	if err != nil {
		return ""
	}
	example, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(example)
}

// Entire swagger schema
// https://fndn.fortinet.net/index.php?/fortiapi/1-fortios/&do=downloadApiJSONFile&handlerId=974&path=&productId=1&csrfKey=c10ec59a59015e6c2538adec6796f1ca
