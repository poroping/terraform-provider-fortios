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

func dataSourceVpnsslAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVpnsslAuthenticationRuleRead,
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

func dataSourceVpnsslAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("error describing VpnsslAuthenticationRule: type error")
	}

	o, err := c.ReadVpnsslAuthenticationRule(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing VpnsslAuthenticationRule: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectVpnsslAuthenticationRule(d, o)
	if err != nil {
		return fmt.Errorf("error describing VpnsslAuthenticationRule from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenVpnsslAuthenticationRuleAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslAuthenticationRuleGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslAuthenticationRuleGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRulePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslAuthenticationRuleSourceAddressName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslAuthenticationRuleSourceAddress6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslAuthenticationRuleSourceInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslAuthenticationRuleSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleUserPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnsslAuthenticationRuleUsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenVpnsslAuthenticationRuleUsersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnsslAuthenticationRuleUsersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectVpnsslAuthenticationRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth", dataSourceFlattenVpnsslAuthenticationRuleAuth(o["auth"], d, "auth")); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("error reading auth: %v", err)
		}
	}

	if err = d.Set("cipher", dataSourceFlattenVpnsslAuthenticationRuleCipher(o["cipher"], d, "cipher")); err != nil {
		if !fortiAPIPatch(o["cipher"]) {
			return fmt.Errorf("error reading cipher: %v", err)
		}
	}

	if err = d.Set("client_cert", dataSourceFlattenVpnsslAuthenticationRuleClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("error reading client_cert: %v", err)
		}
	}

	if err = d.Set("groups", dataSourceFlattenVpnsslAuthenticationRuleGroups(o["groups"], d, "groups")); err != nil {
		if !fortiAPIPatch(o["groups"]) {
			return fmt.Errorf("error reading groups: %v", err)
		}
	}

	if err = d.Set("fosid", dataSourceFlattenVpnsslAuthenticationRuleId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("portal", dataSourceFlattenVpnsslAuthenticationRulePortal(o["portal"], d, "portal")); err != nil {
		if !fortiAPIPatch(o["portal"]) {
			return fmt.Errorf("error reading portal: %v", err)
		}
	}

	if err = d.Set("realm", dataSourceFlattenVpnsslAuthenticationRuleRealm(o["realm"], d, "realm")); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("error reading realm: %v", err)
		}
	}

	if err = d.Set("source_address", dataSourceFlattenVpnsslAuthenticationRuleSourceAddress(o["source-address"], d, "source_address")); err != nil {
		if !fortiAPIPatch(o["source-address"]) {
			return fmt.Errorf("error reading source_address: %v", err)
		}
	}

	if err = d.Set("source_address_negate", dataSourceFlattenVpnsslAuthenticationRuleSourceAddressNegate(o["source-address-negate"], d, "source_address_negate")); err != nil {
		if !fortiAPIPatch(o["source-address-negate"]) {
			return fmt.Errorf("error reading source_address_negate: %v", err)
		}
	}

	if err = d.Set("source_address6", dataSourceFlattenVpnsslAuthenticationRuleSourceAddress6(o["source-address6"], d, "source_address6")); err != nil {
		if !fortiAPIPatch(o["source-address6"]) {
			return fmt.Errorf("error reading source_address6: %v", err)
		}
	}

	if err = d.Set("source_address6_negate", dataSourceFlattenVpnsslAuthenticationRuleSourceAddress6Negate(o["source-address6-negate"], d, "source_address6_negate")); err != nil {
		if !fortiAPIPatch(o["source-address6-negate"]) {
			return fmt.Errorf("error reading source_address6_negate: %v", err)
		}
	}

	if err = d.Set("source_interface", dataSourceFlattenVpnsslAuthenticationRuleSourceInterface(o["source-interface"], d, "source_interface")); err != nil {
		if !fortiAPIPatch(o["source-interface"]) {
			return fmt.Errorf("error reading source_interface: %v", err)
		}
	}

	if err = d.Set("user_peer", dataSourceFlattenVpnsslAuthenticationRuleUserPeer(o["user-peer"], d, "user_peer")); err != nil {
		if !fortiAPIPatch(o["user-peer"]) {
			return fmt.Errorf("error reading user_peer: %v", err)
		}
	}

	if err = d.Set("users", dataSourceFlattenVpnsslAuthenticationRuleUsers(o["users"], d, "users")); err != nil {
		if !fortiAPIPatch(o["users"]) {
			return fmt.Errorf("error reading users: %v", err)
		}
	}

	return nil
}
