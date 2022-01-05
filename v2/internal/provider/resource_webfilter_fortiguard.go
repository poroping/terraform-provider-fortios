// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGuard Web Filter service.

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

func resourceWebfilterFortiguard() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGuard Web Filter service.",

		CreateContext: resourceWebfilterFortiguardCreate,
		ReadContext:   resourceWebfilterFortiguardRead,
		UpdateContext: resourceWebfilterFortiguardUpdate,
		DeleteContext: resourceWebfilterFortiguardDelete,

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
			"cache_mem_percent": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),

				Description: "Maximum percentage of available memory allocated to caching (1 - 15%).",
				Optional:    true,
				Computed:    true,
			},
			"cache_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ttl", "db-ver"}, false),

				Description: "Cache entry expiration mode.",
				Optional:    true,
				Computed:    true,
			},
			"cache_prefix_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable prefix matching in the cache.",
				Optional:    true,
				Computed:    true,
			},
			"close_ports": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Close ports used for HTTP/HTTPS override authentication and disable user overrides.",
				Optional:    true,
				Computed:    true,
			},
			"ovrd_auth_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of HTTPS for override authentication.",
				Optional:    true,
				Computed:    true,
			},
			"ovrd_auth_port_http": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port to use for FortiGuard Web Filter HTTP override authentication",
				Optional:    true,
				Computed:    true,
			},
			"ovrd_auth_port_https": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.",
				Optional:    true,
				Computed:    true,
			},
			"ovrd_auth_port_https_flow": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.",
				Optional:    true,
				Computed:    true,
			},
			"ovrd_auth_port_warning": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port to use for FortiGuard Web Filter Warning override authentication.",
				Optional:    true,
				Computed:    true,
			},
			"request_packet_size_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 10000),

				Description: "Limit size of URL request packets sent to FortiGuard server (0 for default).",
				Optional:    true,
				Computed:    true,
			},
			"warn_auth_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of HTTPS for warning and authentication.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebfilterFortiguardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterFortiguard(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebfilterFortiguard(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterFortiguard")
	}

	return resourceWebfilterFortiguardRead(ctx, d, meta)
}

func resourceWebfilterFortiguardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterFortiguard(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebfilterFortiguard(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebfilterFortiguard resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterFortiguard")
	}

	return resourceWebfilterFortiguardRead(ctx, d, meta)
}

func resourceWebfilterFortiguardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWebfilterFortiguard(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWebfilterFortiguard(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebfilterFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFortiguardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterFortiguard(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterFortiguard resource: %v", err)
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

	diags := refreshObjectWebfilterFortiguard(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWebfilterFortiguard(d *schema.ResourceData, o *models.WebfilterFortiguard, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CacheMemPercent != nil {
		v := *o.CacheMemPercent

		if err = d.Set("cache_mem_percent", v); err != nil {
			return diag.Errorf("error reading cache_mem_percent: %v", err)
		}
	}

	if o.CacheMode != nil {
		v := *o.CacheMode

		if err = d.Set("cache_mode", v); err != nil {
			return diag.Errorf("error reading cache_mode: %v", err)
		}
	}

	if o.CachePrefixMatch != nil {
		v := *o.CachePrefixMatch

		if err = d.Set("cache_prefix_match", v); err != nil {
			return diag.Errorf("error reading cache_prefix_match: %v", err)
		}
	}

	if o.ClosePorts != nil {
		v := *o.ClosePorts

		if err = d.Set("close_ports", v); err != nil {
			return diag.Errorf("error reading close_ports: %v", err)
		}
	}

	if o.OvrdAuthHttps != nil {
		v := *o.OvrdAuthHttps

		if err = d.Set("ovrd_auth_https", v); err != nil {
			return diag.Errorf("error reading ovrd_auth_https: %v", err)
		}
	}

	if o.OvrdAuthPortHttp != nil {
		v := *o.OvrdAuthPortHttp

		if err = d.Set("ovrd_auth_port_http", v); err != nil {
			return diag.Errorf("error reading ovrd_auth_port_http: %v", err)
		}
	}

	if o.OvrdAuthPortHttps != nil {
		v := *o.OvrdAuthPortHttps

		if err = d.Set("ovrd_auth_port_https", v); err != nil {
			return diag.Errorf("error reading ovrd_auth_port_https: %v", err)
		}
	}

	if o.OvrdAuthPortHttpsFlow != nil {
		v := *o.OvrdAuthPortHttpsFlow

		if err = d.Set("ovrd_auth_port_https_flow", v); err != nil {
			return diag.Errorf("error reading ovrd_auth_port_https_flow: %v", err)
		}
	}

	if o.OvrdAuthPortWarning != nil {
		v := *o.OvrdAuthPortWarning

		if err = d.Set("ovrd_auth_port_warning", v); err != nil {
			return diag.Errorf("error reading ovrd_auth_port_warning: %v", err)
		}
	}

	if o.RequestPacketSizeLimit != nil {
		v := *o.RequestPacketSizeLimit

		if err = d.Set("request_packet_size_limit", v); err != nil {
			return diag.Errorf("error reading request_packet_size_limit: %v", err)
		}
	}

	if o.WarnAuthHttps != nil {
		v := *o.WarnAuthHttps

		if err = d.Set("warn_auth_https", v); err != nil {
			return diag.Errorf("error reading warn_auth_https: %v", err)
		}
	}

	return nil
}

func getObjectWebfilterFortiguard(d *schema.ResourceData, sv string) (*models.WebfilterFortiguard, diag.Diagnostics) {
	obj := models.WebfilterFortiguard{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("cache_mem_percent"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_mem_percent", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CacheMemPercent = &tmp
		}
	}
	if v1, ok := d.GetOk("cache_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_mode", sv)
				diags = append(diags, e)
			}
			obj.CacheMode = &v2
		}
	}
	if v1, ok := d.GetOk("cache_prefix_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_prefix_match", sv)
				diags = append(diags, e)
			}
			obj.CachePrefixMatch = &v2
		}
	}
	if v1, ok := d.GetOk("close_ports"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("close_ports", sv)
				diags = append(diags, e)
			}
			obj.ClosePorts = &v2
		}
	}
	if v1, ok := d.GetOk("ovrd_auth_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ovrd_auth_https", sv)
				diags = append(diags, e)
			}
			obj.OvrdAuthHttps = &v2
		}
	}
	if v1, ok := d.GetOk("ovrd_auth_port_http"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ovrd_auth_port_http", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OvrdAuthPortHttp = &tmp
		}
	}
	if v1, ok := d.GetOk("ovrd_auth_port_https"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ovrd_auth_port_https", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OvrdAuthPortHttps = &tmp
		}
	}
	if v1, ok := d.GetOk("ovrd_auth_port_https_flow"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ovrd_auth_port_https_flow", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OvrdAuthPortHttpsFlow = &tmp
		}
	}
	if v1, ok := d.GetOk("ovrd_auth_port_warning"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ovrd_auth_port_warning", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OvrdAuthPortWarning = &tmp
		}
	}
	if v1, ok := d.GetOk("request_packet_size_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("request_packet_size_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RequestPacketSizeLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("warn_auth_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("warn_auth_https", sv)
				diags = append(diags, e)
			}
			obj.WarnAuthHttps = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWebfilterFortiguard(d *schema.ResourceData, sv string) (*models.WebfilterFortiguard, diag.Diagnostics) {
	obj := models.WebfilterFortiguard{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
