// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure RIP.

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

func resourceRouterRip() *schema.Resource {
	return &schema.Resource{
		Description: "Configure RIP.",

		CreateContext: resourceRouterRipCreate,
		ReadContext:   resourceRouterRipRead,
		UpdateContext: resourceRouterRipUpdate,
		DeleteContext: resourceRouterRipDelete,

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
						"access_list": {
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
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

							Description: "Distance prefix.",
							Optional:    true,
							Computed:    true,
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

				Description: "Garbage timer in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "RIP interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_keychain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Authentication key-chain name.",
							Optional:    true,
							Computed:    true,
						},
						"auth_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "md5"}, false),

							Description: "Authentication mode.",
							Optional:    true,
							Computed:    true,
						},
						"auth_string": {
							Type: schema.TypeString,

							Description: "Authentication string/password.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
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
						"receive_version": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Receive version.",
							Optional:         true,
							Computed:         true,
						},
						"send_version": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Send version.",
							Optional:         true,
							Computed:         true,
						},
						"send_version2_broadcast": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable broadcast version 1 compatible packets.",
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
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address.",
							Optional:    true,
							Computed:    true,
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
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Network prefix.",
							Optional:    true,
							Computed:    true,
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
						"access_list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Access list name.",
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
			"recv_buffer_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(8129, 2147483647),

				Description: "Receiving buffer size.",
				Optional:    true,
				Computed:    true,
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

				Description: "Timeout timer in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"update_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483647),

				Description: "Update timer in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

				Description: "RIP version.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterRip(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterRip(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRip")
	}

	return resourceRouterRipRead(ctx, d, meta)
}

func resourceRouterRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterRip(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterRip(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterRip resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRip")
	}

	return resourceRouterRipRead(ctx, d, meta)
}

func resourceRouterRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterRip(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterRip(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterRip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterRip(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterRip resource: %v", err)
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

	diags := refreshObjectRouterRip(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterRipDistance(d *schema.ResourceData, v *[]models.RouterRipDistance, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AccessList; tmp != nil {
				v["access_list"] = *tmp
			}

			if tmp := cfg.Distance; tmp != nil {
				v["distance"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.prefix", prefix, i), *tmp); tmp != nil {
					v["prefix"] = *tmp
				}
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipDistributeList(d *schema.ResourceData, v *[]models.RouterRipDistributeList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenRouterRipInterface(d *schema.ResourceData, v *[]models.RouterRipInterface, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AuthKeychain; tmp != nil {
				v["auth_keychain"] = *tmp
			}

			if tmp := cfg.AuthMode; tmp != nil {
				v["auth_mode"] = *tmp
			}

			if tmp := cfg.AuthString; tmp != nil {
				v["auth_string"] = *tmp
			}

			if tmp := cfg.Flags; tmp != nil {
				v["flags"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.ReceiveVersion; tmp != nil {
				v["receive_version"] = *tmp
			}

			if tmp := cfg.SendVersion; tmp != nil {
				v["send_version"] = *tmp
			}

			if tmp := cfg.SendVersion2Broadcast; tmp != nil {
				v["send_version2_broadcast"] = *tmp
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

func flattenRouterRipNeighbor(d *schema.ResourceData, v *[]models.RouterRipNeighbor, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipNetwork(d *schema.ResourceData, v *[]models.RouterRipNetwork, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.prefix", prefix, i), *tmp); tmp != nil {
					v["prefix"] = *tmp
				}
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRipOffsetList(d *schema.ResourceData, v *[]models.RouterRipOffsetList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AccessList; tmp != nil {
				v["access_list"] = *tmp
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

func flattenRouterRipPassiveInterface(d *schema.ResourceData, v *[]models.RouterRipPassiveInterface, prefix string, sort bool) interface{} {
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

func flattenRouterRipRedistribute(d *schema.ResourceData, v *[]models.RouterRipRedistribute, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func refreshObjectRouterRip(d *schema.ResourceData, o *models.RouterRip, sv string, sort bool) diag.Diagnostics {
	var err error

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
		if err = d.Set("distance", flattenRouterRipDistance(d, o.Distance, "distance", sort)); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.DistributeList != nil {
		if err = d.Set("distribute_list", flattenRouterRipDistributeList(d, o.DistributeList, "distribute_list", sort)); err != nil {
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
		if err = d.Set("interface", flattenRouterRipInterface(d, o.Interface, "interface", sort)); err != nil {
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
		if err = d.Set("neighbor", flattenRouterRipNeighbor(d, o.Neighbor, "neighbor", sort)); err != nil {
			return diag.Errorf("error reading neighbor: %v", err)
		}
	}

	if o.Network != nil {
		if err = d.Set("network", flattenRouterRipNetwork(d, o.Network, "network", sort)); err != nil {
			return diag.Errorf("error reading network: %v", err)
		}
	}

	if o.OffsetList != nil {
		if err = d.Set("offset_list", flattenRouterRipOffsetList(d, o.OffsetList, "offset_list", sort)); err != nil {
			return diag.Errorf("error reading offset_list: %v", err)
		}
	}

	if o.PassiveInterface != nil {
		if err = d.Set("passive_interface", flattenRouterRipPassiveInterface(d, o.PassiveInterface, "passive_interface", sort)); err != nil {
			return diag.Errorf("error reading passive_interface: %v", err)
		}
	}

	if o.RecvBufferSize != nil {
		v := *o.RecvBufferSize

		if err = d.Set("recv_buffer_size", v); err != nil {
			return diag.Errorf("error reading recv_buffer_size: %v", err)
		}
	}

	if o.Redistribute != nil {
		if err = d.Set("redistribute", flattenRouterRipRedistribute(d, o.Redistribute, "redistribute", sort)); err != nil {
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

	if o.Version != nil {
		v := *o.Version

		if err = d.Set("version", v); err != nil {
			return diag.Errorf("error reading version: %v", err)
		}
	}

	return nil
}

func expandRouterRipDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipDistance, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipDistance

	for i := range l {
		tmp := models.RouterRipDistance{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessList = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Distance = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
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

func expandRouterRipDistributeList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipDistributeList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipDistributeList

	for i := range l {
		tmp := models.RouterRipDistributeList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
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

func expandRouterRipInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipInterface

	for i := range l {
		tmp := models.RouterRipInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_keychain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthKeychain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.flags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Flags = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.receive_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ReceiveVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_version2_broadcast", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendVersion2Broadcast = &v2
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

func expandRouterRipNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipNeighbor

	for i := range l {
		tmp := models.RouterRipNeighbor{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRipNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipNetwork, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipNetwork

	for i := range l {
		tmp := models.RouterRipNetwork{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
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

func expandRouterRipOffsetList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipOffsetList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipOffsetList

	for i := range l {
		tmp := models.RouterRipOffsetList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessList = &v2
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

		pre_append = fmt.Sprintf("%s.%d.offset", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Offset = &v3
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

func expandRouterRipPassiveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipPassiveInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipPassiveInterface

	for i := range l {
		tmp := models.RouterRipPassiveInterface{}
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

func expandRouterRipRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRipRedistribute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRipRedistribute

	for i := range l {
		tmp := models.RouterRipRedistribute{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Metric = &v3
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

func getObjectRouterRip(d *schema.ResourceData, sv string) (*models.RouterRip, diag.Diagnostics) {
	obj := models.RouterRip{}
	diags := diag.Diagnostics{}

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
		t, err := expandRouterRipDistance(d, v, "distance", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Distance = t
		}
	} else if d.HasChange("distance") {
		old, new := d.GetChange("distance")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Distance = &[]models.RouterRipDistance{}
		}
	}
	if v, ok := d.GetOk("distribute_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("distribute_list", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipDistributeList(d, v, "distribute_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DistributeList = t
		}
	} else if d.HasChange("distribute_list") {
		old, new := d.GetChange("distribute_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DistributeList = &[]models.RouterRipDistributeList{}
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
		t, err := expandRouterRipInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Interface = t
		}
	} else if d.HasChange("interface") {
		old, new := d.GetChange("interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Interface = &[]models.RouterRipInterface{}
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
		t, err := expandRouterRipNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.RouterRipNeighbor{}
		}
	}
	if v, ok := d.GetOk("network"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("network", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipNetwork(d, v, "network", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Network = t
		}
	} else if d.HasChange("network") {
		old, new := d.GetChange("network")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Network = &[]models.RouterRipNetwork{}
		}
	}
	if v, ok := d.GetOk("offset_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("offset_list", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipOffsetList(d, v, "offset_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OffsetList = t
		}
	} else if d.HasChange("offset_list") {
		old, new := d.GetChange("offset_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OffsetList = &[]models.RouterRipOffsetList{}
		}
	}
	if v, ok := d.GetOk("passive_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("passive_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipPassiveInterface(d, v, "passive_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PassiveInterface = t
		}
	} else if d.HasChange("passive_interface") {
		old, new := d.GetChange("passive_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PassiveInterface = &[]models.RouterRipPassiveInterface{}
		}
	}
	if v1, ok := d.GetOk("recv_buffer_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("recv_buffer_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RecvBufferSize = &tmp
		}
	}
	if v, ok := d.GetOk("redistribute"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRipRedistribute(d, v, "redistribute", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute = t
		}
	} else if d.HasChange("redistribute") {
		old, new := d.GetChange("redistribute")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute = &[]models.RouterRipRedistribute{}
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
	if v1, ok := d.GetOk("version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("version", sv)
				diags = append(diags, e)
			}
			obj.Version = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterRip(d *schema.ResourceData, sv string) (*models.RouterRip, diag.Diagnostics) {
	obj := models.RouterRip{}
	diags := diag.Diagnostics{}

	obj.Distance = &[]models.RouterRipDistance{}
	obj.DistributeList = &[]models.RouterRipDistributeList{}
	obj.Interface = &[]models.RouterRipInterface{}
	obj.Neighbor = &[]models.RouterRipNeighbor{}
	obj.Network = &[]models.RouterRipNetwork{}
	obj.OffsetList = &[]models.RouterRipOffsetList{}
	obj.PassiveInterface = &[]models.RouterRipPassiveInterface{}
	obj.Redistribute = &[]models.RouterRipRedistribute{}

	return &obj, diags
}
