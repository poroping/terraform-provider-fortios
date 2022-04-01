// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure global Web cache settings.

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

func resourceWanoptWebcache() *schema.Resource {
	return &schema.Resource{
		Description: "Configure global Web cache settings.",

		CreateContext: resourceWanoptWebcacheCreate,
		ReadContext:   resourceWanoptWebcacheRead,
		UpdateContext: resourceWanoptWebcacheUpdate,
		DeleteContext: resourceWanoptWebcacheDelete,

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
			"always_revalidate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client.",
				Optional:    true,
				Computed:    true,
			},
			"cache_by_default": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable caching content that lacks explicit caching policies from the server.",
				Optional:    true,
				Computed:    true,
			},
			"cache_cookie": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached.",
				Optional:    true,
				Computed:    true,
			},
			"cache_expired": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable caching type-1 objects that are already expired on arrival.",
				Optional:    true,
				Computed:    true,
			},
			"default_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 5256000),

				Description: "Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.",
				Optional:    true,
				Computed:    true,
			},
			"external": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable external Web caching.",
				Optional:    true,
				Computed:    true,
			},
			"fresh_factor": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.",
				Optional:    true,
				Computed:    true,
			},
			"host_validate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable validating \"Host:\" with original server IP.",
				Optional:    true,
				Computed:    true,
			},
			"ignore_conditional": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable controlling the behavior of cache-control HTTP 1.1 header values.",
				Optional:    true,
				Computed:    true,
			},
			"ignore_ie_reload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header.",
				Optional:    true,
				Computed:    true,
			},
			"ignore_ims": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ignoring the if-modified-since (IMS) header.",
				Optional:    true,
				Computed:    true,
			},
			"ignore_pnc": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ignoring the pragma no-cache (PNC) header.",
				Optional:    true,
				Computed:    true,
			},
			"max_object_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483),

				Description: "Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.",
				Optional:    true,
				Computed:    true,
			},
			"max_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 5256000),

				Description: "Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).",
				Optional:    true,
				Computed:    true,
			},
			"min_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 5256000),

				Description: "Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).",
				Optional:    true,
				Computed:    true,
			},
			"neg_resp_time": {
				Type: schema.TypeInt,

				Description: "Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).",
				Optional:    true,
				Computed:    true,
			},
			"reval_pnc": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWanoptWebcacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptWebcache(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWanoptWebcache(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptWebcache")
	}

	return resourceWanoptWebcacheRead(ctx, d, meta)
}

func resourceWanoptWebcacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptWebcache(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWanoptWebcache(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WanoptWebcache resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptWebcache")
	}

	return resourceWanoptWebcacheRead(ctx, d, meta)
}

func resourceWanoptWebcacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWanoptWebcache(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWanoptWebcache(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WanoptWebcache resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptWebcacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptWebcache(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptWebcache resource: %v", err)
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

	diags := refreshObjectWanoptWebcache(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWanoptWebcache(d *schema.ResourceData, o *models.WanoptWebcache, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AlwaysRevalidate != nil {
		v := *o.AlwaysRevalidate

		if err = d.Set("always_revalidate", v); err != nil {
			return diag.Errorf("error reading always_revalidate: %v", err)
		}
	}

	if o.CacheByDefault != nil {
		v := *o.CacheByDefault

		if err = d.Set("cache_by_default", v); err != nil {
			return diag.Errorf("error reading cache_by_default: %v", err)
		}
	}

	if o.CacheCookie != nil {
		v := *o.CacheCookie

		if err = d.Set("cache_cookie", v); err != nil {
			return diag.Errorf("error reading cache_cookie: %v", err)
		}
	}

	if o.CacheExpired != nil {
		v := *o.CacheExpired

		if err = d.Set("cache_expired", v); err != nil {
			return diag.Errorf("error reading cache_expired: %v", err)
		}
	}

	if o.DefaultTtl != nil {
		v := *o.DefaultTtl

		if err = d.Set("default_ttl", v); err != nil {
			return diag.Errorf("error reading default_ttl: %v", err)
		}
	}

	if o.External != nil {
		v := *o.External

		if err = d.Set("external", v); err != nil {
			return diag.Errorf("error reading external: %v", err)
		}
	}

	if o.FreshFactor != nil {
		v := *o.FreshFactor

		if err = d.Set("fresh_factor", v); err != nil {
			return diag.Errorf("error reading fresh_factor: %v", err)
		}
	}

	if o.HostValidate != nil {
		v := *o.HostValidate

		if err = d.Set("host_validate", v); err != nil {
			return diag.Errorf("error reading host_validate: %v", err)
		}
	}

	if o.IgnoreConditional != nil {
		v := *o.IgnoreConditional

		if err = d.Set("ignore_conditional", v); err != nil {
			return diag.Errorf("error reading ignore_conditional: %v", err)
		}
	}

	if o.IgnoreIeReload != nil {
		v := *o.IgnoreIeReload

		if err = d.Set("ignore_ie_reload", v); err != nil {
			return diag.Errorf("error reading ignore_ie_reload: %v", err)
		}
	}

	if o.IgnoreIms != nil {
		v := *o.IgnoreIms

		if err = d.Set("ignore_ims", v); err != nil {
			return diag.Errorf("error reading ignore_ims: %v", err)
		}
	}

	if o.IgnorePnc != nil {
		v := *o.IgnorePnc

		if err = d.Set("ignore_pnc", v); err != nil {
			return diag.Errorf("error reading ignore_pnc: %v", err)
		}
	}

	if o.MaxObjectSize != nil {
		v := *o.MaxObjectSize

		if err = d.Set("max_object_size", v); err != nil {
			return diag.Errorf("error reading max_object_size: %v", err)
		}
	}

	if o.MaxTtl != nil {
		v := *o.MaxTtl

		if err = d.Set("max_ttl", v); err != nil {
			return diag.Errorf("error reading max_ttl: %v", err)
		}
	}

	if o.MinTtl != nil {
		v := *o.MinTtl

		if err = d.Set("min_ttl", v); err != nil {
			return diag.Errorf("error reading min_ttl: %v", err)
		}
	}

	if o.NegRespTime != nil {
		v := *o.NegRespTime

		if err = d.Set("neg_resp_time", v); err != nil {
			return diag.Errorf("error reading neg_resp_time: %v", err)
		}
	}

	if o.RevalPnc != nil {
		v := *o.RevalPnc

		if err = d.Set("reval_pnc", v); err != nil {
			return diag.Errorf("error reading reval_pnc: %v", err)
		}
	}

	return nil
}

func getObjectWanoptWebcache(d *schema.ResourceData, sv string) (*models.WanoptWebcache, diag.Diagnostics) {
	obj := models.WanoptWebcache{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("always_revalidate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("always_revalidate", sv)
				diags = append(diags, e)
			}
			obj.AlwaysRevalidate = &v2
		}
	}
	if v1, ok := d.GetOk("cache_by_default"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_by_default", sv)
				diags = append(diags, e)
			}
			obj.CacheByDefault = &v2
		}
	}
	if v1, ok := d.GetOk("cache_cookie"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_cookie", sv)
				diags = append(diags, e)
			}
			obj.CacheCookie = &v2
		}
	}
	if v1, ok := d.GetOk("cache_expired"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_expired", sv)
				diags = append(diags, e)
			}
			obj.CacheExpired = &v2
		}
	}
	if v1, ok := d.GetOk("default_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("external"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external", sv)
				diags = append(diags, e)
			}
			obj.External = &v2
		}
	}
	if v1, ok := d.GetOk("fresh_factor"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fresh_factor", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FreshFactor = &tmp
		}
	}
	if v1, ok := d.GetOk("host_validate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host_validate", sv)
				diags = append(diags, e)
			}
			obj.HostValidate = &v2
		}
	}
	if v1, ok := d.GetOk("ignore_conditional"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ignore_conditional", sv)
				diags = append(diags, e)
			}
			obj.IgnoreConditional = &v2
		}
	}
	if v1, ok := d.GetOk("ignore_ie_reload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ignore_ie_reload", sv)
				diags = append(diags, e)
			}
			obj.IgnoreIeReload = &v2
		}
	}
	if v1, ok := d.GetOk("ignore_ims"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ignore_ims", sv)
				diags = append(diags, e)
			}
			obj.IgnoreIms = &v2
		}
	}
	if v1, ok := d.GetOk("ignore_pnc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ignore_pnc", sv)
				diags = append(diags, e)
			}
			obj.IgnorePnc = &v2
		}
	}
	if v1, ok := d.GetOk("max_object_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_object_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxObjectSize = &tmp
		}
	}
	if v1, ok := d.GetOk("max_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("min_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("min_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("neg_resp_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("neg_resp_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NegRespTime = &tmp
		}
	}
	if v1, ok := d.GetOk("reval_pnc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reval_pnc", sv)
				diags = append(diags, e)
			}
			obj.RevalPnc = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWanoptWebcache(d *schema.ResourceData, sv string) (*models.WanoptWebcache, diag.Diagnostics) {
	obj := models.WanoptWebcache{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
