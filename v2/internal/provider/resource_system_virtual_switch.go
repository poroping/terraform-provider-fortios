// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure virtual hardware switch interfaces.

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

func resourceSystemVirtualSwitch() *schema.Resource {
	return &schema.Resource{
		Description: "Configure virtual hardware switch interfaces.",

		CreateContext: resourceSystemVirtualSwitchCreate,
		ReadContext:   resourceSystemVirtualSwitchRead,
		UpdateContext: resourceSystemVirtualSwitchUpdate,
		DeleteContext: resourceSystemVirtualSwitchDelete,

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
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Name of the virtual switch.",
				ForceNew:    true,
				Required:    true,
			},
			"physical_switch": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Physical switch parent.",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeList,
				Description: "Configure member ports.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alias": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 25),

							Description: "Alias.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Physical interface name.",
							Optional:    true,
							Computed:    true,
						},
						"speed": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "10full", "10half", "100full", "100half", "1000full", "1000half", "1000auto", "10000full", "10000auto", "40000full", "25000full"}, false),

							Description: "Interface speed.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"up", "down"}, false),

							Description: "Interface status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"span": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable SPAN.",
				Optional:    true,
				Computed:    true,
			},
			"span_dest_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "SPAN destination port.",
				Optional:    true,
				Computed:    true,
			},
			"span_direction": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"rx", "tx", "both"}, false),

				Description: "SPAN direction.",
				Optional:    true,
				Computed:    true,
			},
			"span_source_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "SPAN source port.",
				Optional:    true,
				Computed:    true,
			},
			"vlan": {
				Type: schema.TypeInt,

				Description: "VLAN.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemVirtualSwitchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemVirtualSwitch resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemVirtualSwitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVirtualSwitch(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVirtualSwitch")
	}

	return resourceSystemVirtualSwitchRead(ctx, d, meta)
}

func resourceSystemVirtualSwitchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVirtualSwitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVirtualSwitch(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVirtualSwitch resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVirtualSwitch")
	}

	return resourceSystemVirtualSwitchRead(ctx, d, meta)
}

func resourceSystemVirtualSwitchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemVirtualSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVirtualSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVirtualSwitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVirtualSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVirtualSwitch resource: %v", err)
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

	diags := refreshObjectSystemVirtualSwitch(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemVirtualSwitchPort(v *[]models.SystemVirtualSwitchPort, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Alias; tmp != nil {
				v["alias"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Speed; tmp != nil {
				v["speed"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemVirtualSwitch(d *schema.ResourceData, o *models.SystemVirtualSwitch, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PhysicalSwitch != nil {
		v := *o.PhysicalSwitch

		if err = d.Set("physical_switch", v); err != nil {
			return diag.Errorf("error reading physical_switch: %v", err)
		}
	}

	if o.Port != nil {
		if err = d.Set("port", flattenSystemVirtualSwitchPort(o.Port, sort)); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Span != nil {
		v := *o.Span

		if err = d.Set("span", v); err != nil {
			return diag.Errorf("error reading span: %v", err)
		}
	}

	if o.SpanDestPort != nil {
		v := *o.SpanDestPort

		if err = d.Set("span_dest_port", v); err != nil {
			return diag.Errorf("error reading span_dest_port: %v", err)
		}
	}

	if o.SpanDirection != nil {
		v := *o.SpanDirection

		if err = d.Set("span_direction", v); err != nil {
			return diag.Errorf("error reading span_direction: %v", err)
		}
	}

	if o.SpanSourcePort != nil {
		v := *o.SpanSourcePort

		if err = d.Set("span_source_port", v); err != nil {
			return diag.Errorf("error reading span_source_port: %v", err)
		}
	}

	if o.Vlan != nil {
		v := *o.Vlan

		if err = d.Set("vlan", v); err != nil {
			return diag.Errorf("error reading vlan: %v", err)
		}
	}

	return nil
}

func expandSystemVirtualSwitchPort(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualSwitchPort, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualSwitchPort

	for i := range l {
		tmp := models.SystemVirtualSwitchPort{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.alias", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Alias = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.speed", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Speed = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemVirtualSwitch(d *schema.ResourceData, sv string) (*models.SystemVirtualSwitch, diag.Diagnostics) {
	obj := models.SystemVirtualSwitch{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("physical_switch"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("physical_switch", sv)
				diags = append(diags, e)
			}
			obj.PhysicalSwitch = &v2
		}
	}
	if v, ok := d.GetOk("port"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("port", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVirtualSwitchPort(d, v, "port", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Port = t
		}
	} else if d.HasChange("port") {
		old, new := d.GetChange("port")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Port = &[]models.SystemVirtualSwitchPort{}
		}
	}
	if v1, ok := d.GetOk("span"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("span", sv)
				diags = append(diags, e)
			}
			obj.Span = &v2
		}
	}
	if v1, ok := d.GetOk("span_dest_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("span_dest_port", sv)
				diags = append(diags, e)
			}
			obj.SpanDestPort = &v2
		}
	}
	if v1, ok := d.GetOk("span_direction"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("span_direction", sv)
				diags = append(diags, e)
			}
			obj.SpanDirection = &v2
		}
	}
	if v1, ok := d.GetOk("span_source_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("span_source_port", sv)
				diags = append(diags, e)
			}
			obj.SpanSourcePort = &v2
		}
	}
	if v1, ok := d.GetOk("vlan"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vlan = &tmp
		}
	}
	return &obj, diags
}
