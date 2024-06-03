// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure QoS map set.

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

func resourceWirelessControllerHotspot20QosMap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure QoS map set.",

		CreateContext: resourceWirelessControllerHotspot20QosMapCreate,
		ReadContext:   resourceWirelessControllerHotspot20QosMapRead,
		UpdateContext: resourceWirelessControllerHotspot20QosMapUpdate,
		DeleteContext: resourceWirelessControllerHotspot20QosMapDelete,

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
			"dscp_except": {
				Type:        schema.TypeList,
				Description: "Differentiated Services Code Point (DSCP) exceptions.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "DSCP value.",
							Optional:    true,
							Computed:    true,
						},
						"index": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 21),

							Description: "DSCP exception index.",
							Optional:    true,
							Computed:    true,
						},
						"up": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "User priority.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dscp_range": {
				Type:        schema.TypeList,
				Description: "Differentiated Services Code Point (DSCP) ranges.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "DSCP high value.",
							Optional:    true,
							Computed:    true,
						},
						"index": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 8),

							Description: "DSCP range index.",
							Optional:    true,
							Computed:    true,
						},
						"low": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "DSCP low value.",
							Optional:    true,
							Computed:    true,
						},
						"up": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "User priority.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "QOS-MAP name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20QosMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20QosMap resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20QosMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20QosMap(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20QosMap")
	}

	return resourceWirelessControllerHotspot20QosMapRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20QosMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20QosMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20QosMap(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20QosMap resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20QosMap")
	}

	return resourceWirelessControllerHotspot20QosMapRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20QosMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20QosMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20QosMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20QosMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20QosMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20QosMap resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20QosMap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData, v *[]models.WirelessControllerHotspot20QosMapDscpExcept, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dscp; tmp != nil {
				v["dscp"] = *tmp
			}

			if tmp := cfg.Index; tmp != nil {
				v["index"] = *tmp
			}

			if tmp := cfg.Up; tmp != nil {
				v["up"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "index")
	}

	return flat
}

func flattenWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData, v *[]models.WirelessControllerHotspot20QosMapDscpRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.High; tmp != nil {
				v["high"] = *tmp
			}

			if tmp := cfg.Index; tmp != nil {
				v["index"] = *tmp
			}

			if tmp := cfg.Low; tmp != nil {
				v["low"] = *tmp
			}

			if tmp := cfg.Up; tmp != nil {
				v["up"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "index")
	}

	return flat
}

func refreshObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData, o *models.WirelessControllerHotspot20QosMap, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DscpExcept != nil {
		if err = d.Set("dscp_except", flattenWirelessControllerHotspot20QosMapDscpExcept(d, o.DscpExcept, "dscp_except", sort)); err != nil {
			return diag.Errorf("error reading dscp_except: %v", err)
		}
	}

	if o.DscpRange != nil {
		if err = d.Set("dscp_range", flattenWirelessControllerHotspot20QosMapDscpRange(d, o.DscpRange, "dscp_range", sort)); err != nil {
			return diag.Errorf("error reading dscp_range: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20QosMapDscpExcept, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20QosMapDscpExcept

	for i := range l {
		tmp := models.WirelessControllerHotspot20QosMapDscpExcept{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dscp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dscp = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.index", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Index = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.up", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Up = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20QosMapDscpRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20QosMapDscpRange

	for i := range l {
		tmp := models.WirelessControllerHotspot20QosMapDscpRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.high", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.High = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.index", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Index = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.low", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Low = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.up", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Up = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20QosMap, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20QosMap{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("dscp_except"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dscp_except", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20QosMapDscpExcept(d, v, "dscp_except", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DscpExcept = t
		}
	} else if d.HasChange("dscp_except") {
		old, new := d.GetChange("dscp_except")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DscpExcept = &[]models.WirelessControllerHotspot20QosMapDscpExcept{}
		}
	}
	if v, ok := d.GetOk("dscp_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dscp_range", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20QosMapDscpRange(d, v, "dscp_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DscpRange = t
		}
	} else if d.HasChange("dscp_range") {
		old, new := d.GetChange("dscp_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DscpRange = &[]models.WirelessControllerHotspot20QosMapDscpRange{}
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
	return &obj, diags
}
