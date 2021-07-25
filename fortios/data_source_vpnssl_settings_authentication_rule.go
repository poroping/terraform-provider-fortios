// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.0 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Authentication rule for SSL VPN.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceVpnsslsettingsAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVpnsslsettingsAuthenticationRuleRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth": {
				Type:        schema.TypeString,
				Description: "SSL VPN authentication method restriction.",
				Computed:    true,
			},
			"cipher": {
				Type:        schema.TypeString,
				Description: "SSL VPN cipher strength.",
				Computed:    true,
			},
			"client_cert": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL VPN client certificate restrictive.",
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "User groups.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Group name.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID (0 - 4294967295).",
				Computed:    true,
			},
			"portal": {
				Type:        schema.TypeString,
				Description: "SSL VPN portal.",
				Computed:    true,
			},
			"realm": {
				Type:        schema.TypeString,
				Description: "SSL VPN realm.",
				Computed:    true,
			},
			"source_address": {
				Type:        schema.TypeList,
				Description: "Source address of incoming traffic.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"source_address_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negated source address match.",
				Computed:    true,
			},
			"source_address6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address of incoming traffic.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IPv6 address name.",
							Computed:    true,
						},
					},
				},
			},
			"source_address6_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negated source IPv6 address match.",
				Computed:    true,
			},
			"source_interface": {
				Type:        schema.TypeList,
				Description: "SSL VPN source interface of incoming traffic.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"user_peer": {
				Type:        schema.TypeString,
				Description: "Name of user peer.",
				Computed:    true,
			},
			"users": {
				Type:        schema.TypeList,
				Description: "User name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVpnsslsettingsAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""
	key := "id"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing VpnsslsettingsAuthenticationRule: type error")
	}

	o, err := c.ReadVpnsslsettingsAuthenticationRule(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing VpnsslsettingsAuthenticationRule: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectVpnsslsettingsAuthenticationRule(d, o)
	if err != nil {
		return fmt.Errorf("error describing VpnsslsettingsAuthenticationRule from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslsettingsAuthenticationRuleGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRulePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddressName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleUserPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleUsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslsettingsAuthenticationRuleUsersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslsettingsAuthenticationRuleUsersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectVpnsslsettingsAuthenticationRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth", dataSourceFlattenVpnsslsettingsAuthenticationRuleAuth(o["auth"], d, "auth")); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("error reading auth: %v", err)
		}
	}

	if err = d.Set("cipher", dataSourceFlattenVpnsslsettingsAuthenticationRuleCipher(o["cipher"], d, "cipher")); err != nil {
		if !fortiAPIPatch(o["cipher"]) {
			return fmt.Errorf("error reading cipher: %v", err)
		}
	}

	if err = d.Set("client_cert", dataSourceFlattenVpnsslsettingsAuthenticationRuleClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("error reading client_cert: %v", err)
		}
	}

	if err = d.Set("groups", dataSourceFlattenVpnsslsettingsAuthenticationRuleGroups(o["groups"], d, "groups")); err != nil {
		if !fortiAPIPatch(o["groups"]) {
			return fmt.Errorf("error reading groups: %v", err)
		}
	}

	if err = d.Set("fosid", dataSourceFlattenVpnsslsettingsAuthenticationRuleId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("portal", dataSourceFlattenVpnsslsettingsAuthenticationRulePortal(o["portal"], d, "portal")); err != nil {
		if !fortiAPIPatch(o["portal"]) {
			return fmt.Errorf("error reading portal: %v", err)
		}
	}

	if err = d.Set("realm", dataSourceFlattenVpnsslsettingsAuthenticationRuleRealm(o["realm"], d, "realm")); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("error reading realm: %v", err)
		}
	}

	if err = d.Set("source_address", dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress(o["source-address"], d, "source_address")); err != nil {
		if !fortiAPIPatch(o["source-address"]) {
			return fmt.Errorf("error reading source_address: %v", err)
		}
	}

	if err = d.Set("source_address_negate", dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddressNegate(o["source-address-negate"], d, "source_address_negate")); err != nil {
		if !fortiAPIPatch(o["source-address-negate"]) {
			return fmt.Errorf("error reading source_address_negate: %v", err)
		}
	}

	if err = d.Set("source_address6", dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress6(o["source-address6"], d, "source_address6")); err != nil {
		if !fortiAPIPatch(o["source-address6"]) {
			return fmt.Errorf("error reading source_address6: %v", err)
		}
	}

	if err = d.Set("source_address6_negate", dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceAddress6Negate(o["source-address6-negate"], d, "source_address6_negate")); err != nil {
		if !fortiAPIPatch(o["source-address6-negate"]) {
			return fmt.Errorf("error reading source_address6_negate: %v", err)
		}
	}

	if err = d.Set("source_interface", dataSourceFlattenVpnsslsettingsAuthenticationRuleSourceInterface(o["source-interface"], d, "source_interface")); err != nil {
		if !fortiAPIPatch(o["source-interface"]) {
			return fmt.Errorf("error reading source_interface: %v", err)
		}
	}

	if err = d.Set("user_peer", dataSourceFlattenVpnsslsettingsAuthenticationRuleUserPeer(o["user-peer"], d, "user_peer")); err != nil {
		if !fortiAPIPatch(o["user-peer"]) {
			return fmt.Errorf("error reading user_peer: %v", err)
		}
	}

	if err = d.Set("users", dataSourceFlattenVpnsslsettingsAuthenticationRuleUsers(o["users"], d, "users")); err != nil {
		if !fortiAPIPatch(o["users"]) {
			return fmt.Errorf("error reading users: %v", err)
		}
	}

	return nil
}
