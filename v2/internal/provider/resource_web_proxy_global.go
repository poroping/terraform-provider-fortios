// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Web proxy global settings.

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

func resourceWebProxyGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Web proxy global settings.",

		CreateContext: resourceWebProxyGlobalCreate,
		ReadContext:   resourceWebProxyGlobalRead,
		UpdateContext: resourceWebProxyGlobalUpdate,
		DeleteContext: resourceWebProxyGlobalDelete,

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
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fast_policy_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable fast matching algorithm for explicit and transparent proxy policy.",
				Optional:    true,
				Computed:    true,
			},
			"forward_proxy_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable forwarding proxy authentication headers.",
				Optional:    true,
				Computed:    true,
			},
			"forward_server_affinity_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 60),

				Description: "Period of time before the source IP's traffic is no longer assigned to the forwarding server (6 - 60 min, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"ldap_user_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable LDAP user cache for explicit and transparent proxy user.",
				Optional:    true,
				Computed:    true,
			},
			"learn_client_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable learning the client's IP address from headers.",
				Optional:    true,
				Computed:    true,
			},
			"learn_client_ip_from_header": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Learn client IP address from the specified headers.",
				Optional:         true,
				Computed:         true,
			},
			"learn_client_ip_srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name (srcaddr or srcaddr6 must be set).",
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
			"learn_client_ip_srcaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 Source address name (srcaddr or srcaddr6 must be set).",
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
			"max_message_length": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(16, 256),

				Description: "Maximum length of HTTP message, not including body (16 - 256 Kbytes, default = 32).",
				Optional:    true,
				Computed:    true,
			},
			"max_request_length": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 64),

				Description: "Maximum length of HTTP request line (2 - 64 Kbytes, default = 8).",
				Optional:    true,
				Computed:    true,
			},
			"max_waf_body_cache_length": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1024),

				Description: "Maximum length of HTTP messages processed by Web Application Firewall (WAF) (10 - 1024 Kbytes, default = 32).",
				Optional:    true,
				Computed:    true,
			},
			"proxy_fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fully Qualified Domain Name (FQDN) that clients connect to (default = default.fqdn) to connect to the explicit web proxy.",
				Optional:    true,
				Computed:    true,
			},
			"src_affinity_exempt_addr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 source addresses to exempt proxy affinity.",
				Optional:    true,
				Computed:    true,
			},
			"src_affinity_exempt_addr6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 source addresses to exempt proxy affinity.",
				Optional:         true,
				Computed:         true,
			},
			"ssl_ca_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSL CA certificate for SSL interception.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSL certificate for SSL interception.",
				Optional:    true,
				Computed:    true,
			},
			"strict_web_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable strict web checking to block web sites that send incorrect headers that don't conform to HTTP 1.1.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_non_http": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing non-HTTP traffic. Allowed non-HTTP traffic is tunneled.",
				Optional:    true,
				Computed:    true,
			},
			"unknown_http_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"reject", "tunnel", "best-effort"}, false),

				Description: "Action to take when an unknown version of HTTP is encountered: reject, allow (tunnel), or proceed with best-effort.",
				Optional:    true,
				Computed:    true,
			},
			"webproxy_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name of the web proxy profile to apply when explicit proxy traffic is allowed by default and traffic is accepted that does not match an explicit proxy policy.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebProxyGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebProxyGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebProxyGlobal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyGlobal")
	}

	return resourceWebProxyGlobalRead(ctx, d, meta)
}

func resourceWebProxyGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebProxyGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebProxyGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebProxyGlobal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebProxyGlobal")
	}

	return resourceWebProxyGlobalRead(ctx, d, meta)
}

func resourceWebProxyGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebProxyGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebProxyGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyGlobal resource: %v", err)
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

	diags := refreshObjectWebProxyGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWebProxyGlobalLearnClientIpSrcaddr(v *[]models.WebProxyGlobalLearnClientIpSrcaddr, sort bool) interface{} {
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

func flattenWebProxyGlobalLearnClientIpSrcaddr6(v *[]models.WebProxyGlobalLearnClientIpSrcaddr6, sort bool) interface{} {
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

func refreshObjectWebProxyGlobal(d *schema.ResourceData, o *models.WebProxyGlobal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FastPolicyMatch != nil {
		v := *o.FastPolicyMatch

		if err = d.Set("fast_policy_match", v); err != nil {
			return diag.Errorf("error reading fast_policy_match: %v", err)
		}
	}

	if o.ForwardProxyAuth != nil {
		v := *o.ForwardProxyAuth

		if err = d.Set("forward_proxy_auth", v); err != nil {
			return diag.Errorf("error reading forward_proxy_auth: %v", err)
		}
	}

	if o.ForwardServerAffinityTimeout != nil {
		v := *o.ForwardServerAffinityTimeout

		if err = d.Set("forward_server_affinity_timeout", v); err != nil {
			return diag.Errorf("error reading forward_server_affinity_timeout: %v", err)
		}
	}

	if o.LdapUserCache != nil {
		v := *o.LdapUserCache

		if err = d.Set("ldap_user_cache", v); err != nil {
			return diag.Errorf("error reading ldap_user_cache: %v", err)
		}
	}

	if o.LearnClientIp != nil {
		v := *o.LearnClientIp

		if err = d.Set("learn_client_ip", v); err != nil {
			return diag.Errorf("error reading learn_client_ip: %v", err)
		}
	}

	if o.LearnClientIpFromHeader != nil {
		v := *o.LearnClientIpFromHeader

		if err = d.Set("learn_client_ip_from_header", v); err != nil {
			return diag.Errorf("error reading learn_client_ip_from_header: %v", err)
		}
	}

	if o.LearnClientIpSrcaddr != nil {
		if err = d.Set("learn_client_ip_srcaddr", flattenWebProxyGlobalLearnClientIpSrcaddr(o.LearnClientIpSrcaddr, sort)); err != nil {
			return diag.Errorf("error reading learn_client_ip_srcaddr: %v", err)
		}
	}

	if o.LearnClientIpSrcaddr6 != nil {
		if err = d.Set("learn_client_ip_srcaddr6", flattenWebProxyGlobalLearnClientIpSrcaddr6(o.LearnClientIpSrcaddr6, sort)); err != nil {
			return diag.Errorf("error reading learn_client_ip_srcaddr6: %v", err)
		}
	}

	if o.MaxMessageLength != nil {
		v := *o.MaxMessageLength

		if err = d.Set("max_message_length", v); err != nil {
			return diag.Errorf("error reading max_message_length: %v", err)
		}
	}

	if o.MaxRequestLength != nil {
		v := *o.MaxRequestLength

		if err = d.Set("max_request_length", v); err != nil {
			return diag.Errorf("error reading max_request_length: %v", err)
		}
	}

	if o.MaxWafBodyCacheLength != nil {
		v := *o.MaxWafBodyCacheLength

		if err = d.Set("max_waf_body_cache_length", v); err != nil {
			return diag.Errorf("error reading max_waf_body_cache_length: %v", err)
		}
	}

	if o.ProxyFqdn != nil {
		v := *o.ProxyFqdn

		if err = d.Set("proxy_fqdn", v); err != nil {
			return diag.Errorf("error reading proxy_fqdn: %v", err)
		}
	}

	if o.SrcAffinityExemptAddr != nil {
		v := *o.SrcAffinityExemptAddr

		if err = d.Set("src_affinity_exempt_addr", v); err != nil {
			return diag.Errorf("error reading src_affinity_exempt_addr: %v", err)
		}
	}

	if o.SrcAffinityExemptAddr6 != nil {
		v := *o.SrcAffinityExemptAddr6

		if err = d.Set("src_affinity_exempt_addr6", v); err != nil {
			return diag.Errorf("error reading src_affinity_exempt_addr6: %v", err)
		}
	}

	if o.SslCaCert != nil {
		v := *o.SslCaCert

		if err = d.Set("ssl_ca_cert", v); err != nil {
			return diag.Errorf("error reading ssl_ca_cert: %v", err)
		}
	}

	if o.SslCert != nil {
		v := *o.SslCert

		if err = d.Set("ssl_cert", v); err != nil {
			return diag.Errorf("error reading ssl_cert: %v", err)
		}
	}

	if o.StrictWebCheck != nil {
		v := *o.StrictWebCheck

		if err = d.Set("strict_web_check", v); err != nil {
			return diag.Errorf("error reading strict_web_check: %v", err)
		}
	}

	if o.TunnelNonHttp != nil {
		v := *o.TunnelNonHttp

		if err = d.Set("tunnel_non_http", v); err != nil {
			return diag.Errorf("error reading tunnel_non_http: %v", err)
		}
	}

	if o.UnknownHttpVersion != nil {
		v := *o.UnknownHttpVersion

		if err = d.Set("unknown_http_version", v); err != nil {
			return diag.Errorf("error reading unknown_http_version: %v", err)
		}
	}

	if o.WebproxyProfile != nil {
		v := *o.WebproxyProfile

		if err = d.Set("webproxy_profile", v); err != nil {
			return diag.Errorf("error reading webproxy_profile: %v", err)
		}
	}

	return nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyGlobalLearnClientIpSrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyGlobalLearnClientIpSrcaddr

	for i := range l {
		tmp := models.WebProxyGlobalLearnClientIpSrcaddr{}
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

func expandWebProxyGlobalLearnClientIpSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebProxyGlobalLearnClientIpSrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebProxyGlobalLearnClientIpSrcaddr6

	for i := range l {
		tmp := models.WebProxyGlobalLearnClientIpSrcaddr6{}
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

func getObjectWebProxyGlobal(d *schema.ResourceData, sv string) (*models.WebProxyGlobal, diag.Diagnostics) {
	obj := models.WebProxyGlobal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fast_policy_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fast_policy_match", sv)
				diags = append(diags, e)
			}
			obj.FastPolicyMatch = &v2
		}
	}
	if v1, ok := d.GetOk("forward_proxy_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_proxy_auth", sv)
				diags = append(diags, e)
			}
			obj.ForwardProxyAuth = &v2
		}
	}
	if v1, ok := d.GetOk("forward_server_affinity_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_server_affinity_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ForwardServerAffinityTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ldap_user_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ldap_user_cache", sv)
				diags = append(diags, e)
			}
			obj.LdapUserCache = &v2
		}
	}
	if v1, ok := d.GetOk("learn_client_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("learn_client_ip", sv)
				diags = append(diags, e)
			}
			obj.LearnClientIp = &v2
		}
	}
	if v1, ok := d.GetOk("learn_client_ip_from_header"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("learn_client_ip_from_header", sv)
				diags = append(diags, e)
			}
			obj.LearnClientIpFromHeader = &v2
		}
	}
	if v, ok := d.GetOk("learn_client_ip_srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("learn_client_ip_srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandWebProxyGlobalLearnClientIpSrcaddr(d, v, "learn_client_ip_srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LearnClientIpSrcaddr = t
		}
	} else if d.HasChange("learn_client_ip_srcaddr") {
		old, new := d.GetChange("learn_client_ip_srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LearnClientIpSrcaddr = &[]models.WebProxyGlobalLearnClientIpSrcaddr{}
		}
	}
	if v, ok := d.GetOk("learn_client_ip_srcaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("learn_client_ip_srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandWebProxyGlobalLearnClientIpSrcaddr6(d, v, "learn_client_ip_srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LearnClientIpSrcaddr6 = t
		}
	} else if d.HasChange("learn_client_ip_srcaddr6") {
		old, new := d.GetChange("learn_client_ip_srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LearnClientIpSrcaddr6 = &[]models.WebProxyGlobalLearnClientIpSrcaddr6{}
		}
	}
	if v1, ok := d.GetOk("max_message_length"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_message_length", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxMessageLength = &tmp
		}
	}
	if v1, ok := d.GetOk("max_request_length"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_request_length", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxRequestLength = &tmp
		}
	}
	if v1, ok := d.GetOk("max_waf_body_cache_length"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_waf_body_cache_length", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxWafBodyCacheLength = &tmp
		}
	}
	if v1, ok := d.GetOk("proxy_fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_fqdn", sv)
				diags = append(diags, e)
			}
			obj.ProxyFqdn = &v2
		}
	}
	if v1, ok := d.GetOk("src_affinity_exempt_addr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("src_affinity_exempt_addr", sv)
				diags = append(diags, e)
			}
			obj.SrcAffinityExemptAddr = &v2
		}
	}
	if v1, ok := d.GetOk("src_affinity_exempt_addr6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("src_affinity_exempt_addr6", sv)
				diags = append(diags, e)
			}
			obj.SrcAffinityExemptAddr6 = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_ca_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_ca_cert", sv)
				diags = append(diags, e)
			}
			obj.SslCaCert = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_cert", sv)
				diags = append(diags, e)
			}
			obj.SslCert = &v2
		}
	}
	if v1, ok := d.GetOk("strict_web_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strict_web_check", sv)
				diags = append(diags, e)
			}
			obj.StrictWebCheck = &v2
		}
	}
	if v1, ok := d.GetOk("tunnel_non_http"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("tunnel_non_http", sv)
				diags = append(diags, e)
			}
			obj.TunnelNonHttp = &v2
		}
	}
	if v1, ok := d.GetOk("unknown_http_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("unknown_http_version", sv)
				diags = append(diags, e)
			}
			obj.UnknownHttpVersion = &v2
		}
	}
	if v1, ok := d.GetOk("webproxy_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webproxy_profile", sv)
				diags = append(diags, e)
			}
			obj.WebproxyProfile = &v2
		}
	}
	return &obj, diags
}
