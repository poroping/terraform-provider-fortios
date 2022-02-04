// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web proxy profiles.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWebProxyProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web proxy profiles.",

		CreateContext: resourceWebProxyProfileCreate,
		ReadContext:   resourceWebProxyProfileRead,
		UpdateContext: resourceWebProxyProfileUpdate,
		DeleteContext: resourceWebProxyProfileDelete,

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
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"header_client_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"header_front_end_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"header_via_request": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"header_via_response": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"header_x_authenticated_groups": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"header_x_authenticated_user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"header_x_forwarded_client_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"header_x_forwarded_for": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"headers": {
				Type:        schema.TypeList,
				Description: "Configure HTTP forwarded requests headers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"add-to-request", "add-to-response", "remove-from-request", "remove-from-response"}, false),

							Description: "Action when the HTTP header is forwarded.",
							Optional:    true,
							Computed:    true,
						},
						"add_option": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"append", "new-on-not-found", "new"}, false),

							Description: "Configure options to append content to existing HTTP header or add new HTTP header.",
							Optional:    true,
							Computed:    true,
						},
						"base64_encoding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable use of base64 encoding of HTTP content.",
							Optional:    true,
							Computed:    true,
						},
						"content": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "HTTP header content.",
							Optional:    true,
							Computed:    true,
						},
						"dstaddr": {
							Type:        schema.TypeList,
							Description: "Destination address and address group names.",
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
							Description: "Destination address and address group names (IPv6).",
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
						"id": {
							Type: schema.TypeInt,

							Description: "HTTP forwarded header id.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "HTTP forwarded header name.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both).",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"log_header_change": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging HTTP header changes.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"strip_encoding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable stripping unsupported encoding from the request header.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebProxyProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WebProxyProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWebProxyProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebProxyProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyProfile")
	}

	return resourceWebProxyProfileRead(ctx, d, meta)
}

func resourceWebProxyProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebProxyProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebProxyProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebProxyProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyProfile")
	}

	return resourceWebProxyProfileRead(ctx, d, meta)
}

func resourceWebProxyProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebProxyProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebProxyProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyProfile resource: %v", err)
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

	diags := refreshObjectWebProxyProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWebProxyProfileHeaders(v *[]models.WebProxyProfileHeaders, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.AddOption; tmp != nil {
				v["add_option"] = *tmp
			}

			if tmp := cfg.Base64Encoding; tmp != nil {
				v["base64_encoding"] = *tmp
			}

			if tmp := cfg.Content; tmp != nil {
				v["content"] = *tmp
			}

			if tmp := cfg.Dstaddr; tmp != nil {
				v["dstaddr"] = flattenWebProxyProfileHeadersDstaddr(tmp, sort)
			}

			if tmp := cfg.Dstaddr6; tmp != nil {
				v["dstaddr6"] = flattenWebProxyProfileHeadersDstaddr6(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWebProxyProfileHeadersDstaddr(v *[]models.WebProxyProfileHeadersDstaddr, sort bool) interface{} {
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

func flattenWebProxyProfileHeadersDstaddr6(v *[]models.WebProxyProfileHeadersDstaddr6, sort bool) interface{} {
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

func refreshObjectWebProxyProfile(d *schema.ResourceData, o *models.WebProxyProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.HeaderClientIp != nil {
		v := *o.HeaderClientIp

		if err = d.Set("header_client_ip", v); err != nil {
			return diag.Errorf("error reading header_client_ip: %v", err)
		}
	}

	if o.HeaderFrontEndHttps != nil {
		v := *o.HeaderFrontEndHttps

		if err = d.Set("header_front_end_https", v); err != nil {
			return diag.Errorf("error reading header_front_end_https: %v", err)
		}
	}

	if o.HeaderViaRequest != nil {
		v := *o.HeaderViaRequest

		if err = d.Set("header_via_request", v); err != nil {
			return diag.Errorf("error reading header_via_request: %v", err)
		}
	}

	if o.HeaderViaResponse != nil {
		v := *o.HeaderViaResponse

		if err = d.Set("header_via_response", v); err != nil {
			return diag.Errorf("error reading header_via_response: %v", err)
		}
	}

	if o.HeaderXAuthenticatedGroups != nil {
		v := *o.HeaderXAuthenticatedGroups

		if err = d.Set("header_x_authenticated_groups", v); err != nil {
			return diag.Errorf("error reading header_x_authenticated_groups: %v", err)
		}
	}

	if o.HeaderXAuthenticatedUser != nil {
		v := *o.HeaderXAuthenticatedUser

		if err = d.Set("header_x_authenticated_user", v); err != nil {
			return diag.Errorf("error reading header_x_authenticated_user: %v", err)
		}
	}

	if o.HeaderXForwardedClientCert != nil {
		v := *o.HeaderXForwardedClientCert

		if err = d.Set("header_x_forwarded_client_cert", v); err != nil {
			return diag.Errorf("error reading header_x_forwarded_client_cert: %v", err)
		}
	}

	if o.HeaderXForwardedFor != nil {
		v := *o.HeaderXForwardedFor

		if err = d.Set("header_x_forwarded_for", v); err != nil {
			return diag.Errorf("error reading header_x_forwarded_for: %v", err)
		}
	}

	if o.Headers != nil {
		if err = d.Set("headers", flattenWebProxyProfileHeaders(o.Headers, sort)); err != nil {
			return diag.Errorf("error reading headers: %v", err)
		}
	}

	if o.LogHeaderChange != nil {
		v := *o.LogHeaderChange

		if err = d.Set("log_header_change", v); err != nil {
			return diag.Errorf("error reading log_header_change: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.StripEncoding != nil {
		v := *o.StripEncoding

		if err = d.Set("strip_encoding", v); err != nil {
			return diag.Errorf("error reading strip_encoding: %v", err)
		}
	}

	return nil
}

func expandWebProxyProfileHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyProfileHeaders, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyProfileHeaders

	for i := range l {
		tmp := models.WebProxyProfileHeaders{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.add_option", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddOption = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.base64_encoding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Base64Encoding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Content = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebProxyProfileHeadersDstaddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebProxyProfileHeadersDstaddr
			// 	}
			tmp.Dstaddr = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWebProxyProfileHeadersDstaddr6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WebProxyProfileHeadersDstaddr6
			// 	}
			tmp.Dstaddr6 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWebProxyProfileHeadersDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyProfileHeadersDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyProfileHeadersDstaddr

	for i := range l {
		tmp := models.WebProxyProfileHeadersDstaddr{}
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

func expandWebProxyProfileHeadersDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyProfileHeadersDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyProfileHeadersDstaddr6

	for i := range l {
		tmp := models.WebProxyProfileHeadersDstaddr6{}
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

func getObjectWebProxyProfile(d *schema.ResourceData, sv string) (*models.WebProxyProfile, diag.Diagnostics) {
	obj := models.WebProxyProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("header_client_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_client_ip", sv)
				diags = append(diags, e)
			}
			obj.HeaderClientIp = &v2
		}
	}
	if v1, ok := d.GetOk("header_front_end_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_front_end_https", sv)
				diags = append(diags, e)
			}
			obj.HeaderFrontEndHttps = &v2
		}
	}
	if v1, ok := d.GetOk("header_via_request"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_via_request", sv)
				diags = append(diags, e)
			}
			obj.HeaderViaRequest = &v2
		}
	}
	if v1, ok := d.GetOk("header_via_response"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_via_response", sv)
				diags = append(diags, e)
			}
			obj.HeaderViaResponse = &v2
		}
	}
	if v1, ok := d.GetOk("header_x_authenticated_groups"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_x_authenticated_groups", sv)
				diags = append(diags, e)
			}
			obj.HeaderXAuthenticatedGroups = &v2
		}
	}
	if v1, ok := d.GetOk("header_x_authenticated_user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_x_authenticated_user", sv)
				diags = append(diags, e)
			}
			obj.HeaderXAuthenticatedUser = &v2
		}
	}
	if v1, ok := d.GetOk("header_x_forwarded_client_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("header_x_forwarded_client_cert", sv)
				diags = append(diags, e)
			}
			obj.HeaderXForwardedClientCert = &v2
		}
	}
	if v1, ok := d.GetOk("header_x_forwarded_for"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_x_forwarded_for", sv)
				diags = append(diags, e)
			}
			obj.HeaderXForwardedFor = &v2
		}
	}
	if v, ok := d.GetOk("headers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("headers", sv)
			diags = append(diags, e)
		}
		t, err := expandWebProxyProfileHeaders(d, v, "headers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Headers = t
		}
	} else if d.HasChange("headers") {
		old, new := d.GetChange("headers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Headers = &[]models.WebProxyProfileHeaders{}
		}
	}
	if v1, ok := d.GetOk("log_header_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_header_change", sv)
				diags = append(diags, e)
			}
			obj.LogHeaderChange = &v2
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
	if v1, ok := d.GetOk("strip_encoding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strip_encoding", sv)
				diags = append(diags, e)
			}
			obj.StripEncoding = &v2
		}
	}
	return &obj, diags
}
