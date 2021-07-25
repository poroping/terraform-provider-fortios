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
)

func main() {
	schemas, _ := ioutil.ReadDir("./apischema/")
	for _, schema := range schemas {
		fname := schema.Name()
		funcMap := template.FuncMap{
			"replace":     replace,
			"typelookup":  typelookup,
			"flatten":     flatten,
			"valilookup":  valilookup,
			"title":       title,
			"subcategory": subcategory,
		}
		t := template.Must(template.New("main").Funcs(funcMap).ParseGlob("./templates/*.gotmpl"))

		m := map[string]interface{}{}
		jsondata := getSchema(fname)
		if err := json.Unmarshal([]byte(jsondata), &m); err != nil {
			panic(err)
		}

		// Add relative path to json struct for use in templates
		n := addPaths(m)

		r := addResourceInfo(n)

		var buf bytes.Buffer
		sname := "resource"
		if err := t.ExecuteTemplate(&buf, sname, r); err != nil {
			panic(err)
		}

		// lint
		f, err := format.Source(buf.Bytes())
		if err != nil {
			panic(err)
		}
		// Debug
		// f := buf.Bytes()
		// fmt.Println(f)

		perm := int(0755)
		resname := sname + "_" + fname
		resname = strings.TrimSuffix(resname, ".json")
		resname += ".go"
		os.WriteFile("./output/"+resname, f, os.FileMode(perm))

		d := addDataSourceInfo(n)

		var buf2 bytes.Buffer
		sname = "data_source"
		if err := t.ExecuteTemplate(&buf2, sname, d); err != nil {
			panic(err)
		}

		// lint
		f, err = format.Source(buf2.Bytes())
		if err != nil {
			panic(err)
		}

		dsname := fname
		dsname = sname + "_" + dsname
		dsname = strings.TrimSuffix(dsname, ".json")
		dsname += ".go"
		os.WriteFile("./output/"+dsname, f, os.FileMode(perm))

		var buf3 bytes.Buffer
		sname = "data_source_docs"
		if err := t.ExecuteTemplate(&buf3, sname, d); err != nil {
			panic(err)
		}

		f = buf3.Bytes()

		dsdocsname := fname
		dsdocsname = "fortios_" + dsdocsname
		dsdocsname = strings.TrimSuffix(dsdocsname, ".json")
		dsdocsname += ".html.markdown"
		os.WriteFile("./output/d/"+dsdocsname, f, os.FileMode(perm))

		var buf4 bytes.Buffer
		sname = "resource_docs"
		if err := t.ExecuteTemplate(&buf4, sname, d); err != nil {
			panic(err)
		}

		f = buf4.Bytes()

		resdocsname := fname
		resdocsname = "fortios_" + resdocsname
		resdocsname = strings.TrimSuffix(resdocsname, ".json")
		resdocsname += ".html.markdown"
		os.WriteFile("./output/r/"+resdocsname, f, os.FileMode(perm))
	}
}

func addPaths(m map[string]interface{}) map[string]interface{} {
	path := flattenTitle(m["path"].(string))
	if _, ok := m["child_path"].(string); ok {
		name, _ := m["name"].(string)
		path += name
	}
	resource := flattenTitle(m["results"].(map[string]interface{})["name"].(string))
	tmp := m["results"]
	m["fpath"] = path + resource
	recursePaths(tmp, path+resource)
	return m
}

func recursePaths(v interface{}, pre string) interface{} {
	child, ok := v.(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			name := flattenTitle(v.(map[string]interface{})["name"].(string))
			v.(map[string]interface{})["fpath"] = pre + name
			recursePaths(v, pre+name)
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
	child, ok := m["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			if s, ok := v.(map[string]interface{})["type"].(string); ok {
				if strings.Contains(s, "password") {
					v.(map[string]interface{})["sensitive"] = true
				}
			}
			recurseSensitive(v.(map[string]interface{}))
		}
	} else {
		return m
	}
	return m
}

func addDataSourceInfo(m map[string]interface{}) map[string]interface{} {
	m = replaceTopLevelId(m)
	m = addSchemaRequired(m)
	m = addSensitive(m)
	return m
}

func addSchemaRequired(m map[string]interface{}) map[string]interface{} {
	mkey := m["results"].(map[string]interface{})["mkey"].(string)
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			if name == mkey {
				v.(map[string]interface{})["schema_required"] = true
			}

		}
	}
	return m
}

func addResourceInfo(m map[string]interface{}) map[string]interface{} {
	m = addDynSortTable(m)
	m = replaceTopLevelId(m)
	m = addSchemaRequired(m)
	m = addSensitive(m)
	return m
}

func replaceTopLevelId(m map[string]interface{}) map[string]interface{} {
	child, ok := m["results"].(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, v := range child {
			name := v.(map[string]interface{})["name"].(string)
			if name == "id" {
				v.(map[string]interface{})["fosid"] = true
				v.(map[string]interface{})["name"] = "fosid"
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

func getSchema(name string) (content []byte) {
	file := "./apischema/"
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
	return s
}

func typelookup(s string) string {
	m := map[string]string{
		"string":             "TypeString",
		"option":             "TypeString",
		"ipv4-address":       "TypeString",
		"ipv4-address-any":   "TypeString",
		"ipv4-classnet":      "TypeString",
		"ipv4-classnet-host": "TypeString",
		"ipv4-classnet-any":  "TypeString",
		"ipv4-netmask":       "TypeString",
		"ipv6-address":       "TypeString",
		"ipv6-prefix":        "TypeString",
		"ipv6-network":       "TypeString",
		"var-string":         "TypeString",
		"password":           "TypeString",
		"integer":            "TypeInt",
		"user":               "TypeString",
		"password-3":         "TypeString",
	}
	s, ok := m[s]
	if !ok {
		s = "TYPE-ERROR"
	}
	return fmt.Sprintf("Type:         schema.%s,", s)
}

func valilookup(values map[string]interface{}) string {
	vtype := values["type"].(string)
	size, _ := values["size"].(float64)
	multi_val, _ := values["multiple_values"].(bool)
	o, _ := values["options"].([]interface{})
	min, _ := values["min-value"].(float64)
	max, _ := values["max-value"].(float64)
	m := map[string]string{
		"string":             valiStringLen(size),
		"option":             valiOptions(o, multi_val),
		"ipv4-address":       "validation.IsIPv4Address",
		"ipv4-address-any":   "validation.IsIPv4Address",
		"ipv4-classnet":      "fortiValidateIPv4Classnet",
		"ipv4-classnet-host": "fortiValidateIPv4ClassnetHost",
		"ipv4-classnet-any":  "fortiValidateIPv4ClassnetAny",
		"ipv4-netmask":       "fortiValidateIPv4Netmask",
		"ipv6-address":       "validation.IsIPv6Address",
		"ipv6-prefix":        "fortiValidateIPv6Prefix",
		"ipv6-network":       "fortiValidateIPv6Network",
		"var-string":         valiStringLen(size),
		"password":           "",
		"integer":            valiInt(int(min), int(max)),
		"user":               "",
		"password-3":         "",
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
	if s == `"enable", "disable"` || s == `"disable", "enable"` {
		return "fortiValidateEnableDisable()"
	}
	return fmt.Sprintf("fortiValidateEnum([]string{%s})", s)
}

func flattenTitle(v string) string {
	v = strings.ReplaceAll(v, ".", "")
	v = strings.ReplaceAll(v, "-", " ")
	return strings.ReplaceAll(strings.Title(v), " ", "")
}

func title(v string) string {
	return strings.Title(v)
}

func subcategory(input string) string {
	s := strings.Split(input, ".")
	return s[0]
}
