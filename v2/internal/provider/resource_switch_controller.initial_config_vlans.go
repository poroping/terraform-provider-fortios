// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure initial template for auto-generated VLAN interfaces.

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

func resourceSwitchControllerInitialConfigVlans() *schema.Resource {
	return &schema.Resource{
		Description: "Configure initial template for auto-generated VLAN interfaces.",

		CreateContext: resourceSwitchControllerInitialConfigVlansCreate,
		ReadContext:   resourceSwitchControllerInitialConfigVlansRead,
		UpdateContext: resourceSwitchControllerInitialConfigVlansUpdate,
		DeleteContext: resourceSwitchControllerInitialConfigVlansDelete,

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
			"default_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Default VLAN (native) assigned to all switch ports upon discovery.",
				Optional:    true,
				Computed:    true,
			},
			"nac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "VLAN for NAC onboarding devices.",
				Optional:    true,
				Computed:    true,
			},
			"nac_segment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "VLAN for NAC segment primary interface.",
				Optional:    true,
				Computed:    true,
			},
			"quarantine": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "VLAN for quarantined traffic.",
				Optional:    true,
				Computed:    true,
			},
			"rspan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "VLAN for RSPAN/ERSPAN mirrored traffic.",
				Optional:    true,
				Computed:    true,
			},
			"video": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "VLAN dedicated for video devices.",
				Optional:    true,
				Computed:    true,
			},
			"voice": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "VLAN dedicated for voice devices.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerInitialConfigVlansCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerInitialConfigVlans(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerInitialConfigVlans(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerInitialConfigVlans")
	}

	return resourceSwitchControllerInitialConfigVlansRead(ctx, d, meta)
}

func resourceSwitchControllerInitialConfigVlansUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerInitialConfigVlans(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerInitialConfigVlans(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerInitialConfigVlans resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerInitialConfigVlans")
	}

	return resourceSwitchControllerInitialConfigVlansRead(ctx, d, meta)
}

func resourceSwitchControllerInitialConfigVlansDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSwitchControllerInitialConfigVlans(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSwitchControllerInitialConfigVlans(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerInitialConfigVlans resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerInitialConfigVlansRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerInitialConfigVlans(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerInitialConfigVlans resource: %v", err)
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

	diags := refreshObjectSwitchControllerInitialConfigVlans(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerInitialConfigVlans(d *schema.ResourceData, o *models.SwitchControllerInitialConfigVlans, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DefaultVlan != nil {
		v := *o.DefaultVlan

		if err = d.Set("default_vlan", v); err != nil {
			return diag.Errorf("error reading default_vlan: %v", err)
		}
	}

	if o.Nac != nil {
		v := *o.Nac

		if err = d.Set("nac", v); err != nil {
			return diag.Errorf("error reading nac: %v", err)
		}
	}

	if o.NacSegment != nil {
		v := *o.NacSegment

		if err = d.Set("nac_segment", v); err != nil {
			return diag.Errorf("error reading nac_segment: %v", err)
		}
	}

	if o.Quarantine != nil {
		v := *o.Quarantine

		if err = d.Set("quarantine", v); err != nil {
			return diag.Errorf("error reading quarantine: %v", err)
		}
	}

	if o.Rspan != nil {
		v := *o.Rspan

		if err = d.Set("rspan", v); err != nil {
			return diag.Errorf("error reading rspan: %v", err)
		}
	}

	if o.Video != nil {
		v := *o.Video

		if err = d.Set("video", v); err != nil {
			return diag.Errorf("error reading video: %v", err)
		}
	}

	if o.Voice != nil {
		v := *o.Voice

		if err = d.Set("voice", v); err != nil {
			return diag.Errorf("error reading voice: %v", err)
		}
	}

	return nil
}

func getObjectSwitchControllerInitialConfigVlans(d *schema.ResourceData, sv string) (*models.SwitchControllerInitialConfigVlans, diag.Diagnostics) {
	obj := models.SwitchControllerInitialConfigVlans{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("default_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_vlan", sv)
				diags = append(diags, e)
			}
			obj.DefaultVlan = &v2
		}
	}
	if v1, ok := d.GetOk("nac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nac", sv)
				diags = append(diags, e)
			}
			obj.Nac = &v2
		}
	}
	if v1, ok := d.GetOk("nac_segment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("nac_segment", sv)
				diags = append(diags, e)
			}
			obj.NacSegment = &v2
		}
	}
	if v1, ok := d.GetOk("quarantine"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("quarantine", sv)
				diags = append(diags, e)
			}
			obj.Quarantine = &v2
		}
	}
	if v1, ok := d.GetOk("rspan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rspan", sv)
				diags = append(diags, e)
			}
			obj.Rspan = &v2
		}
	}
	if v1, ok := d.GetOk("video"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("video", sv)
				diags = append(diags, e)
			}
			obj.Video = &v2
		}
	}
	if v1, ok := d.GetOk("voice"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voice", sv)
				diags = append(diags, e)
			}
			obj.Voice = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSwitchControllerInitialConfigVlans(d *schema.ResourceData, sv string) (*models.SwitchControllerInitialConfigVlans, diag.Diagnostics) {
	obj := models.SwitchControllerInitialConfigVlans{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
