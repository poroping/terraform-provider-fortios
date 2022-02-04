// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system probe response.

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

func resourceSystemProbeResponse() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system probe response.",

		CreateContext: resourceSystemProbeResponseCreate,
		ReadContext:   resourceSystemProbeResponseRead,
		UpdateContext: resourceSystemProbeResponseUpdate,
		DeleteContext: resourceSystemProbeResponseDelete,

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
			"http_probe_value": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "Value to respond to the monitoring server.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "http-probe", "twamp"}, false),

				Description: "SLA response mode.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "TWAMP responder password in authentication mode.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port number to response.",
				Optional:    true,
				Computed:    true,
			},
			"security_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "authentication"}, false),

				Description: "TWAMP responder security mode.",
				Optional:    true,
				Computed:    true,
			},
			"timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),

				Description: "An inactivity timer for a twamp test session.",
				Optional:    true,
				Computed:    true,
			},
			"ttl_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"reinit", "decrease", "retain"}, false),

				Description: "Mode for TWAMP packet TTL modification.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemProbeResponseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemProbeResponse(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemProbeResponse(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemProbeResponse")
	}

	return resourceSystemProbeResponseRead(ctx, d, meta)
}

func resourceSystemProbeResponseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemProbeResponse(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemProbeResponse(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemProbeResponse resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemProbeResponse")
	}

	return resourceSystemProbeResponseRead(ctx, d, meta)
}

func resourceSystemProbeResponseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemProbeResponse(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemProbeResponse(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemProbeResponse resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemProbeResponseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemProbeResponse(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemProbeResponse resource: %v", err)
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

	diags := refreshObjectSystemProbeResponse(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemProbeResponse(d *schema.ResourceData, o *models.SystemProbeResponse, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.HttpProbeValue != nil {
		v := *o.HttpProbeValue

		if err = d.Set("http_probe_value", v); err != nil {
			return diag.Errorf("error reading http_probe_value: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.SecurityMode != nil {
		v := *o.SecurityMode

		if err = d.Set("security_mode", v); err != nil {
			return diag.Errorf("error reading security_mode: %v", err)
		}
	}

	if o.Timeout != nil {
		v := *o.Timeout

		if err = d.Set("timeout", v); err != nil {
			return diag.Errorf("error reading timeout: %v", err)
		}
	}

	if o.TtlMode != nil {
		v := *o.TtlMode

		if err = d.Set("ttl_mode", v); err != nil {
			return diag.Errorf("error reading ttl_mode: %v", err)
		}
	}

	return nil
}

func getObjectSystemProbeResponse(d *schema.ResourceData, sv string) (*models.SystemProbeResponse, diag.Diagnostics) {
	obj := models.SystemProbeResponse{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("http_probe_value"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_probe_value", sv)
				diags = append(diags, e)
			}
			obj.HttpProbeValue = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
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
	if v1, ok := d.GetOk("security_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_mode", sv)
				diags = append(diags, e)
			}
			obj.SecurityMode = &v2
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
	if v1, ok := d.GetOk("ttl_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ttl_mode", sv)
				diags = append(diags, e)
			}
			obj.TtlMode = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemProbeResponse(d *schema.ResourceData, sv string) (*models.SystemProbeResponse, diag.Diagnostics) {
	obj := models.SystemProbeResponse{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
