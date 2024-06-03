// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 multicast.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceRouterMulticast6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 multicast.",

		CreateContext: resourceRouterMulticast6Create,
		ReadContext:   resourceRouterMulticast6Read,
		UpdateContext: resourceRouterMulticast6Update,
		DeleteContext: resourceRouterMulticast6Delete,

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
			"interface": {
				Type:        schema.TypeList,
				Description: "Protocol Independent Multicast (PIM) interfaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_holdtime": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Time before old neighbor information expires in seconds (1 - 65535, default = 105).",
							Optional:    true,
							Computed:    true,
						},
						"hello_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Interval between sending PIM hello messages in seconds (1 - 65535, default = 30).",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"multicast_pmtu": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PMTU for IPv6 multicast.",
				Optional:    true,
				Computed:    true,
			},
			"multicast_routing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 multicast routing.",
				Optional:    true,
				Computed:    true,
			},
			"pim_sm_global": {
				Type:        schema.TypeList,
				Description: "PIM sparse-mode global settings.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"register_rate_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Limit of packets/sec per source registered through this RP (0 means unlimited).",
							Optional:    true,
							Computed:    true,
						},
						"rp_address": {
							Type:        schema.TypeList,
							Description: "Statically configured RP addresses.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "ID of the entry.",
										Optional:    true,
										Computed:    true,
									},
									"ip6_address": {
										Type:             schema.TypeString,
										ValidateFunc:     validation.IsIPv6Address,
										DiffSuppressFunc: suppressors.DiffIPEqual,
										Description:      "RP router IPv6 address.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceRouterMulticast6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterMulticast6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterMulticast6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterMulticast6")
	}

	return resourceRouterMulticast6Read(ctx, d, meta)
}

func resourceRouterMulticast6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterMulticast6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterMulticast6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterMulticast6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterMulticast6")
	}

	return resourceRouterMulticast6Read(ctx, d, meta)
}

func resourceRouterMulticast6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterMulticast6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterMulticast6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterMulticast6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterMulticast6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterMulticast6 resource: %v", err)
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

	diags := refreshObjectRouterMulticast6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterMulticast6Interface(d *schema.ResourceData, v *[]models.RouterMulticast6Interface, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.HelloHoldtime; tmp != nil {
				v["hello_holdtime"] = *tmp
			}

			if tmp := cfg.HelloInterval; tmp != nil {
				v["hello_interval"] = *tmp
			}

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

func flattenRouterMulticast6PimSmGlobal(d *schema.ResourceData, v *models.RouterMulticast6PimSmGlobal, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.RouterMulticast6PimSmGlobal{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.RegisterRateLimit; tmp != nil {
				v["register_rate_limit"] = *tmp
			}

			if tmp := cfg.RpAddress; tmp != nil {
				v["rp_address"] = flattenRouterMulticast6PimSmGlobalRpAddress(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "rp_address"), sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenRouterMulticast6PimSmGlobalRpAddress(d *schema.ResourceData, v *[]models.RouterMulticast6PimSmGlobalRpAddress, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip6Address; tmp != nil {
				v["ip6_address"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterMulticast6(d *schema.ResourceData, o *models.RouterMulticast6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Interface != nil {
		if err = d.Set("interface", flattenRouterMulticast6Interface(d, o.Interface, "interface", sort)); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.MulticastPmtu != nil {
		v := *o.MulticastPmtu

		if err = d.Set("multicast_pmtu", v); err != nil {
			return diag.Errorf("error reading multicast_pmtu: %v", err)
		}
	}

	if o.MulticastRouting != nil {
		v := *o.MulticastRouting

		if err = d.Set("multicast_routing", v); err != nil {
			return diag.Errorf("error reading multicast_routing: %v", err)
		}
	}

	if _, ok := d.GetOk("pim_sm_global"); ok {
		if o.PimSmGlobal != nil {
			if err = d.Set("pim_sm_global", flattenRouterMulticast6PimSmGlobal(d, o.PimSmGlobal, "pim_sm_global", sort)); err != nil {
				return diag.Errorf("error reading pim_sm_global: %v", err)
			}
		}
	}

	return nil
}

func expandRouterMulticast6Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterMulticast6Interface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticast6Interface

	for i := range l {
		tmp := models.RouterMulticast6Interface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.hello_holdtime", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HelloHoldtime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HelloInterval = &v3
			}
		}

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

func expandRouterMulticast6PimSmGlobal(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.RouterMulticast6PimSmGlobal, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticast6PimSmGlobal

	for i := range l {
		tmp := models.RouterMulticast6PimSmGlobal{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.register_rate_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RegisterRateLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rp_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterMulticast6PimSmGlobalRpAddress(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterMulticast6PimSmGlobalRpAddress
			// 	}
			tmp.RpAddress = v2

		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandRouterMulticast6PimSmGlobalRpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterMulticast6PimSmGlobalRpAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticast6PimSmGlobalRpAddress

	for i := range l {
		tmp := models.RouterMulticast6PimSmGlobalRpAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6Address = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterMulticast6(d *schema.ResourceData, sv string) (*models.RouterMulticast6, diag.Diagnostics) {
	obj := models.RouterMulticast6{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterMulticast6Interface(d, v, "interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Interface = t
		}
	} else if d.HasChange("interface") {
		old, new := d.GetChange("interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Interface = &[]models.RouterMulticast6Interface{}
		}
	}
	if v1, ok := d.GetOk("multicast_pmtu"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_pmtu", sv)
				diags = append(diags, e)
			}
			obj.MulticastPmtu = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_routing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_routing", sv)
				diags = append(diags, e)
			}
			obj.MulticastRouting = &v2
		}
	}
	if v, ok := d.GetOk("pim_sm_global"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("pim_sm_global", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterMulticast6PimSmGlobal(d, v, "pim_sm_global", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PimSmGlobal = t
		}
	} else if d.HasChange("pim_sm_global") {
		old, new := d.GetChange("pim_sm_global")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PimSmGlobal = &models.RouterMulticast6PimSmGlobal{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterMulticast6(d *schema.ResourceData, sv string) (*models.RouterMulticast6, diag.Diagnostics) {
	obj := models.RouterMulticast6{}
	diags := diag.Diagnostics{}

	obj.Interface = &[]models.RouterMulticast6Interface{}
	obj.PimSmGlobal = &models.RouterMulticast6PimSmGlobal{}

	return &obj, diags
}
