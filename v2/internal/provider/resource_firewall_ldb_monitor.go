// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure server load balancing health monitors.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallLdbMonitor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure server load balancing health monitors.",

		CreateContext: resourceFirewallLdbMonitorCreate,
		ReadContext:   resourceFirewallLdbMonitorRead,
		UpdateContext: resourceFirewallLdbMonitorUpdate,
		DeleteContext: resourceFirewallLdbMonitorDelete,

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
			"dns_match_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Response IP expected from DNS server.",
				Optional:    true,
				Computed:    true,
			},
			"dns_protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"udp", "tcp"}, false),

				Description: "Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP).",
				Optional:    true,
				Computed:    true,
			},
			"dns_request_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fully qualified domain name to resolve for the DNS probe.",
				Optional:    true,
				Computed:    true,
			},
			"http_get": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "URL used to send a GET request to check the health of an HTTP server.",
				Optional:    true,
				Computed:    true,
			},
			"http_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "String to match the value expected in response to an HTTP-GET request.",
				Optional:    true,
				Computed:    true,
			},
			"http_max_redirects": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 5),

				Description: "The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 65535),

				Description: "Time between health checks (5 - 65535 sec, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Monitor name.",
				ForceNew:    true,
				Required:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65535, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"retry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Number health check attempts before the server is considered down (1 - 255, default = 3).",
				Optional:    true,
				Computed:    true,
			},
			"src_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP for ldb-monitor.",
				Optional:    true,
				Computed:    true,
			},
			"timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ping", "tcp", "http", "https", "dns"}, false),

				Description: "Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP | HTTPS | DNS).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallLdbMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallLdbMonitor resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallLdbMonitor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallLdbMonitor(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallLdbMonitor")
	}

	return resourceFirewallLdbMonitorRead(ctx, d, meta)
}

func resourceFirewallLdbMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallLdbMonitor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallLdbMonitor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallLdbMonitor resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallLdbMonitor")
	}

	return resourceFirewallLdbMonitorRead(ctx, d, meta)
}

func resourceFirewallLdbMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallLdbMonitor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallLdbMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallLdbMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallLdbMonitor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallLdbMonitor resource: %v", err)
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

	diags := refreshObjectFirewallLdbMonitor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallLdbMonitor(d *schema.ResourceData, o *models.FirewallLdbMonitor, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DnsMatchIp != nil {
		v := *o.DnsMatchIp

		if err = d.Set("dns_match_ip", v); err != nil {
			return diag.Errorf("error reading dns_match_ip: %v", err)
		}
	}

	if o.DnsProtocol != nil {
		v := *o.DnsProtocol

		if err = d.Set("dns_protocol", v); err != nil {
			return diag.Errorf("error reading dns_protocol: %v", err)
		}
	}

	if o.DnsRequestDomain != nil {
		v := *o.DnsRequestDomain

		if err = d.Set("dns_request_domain", v); err != nil {
			return diag.Errorf("error reading dns_request_domain: %v", err)
		}
	}

	if o.HttpGet != nil {
		v := *o.HttpGet

		if err = d.Set("http_get", v); err != nil {
			return diag.Errorf("error reading http_get: %v", err)
		}
	}

	if o.HttpMatch != nil {
		v := *o.HttpMatch

		if err = d.Set("http_match", v); err != nil {
			return diag.Errorf("error reading http_match: %v", err)
		}
	}

	if o.HttpMaxRedirects != nil {
		v := *o.HttpMaxRedirects

		if err = d.Set("http_max_redirects", v); err != nil {
			return diag.Errorf("error reading http_max_redirects: %v", err)
		}
	}

	if o.Interval != nil {
		v := *o.Interval

		if err = d.Set("interval", v); err != nil {
			return diag.Errorf("error reading interval: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Retry != nil {
		v := *o.Retry

		if err = d.Set("retry", v); err != nil {
			return diag.Errorf("error reading retry: %v", err)
		}
	}

	if o.SrcIp != nil {
		v := *o.SrcIp

		if err = d.Set("src_ip", v); err != nil {
			return diag.Errorf("error reading src_ip: %v", err)
		}
	}

	if o.Timeout != nil {
		v := *o.Timeout

		if err = d.Set("timeout", v); err != nil {
			return diag.Errorf("error reading timeout: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	return nil
}

func getObjectFirewallLdbMonitor(d *schema.ResourceData, sv string) (*models.FirewallLdbMonitor, diag.Diagnostics) {
	obj := models.FirewallLdbMonitor{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("dns_match_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dns_match_ip", sv)
				diags = append(diags, e)
			}
			obj.DnsMatchIp = &v2
		}
	}
	if v1, ok := d.GetOk("dns_protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dns_protocol", sv)
				diags = append(diags, e)
			}
			obj.DnsProtocol = &v2
		}
	}
	if v1, ok := d.GetOk("dns_request_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dns_request_domain", sv)
				diags = append(diags, e)
			}
			obj.DnsRequestDomain = &v2
		}
	}
	if v1, ok := d.GetOk("http_get"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_get", sv)
				diags = append(diags, e)
			}
			obj.HttpGet = &v2
		}
	}
	if v1, ok := d.GetOk("http_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_match", sv)
				diags = append(diags, e)
			}
			obj.HttpMatch = &v2
		}
	}
	if v1, ok := d.GetOk("http_max_redirects"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_max_redirects", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HttpMaxRedirects = &tmp
		}
	}
	if v1, ok := d.GetOk("interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Interval = &tmp
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
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
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
	if v1, ok := d.GetOk("src_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("src_ip", sv)
				diags = append(diags, e)
			}
			obj.SrcIp = &v2
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
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	return &obj, diags
}
