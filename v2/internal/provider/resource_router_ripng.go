// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure RIPng.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterRipng() *schema.Resource {
	return &schema.Resource{
		Description: "Configure RIPng.",

		CreateContext: resourceRouterRipngCreate,
		ReadContext:   resourceRouterRipngRead,
		UpdateContext: resourceRouterRipngUpdate,
		DeleteContext: resourceRouterRipngDelete,

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
			"aggregate_address": {
				Type:        schema.TypeList,
				Description: "Aggregate address.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Aggregate address entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "Aggregate address prefix.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"default_information_originate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable generation of default route.",
				Optional:    true,
				Computed:    true,
			},
			"default_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16),

				Description: "Default metric.",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeList,
				Description: "Distance.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Access list for route destination.",
							Optional:    true,
							Computed:    true,
						},
						"distance": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Distance (1 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Distance ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "Distance prefix6.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"distribute_list": {
				Type:        schema.TypeList,
				Description: "Distribute list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),

							Description: "Distribute list direction.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Distribute list ID.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Distribute list interface name.",
							Optional:    true,
							Computed:    true,
						},
						"listname": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Distribute access/prefix list name.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"garbage_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 2147483647),

				Description: "Garbage timer.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "RIPng interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flags": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Flags.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
						"split_horizon": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"poisoned", "regular"}, false),

							Description: "Enable/disable split horizon.",
							Optional:    true,
							Computed:    true,
						},
						"split_horizon_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable split horizon.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"max_out_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),

				Description: "Maximum metric allowed to output(0 means 'not set').",
				Optional:    true,
				Computed:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "Neighbor.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Neighbor entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
						"ip6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 link-local address.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "Network.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Network entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "Network IPv6 link-local prefix.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"offset_list": {
				Type:        schema.TypeList,
				Description: "Offset list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 access list name.",
							Optional:    true,
							Computed:    true,
						},
						"direction": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),

							Description: "Offset list direction.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Offset-list ID.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
						"offset": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),

							Description: "Offset.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"passive_interface": {
				Type:        schema.TypeList,
				Description: "Passive interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Passive interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "Redistribute configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),

							Description: "Redistribute metric setting.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Redistribute name.",
							Optional:    true,
							Computed:    true,
						},
						"routemap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map name.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"timeout_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 2147483647),

				Description: "Timeout timer.",
				Optional:    true,
				Computed:    true,
			},
			"update_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 2147483647),

				Description: "Update timer.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterRipngCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterRipng(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterRipng(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRipng")
	}

	return resourceRouterRipngRead(ctx, d, meta)
}

func resourceRouterRipngUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterRipng(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterRipng(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterRipng resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRipng")
	}

	return resourceRouterRipngRead(ctx, d, meta)
}

func resourceRouterRipngDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterRipng(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterRipng(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterRipng resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipngRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterRipng(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterRipng resource: %v", err)
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

	diags := refreshObjectRouterRipng(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterRipngAggregateAddress(v *[]models.RouterRipngAggregateAddress, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix6; tmp != nil {
				v["prefix6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipngDistance(v *[]models.RouterRipngDistance, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AccessList6; tmp != nil {
				v["access_list6"] = *tmp
			}

			if tmp := cfg.Distance; tmp != nil {
				v["distance"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix6; tmp != nil {
				v["prefix6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipngDistributeList(v *[]models.RouterRipngDistributeList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Direction; tmp != nil {
				v["direction"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Listname; tmp != nil {
				v["listname"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipngInterface(v *[]models.RouterRipngInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Flags; tmp != nil {
				v["flags"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.SplitHorizon; tmp != nil {
				v["split_horizon"] = *tmp
			}

			if tmp := cfg.SplitHorizonStatus; tmp != nil {
				v["split_horizon_status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterRipngNeighbor(v *[]models.RouterRipngNeighbor, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Ip6; tmp != nil {
				v["ip6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipngNetwork(v *[]models.RouterRipngNetwork, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipngOffsetList(v *[]models.RouterRipngOffsetList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AccessList6; tmp != nil {
				v["access_list6"] = *tmp
			}

			if tmp := cfg.Direction; tmp != nil {
				v["direction"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Offset; tmp != nil {
				v["offset"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipngPassiveInterface(v *[]models.RouterRipngPassiveInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenRouterRipngRedistribute(v *[]models.RouterRipngRedistribute, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Metric; tmp != nil {
				v["metric"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Routemap; tmp != nil {
				v["routemap"] = *tmp
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

func refreshObjectRouterRipng(d *schema.ResourceData, o *models.RouterRipng, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AggregateAddress != nil {
		if err = d.Set("aggregate_address", flattenRouterRipngAggregateAddress(o.AggregateAddress, sort)); err != nil {
			return diag.Errorf("error reading aggregate_address: %v", err)
		}
	}

	if o.DefaultInformationOriginate != nil {
		v := *o.DefaultInformationOriginate

		if err = d.Set("default_information_originate", v); err != nil {
			return diag.Errorf("error reading default_information_originate: %v", err)
		}
	}

	if o.DefaultMetric != nil {
		v := *o.DefaultMetric

		if err = d.Set("default_metric", v); err != nil {
			return diag.Errorf("error reading default_metric: %v", err)
		}
	}

	if o.Distance != nil {
		if err = d.Set("distance", flattenRouterRipngDistance(o.Distance, sort)); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.DistributeList != nil {
		if err = d.Set("distribute_list", flattenRouterRipngDistributeList(o.DistributeList, sort)); err != nil {
			return diag.Errorf("error reading distribute_list: %v", err)
		}
	}

	if o.GarbageTimer != nil {
		v := *o.GarbageTimer

		if err = d.Set("garbage_timer", v); err != nil {
			return diag.Errorf("error reading garbage_timer: %v", err)
		}
	}

	if o.Interface != nil {
		if err = d.Set("interface", flattenRouterRipngInterface(o.Interface, sort)); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.MaxOutMetric != nil {
		v := *o.MaxOutMetric

		if err = d.Set("max_out_metric", v); err != nil {
			return diag.Errorf("error reading max_out_metric: %v", err)
		}
	}

	if o.Neighbor != nil {
		if err = d.Set("neighbor", flattenRouterRipngNeighbor(o.Neighbor, sort)); err != nil {
			return diag.Errorf("error reading neighbor: %v", err)
		}
	}

	if o.Network != nil {
		if err = d.Set("network", flattenRouterRipngNetwork(o.Network, sort)); err != nil {
			return diag.Errorf("error reading network: %v", err)
		}
	}

	if o.OffsetList != nil {
		if err = d.Set("offset_list", flattenRouterRipngOffsetList(o.OffsetList, sort)); err != nil {
			return diag.Errorf("error reading offset_list: %v", err)
		}
	}

	if o.PassiveInterface != nil {
		if err = d.Set("passive_interface", flattenRouterRipngPassiveInterface(o.PassiveInterface, sort)); err != nil {
			return diag.Errorf("error reading passive_interface: %v", err)
		}
	}

	if o.Redistribute != nil {
		if err = d.Set("redistribute", flattenRouterRipngRedistribute(o.Redistribute, sort)); err != nil {
			return diag.Errorf("error reading redistribute: %v", err)
		}
	}

	if o.TimeoutTimer != nil {
		v := *o.TimeoutTimer

		if err = d.Set("timeout_timer", v); err != nil {
			return diag.Errorf("error reading timeout_timer: %v", err)
		}
	}

	if o.UpdateTimer != nil {
		v := *o.UpdateTimer

		if err = d.Set("update_timer", v); err != nil {
			return diag.Errorf("error reading update_timer: %v", err)
		}
	}

	return nil
}

func expandRouterRipngAggregateAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngAggregateAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngAggregateAddress

	for i := range l {
		tmp := models.RouterRipngAggregateAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRipngDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngDistance, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngDistance

	for i := range l {
		tmp := models.RouterRipngDistance{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_list6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessList6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Distance = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRipngDistributeList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngDistributeList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngDistributeList

	for i := range l {
		tmp := models.RouterRipngDistributeList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.listname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Listname = &v2
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

func expandRouterRipngInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngInterface

	for i := range l {
		tmp := models.RouterRipngInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.flags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Flags = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.split_horizon", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SplitHorizon = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.split_horizon_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SplitHorizonStatus = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRipngNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngNeighbor

	for i := range l {
		tmp := models.RouterRipngNeighbor{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRipngNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngNetwork, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngNetwork

	for i := range l {
		tmp := models.RouterRipngNetwork{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRipngOffsetList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngOffsetList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngOffsetList

	for i := range l {
		tmp := models.RouterRipngOffsetList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_list6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessList6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.offset", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Offset = &v2
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

func expandRouterRipngPassiveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngPassiveInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngPassiveInterface

	for i := range l {
		tmp := models.RouterRipngPassiveInterface{}
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

func expandRouterRipngRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipngRedistribute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipngRedistribute

	for i := range l {
		tmp := models.RouterRipngRedistribute{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Metric = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Routemap = &v2
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

func getObjectRouterRipng(d *schema.ResourceData, sv string) (*models.RouterRipng, diag.Diagnostics) {
	obj := models.RouterRipng{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("aggregate_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("aggregate_address", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngAggregateAddress(d, v, "aggregate_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AggregateAddress = t
		}
	} else if d.HasChange("aggregate_address") {
		old, new := d.GetChange("aggregate_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AggregateAddress = &[]models.RouterRipngAggregateAddress{}
		}
	}
	if v1, ok := d.GetOk("default_information_originate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_originate", sv)
				diags = append(diags, e)
			}
			obj.DefaultInformationOriginate = &v2
		}
	}
	if v1, ok := d.GetOk("default_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultMetric = &tmp
		}
	}
	if v, ok := d.GetOk("distance"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("distance", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngDistance(d, v, "distance", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Distance = t
		}
	} else if d.HasChange("distance") {
		old, new := d.GetChange("distance")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Distance = &[]models.RouterRipngDistance{}
		}
	}
	if v, ok := d.GetOk("distribute_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("distribute_list", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngDistributeList(d, v, "distribute_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DistributeList = t
		}
	} else if d.HasChange("distribute_list") {
		old, new := d.GetChange("distribute_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DistributeList = &[]models.RouterRipngDistributeList{}
		}
	}
	if v1, ok := d.GetOk("garbage_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("garbage_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GarbageTimer = &tmp
		}
	}
	if v, ok := d.GetOk("interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Interface = t
		}
	} else if d.HasChange("interface") {
		old, new := d.GetChange("interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Interface = &[]models.RouterRipngInterface{}
		}
	}
	if v1, ok := d.GetOk("max_out_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_out_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxOutMetric = &tmp
		}
	}
	if v, ok := d.GetOk("neighbor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.RouterRipngNeighbor{}
		}
	}
	if v, ok := d.GetOk("network"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("network", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngNetwork(d, v, "network", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Network = t
		}
	} else if d.HasChange("network") {
		old, new := d.GetChange("network")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Network = &[]models.RouterRipngNetwork{}
		}
	}
	if v, ok := d.GetOk("offset_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("offset_list", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngOffsetList(d, v, "offset_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OffsetList = t
		}
	} else if d.HasChange("offset_list") {
		old, new := d.GetChange("offset_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OffsetList = &[]models.RouterRipngOffsetList{}
		}
	}
	if v, ok := d.GetOk("passive_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("passive_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngPassiveInterface(d, v, "passive_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PassiveInterface = t
		}
	} else if d.HasChange("passive_interface") {
		old, new := d.GetChange("passive_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PassiveInterface = &[]models.RouterRipngPassiveInterface{}
		}
	}
	if v, ok := d.GetOk("redistribute"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipngRedistribute(d, v, "redistribute", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute = t
		}
	} else if d.HasChange("redistribute") {
		old, new := d.GetChange("redistribute")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute = &[]models.RouterRipngRedistribute{}
		}
	}
	if v1, ok := d.GetOk("timeout_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timeout_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TimeoutTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("update_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateTimer = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterRipng(d *schema.ResourceData, sv string) (*models.RouterRipng, diag.Diagnostics) {
	obj := models.RouterRipng{}
	diags := diag.Diagnostics{}

	obj.AggregateAddress = &[]models.RouterRipngAggregateAddress{}
	obj.Distance = &[]models.RouterRipngDistance{}
	obj.DistributeList = &[]models.RouterRipngDistributeList{}
	obj.Interface = &[]models.RouterRipngInterface{}
	obj.Neighbor = &[]models.RouterRipngNeighbor{}
	obj.Network = &[]models.RouterRipngNetwork{}
	obj.OffsetList = &[]models.RouterRipngOffsetList{}
	obj.PassiveInterface = &[]models.RouterRipngPassiveInterface{}
	obj.Redistribute = &[]models.RouterRipngRedistribute{}

	return &obj, diags
}
