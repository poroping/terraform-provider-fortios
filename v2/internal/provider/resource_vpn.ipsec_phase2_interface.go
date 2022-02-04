// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VPN autokey tunnel.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceVpnIpsecPhase2Interface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VPN autokey tunnel.",

		CreateContext: resourceVpnIpsecPhase2InterfaceCreate,
		ReadContext:   resourceVpnIpsecPhase2InterfaceRead,
		UpdateContext: resourceVpnIpsecPhase2InterfaceUpdate,
		DeleteContext: resourceVpnIpsecPhase2InterfaceDelete,

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
			"add_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"phase1", "enable", "disable"}, false),

				Description: "Enable/disable automatic route addition.",
				Optional:    true,
				Computed:    true,
			},
			"auto_discovery_forwarder": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"phase1", "enable", "disable"}, false),

				Description: "Enable/disable forwarding short-cut messages.",
				Optional:    true,
				Computed:    true,
			},
			"auto_discovery_sender": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"phase1", "enable", "disable"}, false),

				Description: "Enable/disable sending short-cut messages.",
				Optional:    true,
				Computed:    true,
			},
			"auto_negotiate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPsec SA auto-negotiation.",
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
			"dhcp_ipsec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DHCP-IPsec.",
				Optional:    true,
				Computed:    true,
			},
			"dhgrp": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Phase2 DH group.",
				Optional:         true,
				Computed:         true,
			},
			"diffserv": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable applying DSCP value to the IPsec tunnel outer IP header.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode": {
				Type: schema.TypeString,

				Description: "DSCP value to be applied to the IPsec tunnel outer IP header.",
				Optional:    true,
				Computed:    true,
			},
			"dst_addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"subnet", "range", "ip", "name", "subnet6", "range6", "ip6", "name6"}, false),

				Description: "Remote proxy ID type.",
				Optional:    true,
				Computed:    true,
			},
			"dst_end_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Remote proxy ID IPv4 end.",
				Optional:    true,
				Computed:    true,
			},
			"dst_end_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Remote proxy ID IPv6 end.",
				Optional:         true,
				Computed:         true,
			},
			"dst_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Remote proxy ID name.",
				Optional:    true,
				Computed:    true,
			},
			"dst_name6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Remote proxy ID name.",
				Optional:    true,
				Computed:    true,
			},
			"dst_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Quick mode destination port (1 - 65535 or 0 for all).",
				Optional:    true,
				Computed:    true,
			},
			"dst_start_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Remote proxy ID IPv4 start.",
				Optional:    true,
				Computed:    true,
			},
			"dst_start_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Remote proxy ID IPv6 start.",
				Optional:         true,
				Computed:         true,
			},
			"dst_subnet": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Remote proxy ID IPv4 subnet.",
				Optional:    true,
				Computed:    true,
			},
			"dst_subnet6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Remote proxy ID IPv6 subnet.",
				Optional:         true,
				Computed:         true,
			},
			"encapsulation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tunnel-mode", "transport-mode"}, false),

				Description: "ESP encapsulation mode.",
				Optional:    true,
				Computed:    true,
			},
			"initiator_ts_narrow": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable traffic selector narrowing for IKEv2 initiator.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_df": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable setting and resetting of IPv4 'Don't Fragment' bit.",
				Optional:    true,
				Computed:    true,
			},
			"keepalive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable keep alive.",
				Optional:    true,
				Computed:    true,
			},
			"keylife_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"seconds", "kbs", "both"}, false),

				Description: "Keylife type.",
				Optional:    true,
				Computed:    true,
			},
			"keylifekbs": {
				Type: schema.TypeInt,

				Description: "Phase2 key life in number of kilobytes of traffic (5120 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"keylifeseconds": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 172800),

				Description: "Phase2 key life in time in seconds (120 - 172800).",
				Optional:    true,
				Computed:    true,
			},
			"l2tp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable L2TP over IPsec.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IPsec tunnel name.",
				ForceNew:    true,
				Required:    true,
			},
			"pfs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PFS feature.",
				Optional:    true,
				Computed:    true,
			},
			"phase1name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Phase 1 determines the options required for phase 2.",
				Optional:    true,
				Computed:    true,
			},
			"proposal": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Phase2 proposal.",
				Optional:         true,
				Computed:         true,
			},
			"protocol": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Quick mode protocol selector (1 - 255 or 0 for all).",
				Optional:    true,
				Computed:    true,
			},
			"replay": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable replay detection.",
				Optional:    true,
				Computed:    true,
			},
			"route_overlap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"use-old", "use-new", "allow"}, false),

				Description: "Action for overlapping routes.",
				Optional:    true,
				Computed:    true,
			},
			"single_source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable single source IP restriction.",
				Optional:    true,
				Computed:    true,
			},
			"src_addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"subnet", "range", "ip", "name", "subnet6", "range6", "ip6", "name6"}, false),

				Description: "Local proxy ID type.",
				Optional:    true,
				Computed:    true,
			},
			"src_end_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Local proxy ID end.",
				Optional:    true,
				Computed:    true,
			},
			"src_end_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Local proxy ID IPv6 end.",
				Optional:         true,
				Computed:         true,
			},
			"src_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Local proxy ID name.",
				Optional:    true,
				Computed:    true,
			},
			"src_name6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Local proxy ID name.",
				Optional:    true,
				Computed:    true,
			},
			"src_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Quick mode source port (1 - 65535 or 0 for all).",
				Optional:    true,
				Computed:    true,
			},
			"src_start_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Local proxy ID start.",
				Optional:    true,
				Computed:    true,
			},
			"src_start_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Local proxy ID IPv6 start.",
				Optional:         true,
				Computed:         true,
			},
			"src_subnet": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Local proxy ID subnet.",
				Optional:    true,
				Computed:    true,
			},
			"src_subnet6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Local proxy ID IPv6 subnet.",
				Optional:         true,
				Computed:         true,
			},
		},
	}
}

func resourceVpnIpsecPhase2InterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnIpsecPhase2Interface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnIpsecPhase2Interface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnIpsecPhase2Interface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecPhase2Interface")
	}

	return resourceVpnIpsecPhase2InterfaceRead(ctx, d, meta)
}

func resourceVpnIpsecPhase2InterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnIpsecPhase2Interface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnIpsecPhase2Interface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnIpsecPhase2Interface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecPhase2Interface")
	}

	return resourceVpnIpsecPhase2InterfaceRead(ctx, d, meta)
}

func resourceVpnIpsecPhase2InterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnIpsecPhase2Interface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnIpsecPhase2Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase2InterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecPhase2Interface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecPhase2Interface resource: %v", err)
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

	diags := refreshObjectVpnIpsecPhase2Interface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnIpsecPhase2Interface(d *schema.ResourceData, o *models.VpnIpsecPhase2Interface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddRoute != nil {
		v := *o.AddRoute

		if err = d.Set("add_route", v); err != nil {
			return diag.Errorf("error reading add_route: %v", err)
		}
	}

	if o.AutoDiscoveryForwarder != nil {
		v := *o.AutoDiscoveryForwarder

		if err = d.Set("auto_discovery_forwarder", v); err != nil {
			return diag.Errorf("error reading auto_discovery_forwarder: %v", err)
		}
	}

	if o.AutoDiscoverySender != nil {
		v := *o.AutoDiscoverySender

		if err = d.Set("auto_discovery_sender", v); err != nil {
			return diag.Errorf("error reading auto_discovery_sender: %v", err)
		}
	}

	if o.AutoNegotiate != nil {
		v := *o.AutoNegotiate

		if err = d.Set("auto_negotiate", v); err != nil {
			return diag.Errorf("error reading auto_negotiate: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.DhcpIpsec != nil {
		v := *o.DhcpIpsec

		if err = d.Set("dhcp_ipsec", v); err != nil {
			return diag.Errorf("error reading dhcp_ipsec: %v", err)
		}
	}

	if o.Dhgrp != nil {
		v := *o.Dhgrp

		if err = d.Set("dhgrp", v); err != nil {
			return diag.Errorf("error reading dhgrp: %v", err)
		}
	}

	if o.Diffserv != nil {
		v := *o.Diffserv

		if err = d.Set("diffserv", v); err != nil {
			return diag.Errorf("error reading diffserv: %v", err)
		}
	}

	if o.Diffservcode != nil {
		v := *o.Diffservcode

		if err = d.Set("diffservcode", v); err != nil {
			return diag.Errorf("error reading diffservcode: %v", err)
		}
	}

	if o.DstAddrType != nil {
		v := *o.DstAddrType

		if err = d.Set("dst_addr_type", v); err != nil {
			return diag.Errorf("error reading dst_addr_type: %v", err)
		}
	}

	if o.DstEndIp != nil {
		v := *o.DstEndIp

		if err = d.Set("dst_end_ip", v); err != nil {
			return diag.Errorf("error reading dst_end_ip: %v", err)
		}
	}

	if o.DstEndIp6 != nil {
		v := *o.DstEndIp6

		if err = d.Set("dst_end_ip6", v); err != nil {
			return diag.Errorf("error reading dst_end_ip6: %v", err)
		}
	}

	if o.DstName != nil {
		v := *o.DstName

		if err = d.Set("dst_name", v); err != nil {
			return diag.Errorf("error reading dst_name: %v", err)
		}
	}

	if o.DstName6 != nil {
		v := *o.DstName6

		if err = d.Set("dst_name6", v); err != nil {
			return diag.Errorf("error reading dst_name6: %v", err)
		}
	}

	if o.DstPort != nil {
		v := *o.DstPort

		if err = d.Set("dst_port", v); err != nil {
			return diag.Errorf("error reading dst_port: %v", err)
		}
	}

	if o.DstStartIp != nil {
		v := *o.DstStartIp

		if err = d.Set("dst_start_ip", v); err != nil {
			return diag.Errorf("error reading dst_start_ip: %v", err)
		}
	}

	if o.DstStartIp6 != nil {
		v := *o.DstStartIp6

		if err = d.Set("dst_start_ip6", v); err != nil {
			return diag.Errorf("error reading dst_start_ip6: %v", err)
		}
	}

	if o.DstSubnet != nil {
		v := *o.DstSubnet
		if current, ok := d.GetOk("dst_subnet"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("dst_subnet", v); err != nil {
			return diag.Errorf("error reading dst_subnet: %v", err)
		}
	}

	if o.DstSubnet6 != nil {
		v := *o.DstSubnet6

		if err = d.Set("dst_subnet6", v); err != nil {
			return diag.Errorf("error reading dst_subnet6: %v", err)
		}
	}

	if o.Encapsulation != nil {
		v := *o.Encapsulation

		if err = d.Set("encapsulation", v); err != nil {
			return diag.Errorf("error reading encapsulation: %v", err)
		}
	}

	if o.InitiatorTsNarrow != nil {
		v := *o.InitiatorTsNarrow

		if err = d.Set("initiator_ts_narrow", v); err != nil {
			return diag.Errorf("error reading initiator_ts_narrow: %v", err)
		}
	}

	if o.Ipv4Df != nil {
		v := *o.Ipv4Df

		if err = d.Set("ipv4_df", v); err != nil {
			return diag.Errorf("error reading ipv4_df: %v", err)
		}
	}

	if o.Keepalive != nil {
		v := *o.Keepalive

		if err = d.Set("keepalive", v); err != nil {
			return diag.Errorf("error reading keepalive: %v", err)
		}
	}

	if o.KeylifeType != nil {
		v := *o.KeylifeType

		if err = d.Set("keylife_type", v); err != nil {
			return diag.Errorf("error reading keylife_type: %v", err)
		}
	}

	if o.Keylifekbs != nil {
		v := *o.Keylifekbs

		if err = d.Set("keylifekbs", v); err != nil {
			return diag.Errorf("error reading keylifekbs: %v", err)
		}
	}

	if o.Keylifeseconds != nil {
		v := *o.Keylifeseconds

		if err = d.Set("keylifeseconds", v); err != nil {
			return diag.Errorf("error reading keylifeseconds: %v", err)
		}
	}

	if o.L2tp != nil {
		v := *o.L2tp

		if err = d.Set("l2tp", v); err != nil {
			return diag.Errorf("error reading l2tp: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Pfs != nil {
		v := *o.Pfs

		if err = d.Set("pfs", v); err != nil {
			return diag.Errorf("error reading pfs: %v", err)
		}
	}

	if o.Phase1name != nil {
		v := *o.Phase1name

		if err = d.Set("phase1name", v); err != nil {
			return diag.Errorf("error reading phase1name: %v", err)
		}
	}

	if o.Proposal != nil {
		v := *o.Proposal

		if err = d.Set("proposal", v); err != nil {
			return diag.Errorf("error reading proposal: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Replay != nil {
		v := *o.Replay

		if err = d.Set("replay", v); err != nil {
			return diag.Errorf("error reading replay: %v", err)
		}
	}

	if o.RouteOverlap != nil {
		v := *o.RouteOverlap

		if err = d.Set("route_overlap", v); err != nil {
			return diag.Errorf("error reading route_overlap: %v", err)
		}
	}

	if o.SingleSource != nil {
		v := *o.SingleSource

		if err = d.Set("single_source", v); err != nil {
			return diag.Errorf("error reading single_source: %v", err)
		}
	}

	if o.SrcAddrType != nil {
		v := *o.SrcAddrType

		if err = d.Set("src_addr_type", v); err != nil {
			return diag.Errorf("error reading src_addr_type: %v", err)
		}
	}

	if o.SrcEndIp != nil {
		v := *o.SrcEndIp

		if err = d.Set("src_end_ip", v); err != nil {
			return diag.Errorf("error reading src_end_ip: %v", err)
		}
	}

	if o.SrcEndIp6 != nil {
		v := *o.SrcEndIp6

		if err = d.Set("src_end_ip6", v); err != nil {
			return diag.Errorf("error reading src_end_ip6: %v", err)
		}
	}

	if o.SrcName != nil {
		v := *o.SrcName

		if err = d.Set("src_name", v); err != nil {
			return diag.Errorf("error reading src_name: %v", err)
		}
	}

	if o.SrcName6 != nil {
		v := *o.SrcName6

		if err = d.Set("src_name6", v); err != nil {
			return diag.Errorf("error reading src_name6: %v", err)
		}
	}

	if o.SrcPort != nil {
		v := *o.SrcPort

		if err = d.Set("src_port", v); err != nil {
			return diag.Errorf("error reading src_port: %v", err)
		}
	}

	if o.SrcStartIp != nil {
		v := *o.SrcStartIp

		if err = d.Set("src_start_ip", v); err != nil {
			return diag.Errorf("error reading src_start_ip: %v", err)
		}
	}

	if o.SrcStartIp6 != nil {
		v := *o.SrcStartIp6

		if err = d.Set("src_start_ip6", v); err != nil {
			return diag.Errorf("error reading src_start_ip6: %v", err)
		}
	}

	if o.SrcSubnet != nil {
		v := *o.SrcSubnet
		if current, ok := d.GetOk("src_subnet"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("src_subnet", v); err != nil {
			return diag.Errorf("error reading src_subnet: %v", err)
		}
	}

	if o.SrcSubnet6 != nil {
		v := *o.SrcSubnet6

		if err = d.Set("src_subnet6", v); err != nil {
			return diag.Errorf("error reading src_subnet6: %v", err)
		}
	}

	return nil
}

func getObjectVpnIpsecPhase2Interface(d *schema.ResourceData, sv string) (*models.VpnIpsecPhase2Interface, diag.Diagnostics) {
	obj := models.VpnIpsecPhase2Interface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("add_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("add_route", sv)
				diags = append(diags, e)
			}
			obj.AddRoute = &v2
		}
	}
	if v1, ok := d.GetOk("auto_discovery_forwarder"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_discovery_forwarder", sv)
				diags = append(diags, e)
			}
			obj.AutoDiscoveryForwarder = &v2
		}
	}
	if v1, ok := d.GetOk("auto_discovery_sender"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_discovery_sender", sv)
				diags = append(diags, e)
			}
			obj.AutoDiscoverySender = &v2
		}
	}
	if v1, ok := d.GetOk("auto_negotiate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_negotiate", sv)
				diags = append(diags, e)
			}
			obj.AutoNegotiate = &v2
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
	if v1, ok := d.GetOk("dhcp_ipsec"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_ipsec", sv)
				diags = append(diags, e)
			}
			obj.DhcpIpsec = &v2
		}
	}
	if v1, ok := d.GetOk("dhgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhgrp", sv)
				diags = append(diags, e)
			}
			obj.Dhgrp = &v2
		}
	}
	if v1, ok := d.GetOk("diffserv"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("diffserv", sv)
				diags = append(diags, e)
			}
			obj.Diffserv = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("diffservcode", sv)
				diags = append(diags, e)
			}
			obj.Diffservcode = &v2
		}
	}
	if v1, ok := d.GetOk("dst_addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_addr_type", sv)
				diags = append(diags, e)
			}
			obj.DstAddrType = &v2
		}
	}
	if v1, ok := d.GetOk("dst_end_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_end_ip", sv)
				diags = append(diags, e)
			}
			obj.DstEndIp = &v2
		}
	}
	if v1, ok := d.GetOk("dst_end_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_end_ip6", sv)
				diags = append(diags, e)
			}
			obj.DstEndIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("dst_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_name", sv)
				diags = append(diags, e)
			}
			obj.DstName = &v2
		}
	}
	if v1, ok := d.GetOk("dst_name6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_name6", sv)
				diags = append(diags, e)
			}
			obj.DstName6 = &v2
		}
	}
	if v1, ok := d.GetOk("dst_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DstPort = &tmp
		}
	}
	if v1, ok := d.GetOk("dst_start_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_start_ip", sv)
				diags = append(diags, e)
			}
			obj.DstStartIp = &v2
		}
	}
	if v1, ok := d.GetOk("dst_start_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_start_ip6", sv)
				diags = append(diags, e)
			}
			obj.DstStartIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("dst_subnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_subnet", sv)
				diags = append(diags, e)
			}
			obj.DstSubnet = &v2
		}
	}
	if v1, ok := d.GetOk("dst_subnet6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_subnet6", sv)
				diags = append(diags, e)
			}
			obj.DstSubnet6 = &v2
		}
	}
	if v1, ok := d.GetOk("encapsulation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("encapsulation", sv)
				diags = append(diags, e)
			}
			obj.Encapsulation = &v2
		}
	}
	if v1, ok := d.GetOk("initiator_ts_narrow"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("initiator_ts_narrow", sv)
				diags = append(diags, e)
			}
			obj.InitiatorTsNarrow = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_df"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_df", sv)
				diags = append(diags, e)
			}
			obj.Ipv4Df = &v2
		}
	}
	if v1, ok := d.GetOk("keepalive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keepalive", sv)
				diags = append(diags, e)
			}
			obj.Keepalive = &v2
		}
	}
	if v1, ok := d.GetOk("keylife_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keylife_type", sv)
				diags = append(diags, e)
			}
			obj.KeylifeType = &v2
		}
	}
	if v1, ok := d.GetOk("keylifekbs"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keylifekbs", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Keylifekbs = &tmp
		}
	}
	if v1, ok := d.GetOk("keylifeseconds"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keylifeseconds", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Keylifeseconds = &tmp
		}
	}
	if v1, ok := d.GetOk("l2tp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("l2tp", sv)
				diags = append(diags, e)
			}
			obj.L2tp = &v2
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
	if v1, ok := d.GetOk("pfs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pfs", sv)
				diags = append(diags, e)
			}
			obj.Pfs = &v2
		}
	}
	if v1, ok := d.GetOk("phase1name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("phase1name", sv)
				diags = append(diags, e)
			}
			obj.Phase1name = &v2
		}
	}
	if v1, ok := d.GetOk("proposal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proposal", sv)
				diags = append(diags, e)
			}
			obj.Proposal = &v2
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Protocol = &tmp
		}
	}
	if v1, ok := d.GetOk("replay"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replay", sv)
				diags = append(diags, e)
			}
			obj.Replay = &v2
		}
	}
	if v1, ok := d.GetOk("route_overlap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_overlap", sv)
				diags = append(diags, e)
			}
			obj.RouteOverlap = &v2
		}
	}
	if v1, ok := d.GetOk("single_source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("single_source", sv)
				diags = append(diags, e)
			}
			obj.SingleSource = &v2
		}
	}
	if v1, ok := d.GetOk("src_addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_addr_type", sv)
				diags = append(diags, e)
			}
			obj.SrcAddrType = &v2
		}
	}
	if v1, ok := d.GetOk("src_end_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_end_ip", sv)
				diags = append(diags, e)
			}
			obj.SrcEndIp = &v2
		}
	}
	if v1, ok := d.GetOk("src_end_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_end_ip6", sv)
				diags = append(diags, e)
			}
			obj.SrcEndIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("src_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_name", sv)
				diags = append(diags, e)
			}
			obj.SrcName = &v2
		}
	}
	if v1, ok := d.GetOk("src_name6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_name6", sv)
				diags = append(diags, e)
			}
			obj.SrcName6 = &v2
		}
	}
	if v1, ok := d.GetOk("src_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SrcPort = &tmp
		}
	}
	if v1, ok := d.GetOk("src_start_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_start_ip", sv)
				diags = append(diags, e)
			}
			obj.SrcStartIp = &v2
		}
	}
	if v1, ok := d.GetOk("src_start_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_start_ip6", sv)
				diags = append(diags, e)
			}
			obj.SrcStartIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("src_subnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_subnet", sv)
				diags = append(diags, e)
			}
			obj.SrcSubnet = &v2
		}
	}
	if v1, ok := d.GetOk("src_subnet6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_subnet6", sv)
				diags = append(diags, e)
			}
			obj.SrcSubnet6 = &v2
		}
	}
	return &obj, diags
}
