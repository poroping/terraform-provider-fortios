// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Onetime schedule configuration.

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

func resourceFirewallScheduleOnetime() *schema.Resource {
	return &schema.Resource{
		Description: "Onetime schedule configuration.",

		CreateContext: resourceFirewallScheduleOnetimeCreate,
		ReadContext:   resourceFirewallScheduleOnetimeRead,
		UpdateContext: resourceFirewallScheduleOnetimeUpdate,
		DeleteContext: resourceFirewallScheduleOnetimeDelete,

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
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"end": {
				Type: schema.TypeString,

				Description: "Schedule end date and time, format hh:mm yyyy/mm/dd.",
				Optional:    true,
				Computed:    true,
			},
			"end_utc": {
				Type: schema.TypeString,

				Description: "Schedule end date and time, in epoch format.",
				Optional:    true,
				Computed:    true,
			},
			"expiration_days": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Write an event log message this many days before the schedule expires.",
				Optional:    true,
				Computed:    true,
			},
			"fabric_object": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Security Fabric global object setting.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Onetime schedule name.",
				ForceNew:    true,
				Required:    true,
			},
			"start": {
				Type: schema.TypeString,

				Description: "Schedule start date and time, format hh:mm yyyy/mm/dd.",
				Optional:    true,
				Computed:    true,
			},
			"start_utc": {
				Type: schema.TypeString,

				Description: "Schedule start date and time, in epoch format.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallScheduleOnetimeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallScheduleOnetime resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallScheduleOnetime(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallScheduleOnetime(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallScheduleOnetime")
	}

	return resourceFirewallScheduleOnetimeRead(ctx, d, meta)
}

func resourceFirewallScheduleOnetimeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallScheduleOnetime(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallScheduleOnetime(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallScheduleOnetime resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallScheduleOnetime")
	}

	return resourceFirewallScheduleOnetimeRead(ctx, d, meta)
}

func resourceFirewallScheduleOnetimeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallScheduleOnetime(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallScheduleOnetime resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallScheduleOnetimeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallScheduleOnetime(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallScheduleOnetime resource: %v", err)
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

	diags := refreshObjectFirewallScheduleOnetime(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallScheduleOnetime(d *schema.ResourceData, o *models.FirewallScheduleOnetime, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.End != nil {
		v := *o.End

		if err = d.Set("end", v); err != nil {
			return diag.Errorf("error reading end: %v", err)
		}
	}

	if o.EndUtc != nil {
		v := *o.EndUtc

		if err = d.Set("end_utc", v); err != nil {
			return diag.Errorf("error reading end_utc: %v", err)
		}
	}

	if o.ExpirationDays != nil {
		v := *o.ExpirationDays

		if err = d.Set("expiration_days", v); err != nil {
			return diag.Errorf("error reading expiration_days: %v", err)
		}
	}

	if o.FabricObject != nil {
		v := *o.FabricObject

		if err = d.Set("fabric_object", v); err != nil {
			return diag.Errorf("error reading fabric_object: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Start != nil {
		v := *o.Start

		if err = d.Set("start", v); err != nil {
			return diag.Errorf("error reading start: %v", err)
		}
	}

	if o.StartUtc != nil {
		v := *o.StartUtc

		if err = d.Set("start_utc", v); err != nil {
			return diag.Errorf("error reading start_utc: %v", err)
		}
	}

	return nil
}

func getObjectFirewallScheduleOnetime(d *schema.ResourceData, sv string) (*models.FirewallScheduleOnetime, diag.Diagnostics) {
	obj := models.FirewallScheduleOnetime{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("color"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Color = &tmp
		}
	}
	if v1, ok := d.GetOk("end"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("end", sv)
				diags = append(diags, e)
			}
			obj.End = &v2
		}
	}
	if v1, ok := d.GetOk("end_utc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("end_utc", sv)
				diags = append(diags, e)
			}
			obj.EndUtc = &v2
		}
	}
	if v1, ok := d.GetOk("expiration_days"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expiration_days", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ExpirationDays = &tmp
		}
	}
	if v1, ok := d.GetOk("fabric_object"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("fabric_object", sv)
				diags = append(diags, e)
			}
			obj.FabricObject = &v2
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
	if v1, ok := d.GetOk("start"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("start", sv)
				diags = append(diags, e)
			}
			obj.Start = &v2
		}
	}
	if v1, ok := d.GetOk("start_utc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("start_utc", sv)
				diags = append(diags, e)
			}
			obj.StartUtc = &v2
		}
	}
	return &obj, diags
}
