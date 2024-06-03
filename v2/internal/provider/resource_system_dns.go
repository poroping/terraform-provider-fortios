// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DNS.

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

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DNS.",

		CreateContext: resourceSystemDnsCreate,
		ReadContext:   resourceSystemDnsRead,
		UpdateContext: resourceSystemDnsUpdate,
		DeleteContext: resourceSystemDnsDelete,

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
			"alt_primary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Alternate primary DNS server. This is not used as a failover DNS server.",
				Optional:    true,
				Computed:    true,
			},
			"alt_secondary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Alternate secondary DNS server. This is not used as a failover DNS server.",
				Optional:    true,
				Computed:    true,
			},
			"cache_notfound_responses": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable response from the DNS server when a record is not in cache.",
				Optional:    true,
				Computed:    true,
			},
			"dns_cache_limit": {
				Type: schema.TypeInt,

				Description: "Maximum number of records in the DNS cache.",
				Optional:    true,
				Computed:    true,
			},
			"dns_cache_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),

				Description: "Duration in seconds that the DNS cache retains information.",
				Optional:    true,
				Computed:    true,
			},
			"dns_over_tls": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "enforce"}, false),

				Description: "Enable/disable/enforce DNS over TLS.",
				Optional:    true,
				Computed:    true,
			},
			"domain": {
				Type:        schema.TypeList,
				Description: "Search suffix list for hostname lookup.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "DNS search domain list separated by space (maximum 8 domains).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fqdn_cache_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "FQDN cache time to live in seconds (0 - 86400, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"fqdn_min_refresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),

				Description: "FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"ip6_primary": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Primary DNS server IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_secondary": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Secondary DNS server IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "error", "all"}, false),

				Description: "Local DNS log setting.",
				Optional:    true,
				Computed:    true,
			},
			"primary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Primary DNS server IP address.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "DNS transport protocols.",
				Optional:         true,
				Computed:         true,
			},
			"retry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 5),

				Description: "Number of times to retry (0 - 5).",
				Optional:    true,
				Computed:    true,
			},
			"secondary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Secondary DNS server IP address.",
				Optional:    true,
				Computed:    true,
			},
			"server_hostname": {
				Type:        schema.TypeList,
				Description: "DNS server host name list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "DNS server host name list separated by space (maximum 4 domains).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"least-rtt", "failover"}, false),

				Description: "Specify how configured servers are prioritized.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address used by the DNS server as its source IP.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of local certificate for SSL connections.",
				Optional:    true,
				Computed:    true,
			},
			"timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "DNS query timeout interval in seconds (1 - 10).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemDns(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDns")
	}

	return resourceSystemDnsRead(ctx, d, meta)
}

func resourceSystemDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemDns(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemDns resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDns")
	}

	return resourceSystemDnsRead(ctx, d, meta)
}

func resourceSystemDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemDns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemDns(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDns(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDns resource: %v", err)
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

	diags := refreshObjectSystemDns(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemDnsDomain(d *schema.ResourceData, v *[]models.SystemDnsDomain, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Domain; tmp != nil {
				v["domain"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "domain")
	}

	return flat
}

func flattenSystemDnsServerHostname(d *schema.ResourceData, v *[]models.SystemDnsServerHostname, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Hostname; tmp != nil {
				v["hostname"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "hostname")
	}

	return flat
}

func refreshObjectSystemDns(d *schema.ResourceData, o *models.SystemDns, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AltPrimary != nil {
		v := *o.AltPrimary

		if err = d.Set("alt_primary", v); err != nil {
			return diag.Errorf("error reading alt_primary: %v", err)
		}
	}

	if o.AltSecondary != nil {
		v := *o.AltSecondary

		if err = d.Set("alt_secondary", v); err != nil {
			return diag.Errorf("error reading alt_secondary: %v", err)
		}
	}

	if o.CacheNotfoundResponses != nil {
		v := *o.CacheNotfoundResponses

		if err = d.Set("cache_notfound_responses", v); err != nil {
			return diag.Errorf("error reading cache_notfound_responses: %v", err)
		}
	}

	if o.DnsCacheLimit != nil {
		v := *o.DnsCacheLimit

		if err = d.Set("dns_cache_limit", v); err != nil {
			return diag.Errorf("error reading dns_cache_limit: %v", err)
		}
	}

	if o.DnsCacheTtl != nil {
		v := *o.DnsCacheTtl

		if err = d.Set("dns_cache_ttl", v); err != nil {
			return diag.Errorf("error reading dns_cache_ttl: %v", err)
		}
	}

	if o.DnsOverTls != nil {
		v := *o.DnsOverTls

		if err = d.Set("dns_over_tls", v); err != nil {
			return diag.Errorf("error reading dns_over_tls: %v", err)
		}
	}

	if o.Domain != nil {
		if err = d.Set("domain", flattenSystemDnsDomain(d, o.Domain, "domain", sort)); err != nil {
			return diag.Errorf("error reading domain: %v", err)
		}
	}

	if o.FqdnCacheTtl != nil {
		v := *o.FqdnCacheTtl

		if err = d.Set("fqdn_cache_ttl", v); err != nil {
			return diag.Errorf("error reading fqdn_cache_ttl: %v", err)
		}
	}

	if o.FqdnMinRefresh != nil {
		v := *o.FqdnMinRefresh

		if err = d.Set("fqdn_min_refresh", v); err != nil {
			return diag.Errorf("error reading fqdn_min_refresh: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.Ip6Primary != nil {
		v := *o.Ip6Primary

		if err = d.Set("ip6_primary", v); err != nil {
			return diag.Errorf("error reading ip6_primary: %v", err)
		}
	}

	if o.Ip6Secondary != nil {
		v := *o.Ip6Secondary

		if err = d.Set("ip6_secondary", v); err != nil {
			return diag.Errorf("error reading ip6_secondary: %v", err)
		}
	}

	if o.Log != nil {
		v := *o.Log

		if err = d.Set("log", v); err != nil {
			return diag.Errorf("error reading log: %v", err)
		}
	}

	if o.Primary != nil {
		v := *o.Primary

		if err = d.Set("primary", v); err != nil {
			return diag.Errorf("error reading primary: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Retry != nil {
		v := *o.Retry

		if err = d.Set("retry", v); err != nil {
			return diag.Errorf("error reading retry: %v", err)
		}
	}

	if o.Secondary != nil {
		v := *o.Secondary

		if err = d.Set("secondary", v); err != nil {
			return diag.Errorf("error reading secondary: %v", err)
		}
	}

	if o.ServerHostname != nil {
		if err = d.Set("server_hostname", flattenSystemDnsServerHostname(d, o.ServerHostname, "server_hostname", sort)); err != nil {
			return diag.Errorf("error reading server_hostname: %v", err)
		}
	}

	if o.ServerSelectMethod != nil {
		v := *o.ServerSelectMethod

		if err = d.Set("server_select_method", v); err != nil {
			return diag.Errorf("error reading server_select_method: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SslCertificate != nil {
		v := *o.SslCertificate

		if err = d.Set("ssl_certificate", v); err != nil {
			return diag.Errorf("error reading ssl_certificate: %v", err)
		}
	}

	if o.Timeout != nil {
		v := *o.Timeout

		if err = d.Set("timeout", v); err != nil {
			return diag.Errorf("error reading timeout: %v", err)
		}
	}

	return nil
}

func expandSystemDnsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDnsDomain, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDnsDomain

	for i := range l {
		tmp := models.SystemDnsDomain{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.domain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Domain = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDnsServerHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDnsServerHostname, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDnsServerHostname

	for i := range l {
		tmp := models.SystemDnsServerHostname{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.hostname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Hostname = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemDns(d *schema.ResourceData, sv string) (*models.SystemDns, diag.Diagnostics) {
	obj := models.SystemDns{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("alt_primary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("alt_primary", sv)
				diags = append(diags, e)
			}
			obj.AltPrimary = &v2
		}
	}
	if v1, ok := d.GetOk("alt_secondary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("alt_secondary", sv)
				diags = append(diags, e)
			}
			obj.AltSecondary = &v2
		}
	}
	if v1, ok := d.GetOk("cache_notfound_responses"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_notfound_responses", sv)
				diags = append(diags, e)
			}
			obj.CacheNotfoundResponses = &v2
		}
	}
	if v1, ok := d.GetOk("dns_cache_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_cache_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DnsCacheLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("dns_cache_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_cache_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DnsCacheTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("dns_over_tls"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("dns_over_tls", sv)
				diags = append(diags, e)
			}
			obj.DnsOverTls = &v2
		}
	}
	if v, ok := d.GetOk("domain"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("domain", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDnsDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Domain = t
		}
	} else if d.HasChange("domain") {
		old, new := d.GetChange("domain")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Domain = &[]models.SystemDnsDomain{}
		}
	}
	if v1, ok := d.GetOk("fqdn_cache_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.1", "") {
				e := utils.AttributeVersionWarning("fqdn_cache_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FqdnCacheTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("fqdn_min_refresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.1", "") {
				e := utils.AttributeVersionWarning("fqdn_min_refresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FqdnMinRefresh = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_primary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_primary", sv)
				diags = append(diags, e)
			}
			obj.Ip6Primary = &v2
		}
	}
	if v1, ok := d.GetOk("ip6_secondary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6_secondary", sv)
				diags = append(diags, e)
			}
			obj.Ip6Secondary = &v2
		}
	}
	if v1, ok := d.GetOk("log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("log", sv)
				diags = append(diags, e)
			}
			obj.Log = &v2
		}
	}
	if v1, ok := d.GetOk("primary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("primary", sv)
				diags = append(diags, e)
			}
			obj.Primary = &v2
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v1, ok := d.GetOk("retry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("retry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Retry = &tmp
		}
	}
	if v1, ok := d.GetOk("secondary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary", sv)
				diags = append(diags, e)
			}
			obj.Secondary = &v2
		}
	}
	if v, ok := d.GetOk("server_hostname"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("server_hostname", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDnsServerHostname(d, v, "server_hostname", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerHostname = t
		}
	} else if d.HasChange("server_hostname") {
		old, new := d.GetChange("server_hostname")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerHostname = &[]models.SystemDnsServerHostname{}
		}
	}
	if v1, ok := d.GetOk("server_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("server_select_method", sv)
				diags = append(diags, e)
			}
			obj.ServerSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_certificate", sv)
				diags = append(diags, e)
			}
			obj.SslCertificate = &v2
		}
	}
	if v1, ok := d.GetOk("timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Timeout = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemDns(d *schema.ResourceData, sv string) (*models.SystemDns, diag.Diagnostics) {
	obj := models.SystemDns{}
	diags := diag.Diagnostics{}

	obj.Domain = &[]models.SystemDnsDomain{}
	obj.ServerHostname = &[]models.SystemDnsServerHostname{}

	return &obj, diags
}
