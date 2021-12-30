// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Recurring schedule configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallscheduleRecurring() *schema.Resource {
	return &schema.Resource{
		Description: "Recurring schedule configuration.",

		CreateContext: resourceFirewallscheduleRecurringCreate,
		ReadContext:   resourceFirewallscheduleRecurringRead,
		UpdateContext: resourceFirewallscheduleRecurringUpdate,
		DeleteContext: resourceFirewallscheduleRecurringDelete,

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
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"day": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "One or more days of the week on which the schedule is valid. Separate the names of the days with a space.",
				Optional:         true,
				Computed:         true,
			},
			"end": {
				Type: schema.TypeString,

				Description: "Time of day to end the schedule, format hh:mm.",
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

				Description: "Recurring schedule name.",
				ForceNew:    true,
				Required:    true,
			},
			"start": {
				Type: schema.TypeString,

				Description: "Time of day to start the schedule, format hh:mm.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallscheduleRecurringCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallscheduleRecurring resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallscheduleRecurring(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallscheduleRecurring(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallscheduleRecurring")
	}

	return resourceFirewallscheduleRecurringRead(ctx, d, meta)
}

func resourceFirewallscheduleRecurringUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallscheduleRecurring(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallscheduleRecurring(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallscheduleRecurring resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallscheduleRecurring")
	}

	return resourceFirewallscheduleRecurringRead(ctx, d, meta)
}

func resourceFirewallscheduleRecurringDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallscheduleRecurring(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallscheduleRecurring resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallscheduleRecurringRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallscheduleRecurring(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallscheduleRecurring resource: %v", err)
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

	diags := refreshObjectFirewallscheduleRecurring(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallscheduleRecurring(d *schema.ResourceData, o *models.FirewallscheduleRecurring, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.Day != nil {
		v := *o.Day

		if err = d.Set("day", v); err != nil {
			return diag.Errorf("error reading day: %v", err)
		}
	}

	if o.End != nil {
		v := *o.End

		if err = d.Set("end", v); err != nil {
			return diag.Errorf("error reading end: %v", err)
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

	return nil
}

func getObjectFirewallscheduleRecurring(d *schema.ResourceData, sv string) (*models.FirewallscheduleRecurring, diag.Diagnostics) {
	obj := models.FirewallscheduleRecurring{}
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
	if v1, ok := d.GetOk("day"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("day", sv)
				diags = append(diags, e)
			}
			obj.Day = &v2
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
	return &obj, diags
}
