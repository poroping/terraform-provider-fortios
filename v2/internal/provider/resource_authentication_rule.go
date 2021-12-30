// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Authentication Rules.

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

func resourceAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Authentication Rules.",

		CreateContext: resourceAuthenticationRuleCreate,
		ReadContext:   resourceAuthenticationRuleRead,
		UpdateContext: resourceAuthenticationRuleUpdate,
		DeleteContext: resourceAuthenticationRuleDelete,

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
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"active_auth_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Select an active authentication method.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Select an IPv4 destination address from available options. Required for web proxy authentication.",
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
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "Select an IPv6 destination address from available options. Required for web proxy authentication.",
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
			"ip_based": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IP-based authentication. When enabled, previously authenticated users from the same IP address will be exempted.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication rule name.",
				ForceNew:    true,
				Required:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"http", "ftp", "socks", "ssh"}, false),

				Description: "Authentication is required for the selected protocol (default = HTTP).",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Authentication is required for the selected IPv4 source address.",
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
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "Authentication is required for the selected IPv6 source address.",
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
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Incoming (ingress) interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"sso_auth_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Select a single-sign on (SSO) authentication method.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this authentication rule.",
				Optional:    true,
				Computed:    true,
			},
			"transaction_based": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable transaction based authentication (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"web_auth_cookie": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Web authentication cookies (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"web_portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable web portal for proxy transparent policy (default = enable).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceAuthenticationRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating AuthenticationRule resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectAuthenticationRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateAuthenticationRule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AuthenticationRule")
	}

	return resourceAuthenticationRuleRead(ctx, d, meta)
}

func resourceAuthenticationRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAuthenticationRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateAuthenticationRule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating AuthenticationRule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AuthenticationRule")
	}

	return resourceAuthenticationRuleRead(ctx, d, meta)
}

func resourceAuthenticationRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteAuthenticationRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting AuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAuthenticationRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AuthenticationRule resource: %v", err)
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

	diags := refreshObjectAuthenticationRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenAuthenticationRuleDstaddr(v *[]models.AuthenticationRuleDstaddr, sort bool) interface{} {
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

func flattenAuthenticationRuleDstaddr6(v *[]models.AuthenticationRuleDstaddr6, sort bool) interface{} {
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

func flattenAuthenticationRuleSrcaddr(v *[]models.AuthenticationRuleSrcaddr, sort bool) interface{} {
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

func flattenAuthenticationRuleSrcaddr6(v *[]models.AuthenticationRuleSrcaddr6, sort bool) interface{} {
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

func flattenAuthenticationRuleSrcintf(v *[]models.AuthenticationRuleSrcintf, sort bool) interface{} {
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

func refreshObjectAuthenticationRule(d *schema.ResourceData, o *models.AuthenticationRule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ActiveAuthMethod != nil {
		v := *o.ActiveAuthMethod

		if err = d.Set("active_auth_method", v); err != nil {
			return diag.Errorf("error reading active_auth_method: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenAuthenticationRuleDstaddr(o.Dstaddr, sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.Dstaddr6 != nil {
		if err = d.Set("dstaddr6", flattenAuthenticationRuleDstaddr6(o.Dstaddr6, sort)); err != nil {
			return diag.Errorf("error reading dstaddr6: %v", err)
		}
	}

	if o.IpBased != nil {
		v := *o.IpBased

		if err = d.Set("ip_based", v); err != nil {
			return diag.Errorf("error reading ip_based: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenAuthenticationRuleSrcaddr(o.Srcaddr, sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.Srcaddr6 != nil {
		if err = d.Set("srcaddr6", flattenAuthenticationRuleSrcaddr6(o.Srcaddr6, sort)); err != nil {
			return diag.Errorf("error reading srcaddr6: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenAuthenticationRuleSrcintf(o.Srcintf, sort)); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.SsoAuthMethod != nil {
		v := *o.SsoAuthMethod

		if err = d.Set("sso_auth_method", v); err != nil {
			return diag.Errorf("error reading sso_auth_method: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TransactionBased != nil {
		v := *o.TransactionBased

		if err = d.Set("transaction_based", v); err != nil {
			return diag.Errorf("error reading transaction_based: %v", err)
		}
	}

	if o.WebAuthCookie != nil {
		v := *o.WebAuthCookie

		if err = d.Set("web_auth_cookie", v); err != nil {
			return diag.Errorf("error reading web_auth_cookie: %v", err)
		}
	}

	if o.WebPortal != nil {
		v := *o.WebPortal

		if err = d.Set("web_portal", v); err != nil {
			return diag.Errorf("error reading web_portal: %v", err)
		}
	}

	return nil
}

func expandAuthenticationRuleDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationRuleDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationRuleDstaddr

	for i := range l {
		tmp := models.AuthenticationRuleDstaddr{}
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

func expandAuthenticationRuleDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationRuleDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationRuleDstaddr6

	for i := range l {
		tmp := models.AuthenticationRuleDstaddr6{}
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

func expandAuthenticationRuleSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationRuleSrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationRuleSrcaddr

	for i := range l {
		tmp := models.AuthenticationRuleSrcaddr{}
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

func expandAuthenticationRuleSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationRuleSrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationRuleSrcaddr6

	for i := range l {
		tmp := models.AuthenticationRuleSrcaddr6{}
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

func expandAuthenticationRuleSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationRuleSrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationRuleSrcintf

	for i := range l {
		tmp := models.AuthenticationRuleSrcintf{}
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

func getObjectAuthenticationRule(d *schema.ResourceData, sv string) (*models.AuthenticationRule, diag.Diagnostics) {
	obj := models.AuthenticationRule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("active_auth_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("active_auth_method", sv)
				diags = append(diags, e)
			}
			obj.ActiveAuthMethod = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationRuleDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.AuthenticationRuleDstaddr{}
		}
	}
	if v, ok := d.GetOk("dstaddr6"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("dstaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationRuleDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr6 = t
		}
	} else if d.HasChange("dstaddr6") {
		old, new := d.GetChange("dstaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr6 = &[]models.AuthenticationRuleDstaddr6{}
		}
	}
	if v1, ok := d.GetOk("ip_based"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_based", sv)
				diags = append(diags, e)
			}
			obj.IpBased = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationRuleSrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.AuthenticationRuleSrcaddr{}
		}
	}
	if v, ok := d.GetOk("srcaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationRuleSrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr6 = t
		}
	} else if d.HasChange("srcaddr6") {
		old, new := d.GetChange("srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr6 = &[]models.AuthenticationRuleSrcaddr6{}
		}
	}
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationRuleSrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.AuthenticationRuleSrcintf{}
		}
	}
	if v1, ok := d.GetOk("sso_auth_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sso_auth_method", sv)
				diags = append(diags, e)
			}
			obj.SsoAuthMethod = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("transaction_based"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("transaction_based", sv)
				diags = append(diags, e)
			}
			obj.TransactionBased = &v2
		}
	}
	if v1, ok := d.GetOk("web_auth_cookie"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_auth_cookie", sv)
				diags = append(diags, e)
			}
			obj.WebAuthCookie = &v2
		}
	}
	if v1, ok := d.GetOk("web_portal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_portal", sv)
				diags = append(diags, e)
			}
			obj.WebPortal = &v2
		}
	}
	return &obj, diags
}
