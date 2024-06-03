// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure software switch interfaces by grouping physical and WiFi interfaces.

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

func resourceSystemSwitchInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure software switch interfaces by grouping physical and WiFi interfaces.",

		CreateContext: resourceSystemSwitchInterfaceCreate,
		ReadContext:   resourceSystemSwitchInterfaceRead,
		UpdateContext: resourceSystemSwitchInterfaceUpdate,
		DeleteContext: resourceSystemSwitchInterfaceDelete,

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
			"intra_switch_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"implicit", "explicit"}, false),

				Description: "Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces.",
				Optional:    true,
				Computed:    true,
			},
			"mac_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),

				Description: "Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Names of the interfaces that belong to the virtual switch.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Physical interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).",
				ForceNew:    true,
				Required:    true,
			},
			"span": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port.",
				Optional:    true,
				Computed:    true,
			},
			"span_dest_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.",
				Optional:    true,
				Computed:    true,
			},
			"span_direction": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"rx", "tx", "both"}, false),

				Description: "The direction in which the SPAN port operates, either: rx, tx, or both.",
				Optional:    true,
				Computed:    true,
			},
			"span_source_port": {
				Type:        schema.TypeList,
				Description: "Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Physical interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"switch", "hub"}, false),

				Description: "Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "VDOM that the software switch belongs to.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSwitchInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSwitchInterface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSwitchInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSwitchInterface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSwitchInterface")
	}

	return resourceSystemSwitchInterfaceRead(ctx, d, meta)
}

func resourceSystemSwitchInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSwitchInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSwitchInterface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSwitchInterface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSwitchInterface")
	}

	return resourceSystemSwitchInterfaceRead(ctx, d, meta)
}

func resourceSystemSwitchInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSwitchInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSwitchInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSwitchInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSwitchInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSwitchInterface resource: %v", err)
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

	diags := refreshObjectSystemSwitchInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSwitchInterfaceMember(d *schema.ResourceData, v *[]models.SystemSwitchInterfaceMember, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.InterfaceName; tmp != nil {
				v["interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface_name")
	}

	return flat
}

func flattenSystemSwitchInterfaceSpanSourcePort(d *schema.ResourceData, v *[]models.SystemSwitchInterfaceSpanSourcePort, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.InterfaceName; tmp != nil {
				v["interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface_name")
	}

	return flat
}

func refreshObjectSystemSwitchInterface(d *schema.ResourceData, o *models.SystemSwitchInterface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.IntraSwitchPolicy != nil {
		v := *o.IntraSwitchPolicy

		if err = d.Set("intra_switch_policy", v); err != nil {
			return diag.Errorf("error reading intra_switch_policy: %v", err)
		}
	}

	if o.MacTtl != nil {
		v := *o.MacTtl

		if err = d.Set("mac_ttl", v); err != nil {
			return diag.Errorf("error reading mac_ttl: %v", err)
		}
	}

	if o.Member != nil {
		if err = d.Set("member", flattenSystemSwitchInterfaceMember(d, o.Member, "member", sort)); err != nil {
			return diag.Errorf("error reading member: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
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
		if err = d.Set("span_source_port", flattenSystemSwitchInterfaceSpanSourcePort(d, o.SpanSourcePort, "span_source_port", sort)); err != nil {
			return diag.Errorf("error reading span_source_port: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Vdom != nil {
		v := *o.Vdom

		if err = d.Set("vdom", v); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	return nil
}

func expandSystemSwitchInterfaceMember(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSwitchInterfaceMember, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSwitchInterfaceMember

	for i := range l {
		tmp := models.SystemSwitchInterfaceMember{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSwitchInterfaceSpanSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSwitchInterfaceSpanSourcePort, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSwitchInterfaceSpanSourcePort

	for i := range l {
		tmp := models.SystemSwitchInterfaceSpanSourcePort{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemSwitchInterface(d *schema.ResourceData, sv string) (*models.SystemSwitchInterface, diag.Diagnostics) {
	obj := models.SystemSwitchInterface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("intra_switch_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("intra_switch_policy", sv)
				diags = append(diags, e)
			}
			obj.IntraSwitchPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("mac_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("mac_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MacTtl = &tmp
		}
	}
	if v, ok := d.GetOk("member"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("member", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSwitchInterfaceMember(d, v, "member", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Member = t
		}
	} else if d.HasChange("member") {
		old, new := d.GetChange("member")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Member = &[]models.SystemSwitchInterfaceMember{}
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
	if v1, ok := d.GetOk("span"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("span", sv)
				diags = append(diags, e)
			}
			obj.Span = &v2
		}
	}
	if v1, ok := d.GetOk("span_dest_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("span_dest_port", sv)
				diags = append(diags, e)
			}
			obj.SpanDestPort = &v2
		}
	}
	if v1, ok := d.GetOk("span_direction"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("span_direction", sv)
				diags = append(diags, e)
			}
			obj.SpanDirection = &v2
		}
	}
	if v, ok := d.GetOk("span_source_port"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("span_source_port", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSwitchInterfaceSpanSourcePort(d, v, "span_source_port", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SpanSourcePort = t
		}
	} else if d.HasChange("span_source_port") {
		old, new := d.GetChange("span_source_port")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SpanSourcePort = &[]models.SystemSwitchInterfaceSpanSourcePort{}
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
	if v1, ok := d.GetOk("vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom", sv)
				diags = append(diags, e)
			}
			obj.Vdom = &v2
		}
	}
	return &obj, diags
}
