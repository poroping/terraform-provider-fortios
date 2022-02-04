// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure shaping profiles.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallShapingProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure shaping profiles.",

		CreateContext: resourceFirewallShapingProfileCreate,
		ReadContext:   resourceFirewallShapingProfileRead,
		UpdateContext: resourceFirewallShapingProfileUpdate,
		DeleteContext: resourceFirewallShapingProfileDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"default_class_id": {
				Type: schema.TypeInt,

				Description: "Default class ID to handle unclassified packets (including all local traffic).",
				Optional:    true,
				Computed:    true,
			},
			"profile_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Shaping profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"shaping_entries": {
				Type:        schema.TypeList,
				Description: "Define shaping entries of this shaping profile.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"burst_in_msec": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2000),

							Description: "Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.",
							Optional:    true,
							Computed:    true,
						},
						"cburst_in_msec": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2000),

							Description: "Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.",
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Type: schema.TypeInt,

							Description: "Class ID.",
							Optional:    true,
							Computed:    true,
						},
						"guaranteed_bandwidth_percentage": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Guaranteed bandwidth in percentage.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID number.",
							Optional:    true,
							Computed:    true,
						},
						"limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 10000),

							Description: "Hard limit on the real queue size in packets.",
							Optional:    true,
							Computed:    true,
						},
						"max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 3000),

							Description: "Average queue size in packets at which RED drop probability is maximal.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_bandwidth_percentage": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Maximum bandwidth in percentage.",
							Optional:    true,
							Computed:    true,
						},
						"min": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 3000),

							Description: "Average queue size in packets at which RED drop becomes a possibility.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"top", "critical", "high", "medium", "low"}, false),

							Description: "Priority.",
							Optional:    true,
							Computed:    true,
						},
						"red_probability": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 20),

							Description: "Maximum probability (in percentage) for RED marking.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"policing", "queuing"}, false),

				Description: "Select shaping profile type: policing / queuing.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallShapingProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "profile_name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallShapingProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallShapingProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallShapingProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShapingProfile")
	}

	return resourceFirewallShapingProfileRead(ctx, d, meta)
}

func resourceFirewallShapingProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallShapingProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallShapingProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallShapingProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShapingProfile")
	}

	return resourceFirewallShapingProfileRead(ctx, d, meta)
}

func resourceFirewallShapingProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallShapingProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallShapingProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShapingProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallShapingProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShapingProfile resource: %v", err)
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

	diags := refreshObjectFirewallShapingProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallShapingProfileShapingEntries(v *[]models.FirewallShapingProfileShapingEntries, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.BurstInMsec; tmp != nil {
				v["burst_in_msec"] = *tmp
			}

			if tmp := cfg.CburstInMsec; tmp != nil {
				v["cburst_in_msec"] = *tmp
			}

			if tmp := cfg.ClassId; tmp != nil {
				v["class_id"] = *tmp
			}

			if tmp := cfg.GuaranteedBandwidthPercentage; tmp != nil {
				v["guaranteed_bandwidth_percentage"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Limit; tmp != nil {
				v["limit"] = *tmp
			}

			if tmp := cfg.Max; tmp != nil {
				v["max"] = *tmp
			}

			if tmp := cfg.MaximumBandwidthPercentage; tmp != nil {
				v["maximum_bandwidth_percentage"] = *tmp
			}

			if tmp := cfg.Min; tmp != nil {
				v["min"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.RedProbability; tmp != nil {
				v["red_probability"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectFirewallShapingProfile(d *schema.ResourceData, o *models.FirewallShapingProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.DefaultClassId != nil {
		v := *o.DefaultClassId

		if err = d.Set("default_class_id", v); err != nil {
			return diag.Errorf("error reading default_class_id: %v", err)
		}
	}

	if o.ProfileName != nil {
		v := *o.ProfileName

		if err = d.Set("profile_name", v); err != nil {
			return diag.Errorf("error reading profile_name: %v", err)
		}
	}

	if o.ShapingEntries != nil {
		if err = d.Set("shaping_entries", flattenFirewallShapingProfileShapingEntries(o.ShapingEntries, sort)); err != nil {
			return diag.Errorf("error reading shaping_entries: %v", err)
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

func expandFirewallShapingProfileShapingEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingProfileShapingEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingProfileShapingEntries

	for i := range l {
		tmp := models.FirewallShapingProfileShapingEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.burst_in_msec", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.BurstInMsec = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cburst_in_msec", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.CburstInMsec = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.class_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ClassId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.guaranteed_bandwidth_percentage", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.GuaranteedBandwidthPercentage = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Limit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Max = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_bandwidth_percentage", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaximumBandwidthPercentage = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.min", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Min = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.red_probability", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RedProbability = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallShapingProfile(d *schema.ResourceData, sv string) (*models.FirewallShapingProfile, diag.Diagnostics) {
	obj := models.FirewallShapingProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("default_class_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_class_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultClassId = &tmp
		}
	}
	if v1, ok := d.GetOk("profile_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_name", sv)
				diags = append(diags, e)
			}
			obj.ProfileName = &v2
		}
	}
	if v, ok := d.GetOk("shaping_entries"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("shaping_entries", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingProfileShapingEntries(d, v, "shaping_entries", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ShapingEntries = t
		}
	} else if d.HasChange("shaping_entries") {
		old, new := d.GetChange("shaping_entries")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ShapingEntries = &[]models.FirewallShapingProfileShapingEntries{}
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
