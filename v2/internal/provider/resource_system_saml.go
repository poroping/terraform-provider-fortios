// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global settings for SAML authentication.

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

func resourceSystemSaml() *schema.Resource {
	return &schema.Resource{
		Description: "Global settings for SAML authentication.",

		CreateContext: resourceSystemSamlCreate,
		ReadContext:   resourceSystemSamlRead,
		UpdateContext: resourceSystemSamlUpdate,
		DeleteContext: resourceSystemSamlDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"binding_protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"post", "redirect"}, false),

				Description: "IdP Binding protocol.",
				Optional:    true,
				Computed:    true,
			},
			"cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Certificate to sign SAML messages.",
				Optional:    true,
				Computed:    true,
			},
			"default_login_page": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"normal", "sso"}, false),

				Description: "Choose default login page.",
				Optional:    true,
				Computed:    true,
			},
			"default_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Default profile for new SSO admin.",
				Optional:    true,
				Computed:    true,
			},
			"entity_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SP entity ID.",
				Optional:    true,
				Computed:    true,
			},
			"idp_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IDP certificate name.",
				Optional:    true,
				Computed:    true,
			},
			"idp_entity_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IDP entity ID.",
				Optional:    true,
				Computed:    true,
			},
			"idp_single_logout_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IDP single logout URL.",
				Optional:    true,
				Computed:    true,
			},
			"idp_single_sign_on_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IDP single sign-on URL.",
				Optional:    true,
				Computed:    true,
			},
			"life": {
				Type: schema.TypeInt,

				Description: "Length of the range of time when the assertion is valid (in minutes).",
				Optional:    true,
				Computed:    true,
			},
			"portal_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SP portal URL.",
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"identity-provider", "service-provider"}, false),

				Description: "SAML role.",
				Optional:    true,
				Computed:    true,
			},
			"server_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Server address.",
				Optional:    true,
				Computed:    true,
			},
			"service_providers": {
				Type:        schema.TypeList,
				Description: "Authorized service providers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assertion_attributes": {
							Type:        schema.TypeList,
							Description: "Customized SAML attributes to send along with assertion.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Name.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"username", "email", "profile-name"}, false),

										Description: "Type.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"idp_entity_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "IDP entity ID.",
							Optional:    true,
							Computed:    true,
						},
						"idp_single_logout_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "IDP single logout URL.",
							Optional:    true,
							Computed:    true,
						},
						"idp_single_sign_on_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "IDP single sign-on URL.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Prefix.",
							Optional:    true,
							Computed:    true,
						},
						"sp_binding_protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"post", "redirect"}, false),

							Description: "SP binding protocol.",
							Optional:    true,
							Computed:    true,
						},
						"sp_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SP certificate name.",
							Optional:    true,
							Computed:    true,
						},
						"sp_entity_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SP entity ID.",
							Optional:    true,
							Computed:    true,
						},
						"sp_portal_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SP portal URL.",
							Optional:    true,
							Computed:    true,
						},
						"sp_single_logout_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SP single logout URL.",
							Optional:    true,
							Computed:    true,
						},
						"sp_single_sign_on_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SP single sign-on URL.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"single_logout_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SP single logout URL.",
				Optional:    true,
				Computed:    true,
			},
			"single_sign_on_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SP single sign-on URL.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SAML authentication (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"tolerance": {
				Type: schema.TypeInt,

				Description: "Tolerance to the range of time when the assertion is valid (in minutes).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSamlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSaml(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSaml(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSaml")
	}

	return resourceSystemSamlRead(ctx, d, meta)
}

func resourceSystemSamlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSaml(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSaml(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSaml resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSaml")
	}

	return resourceSystemSamlRead(ctx, d, meta)
}

func resourceSystemSamlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemSaml(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemSaml(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSamlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadSystemSaml(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSaml resource: %v", err)
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

	diags := refreshObjectSystemSaml(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSamlServiceProviders(d *schema.ResourceData, v *[]models.SystemSamlServiceProviders, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AssertionAttributes; tmp != nil {
				v["assertion_attributes"] = flattenSystemSamlServiceProvidersAssertionAttributes(d, tmp, prefix+"assertion_attributes", sort)
			}

			if tmp := cfg.IdpEntityId; tmp != nil {
				v["idp_entity_id"] = *tmp
			}

			if tmp := cfg.IdpSingleLogoutUrl; tmp != nil {
				v["idp_single_logout_url"] = *tmp
			}

			if tmp := cfg.IdpSingleSignOnUrl; tmp != nil {
				v["idp_single_sign_on_url"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			if tmp := cfg.SpBindingProtocol; tmp != nil {
				v["sp_binding_protocol"] = *tmp
			}

			if tmp := cfg.SpCert; tmp != nil {
				v["sp_cert"] = *tmp
			}

			if tmp := cfg.SpEntityId; tmp != nil {
				v["sp_entity_id"] = *tmp
			}

			if tmp := cfg.SpPortalUrl; tmp != nil {
				v["sp_portal_url"] = *tmp
			}

			if tmp := cfg.SpSingleLogoutUrl; tmp != nil {
				v["sp_single_logout_url"] = *tmp
			}

			if tmp := cfg.SpSingleSignOnUrl; tmp != nil {
				v["sp_single_sign_on_url"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemSamlServiceProvidersAssertionAttributes(d *schema.ResourceData, v *[]models.SystemSamlServiceProvidersAssertionAttributes, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemSaml(d *schema.ResourceData, o *models.SystemSaml, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BindingProtocol != nil {
		v := *o.BindingProtocol

		if err = d.Set("binding_protocol", v); err != nil {
			return diag.Errorf("error reading binding_protocol: %v", err)
		}
	}

	if o.Cert != nil {
		v := *o.Cert

		if err = d.Set("cert", v); err != nil {
			return diag.Errorf("error reading cert: %v", err)
		}
	}

	if o.DefaultLoginPage != nil {
		v := *o.DefaultLoginPage

		if err = d.Set("default_login_page", v); err != nil {
			return diag.Errorf("error reading default_login_page: %v", err)
		}
	}

	if o.DefaultProfile != nil {
		v := *o.DefaultProfile

		if err = d.Set("default_profile", v); err != nil {
			return diag.Errorf("error reading default_profile: %v", err)
		}
	}

	if o.EntityId != nil {
		v := *o.EntityId

		if err = d.Set("entity_id", v); err != nil {
			return diag.Errorf("error reading entity_id: %v", err)
		}
	}

	if o.IdpCert != nil {
		v := *o.IdpCert

		if err = d.Set("idp_cert", v); err != nil {
			return diag.Errorf("error reading idp_cert: %v", err)
		}
	}

	if o.IdpEntityId != nil {
		v := *o.IdpEntityId

		if err = d.Set("idp_entity_id", v); err != nil {
			return diag.Errorf("error reading idp_entity_id: %v", err)
		}
	}

	if o.IdpSingleLogoutUrl != nil {
		v := *o.IdpSingleLogoutUrl

		if err = d.Set("idp_single_logout_url", v); err != nil {
			return diag.Errorf("error reading idp_single_logout_url: %v", err)
		}
	}

	if o.IdpSingleSignOnUrl != nil {
		v := *o.IdpSingleSignOnUrl

		if err = d.Set("idp_single_sign_on_url", v); err != nil {
			return diag.Errorf("error reading idp_single_sign_on_url: %v", err)
		}
	}

	if o.Life != nil {
		v := *o.Life

		if err = d.Set("life", v); err != nil {
			return diag.Errorf("error reading life: %v", err)
		}
	}

	if o.PortalUrl != nil {
		v := *o.PortalUrl

		if err = d.Set("portal_url", v); err != nil {
			return diag.Errorf("error reading portal_url: %v", err)
		}
	}

	if o.Role != nil {
		v := *o.Role

		if err = d.Set("role", v); err != nil {
			return diag.Errorf("error reading role: %v", err)
		}
	}

	if o.ServerAddress != nil {
		v := *o.ServerAddress

		if err = d.Set("server_address", v); err != nil {
			return diag.Errorf("error reading server_address: %v", err)
		}
	}

	if o.ServiceProviders != nil {
		if err = d.Set("service_providers", flattenSystemSamlServiceProviders(d, o.ServiceProviders, "service_providers", sort)); err != nil {
			return diag.Errorf("error reading service_providers: %v", err)
		}
	}

	if o.SingleLogoutUrl != nil {
		v := *o.SingleLogoutUrl

		if err = d.Set("single_logout_url", v); err != nil {
			return diag.Errorf("error reading single_logout_url: %v", err)
		}
	}

	if o.SingleSignOnUrl != nil {
		v := *o.SingleSignOnUrl

		if err = d.Set("single_sign_on_url", v); err != nil {
			return diag.Errorf("error reading single_sign_on_url: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Tolerance != nil {
		v := *o.Tolerance

		if err = d.Set("tolerance", v); err != nil {
			return diag.Errorf("error reading tolerance: %v", err)
		}
	}

	return nil
}

func expandSystemSamlServiceProviders(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSamlServiceProviders, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSamlServiceProviders

	for i := range l {
		tmp := models.SystemSamlServiceProviders{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.assertion_attributes", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSamlServiceProvidersAssertionAttributes(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSamlServiceProvidersAssertionAttributes
			// 	}
			tmp.AssertionAttributes = v2

		}

		pre_append = fmt.Sprintf("%s.%d.idp_entity_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IdpEntityId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.idp_single_logout_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IdpSingleLogoutUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.idp_single_sign_on_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IdpSingleSignOnUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sp_binding_protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpBindingProtocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sp_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sp_entity_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpEntityId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sp_portal_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpPortalUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sp_single_logout_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpSingleLogoutUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sp_single_sign_on_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpSingleSignOnUrl = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSamlServiceProvidersAssertionAttributes(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSamlServiceProvidersAssertionAttributes, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSamlServiceProvidersAssertionAttributes

	for i := range l {
		tmp := models.SystemSamlServiceProvidersAssertionAttributes{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemSaml(d *schema.ResourceData, sv string) (*models.SystemSaml, diag.Diagnostics) {
	obj := models.SystemSaml{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("binding_protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("binding_protocol", sv)
				diags = append(diags, e)
			}
			obj.BindingProtocol = &v2
		}
	}
	if v1, ok := d.GetOk("cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert", sv)
				diags = append(diags, e)
			}
			obj.Cert = &v2
		}
	}
	if v1, ok := d.GetOk("default_login_page"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_login_page", sv)
				diags = append(diags, e)
			}
			obj.DefaultLoginPage = &v2
		}
	}
	if v1, ok := d.GetOk("default_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_profile", sv)
				diags = append(diags, e)
			}
			obj.DefaultProfile = &v2
		}
	}
	if v1, ok := d.GetOk("entity_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("entity_id", sv)
				diags = append(diags, e)
			}
			obj.EntityId = &v2
		}
	}
	if v1, ok := d.GetOk("idp_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_cert", sv)
				diags = append(diags, e)
			}
			obj.IdpCert = &v2
		}
	}
	if v1, ok := d.GetOk("idp_entity_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_entity_id", sv)
				diags = append(diags, e)
			}
			obj.IdpEntityId = &v2
		}
	}
	if v1, ok := d.GetOk("idp_single_logout_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_single_logout_url", sv)
				diags = append(diags, e)
			}
			obj.IdpSingleLogoutUrl = &v2
		}
	}
	if v1, ok := d.GetOk("idp_single_sign_on_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idp_single_sign_on_url", sv)
				diags = append(diags, e)
			}
			obj.IdpSingleSignOnUrl = &v2
		}
	}
	if v1, ok := d.GetOk("life"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("life", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Life = &tmp
		}
	}
	if v1, ok := d.GetOk("portal_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portal_url", sv)
				diags = append(diags, e)
			}
			obj.PortalUrl = &v2
		}
	}
	if v1, ok := d.GetOk("role"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("role", sv)
				diags = append(diags, e)
			}
			obj.Role = &v2
		}
	}
	if v1, ok := d.GetOk("server_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_address", sv)
				diags = append(diags, e)
			}
			obj.ServerAddress = &v2
		}
	}
	if v, ok := d.GetOk("service_providers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service_providers", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSamlServiceProviders(d, v, "service_providers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServiceProviders = t
		}
	} else if d.HasChange("service_providers") {
		old, new := d.GetChange("service_providers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServiceProviders = &[]models.SystemSamlServiceProviders{}
		}
	}
	if v1, ok := d.GetOk("single_logout_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("single_logout_url", sv)
				diags = append(diags, e)
			}
			obj.SingleLogoutUrl = &v2
		}
	}
	if v1, ok := d.GetOk("single_sign_on_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("single_sign_on_url", sv)
				diags = append(diags, e)
			}
			obj.SingleSignOnUrl = &v2
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
	if v1, ok := d.GetOk("tolerance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tolerance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Tolerance = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemSaml(d *schema.ResourceData, sv string) (*models.SystemSaml, diag.Diagnostics) {
	obj := models.SystemSaml{}
	diags := diag.Diagnostics{}

	obj.ServiceProviders = &[]models.SystemSamlServiceProviders{}

	return &obj, diags
}
