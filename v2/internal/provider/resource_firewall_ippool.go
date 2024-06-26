// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 IP pools.

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

func resourceFirewallIppool() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 IP pools.",

		CreateContext: resourceFirewallIppoolCreate,
		ReadContext:   resourceFirewallIppoolRead,
		UpdateContext: resourceFirewallIppoolUpdate,
		DeleteContext: resourceFirewallIppoolDelete,

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
			"add_nat64_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable adding NAT64 route.",
				Optional:    true,
				Computed:    true,
			},
			"arp_intf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Select an interface from available options that will reply to ARP requests. (If blank, any is selected).",
				Optional:    true,
				Computed:    true,
			},
			"arp_reply": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"associated_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Associated interface name.",
				Optional:    true,
				Computed:    true,
			},
			"block_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 4096),

				Description: "Number of addresses in a block (64 - 4096, default = 128).",
				Optional:    true,
				Computed:    true,
			},
			"cgn_block_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 4096),

				Description: "Number of ports in a block(64 to 4096 in unit of 64, default = 128).",
				Optional:    true,
				Computed:    true,
			},
			"cgn_client_endip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Final client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"cgn_client_ipv6shift": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 127),

				Description: "IPv6 shift for fixed-allocation.(default 0)",
				Optional:    true,
				Computed:    true,
			},
			"cgn_client_startip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "First client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"cgn_fixedalloc": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable fixed-allocation mode.",
				Optional:    true,
				Computed:    true,
			},
			"cgn_overload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable overload mode.",
				Optional:    true,
				Computed:    true,
			},
			"cgn_port_end": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 65535),

				Description: "Ending public port can be allocated. ",
				Optional:    true,
				Computed:    true,
			},
			"cgn_port_start": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 65535),

				Description: "Starting public port can be allocated. ",
				Optional:    true,
				Computed:    true,
			},
			"cgn_spa": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable single port allocation mode.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"endip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"endport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5117, 65533),

				Description: "Final port number (inclusive) in the range for the address pool (Default: 65533).",
				Optional:    true,
				Computed:    true,
			},
			"exclude_ip": {
				Type:        schema.TypeList,
				Description: "Exclude IPs x.x.x.x.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Exclude IPs",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IP pool name.",
				ForceNew:    true,
				Required:    true,
			},
			"nat64": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable NAT64.",
				Optional:    true,
				Computed:    true,
			},
			"num_blocks_per_user": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 128),

				Description: "Number of addresses blocks that can be used by a user (1 to 128, default = 8).",
				Optional:    true,
				Computed:    true,
			},
			"pba_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 86400),

				Description: "Port block allocation timeout (seconds).",
				Optional:    true,
				Computed:    true,
			},
			"permit_any_host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable full cone NAT.",
				Optional:    true,
				Computed:    true,
			},
			"port_per_user": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(32, 60417),

				Description: "Number of port for each user (32 - 60416, default = 0, which is auto).",
				Optional:    true,
				Computed:    true,
			},
			"source_endip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"source_startip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "First IPv4 address (inclusive) in the range of the source addresses to be translated (format = xxx.xxx.xxx.xxx, default = 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"startip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Optional:    true,
				Computed:    true,
			},
			"startport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5117, 65533),

				Description: "First port number (inclusive) in the range for the address pool (Default: 5117).",
				Optional:    true,
				Computed:    true,
			},
			"subnet_broadcast_in_ippool": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"overload", "one-to-one", "fixed-port-range", "port-block-allocation", "cgn-resource-allocation"}, false),

				Description: "IP pool type (overload, one-to-one, fixed port range, or port block allocation).",
				Optional:    true,
				Computed:    true,
			},
			"utilization_alarm_clear": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(40, 100),

				Description: "Pool utilization alarm clear threshold (40-100).",
				Optional:    true,
				Computed:    true,
			},
			"utilization_alarm_raise": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 100),

				Description: "Pool utilization alarm raise threshold (50-100).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallIppoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallIppool resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallIppool(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallIppool(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallIppool")
	}

	return resourceFirewallIppoolRead(ctx, d, meta)
}

func resourceFirewallIppoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallIppool(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallIppool(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallIppool resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallIppool")
	}

	return resourceFirewallIppoolRead(ctx, d, meta)
}

func resourceFirewallIppoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallIppool(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallIppool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIppoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallIppool(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallIppool resource: %v", err)
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

	diags := refreshObjectFirewallIppool(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallIppoolExcludeIp(d *schema.ResourceData, v *[]models.FirewallIppoolExcludeIp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Subnet; tmp != nil {
				v["subnet"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "subnet")
	}

	return flat
}

func refreshObjectFirewallIppool(d *schema.ResourceData, o *models.FirewallIppool, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddNat64Route != nil {
		v := *o.AddNat64Route

		if err = d.Set("add_nat64_route", v); err != nil {
			return diag.Errorf("error reading add_nat64_route: %v", err)
		}
	}

	if o.ArpIntf != nil {
		v := *o.ArpIntf

		if err = d.Set("arp_intf", v); err != nil {
			return diag.Errorf("error reading arp_intf: %v", err)
		}
	}

	if o.ArpReply != nil {
		v := *o.ArpReply

		if err = d.Set("arp_reply", v); err != nil {
			return diag.Errorf("error reading arp_reply: %v", err)
		}
	}

	if o.AssociatedInterface != nil {
		v := *o.AssociatedInterface

		if err = d.Set("associated_interface", v); err != nil {
			return diag.Errorf("error reading associated_interface: %v", err)
		}
	}

	if o.BlockSize != nil {
		v := *o.BlockSize

		if err = d.Set("block_size", v); err != nil {
			return diag.Errorf("error reading block_size: %v", err)
		}
	}

	if o.CgnBlockSize != nil {
		v := *o.CgnBlockSize

		if err = d.Set("cgn_block_size", v); err != nil {
			return diag.Errorf("error reading cgn_block_size: %v", err)
		}
	}

	if o.CgnClientEndip != nil {
		v := *o.CgnClientEndip

		if err = d.Set("cgn_client_endip", v); err != nil {
			return diag.Errorf("error reading cgn_client_endip: %v", err)
		}
	}

	if o.CgnClientIpv6shift != nil {
		v := *o.CgnClientIpv6shift

		if err = d.Set("cgn_client_ipv6shift", v); err != nil {
			return diag.Errorf("error reading cgn_client_ipv6shift: %v", err)
		}
	}

	if o.CgnClientStartip != nil {
		v := *o.CgnClientStartip

		if err = d.Set("cgn_client_startip", v); err != nil {
			return diag.Errorf("error reading cgn_client_startip: %v", err)
		}
	}

	if o.CgnFixedalloc != nil {
		v := *o.CgnFixedalloc

		if err = d.Set("cgn_fixedalloc", v); err != nil {
			return diag.Errorf("error reading cgn_fixedalloc: %v", err)
		}
	}

	if o.CgnOverload != nil {
		v := *o.CgnOverload

		if err = d.Set("cgn_overload", v); err != nil {
			return diag.Errorf("error reading cgn_overload: %v", err)
		}
	}

	if o.CgnPortEnd != nil {
		v := *o.CgnPortEnd

		if err = d.Set("cgn_port_end", v); err != nil {
			return diag.Errorf("error reading cgn_port_end: %v", err)
		}
	}

	if o.CgnPortStart != nil {
		v := *o.CgnPortStart

		if err = d.Set("cgn_port_start", v); err != nil {
			return diag.Errorf("error reading cgn_port_start: %v", err)
		}
	}

	if o.CgnSpa != nil {
		v := *o.CgnSpa

		if err = d.Set("cgn_spa", v); err != nil {
			return diag.Errorf("error reading cgn_spa: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Endip != nil {
		v := *o.Endip

		if err = d.Set("endip", v); err != nil {
			return diag.Errorf("error reading endip: %v", err)
		}
	}

	if o.Endport != nil {
		v := *o.Endport

		if err = d.Set("endport", v); err != nil {
			return diag.Errorf("error reading endport: %v", err)
		}
	}

	if o.ExcludeIp != nil {
		if err = d.Set("exclude_ip", flattenFirewallIppoolExcludeIp(d, o.ExcludeIp, "exclude_ip", sort)); err != nil {
			return diag.Errorf("error reading exclude_ip: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Nat64 != nil {
		v := *o.Nat64

		if err = d.Set("nat64", v); err != nil {
			return diag.Errorf("error reading nat64: %v", err)
		}
	}

	if o.NumBlocksPerUser != nil {
		v := *o.NumBlocksPerUser

		if err = d.Set("num_blocks_per_user", v); err != nil {
			return diag.Errorf("error reading num_blocks_per_user: %v", err)
		}
	}

	if o.PbaTimeout != nil {
		v := *o.PbaTimeout

		if err = d.Set("pba_timeout", v); err != nil {
			return diag.Errorf("error reading pba_timeout: %v", err)
		}
	}

	if o.PermitAnyHost != nil {
		v := *o.PermitAnyHost

		if err = d.Set("permit_any_host", v); err != nil {
			return diag.Errorf("error reading permit_any_host: %v", err)
		}
	}

	if o.PortPerUser != nil {
		v := *o.PortPerUser

		if err = d.Set("port_per_user", v); err != nil {
			return diag.Errorf("error reading port_per_user: %v", err)
		}
	}

	if o.SourceEndip != nil {
		v := *o.SourceEndip

		if err = d.Set("source_endip", v); err != nil {
			return diag.Errorf("error reading source_endip: %v", err)
		}
	}

	if o.SourceStartip != nil {
		v := *o.SourceStartip

		if err = d.Set("source_startip", v); err != nil {
			return diag.Errorf("error reading source_startip: %v", err)
		}
	}

	if o.Startip != nil {
		v := *o.Startip

		if err = d.Set("startip", v); err != nil {
			return diag.Errorf("error reading startip: %v", err)
		}
	}

	if o.Startport != nil {
		v := *o.Startport

		if err = d.Set("startport", v); err != nil {
			return diag.Errorf("error reading startport: %v", err)
		}
	}

	if o.SubnetBroadcastInIppool != nil {
		v := *o.SubnetBroadcastInIppool

		if err = d.Set("subnet_broadcast_in_ippool", v); err != nil {
			return diag.Errorf("error reading subnet_broadcast_in_ippool: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.UtilizationAlarmClear != nil {
		v := *o.UtilizationAlarmClear

		if err = d.Set("utilization_alarm_clear", v); err != nil {
			return diag.Errorf("error reading utilization_alarm_clear: %v", err)
		}
	}

	if o.UtilizationAlarmRaise != nil {
		v := *o.UtilizationAlarmRaise

		if err = d.Set("utilization_alarm_raise", v); err != nil {
			return diag.Errorf("error reading utilization_alarm_raise: %v", err)
		}
	}

	return nil
}

func expandFirewallIppoolExcludeIp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallIppoolExcludeIp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallIppoolExcludeIp

	for i := range l {
		tmp := models.FirewallIppoolExcludeIp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Subnet = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallIppool(d *schema.ResourceData, sv string) (*models.FirewallIppool, diag.Diagnostics) {
	obj := models.FirewallIppool{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("add_nat64_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("add_nat64_route", sv)
				diags = append(diags, e)
			}
			obj.AddNat64Route = &v2
		}
	}
	if v1, ok := d.GetOk("arp_intf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arp_intf", sv)
				diags = append(diags, e)
			}
			obj.ArpIntf = &v2
		}
	}
	if v1, ok := d.GetOk("arp_reply"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arp_reply", sv)
				diags = append(diags, e)
			}
			obj.ArpReply = &v2
		}
	}
	if v1, ok := d.GetOk("associated_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("associated_interface", sv)
				diags = append(diags, e)
			}
			obj.AssociatedInterface = &v2
		}
	}
	if v1, ok := d.GetOk("block_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("block_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BlockSize = &tmp
		}
	}
	if v1, ok := d.GetOk("cgn_block_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_block_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CgnBlockSize = &tmp
		}
	}
	if v1, ok := d.GetOk("cgn_client_endip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_client_endip", sv)
				diags = append(diags, e)
			}
			obj.CgnClientEndip = &v2
		}
	}
	if v1, ok := d.GetOk("cgn_client_ipv6shift"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_client_ipv6shift", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CgnClientIpv6shift = &tmp
		}
	}
	if v1, ok := d.GetOk("cgn_client_startip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_client_startip", sv)
				diags = append(diags, e)
			}
			obj.CgnClientStartip = &v2
		}
	}
	if v1, ok := d.GetOk("cgn_fixedalloc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_fixedalloc", sv)
				diags = append(diags, e)
			}
			obj.CgnFixedalloc = &v2
		}
	}
	if v1, ok := d.GetOk("cgn_overload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_overload", sv)
				diags = append(diags, e)
			}
			obj.CgnOverload = &v2
		}
	}
	if v1, ok := d.GetOk("cgn_port_end"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_port_end", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CgnPortEnd = &tmp
		}
	}
	if v1, ok := d.GetOk("cgn_port_start"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_port_start", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CgnPortStart = &tmp
		}
	}
	if v1, ok := d.GetOk("cgn_spa"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("cgn_spa", sv)
				diags = append(diags, e)
			}
			obj.CgnSpa = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("endip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("endip", sv)
				diags = append(diags, e)
			}
			obj.Endip = &v2
		}
	}
	if v1, ok := d.GetOk("endport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("endport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Endport = &tmp
		}
	}
	if v, ok := d.GetOk("exclude_ip"); ok {
		if !utils.CheckVer(sv, "v7.2.8", "") {
			e := utils.AttributeVersionWarning("exclude_ip", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallIppoolExcludeIp(d, v, "exclude_ip", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExcludeIp = t
		}
	} else if d.HasChange("exclude_ip") {
		old, new := d.GetChange("exclude_ip")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExcludeIp = &[]models.FirewallIppoolExcludeIp{}
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
	if v1, ok := d.GetOk("nat64"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("nat64", sv)
				diags = append(diags, e)
			}
			obj.Nat64 = &v2
		}
	}
	if v1, ok := d.GetOk("num_blocks_per_user"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("num_blocks_per_user", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NumBlocksPerUser = &tmp
		}
	}
	if v1, ok := d.GetOk("pba_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pba_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PbaTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("permit_any_host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_any_host", sv)
				diags = append(diags, e)
			}
			obj.PermitAnyHost = &v2
		}
	}
	if v1, ok := d.GetOk("port_per_user"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("port_per_user", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PortPerUser = &tmp
		}
	}
	if v1, ok := d.GetOk("source_endip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_endip", sv)
				diags = append(diags, e)
			}
			obj.SourceEndip = &v2
		}
	}
	if v1, ok := d.GetOk("source_startip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_startip", sv)
				diags = append(diags, e)
			}
			obj.SourceStartip = &v2
		}
	}
	if v1, ok := d.GetOk("startip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("startip", sv)
				diags = append(diags, e)
			}
			obj.Startip = &v2
		}
	}
	if v1, ok := d.GetOk("startport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("startport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Startport = &tmp
		}
	}
	if v1, ok := d.GetOk("subnet_broadcast_in_ippool"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("subnet_broadcast_in_ippool", sv)
				diags = append(diags, e)
			}
			obj.SubnetBroadcastInIppool = &v2
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
	if v1, ok := d.GetOk("utilization_alarm_clear"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("utilization_alarm_clear", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UtilizationAlarmClear = &tmp
		}
	}
	if v1, ok := d.GetOk("utilization_alarm_raise"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("utilization_alarm_raise", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UtilizationAlarmRaise = &tmp
		}
	}
	return &obj, diags
}
