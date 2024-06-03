// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DNS servers for a non-management VDOM.

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

func resourceSystemVdomDns() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DNS servers for a non-management VDOM.",

		CreateContext: resourceSystemVdomDnsCreate,
		ReadContext:   resourceSystemVdomDnsRead,
		UpdateContext: resourceSystemVdomDnsUpdate,
		DeleteContext: resourceSystemVdomDnsDelete,

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
			"dns_over_tls": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "enforce"}, false),

				Description: "Enable/disable/enforce DNS over TLS.",
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
				Description:      "Primary IPv6 DNS server IP address for the VDOM.",
				Optional:         true,
				Computed:         true,
			},
			"ip6_secondary": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Secondary IPv6 DNS server IP address for the VDOM.",
				Optional:         true,
				Computed:         true,
			},
			"primary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Primary DNS server IP address for the VDOM.",
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
			"secondary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Secondary DNS server IP address for the VDOM.",
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

				Description: "Source IP for communications with the DNS server.",
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
			"vdom_dns": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable configuring DNS servers for the current VDOM.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemVdomDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVdomDns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVdomDns(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomDns")
	}

	return resourceSystemVdomDnsRead(ctx, d, meta)
}

func resourceSystemVdomDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVdomDns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVdomDns(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVdomDns resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomDns")
	}

	return resourceSystemVdomDnsRead(ctx, d, meta)
}

func resourceSystemVdomDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemVdomDns(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemVdomDns(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVdomDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVdomDns(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdomDns resource: %v", err)
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

	diags := refreshObjectSystemVdomDns(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemVdomDnsServerHostname(d *schema.ResourceData, v *[]models.SystemVdomDnsServerHostname, prefix string, sort bool) interface{} {
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

func refreshObjectSystemVdomDns(d *schema.ResourceData, o *models.SystemVdomDns, sv string, sort bool) diag.Diagnostics {
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

	if o.DnsOverTls != nil {
		v := *o.DnsOverTls

		if err = d.Set("dns_over_tls", v); err != nil {
			return diag.Errorf("error reading dns_over_tls: %v", err)
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

	if o.Secondary != nil {
		v := *o.Secondary

		if err = d.Set("secondary", v); err != nil {
			return diag.Errorf("error reading secondary: %v", err)
		}
	}

	if o.ServerHostname != nil {
		if err = d.Set("server_hostname", flattenSystemVdomDnsServerHostname(d, o.ServerHostname, "server_hostname", sort)); err != nil {
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

	if o.VdomDns != nil {
		v := *o.VdomDns

		if err = d.Set("vdom_dns", v); err != nil {
			return diag.Errorf("error reading vdom_dns: %v", err)
		}
	}

	return nil
}

func expandSystemVdomDnsServerHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVdomDnsServerHostname, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVdomDnsServerHostname

	for i := range l {
		tmp := models.SystemVdomDnsServerHostname{}
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

func getObjectSystemVdomDns(d *schema.ResourceData, sv string) (*models.SystemVdomDns, diag.Diagnostics) {
	obj := models.SystemVdomDns{}
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
	if v1, ok := d.GetOk("dns_over_tls"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("dns_over_tls", sv)
				diags = append(diags, e)
			}
			obj.DnsOverTls = &v2
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
		t, err := expandSystemVdomDnsServerHostname(d, v, "server_hostname", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerHostname = t
		}
	} else if d.HasChange("server_hostname") {
		old, new := d.GetChange("server_hostname")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerHostname = &[]models.SystemVdomDnsServerHostname{}
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
	if v1, ok := d.GetOk("vdom_dns"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom_dns", sv)
				diags = append(diags, e)
			}
			obj.VdomDns = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemVdomDns(d *schema.ResourceData, sv string) (*models.SystemVdomDns, diag.Diagnostics) {
	obj := models.SystemVdomDns{}
	diags := diag.Diagnostics{}

	obj.ServerHostname = &[]models.SystemVdomDnsServerHostname{}

	return &obj, diags
}
