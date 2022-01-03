// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN metrics.

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

func resourceWirelessControllerHotspot20H2qpWanMetric() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN metrics.",

		CreateContext: resourceWirelessControllerHotspot20H2qpWanMetricCreate,
		ReadContext:   resourceWirelessControllerHotspot20H2qpWanMetricRead,
		UpdateContext: resourceWirelessControllerHotspot20H2qpWanMetricUpdate,
		DeleteContext: resourceWirelessControllerHotspot20H2qpWanMetricDelete,

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
			"downlink_load": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Downlink load.",
				Optional:    true,
				Computed:    true,
			},
			"downlink_speed": {
				Type: schema.TypeInt,

				Description: "Downlink speed (in kilobits/s).",
				Optional:    true,
				Computed:    true,
			},
			"link_at_capacity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Link at capacity.",
				Optional:    true,
				Computed:    true,
			},
			"link_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"up", "down", "in-test"}, false),

				Description: "Link status.",
				Optional:    true,
				Computed:    true,
			},
			"load_measurement_duration": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Load measurement duration (in tenths of a second).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WAN metric name.",
				ForceNew:    true,
				Required:    true,
			},
			"symmetric_wan_link": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"symmetric", "asymmetric"}, false),

				Description: "WAN link symmetry.",
				Optional:    true,
				Computed:    true,
			},
			"uplink_load": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Uplink load.",
				Optional:    true,
				Computed:    true,
			},
			"uplink_speed": {
				Type: schema.TypeInt,

				Description: "Uplink speed (in kilobits/s).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2qpWanMetricCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20H2qpWanMetric resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20H2qpWanMetric(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20H2qpWanMetric(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpWanMetric")
	}

	return resourceWirelessControllerHotspot20H2qpWanMetricRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpWanMetricUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20H2qpWanMetric(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20H2qpWanMetric(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20H2qpWanMetric resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpWanMetric")
	}

	return resourceWirelessControllerHotspot20H2qpWanMetricRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpWanMetricDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20H2qpWanMetric(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20H2qpWanMetric resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2qpWanMetricRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpWanMetric(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpWanMetric resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpWanMetric(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerHotspot20H2qpWanMetric(d *schema.ResourceData, o *models.WirelessControllerHotspot20H2qpWanMetric, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DownlinkLoad != nil {
		v := *o.DownlinkLoad

		if err = d.Set("downlink_load", v); err != nil {
			return diag.Errorf("error reading downlink_load: %v", err)
		}
	}

	if o.DownlinkSpeed != nil {
		v := *o.DownlinkSpeed

		if err = d.Set("downlink_speed", v); err != nil {
			return diag.Errorf("error reading downlink_speed: %v", err)
		}
	}

	if o.LinkAtCapacity != nil {
		v := *o.LinkAtCapacity

		if err = d.Set("link_at_capacity", v); err != nil {
			return diag.Errorf("error reading link_at_capacity: %v", err)
		}
	}

	if o.LinkStatus != nil {
		v := *o.LinkStatus

		if err = d.Set("link_status", v); err != nil {
			return diag.Errorf("error reading link_status: %v", err)
		}
	}

	if o.LoadMeasurementDuration != nil {
		v := *o.LoadMeasurementDuration

		if err = d.Set("load_measurement_duration", v); err != nil {
			return diag.Errorf("error reading load_measurement_duration: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.SymmetricWanLink != nil {
		v := *o.SymmetricWanLink

		if err = d.Set("symmetric_wan_link", v); err != nil {
			return diag.Errorf("error reading symmetric_wan_link: %v", err)
		}
	}

	if o.UplinkLoad != nil {
		v := *o.UplinkLoad

		if err = d.Set("uplink_load", v); err != nil {
			return diag.Errorf("error reading uplink_load: %v", err)
		}
	}

	if o.UplinkSpeed != nil {
		v := *o.UplinkSpeed

		if err = d.Set("uplink_speed", v); err != nil {
			return diag.Errorf("error reading uplink_speed: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerHotspot20H2qpWanMetric(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20H2qpWanMetric, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20H2qpWanMetric{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("downlink_load"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("downlink_load", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DownlinkLoad = &tmp
		}
	}
	if v1, ok := d.GetOk("downlink_speed"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("downlink_speed", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DownlinkSpeed = &tmp
		}
	}
	if v1, ok := d.GetOk("link_at_capacity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_at_capacity", sv)
				diags = append(diags, e)
			}
			obj.LinkAtCapacity = &v2
		}
	}
	if v1, ok := d.GetOk("link_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_status", sv)
				diags = append(diags, e)
			}
			obj.LinkStatus = &v2
		}
	}
	if v1, ok := d.GetOk("load_measurement_duration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("load_measurement_duration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LoadMeasurementDuration = &tmp
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
	if v1, ok := d.GetOk("symmetric_wan_link"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("symmetric_wan_link", sv)
				diags = append(diags, e)
			}
			obj.SymmetricWanLink = &v2
		}
	}
	if v1, ok := d.GetOk("uplink_load"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uplink_load", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UplinkLoad = &tmp
		}
	}
	if v1, ok := d.GetOk("uplink_speed"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uplink_speed", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UplinkSpeed = &tmp
		}
	}
	return &obj, diags
}
