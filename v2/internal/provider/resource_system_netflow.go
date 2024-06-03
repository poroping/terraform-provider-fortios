// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NetFlow.

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

func resourceSystemNetflow() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NetFlow.",

		CreateContext: resourceSystemNetflowCreate,
		ReadContext:   resourceSystemNetflowRead,
		UpdateContext: resourceSystemNetflowUpdate,
		DeleteContext: resourceSystemNetflowDelete,

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
			"active_flow_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 3600),

				Description: "Timeout to report active flows (60 - 3600 sec, default = 1800).",
				Optional:    true,
				Computed:    true,
			},
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
			"inactive_flow_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 600),

				Description: "Timeout for periodic report of finished flows (10 - 600 sec, default = 15).",
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
			"template_tx_counter": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 6000),

				Description: "Counter of flowset records before resending a template flowset record.",
				Optional:    true,
				Computed:    true,
			},
			"template_tx_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),

				Description: "Timeout for periodic template flowset transmission (60 - 86400 sec, default = 1800).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemNetflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNetflow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemNetflow(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNetflow")
	}

	return resourceSystemNetflowRead(ctx, d, meta)
}

func resourceSystemNetflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNetflow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemNetflow(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemNetflow resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNetflow")
	}

	return resourceSystemNetflowRead(ctx, d, meta)
}

func resourceSystemNetflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemNetflow(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemNetflow(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemNetflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNetflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemNetflow(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNetflow resource: %v", err)
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

	diags := refreshObjectSystemNetflow(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemNetflowCollectors(d *schema.ResourceData, v *[]models.SystemNetflowCollectors, prefix string, sort bool) interface{} {
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

func refreshObjectSystemNetflow(d *schema.ResourceData, o *models.SystemNetflow, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ActiveFlowTimeout != nil {
		v := *o.ActiveFlowTimeout

		if err = d.Set("active_flow_timeout", v); err != nil {
			return diag.Errorf("error reading active_flow_timeout: %v", err)
		}
	}

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
		if err = d.Set("collectors", flattenSystemNetflowCollectors(d, o.Collectors, "collectors", sort)); err != nil {
			return diag.Errorf("error reading collectors: %v", err)
		}
	}

	if o.InactiveFlowTimeout != nil {
		v := *o.InactiveFlowTimeout

		if err = d.Set("inactive_flow_timeout", v); err != nil {
			return diag.Errorf("error reading inactive_flow_timeout: %v", err)
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

	if o.TemplateTxCounter != nil {
		v := *o.TemplateTxCounter

		if err = d.Set("template_tx_counter", v); err != nil {
			return diag.Errorf("error reading template_tx_counter: %v", err)
		}
	}

	if o.TemplateTxTimeout != nil {
		v := *o.TemplateTxTimeout

		if err = d.Set("template_tx_timeout", v); err != nil {
			return diag.Errorf("error reading template_tx_timeout: %v", err)
		}
	}

	return nil
}

func expandSystemNetflowCollectors(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNetflowCollectors, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNetflowCollectors

	for i := range l {
		tmp := models.SystemNetflowCollectors{}
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

func getObjectSystemNetflow(d *schema.ResourceData, sv string) (*models.SystemNetflow, diag.Diagnostics) {
	obj := models.SystemNetflow{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("active_flow_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("active_flow_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ActiveFlowTimeout = &tmp
		}
	}
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
		t, err := expandSystemNetflowCollectors(d, v, "collectors", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Collectors = t
		}
	} else if d.HasChange("collectors") {
		old, new := d.GetChange("collectors")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Collectors = &[]models.SystemNetflowCollectors{}
		}
	}
	if v1, ok := d.GetOk("inactive_flow_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inactive_flow_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.InactiveFlowTimeout = &tmp
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
	if v1, ok := d.GetOk("template_tx_counter"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("template_tx_counter", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TemplateTxCounter = &tmp
		}
	}
	if v1, ok := d.GetOk("template_tx_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("template_tx_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TemplateTxTimeout = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemNetflow(d *schema.ResourceData, sv string) (*models.SystemNetflow, diag.Diagnostics) {
	obj := models.SystemNetflow{}
	diags := diag.Diagnostics{}

	obj.Collectors = &[]models.SystemNetflowCollectors{}

	return &obj, diags
}
