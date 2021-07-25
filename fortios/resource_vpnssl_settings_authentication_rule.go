// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.0 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Authentication rule for SSL VPN.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnsslsettingsAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnsslsettingsAuthenticationRuleCreate,
		Read:   resourceVpnsslsettingsAuthenticationRuleRead,
		Update: resourceVpnsslsettingsAuthenticationRuleUpdate,
		Delete: resourceVpnsslsettingsAuthenticationRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
			},
			"dynamic_sort_table": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auth": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"any", "local", "radius", "tacacs+", "ldap"}),
				Description:  "SSL VPN authentication method restriction.",
				Optional:     true,
				Computed:     true,
			},
			"cipher": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"any", "high", "medium"}),
				Description:  "SSL VPN cipher strength.",
				Optional:     true,
				Computed:     true,
			},
			"client_cert": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable SSL VPN client certificate restrictive.",
				Optional:     true,
				Computed:     true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "User groups.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Description:  "Group name.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "SSL VPN portal.",
				Optional:     true,
				Computed:     true,
			},
			"realm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "SSL VPN realm.",
				Optional:     true,
				Computed:     true,
			},
			"source_address": {
				Type:        schema.TypeList,
				Description: "Source address of incoming traffic.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Description:  "Address name.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"source_address_negate": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable negated source address match.",
				Optional:     true,
				Computed:     true,
			},
			"source_address6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address of incoming traffic.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Description:  "IPv6 address name.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"source_address6_negate": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable negated source IPv6 address match.",
				Optional:     true,
				Computed:     true,
			},
			"source_interface": {
				Type:        schema.TypeList,
				Description: "SSL VPN source interface of incoming traffic.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Description:  "Interface name.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"user_peer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Name of user peer.",
				Optional:     true,
				Computed:     true,
			},
			"users": {
				Type:        schema.TypeList,
				Description: "User name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Description:  "User name.",
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceVpnsslsettingsAuthenticationRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectVpnsslsettingsAuthenticationRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating VpnsslsettingsAuthenticationRule resource while getting object: %v", err)
	}

	o, err := c.CreateVpnsslsettingsAuthenticationRule(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating VpnsslsettingsAuthenticationRule resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("VpnsslsettingsAuthenticationRule")
	}

	return resourceVpnsslsettingsAuthenticationRuleRead(d, m)
}

func resourceVpnsslsettingsAuthenticationRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectVpnsslsettingsAuthenticationRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating VpnsslsettingsAuthenticationRule resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnsslsettingsAuthenticationRule(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating VpnsslsettingsAuthenticationRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("VpnsslsettingsAuthenticationRule")
	}

	return resourceVpnsslsettingsAuthenticationRuleRead(d, m)
}

func resourceVpnsslsettingsAuthenticationRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	err := c.DeleteVpnsslsettingsAuthenticationRule(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting VpnsslsettingsAuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnsslsettingsAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	o, err := c.ReadVpnsslsettingsAuthenticationRule(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading VpnsslsettingsAuthenticationRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnsslsettingsAuthenticationRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading VpnsslsettingsAuthenticationRule resource from API: %v", err)
	}
	return nil
}

func flattenVpnsslsettingsAuthenticationRuleAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnsslsettingsAuthenticationRuleGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslsettingsAuthenticationRuleGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRulePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnsslsettingsAuthenticationRuleSourceAddressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnsslsettingsAuthenticationRuleSourceAddress6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleSourceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnsslsettingsAuthenticationRuleSourceInterfaceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslsettingsAuthenticationRuleSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleUserPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslsettingsAuthenticationRuleUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnsslsettingsAuthenticationRuleUsersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslsettingsAuthenticationRuleUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnsslsettingsAuthenticationRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("auth", flattenVpnsslsettingsAuthenticationRuleAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("error reading auth: %v", err)
		}
	}

	if err = d.Set("cipher", flattenVpnsslsettingsAuthenticationRuleCipher(o["cipher"], d, "cipher", sv)); err != nil {
		if !fortiAPIPatch(o["cipher"]) {
			return fmt.Errorf("error reading cipher: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenVpnsslsettingsAuthenticationRuleClientCert(o["client-cert"], d, "client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("error reading client_cert: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("groups", flattenVpnsslsettingsAuthenticationRuleGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenVpnsslsettingsAuthenticationRuleGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("error reading groups: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenVpnsslsettingsAuthenticationRuleId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("portal", flattenVpnsslsettingsAuthenticationRulePortal(o["portal"], d, "portal", sv)); err != nil {
		if !fortiAPIPatch(o["portal"]) {
			return fmt.Errorf("error reading portal: %v", err)
		}
	}

	if err = d.Set("realm", flattenVpnsslsettingsAuthenticationRuleRealm(o["realm"], d, "realm", sv)); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("error reading realm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_address", flattenVpnsslsettingsAuthenticationRuleSourceAddress(o["source-address"], d, "source_address", sv)); err != nil {
			if !fortiAPIPatch(o["source-address"]) {
				return fmt.Errorf("error reading source_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_address"); ok {
			if err = d.Set("source_address", flattenVpnsslsettingsAuthenticationRuleSourceAddress(o["source-address"], d, "source_address", sv)); err != nil {
				if !fortiAPIPatch(o["source-address"]) {
					return fmt.Errorf("error reading source_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("source_address_negate", flattenVpnsslsettingsAuthenticationRuleSourceAddressNegate(o["source-address-negate"], d, "source_address_negate", sv)); err != nil {
		if !fortiAPIPatch(o["source-address-negate"]) {
			return fmt.Errorf("error reading source_address_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_address6", flattenVpnsslsettingsAuthenticationRuleSourceAddress6(o["source-address6"], d, "source_address6", sv)); err != nil {
			if !fortiAPIPatch(o["source-address6"]) {
				return fmt.Errorf("error reading source_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_address6"); ok {
			if err = d.Set("source_address6", flattenVpnsslsettingsAuthenticationRuleSourceAddress6(o["source-address6"], d, "source_address6", sv)); err != nil {
				if !fortiAPIPatch(o["source-address6"]) {
					return fmt.Errorf("error reading source_address6: %v", err)
				}
			}
		}
	}

	if err = d.Set("source_address6_negate", flattenVpnsslsettingsAuthenticationRuleSourceAddress6Negate(o["source-address6-negate"], d, "source_address6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["source-address6-negate"]) {
			return fmt.Errorf("error reading source_address6_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_interface", flattenVpnsslsettingsAuthenticationRuleSourceInterface(o["source-interface"], d, "source_interface", sv)); err != nil {
			if !fortiAPIPatch(o["source-interface"]) {
				return fmt.Errorf("error reading source_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_interface"); ok {
			if err = d.Set("source_interface", flattenVpnsslsettingsAuthenticationRuleSourceInterface(o["source-interface"], d, "source_interface", sv)); err != nil {
				if !fortiAPIPatch(o["source-interface"]) {
					return fmt.Errorf("error reading source_interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("user_peer", flattenVpnsslsettingsAuthenticationRuleUserPeer(o["user-peer"], d, "user_peer", sv)); err != nil {
		if !fortiAPIPatch(o["user-peer"]) {
			return fmt.Errorf("error reading user_peer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("users", flattenVpnsslsettingsAuthenticationRuleUsers(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenVpnsslsettingsAuthenticationRuleUsers(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("error reading users: %v", err)
				}
			}
		}
	}

	return nil
}

func expandVpnsslsettingsAuthenticationRuleAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnsslsettingsAuthenticationRuleGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslsettingsAuthenticationRuleGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRulePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnsslsettingsAuthenticationRuleSourceAddressName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddressNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnsslsettingsAuthenticationRuleSourceAddress6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddress6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddress6Negate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnsslsettingsAuthenticationRuleSourceInterfaceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleUserPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslsettingsAuthenticationRuleUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnsslsettingsAuthenticationRuleUsersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslsettingsAuthenticationRuleUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnsslsettingsAuthenticationRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleAuth(d, v, "auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("cipher"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleCipher(d, v, "cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleClientCert(d, v, "client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("portal"); ok {
		t, err := expandVpnsslsettingsAuthenticationRulePortal(d, v, "portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleRealm(d, v, "realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleSourceAddress(d, v, "source_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	if v, ok := d.GetOk("source_address_negate"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleSourceAddressNegate(d, v, "source_address_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_address6"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleSourceAddress6(d, v, "source_address6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6"] = t
		}
	}

	if v, ok := d.GetOk("source_address6_negate"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleSourceAddress6Negate(d, v, "source_address6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_interface"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleSourceInterface(d, v, "source_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-interface"] = t
		}
	}

	if v, ok := d.GetOk("user_peer"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleUserPeer(d, v, "user_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-peer"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {
		t, err := expandVpnsslsettingsAuthenticationRuleUsers(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	return &obj, nil
}
