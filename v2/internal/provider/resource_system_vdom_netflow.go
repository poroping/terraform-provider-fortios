// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NetFlow per VDOM.

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

func resourceSystemVdomNetflow() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NetFlow per VDOM.",

		CreateContext: resourceSystemVdomNetflowCreate,
		ReadContext:   resourceSystemVdomNetflowRead,
		UpdateContext: resourceSystemVdomNetflowUpdate,
		DeleteContext: resourceSystemVdomNetflowDelete,

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
			"collector_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "NetFlow collector IP address.",
				Optional:    true,
				Computed:    true,
			},
			"collector_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "NetFlow collector port number.",
				Optional:    true,
				Computed:    true,
			},
			"collectors": {
				Type:        schema.TypeList,
				Description: "Netflow collectors.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"collector_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Collector IP.",
							Optional:    true,
							Computed:    true,
						},
						"collector_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "NetFlow collector port number.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 6),

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Specify outgoing interface to reach server.",
							Optional:    true,
							Computed:    true,
						},
						"interface_select_method": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

							Description: "Specify how to select outgoing interface to reach server.",
							Optional:    true,
							Computed:    true,
						},
						"source_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Source IP address for communication with the NetFlow agent.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Source IP address for communication with the NetFlow agent.",
				Optional:    true,
				Computed:    true,
			},
			"vdom_netflow": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NetFlow per VDOM.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemVdomNetflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVdomNetflow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVdomNetflow(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomNetflow")
	}

	return resourceSystemVdomNetflowRead(ctx, d, meta)
}

func resourceSystemVdomNetflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVdomNetflow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVdomNetflow(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVdomNetflow resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVdomNetflow")
	}

	return resourceSystemVdomNetflowRead(ctx, d, meta)
}

func resourceSystemVdomNetflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemVdomNetflow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemVdomNetflow(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVdomNetflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomNetflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVdomNetflow(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdomNetflow resource: %v", err)
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

	diags := refreshObjectSystemVdomNetflow(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemVdomNetflowCollectors(d *schema.ResourceData, v *[]models.SystemVdomNetflowCollectors, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CollectorIp; tmp != nil {
				v["collector_ip"] = *tmp
			}

			if tmp := cfg.CollectorPort; tmp != nil {
				v["collector_port"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.InterfaceSelectMethod; tmp != nil {
				v["interface_select_method"] = *tmp
			}

			if tmp := cfg.SourceIp; tmp != nil {
				v["source_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemVdomNetflow(d *schema.ResourceData, o *models.SystemVdomNetflow, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CollectorIp != nil {
		v := *o.CollectorIp

		if err = d.Set("collector_ip", v); err != nil {
			return diag.Errorf("error reading collector_ip: %v", err)
		}
	}

	if o.CollectorPort != nil {
		v := *o.CollectorPort

		if err = d.Set("collector_port", v); err != nil {
			return diag.Errorf("error reading collector_port: %v", err)
		}
	}

	if o.Collectors != nil {
		if err = d.Set("collectors", flattenSystemVdomNetflowCollectors(d, o.Collectors, "collectors", sort)); err != nil {
			return diag.Errorf("error reading collectors: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.VdomNetflow != nil {
		v := *o.VdomNetflow

		if err = d.Set("vdom_netflow", v); err != nil {
			return diag.Errorf("error reading vdom_netflow: %v", err)
		}
	}

	return nil
}

func expandSystemVdomNetflowCollectors(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVdomNetflowCollectors, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVdomNetflowCollectors

	for i := range l {
		tmp := models.SystemVdomNetflowCollectors{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.collector_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CollectorIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.collector_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CollectorPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface_select_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceSelectMethod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemVdomNetflow(d *schema.ResourceData, sv string) (*models.SystemVdomNetflow, diag.Diagnostics) {
	obj := models.SystemVdomNetflow{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("collector_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.8") {
				e := utils.AttributeVersionWarning("collector_ip", sv)
				diags = append(diags, e)
			}
			obj.CollectorIp = &v2
		}
	}
	if v1, ok := d.GetOk("collector_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.2.8") {
				e := utils.AttributeVersionWarning("collector_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CollectorPort = &tmp
		}
	}
	if v, ok := d.GetOk("collectors"); ok {
		if !utils.CheckVer(sv, "v7.2.8", "") {
			e := utils.AttributeVersionWarning("collectors", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVdomNetflowCollectors(d, v, "collectors", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Collectors = t
		}
	} else if d.HasChange("collectors") {
		old, new := d.GetChange("collectors")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Collectors = &[]models.SystemVdomNetflowCollectors{}
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "v7.2.8") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "v7.2.8") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.8") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("vdom_netflow"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom_netflow", sv)
				diags = append(diags, e)
			}
			obj.VdomNetflow = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemVdomNetflow(d *schema.ResourceData, sv string) (*models.SystemVdomNetflow, diag.Diagnostics) {
	obj := models.SystemVdomNetflow{}
	diags := diag.Diagnostics{}

	obj.Collectors = &[]models.SystemVdomNetflowCollectors{}

	return &obj, diags
}
