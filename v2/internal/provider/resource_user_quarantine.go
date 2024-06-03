// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure quarantine support.

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

func resourceUserQuarantine() *schema.Resource {
	return &schema.Resource{
		Description: "Configure quarantine support.",

		CreateContext: resourceUserQuarantineCreate,
		ReadContext:   resourceUserQuarantineRead,
		UpdateContext: resourceUserQuarantineUpdate,
		DeleteContext: resourceUserQuarantineDelete,

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
			"firewall_groups": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Firewall address group which includes all quarantine MAC address.",
				Optional:    true,
				Computed:    true,
			},
			"quarantine": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable quarantine.",
				Optional:    true,
				Computed:    true,
			},
			"targets": {
				Type:        schema.TypeList,
				Description: "Quarantine entry to hold multiple MACs.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description for the quarantine entry.",
							Optional:    true,
							Computed:    true,
						},
						"entry": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Quarantine entry name.",
							Optional:    true,
							Computed:    true,
						},
						"macs": {
							Type:        schema.TypeList,
							Description: "Quarantine MACs.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Description for the quarantine MAC.",
										Optional:    true,
										Computed:    true,
									},
									"drop": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable/disable dropping of quarantined device traffic.",
										Optional:    true,
										Computed:    true,
									},
									"mac": {
										Type: schema.TypeString,

										Description: "Quarantine MAC.",
										Optional:    true,
										Computed:    true,
									},
									"parent": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Parent entry name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"traffic_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Traffic policy for quarantined MACs.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserQuarantineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserQuarantine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserQuarantine(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserQuarantine")
	}

	return resourceUserQuarantineRead(ctx, d, meta)
}

func resourceUserQuarantineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserQuarantine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserQuarantine(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserQuarantine resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserQuarantine")
	}

	return resourceUserQuarantineRead(ctx, d, meta)
}

func resourceUserQuarantineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectUserQuarantine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateUserQuarantine(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserQuarantineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserQuarantine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserQuarantine resource: %v", err)
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

	diags := refreshObjectUserQuarantine(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserQuarantineTargets(d *schema.ResourceData, v *[]models.UserQuarantineTargets, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Entry; tmp != nil {
				v["entry"] = *tmp
			}

			if tmp := cfg.Macs; tmp != nil {
				v["macs"] = flattenUserQuarantineTargetsMacs(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "macs"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "entry")
	}

	return flat
}

func flattenUserQuarantineTargetsMacs(d *schema.ResourceData, v *[]models.UserQuarantineTargetsMacs, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Drop; tmp != nil {
				v["drop"] = *tmp
			}

			if tmp := cfg.Mac; tmp != nil {
				v["mac"] = *tmp
			}

			if tmp := cfg.Parent; tmp != nil {
				v["parent"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "mac")
	}

	return flat
}

func refreshObjectUserQuarantine(d *schema.ResourceData, o *models.UserQuarantine, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FirewallGroups != nil {
		v := *o.FirewallGroups

		if err = d.Set("firewall_groups", v); err != nil {
			return diag.Errorf("error reading firewall_groups: %v", err)
		}
	}

	if o.Quarantine != nil {
		v := *o.Quarantine

		if err = d.Set("quarantine", v); err != nil {
			return diag.Errorf("error reading quarantine: %v", err)
		}
	}

	if o.Targets != nil {
		if err = d.Set("targets", flattenUserQuarantineTargets(d, o.Targets, "targets", sort)); err != nil {
			return diag.Errorf("error reading targets: %v", err)
		}
	}

	if o.TrafficPolicy != nil {
		v := *o.TrafficPolicy

		if err = d.Set("traffic_policy", v); err != nil {
			return diag.Errorf("error reading traffic_policy: %v", err)
		}
	}

	return nil
}

func expandUserQuarantineTargets(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserQuarantineTargets, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserQuarantineTargets

	for i := range l {
		tmp := models.UserQuarantineTargets{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.entry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Entry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.macs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandUserQuarantineTargetsMacs(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.UserQuarantineTargetsMacs
			// 	}
			tmp.Macs = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandUserQuarantineTargetsMacs(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserQuarantineTargetsMacs, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserQuarantineTargetsMacs

	for i := range l {
		tmp := models.UserQuarantineTargetsMacs{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Drop = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mac = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.parent", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Parent = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectUserQuarantine(d *schema.ResourceData, sv string) (*models.UserQuarantine, diag.Diagnostics) {
	obj := models.UserQuarantine{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("firewall_groups"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("firewall_groups", sv)
				diags = append(diags, e)
			}
			obj.FirewallGroups = &v2
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
	if v, ok := d.GetOk("targets"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("targets", sv)
			diags = append(diags, e)
		}
		t, err := expandUserQuarantineTargets(d, v, "targets", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Targets = t
		}
	} else if d.HasChange("targets") {
		old, new := d.GetChange("targets")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Targets = &[]models.UserQuarantineTargets{}
		}
	}
	if v1, ok := d.GetOk("traffic_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_policy", sv)
				diags = append(diags, e)
			}
			obj.TrafficPolicy = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectUserQuarantine(d *schema.ResourceData, sv string) (*models.UserQuarantine, diag.Diagnostics) {
	obj := models.UserQuarantine{}
	diags := diag.Diagnostics{}

	obj.Targets = &[]models.UserQuarantineTargets{}

	return &obj, diags
}
