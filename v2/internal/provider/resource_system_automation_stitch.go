// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Automation stitches.

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

func resourceSystemAutomationStitch() *schema.Resource {
	return &schema.Resource{
		Description: "Automation stitches.",

		CreateContext: resourceSystemAutomationStitchCreate,
		ReadContext:   resourceSystemAutomationStitchRead,
		UpdateContext: resourceSystemAutomationStitchUpdate,
		DeleteContext: resourceSystemAutomationStitchDelete,

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
			"action": {
				Type:        schema.TypeList,
				Description: "Action names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Action name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"actions": {
				Type:        schema.TypeList,
				Description: "Configure stitch actions.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Action name.",
							Optional:    true,
							Computed:    true,
						},
						"delay": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),

							Description: "Delay before execution (in seconds).",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"required": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Required in action chain.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"destination": {
				Type:        schema.TypeList,
				Description: "Serial number/HA group-name of destination devices.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Destination name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this stitch.",
				Optional:    true,
				Computed:    true,
			},
			"trigger": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Trigger name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAutomationStitchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemAutomationStitch resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemAutomationStitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAutomationStitch(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutomationStitch")
	}

	return resourceSystemAutomationStitchRead(ctx, d, meta)
}

func resourceSystemAutomationStitchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutomationStitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAutomationStitch(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAutomationStitch resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutomationStitch")
	}

	return resourceSystemAutomationStitchRead(ctx, d, meta)
}

func resourceSystemAutomationStitchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemAutomationStitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAutomationStitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationStitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutomationStitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutomationStitch resource: %v", err)
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

	diags := refreshObjectSystemAutomationStitch(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemAutomationStitchAction(d *schema.ResourceData, v *[]models.SystemAutomationStitchAction, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemAutomationStitchActions(d *schema.ResourceData, v *[]models.SystemAutomationStitchActions, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Delay; tmp != nil {
				v["delay"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Required; tmp != nil {
				v["required"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemAutomationStitchDestination(d *schema.ResourceData, v *[]models.SystemAutomationStitchDestination, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemAutomationStitch(d *schema.ResourceData, o *models.SystemAutomationStitch, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		if err = d.Set("action", flattenSystemAutomationStitchAction(d, o.Action, "action", sort)); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.Actions != nil {
		if err = d.Set("actions", flattenSystemAutomationStitchActions(d, o.Actions, "actions", sort)); err != nil {
			return diag.Errorf("error reading actions: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.Destination != nil {
		if err = d.Set("destination", flattenSystemAutomationStitchDestination(d, o.Destination, "destination", sort)); err != nil {
			return diag.Errorf("error reading destination: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Trigger != nil {
		v := *o.Trigger

		if err = d.Set("trigger", v); err != nil {
			return diag.Errorf("error reading trigger: %v", err)
		}
	}

	return nil
}

func expandSystemAutomationStitchAction(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationStitchAction, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationStitchAction

	for i := range l {
		tmp := models.SystemAutomationStitchAction{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAutomationStitchActions(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationStitchActions, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationStitchActions

	for i := range l {
		tmp := models.SystemAutomationStitchActions{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.delay", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Delay = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.required", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Required = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAutomationStitchDestination(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationStitchDestination, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationStitchDestination

	for i := range l {
		tmp := models.SystemAutomationStitchDestination{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemAutomationStitch(d *schema.ResourceData, sv string) (*models.SystemAutomationStitch, diag.Diagnostics) {
	obj := models.SystemAutomationStitch{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("action"); ok {
		if !utils.CheckVer(sv, "", "v7.0.1") {
			e := utils.AttributeVersionWarning("action", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationStitchAction(d, v, "action", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Action = t
		}
	} else if d.HasChange("action") {
		old, new := d.GetChange("action")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Action = &[]models.SystemAutomationStitchAction{}
		}
	}
	if v, ok := d.GetOk("actions"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("actions", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationStitchActions(d, v, "actions", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Actions = t
		}
	} else if d.HasChange("actions") {
		old, new := d.GetChange("actions")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Actions = &[]models.SystemAutomationStitchActions{}
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v, ok := d.GetOk("destination"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("destination", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationStitchDestination(d, v, "destination", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Destination = t
		}
	} else if d.HasChange("destination") {
		old, new := d.GetChange("destination")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Destination = &[]models.SystemAutomationStitchDestination{}
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
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("trigger"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trigger", sv)
				diags = append(diags, e)
			}
			obj.Trigger = &v2
		}
	}
	return &obj, diags
}
