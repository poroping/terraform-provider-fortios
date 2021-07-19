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

func resourceVpnsslAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnsslAuthenticationRuleCreate,
		Read:   resourceVpnsslAuthenticationRuleRead,
		Update: resourceVpnsslAuthenticationRuleUpdate,
		Delete: resourceVpnsslAuthenticationRuleDelete,

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

func resourceVpnsslAuthenticationRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnsslAuthenticationRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating VpnsslAuthenticationRule resource while getting object: %v", err)
	}

	o, err := c.CreateVpnsslAuthenticationRule(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating VpnsslAuthenticationRule resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("VpnsslAuthenticationRule")
	}

	return resourceVpnsslAuthenticationRuleRead(d, m)
}

func resourceVpnsslAuthenticationRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnsslAuthenticationRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating VpnsslAuthenticationRule resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnsslAuthenticationRule(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating VpnsslAuthenticationRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("VpnsslAuthenticationRule")
	}

	return resourceVpnsslAuthenticationRuleRead(d, m)
}

func resourceVpnsslAuthenticationRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteVpnsslAuthenticationRule(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting VpnsslAuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnsslAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnsslAuthenticationRule(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading VpnsslAuthenticationRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnsslAuthenticationRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading VpnsslAuthenticationRule resource from API: %v", err)
	}
	return nil
}

func flattenVpnsslAuthenticationRuleAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnsslAuthenticationRuleGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslAuthenticationRuleGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRulePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleSourceAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnsslAuthenticationRuleSourceAddressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslAuthenticationRuleSourceAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleSourceAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnsslAuthenticationRuleSourceAddress6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslAuthenticationRuleSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleSourceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnsslAuthenticationRuleSourceInterfaceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslAuthenticationRuleSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleUserPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnsslAuthenticationRuleUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnsslAuthenticationRuleUsersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnsslAuthenticationRuleUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnsslAuthenticationRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("auth", flattenVpnsslAuthenticationRuleAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("error reading auth: %v", err)
		}
	}

	if err = d.Set("cipher", flattenVpnsslAuthenticationRuleCipher(o["cipher"], d, "cipher", sv)); err != nil {
		if !fortiAPIPatch(o["cipher"]) {
			return fmt.Errorf("error reading cipher: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenVpnsslAuthenticationRuleClientCert(o["client-cert"], d, "client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("error reading client_cert: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("groups", flattenVpnsslAuthenticationRuleGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenVpnsslAuthenticationRuleGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("error reading groups: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenVpnsslAuthenticationRuleId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("portal", flattenVpnsslAuthenticationRulePortal(o["portal"], d, "portal", sv)); err != nil {
		if !fortiAPIPatch(o["portal"]) {
			return fmt.Errorf("error reading portal: %v", err)
		}
	}

	if err = d.Set("realm", flattenVpnsslAuthenticationRuleRealm(o["realm"], d, "realm", sv)); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("error reading realm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_address", flattenVpnsslAuthenticationRuleSourceAddress(o["source-address"], d, "source_address", sv)); err != nil {
			if !fortiAPIPatch(o["source-address"]) {
				return fmt.Errorf("error reading source_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_address"); ok {
			if err = d.Set("source_address", flattenVpnsslAuthenticationRuleSourceAddress(o["source-address"], d, "source_address", sv)); err != nil {
				if !fortiAPIPatch(o["source-address"]) {
					return fmt.Errorf("error reading source_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("source_address_negate", flattenVpnsslAuthenticationRuleSourceAddressNegate(o["source-address-negate"], d, "source_address_negate", sv)); err != nil {
		if !fortiAPIPatch(o["source-address-negate"]) {
			return fmt.Errorf("error reading source_address_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_address6", flattenVpnsslAuthenticationRuleSourceAddress6(o["source-address6"], d, "source_address6", sv)); err != nil {
			if !fortiAPIPatch(o["source-address6"]) {
				return fmt.Errorf("error reading source_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_address6"); ok {
			if err = d.Set("source_address6", flattenVpnsslAuthenticationRuleSourceAddress6(o["source-address6"], d, "source_address6", sv)); err != nil {
				if !fortiAPIPatch(o["source-address6"]) {
					return fmt.Errorf("error reading source_address6: %v", err)
				}
			}
		}
	}

	if err = d.Set("source_address6_negate", flattenVpnsslAuthenticationRuleSourceAddress6Negate(o["source-address6-negate"], d, "source_address6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["source-address6-negate"]) {
			return fmt.Errorf("error reading source_address6_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_interface", flattenVpnsslAuthenticationRuleSourceInterface(o["source-interface"], d, "source_interface", sv)); err != nil {
			if !fortiAPIPatch(o["source-interface"]) {
				return fmt.Errorf("error reading source_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_interface"); ok {
			if err = d.Set("source_interface", flattenVpnsslAuthenticationRuleSourceInterface(o["source-interface"], d, "source_interface", sv)); err != nil {
				if !fortiAPIPatch(o["source-interface"]) {
					return fmt.Errorf("error reading source_interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("user_peer", flattenVpnsslAuthenticationRuleUserPeer(o["user-peer"], d, "user_peer", sv)); err != nil {
		if !fortiAPIPatch(o["user-peer"]) {
			return fmt.Errorf("error reading user_peer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("users", flattenVpnsslAuthenticationRuleUsers(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenVpnsslAuthenticationRuleUsers(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("error reading users: %v", err)
				}
			}
		}
	}

	return nil
}

func expandVpnsslAuthenticationRuleAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandVpnsslAuthenticationRuleGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslAuthenticationRuleGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRulePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandVpnsslAuthenticationRuleSourceAddressName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslAuthenticationRuleSourceAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleSourceAddressNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleSourceAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandVpnsslAuthenticationRuleSourceAddress6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslAuthenticationRuleSourceAddress6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleSourceAddress6Negate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandVpnsslAuthenticationRuleSourceInterfaceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslAuthenticationRuleSourceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleUserPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnsslAuthenticationRuleUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandVpnsslAuthenticationRuleUsersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnsslAuthenticationRuleUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnsslAuthenticationRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth"); ok {
		t, err := expandVpnsslAuthenticationRuleAuth(d, v, "auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("cipher"); ok {
		t, err := expandVpnsslAuthenticationRuleCipher(d, v, "cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok {
		t, err := expandVpnsslAuthenticationRuleClientCert(d, v, "client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {
		t, err := expandVpnsslAuthenticationRuleGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandVpnsslAuthenticationRuleId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("portal"); ok {
		t, err := expandVpnsslAuthenticationRulePortal(d, v, "portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok {
		t, err := expandVpnsslAuthenticationRuleRealm(d, v, "realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok {
		t, err := expandVpnsslAuthenticationRuleSourceAddress(d, v, "source_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	if v, ok := d.GetOk("source_address_negate"); ok {
		t, err := expandVpnsslAuthenticationRuleSourceAddressNegate(d, v, "source_address_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_address6"); ok {
		t, err := expandVpnsslAuthenticationRuleSourceAddress6(d, v, "source_address6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6"] = t
		}
	}

	if v, ok := d.GetOk("source_address6_negate"); ok {
		t, err := expandVpnsslAuthenticationRuleSourceAddress6Negate(d, v, "source_address6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_interface"); ok {
		t, err := expandVpnsslAuthenticationRuleSourceInterface(d, v, "source_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-interface"] = t
		}
	}

	if v, ok := d.GetOk("user_peer"); ok {
		t, err := expandVpnsslAuthenticationRuleUserPeer(d, v, "user_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-peer"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {
		t, err := expandVpnsslAuthenticationRuleUsers(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	return &obj, nil
}
