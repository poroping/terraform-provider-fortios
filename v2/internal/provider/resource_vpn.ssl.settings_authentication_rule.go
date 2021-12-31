// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Authentication rule for SSL-VPN.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceVpnsslsettingsAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Description: "Authentication rule for SSL-VPN.",

		CreateContext: resourceVpnsslsettingsAuthenticationRuleCreate,
		ReadContext:   resourceVpnsslsettingsAuthenticationRuleRead,
		UpdateContext: resourceVpnsslsettingsAuthenticationRuleUpdate,
		DeleteContext: resourceVpnsslsettingsAuthenticationRuleDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:         schema.TypeBool,
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"any", "local", "radius", "tacacs+", "ldap", "peer"}, false),

				Description: "SSL-VPN authentication method restriction.",
				Optional:    true,
				Computed:    true,
			},
			"cipher": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"any", "high", "medium"}, false),

				Description: "SSL-VPN cipher strength.",
				Optional:    true,
				Computed:    true,
			},
			"client_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL-VPN client certificate restrictive.",
				Optional:    true,
				Computed:    true,
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

							Description: "Group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID (0 - 4294967295).",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSL-VPN portal.",
				Optional:    true,
				Computed:    true,
			},
			"realm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSL-VPN realm.",
				Optional:    true,
				Computed:    true,
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

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"source_address_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negated source address match.",
				Optional:    true,
				Computed:    true,
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

							Description: "IPv6 address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"source_address6_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negated source IPv6 address match.",
				Optional:    true,
				Computed:    true,
			},
			"source_interface": {
				Type:        schema.TypeList,
				Description: "SSL-VPN source interface of incoming traffic.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"user_peer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of user peer.",
				Optional:    true,
				Computed:    true,
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

							Description: "User name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceVpnsslsettingsAuthenticationRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating VpnsslsettingsAuthenticationRule resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnsslsettingsAuthenticationRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnsslsettingsAuthenticationRule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnsslsettingsAuthenticationRule")
	}

	return resourceVpnsslsettingsAuthenticationRuleRead(ctx, d, meta)
}

func resourceVpnsslsettingsAuthenticationRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
	c := meta.(*apiClient).Client
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	obj, diags := getObjectVpnsslsettingsAuthenticationRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnsslsettingsAuthenticationRule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnsslsettingsAuthenticationRule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnsslsettingsAuthenticationRule")
	}

	return resourceVpnsslsettingsAuthenticationRuleRead(ctx, d, meta)
}

func resourceVpnsslsettingsAuthenticationRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	err := c.Cmdb.DeleteVpnsslsettingsAuthenticationRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnsslsettingsAuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnsslsettingsAuthenticationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	o, err := c.Cmdb.ReadVpnsslsettingsAuthenticationRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnsslsettingsAuthenticationRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectVpnsslsettingsAuthenticationRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnsslsettingsAuthenticationRuleGroups(v *[]models.VpnsslsettingsAuthenticationRuleGroups, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddress(v *[]models.VpnsslsettingsAuthenticationRuleSourceAddress, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnsslsettingsAuthenticationRuleSourceAddress6(v *[]models.VpnsslsettingsAuthenticationRuleSourceAddress6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnsslsettingsAuthenticationRuleSourceInterface(v *[]models.VpnsslsettingsAuthenticationRuleSourceInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnsslsettingsAuthenticationRuleUsers(v *[]models.VpnsslsettingsAuthenticationRuleUsers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectVpnsslsettingsAuthenticationRule(d *schema.ResourceData, o *models.VpnsslsettingsAuthenticationRule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Auth != nil {
		v := *o.Auth

		if err = d.Set("auth", v); err != nil {
			return diag.Errorf("error reading auth: %v", err)
		}
	}

	if o.Cipher != nil {
		v := *o.Cipher

		if err = d.Set("cipher", v); err != nil {
			return diag.Errorf("error reading cipher: %v", err)
		}
	}

	if o.ClientCert != nil {
		v := *o.ClientCert

		if err = d.Set("client_cert", v); err != nil {
			return diag.Errorf("error reading client_cert: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenVpnsslsettingsAuthenticationRuleGroups(o.Groups, sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Portal != nil {
		v := *o.Portal

		if err = d.Set("portal", v); err != nil {
			return diag.Errorf("error reading portal: %v", err)
		}
	}

	if o.Realm != nil {
		v := *o.Realm

		if err = d.Set("realm", v); err != nil {
			return diag.Errorf("error reading realm: %v", err)
		}
	}

	if o.SourceAddress != nil {
		if err = d.Set("source_address", flattenVpnsslsettingsAuthenticationRuleSourceAddress(o.SourceAddress, sort)); err != nil {
			return diag.Errorf("error reading source_address: %v", err)
		}
	}

	if o.SourceAddressNegate != nil {
		v := *o.SourceAddressNegate

		if err = d.Set("source_address_negate", v); err != nil {
			return diag.Errorf("error reading source_address_negate: %v", err)
		}
	}

	if o.SourceAddress6 != nil {
		if err = d.Set("source_address6", flattenVpnsslsettingsAuthenticationRuleSourceAddress6(o.SourceAddress6, sort)); err != nil {
			return diag.Errorf("error reading source_address6: %v", err)
		}
	}

	if o.SourceAddress6Negate != nil {
		v := *o.SourceAddress6Negate

		if err = d.Set("source_address6_negate", v); err != nil {
			return diag.Errorf("error reading source_address6_negate: %v", err)
		}
	}

	if o.SourceInterface != nil {
		if err = d.Set("source_interface", flattenVpnsslsettingsAuthenticationRuleSourceInterface(o.SourceInterface, sort)); err != nil {
			return diag.Errorf("error reading source_interface: %v", err)
		}
	}

	if o.UserPeer != nil {
		v := *o.UserPeer

		if err = d.Set("user_peer", v); err != nil {
			return diag.Errorf("error reading user_peer: %v", err)
		}
	}

	if o.Users != nil {
		if err = d.Set("users", flattenVpnsslsettingsAuthenticationRuleUsers(o.Users, sort)); err != nil {
			return diag.Errorf("error reading users: %v", err)
		}
	}

	return nil
}

func expandVpnsslsettingsAuthenticationRuleGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnsslsettingsAuthenticationRuleGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnsslsettingsAuthenticationRuleGroups

	for i := range l {
		tmp := models.VpnsslsettingsAuthenticationRuleGroups{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnsslsettingsAuthenticationRuleSourceAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnsslsettingsAuthenticationRuleSourceAddress

	for i := range l {
		tmp := models.VpnsslsettingsAuthenticationRuleSourceAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnsslsettingsAuthenticationRuleSourceAddress6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnsslsettingsAuthenticationRuleSourceAddress6

	for i := range l {
		tmp := models.VpnsslsettingsAuthenticationRuleSourceAddress6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnsslsettingsAuthenticationRuleSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnsslsettingsAuthenticationRuleSourceInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnsslsettingsAuthenticationRuleSourceInterface

	for i := range l {
		tmp := models.VpnsslsettingsAuthenticationRuleSourceInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnsslsettingsAuthenticationRuleUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnsslsettingsAuthenticationRuleUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnsslsettingsAuthenticationRuleUsers

	for i := range l {
		tmp := models.VpnsslsettingsAuthenticationRuleUsers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectVpnsslsettingsAuthenticationRule(d *schema.ResourceData, sv string) (*models.VpnsslsettingsAuthenticationRule, diag.Diagnostics) {
	obj := models.VpnsslsettingsAuthenticationRule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth", sv)
				diags = append(diags, e)
			}
			obj.Auth = &v2
		}
	}
	if v1, ok := d.GetOk("cipher"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cipher", sv)
				diags = append(diags, e)
			}
			obj.Cipher = &v2
		}
	}
	if v1, ok := d.GetOk("client_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_cert", sv)
				diags = append(diags, e)
			}
			obj.ClientCert = &v2
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnsslsettingsAuthenticationRuleGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.VpnsslsettingsAuthenticationRuleGroups{}
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("portal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portal", sv)
				diags = append(diags, e)
			}
			obj.Portal = &v2
		}
	}
	if v1, ok := d.GetOk("realm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("realm", sv)
				diags = append(diags, e)
			}
			obj.Realm = &v2
		}
	}
	if v, ok := d.GetOk("source_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("source_address", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnsslsettingsAuthenticationRuleSourceAddress(d, v, "source_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SourceAddress = t
		}
	} else if d.HasChange("source_address") {
		old, new := d.GetChange("source_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SourceAddress = &[]models.VpnsslsettingsAuthenticationRuleSourceAddress{}
		}
	}
	if v1, ok := d.GetOk("source_address_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_address_negate", sv)
				diags = append(diags, e)
			}
			obj.SourceAddressNegate = &v2
		}
	}
	if v, ok := d.GetOk("source_address6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("source_address6", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnsslsettingsAuthenticationRuleSourceAddress6(d, v, "source_address6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SourceAddress6 = t
		}
	} else if d.HasChange("source_address6") {
		old, new := d.GetChange("source_address6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SourceAddress6 = &[]models.VpnsslsettingsAuthenticationRuleSourceAddress6{}
		}
	}
	if v1, ok := d.GetOk("source_address6_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_address6_negate", sv)
				diags = append(diags, e)
			}
			obj.SourceAddress6Negate = &v2
		}
	}
	if v, ok := d.GetOk("source_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("source_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnsslsettingsAuthenticationRuleSourceInterface(d, v, "source_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SourceInterface = t
		}
	} else if d.HasChange("source_interface") {
		old, new := d.GetChange("source_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SourceInterface = &[]models.VpnsslsettingsAuthenticationRuleSourceInterface{}
		}
	}
	if v1, ok := d.GetOk("user_peer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_peer", sv)
				diags = append(diags, e)
			}
			obj.UserPeer = &v2
		}
	}
	if v, ok := d.GetOk("users"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("users", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnsslsettingsAuthenticationRuleUsers(d, v, "users", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Users = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Users = &[]models.VpnsslsettingsAuthenticationRuleUsers{}
		}
	}
	return &obj, diags
}
