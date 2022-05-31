// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless controller global settings.

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
)

func resourceWirelessControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless controller global settings.",

		CreateContext: resourceWirelessControllerGlobalCreate,
		ReadContext:   resourceWirelessControllerGlobalRead,
		UpdateContext: resourceWirelessControllerGlobalUpdate,
		DeleteContext: resourceWirelessControllerGlobalDelete,

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
			"ap_log_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable configuring FortiGate to redirect wireless event log messages or FortiAPs to send UTM log messages to a syslog server (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"ap_log_server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address that FortiGate or FortiAPs send log messages to.",
				Optional:    true,
				Computed:    true,
			},
			"ap_log_server_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port that FortiGate or FortiAPs send log messages to.",
				Optional:    true,
				Computed:    true,
			},
			"control_message_offload": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Configure CAPWAP control message data channel offload.",
				Optional:         true,
				Computed:         true,
			},
			"data_ethernet_ii": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"discovery_mc_addr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Multicast IP address for AP discovery (default = 244.0.1.140).",
				Optional:    true,
				Computed:    true,
			},
			"fiapp_eth_type": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).",
				Optional:    true,
				Computed:    true,
			},
			"image_download": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WTP image download at join time.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_base_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).",
				Optional:    true,
				Computed:    true,
			},
			"link_aggregation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Description of the location of the wireless controller.",
				Optional:    true,
				Computed:    true,
			},
			"max_clients": {
				Type: schema.TypeInt,

				Description: "Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).",
				Optional:    true,
				Computed:    true,
			},
			"max_retransmit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 64),

				Description: "Maximum number of tunnel packet retransmissions (0 - 64, default = 3).",
				Optional:    true,
				Computed:    true,
			},
			"mesh_eth_type": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).",
				Optional:    true,
				Computed:    true,
			},
			"nac_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 600),

				Description: "Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the wireless controller.",
				Optional:    true,
				Computed:    true,
			},
			"rogue_scan_mac_adjacency": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31),

				Description: "Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"compatible", "strict"}, false),

				Description: "Compatible/strict tunnel mode.",
				Optional:    true,
				Computed:    true,
			},
			"wtp_share": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sharing of WTPs between VDOMs.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerGlobal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerGlobal")
	}

	return resourceWirelessControllerGlobalRead(ctx, d, meta)
}

func resourceWirelessControllerGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerGlobal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerGlobal")
	}

	return resourceWirelessControllerGlobalRead(ctx, d, meta)
}

func resourceWirelessControllerGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWirelessControllerGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWirelessControllerGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerGlobal resource: %v", err)
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

	diags := refreshObjectWirelessControllerGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerGlobal(d *schema.ResourceData, o *models.WirelessControllerGlobal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ApLogServer != nil {
		v := *o.ApLogServer

		if err = d.Set("ap_log_server", v); err != nil {
			return diag.Errorf("error reading ap_log_server: %v", err)
		}
	}

	if o.ApLogServerIp != nil {
		v := *o.ApLogServerIp

		if err = d.Set("ap_log_server_ip", v); err != nil {
			return diag.Errorf("error reading ap_log_server_ip: %v", err)
		}
	}

	if o.ApLogServerPort != nil {
		v := *o.ApLogServerPort

		if err = d.Set("ap_log_server_port", v); err != nil {
			return diag.Errorf("error reading ap_log_server_port: %v", err)
		}
	}

	if o.ControlMessageOffload != nil {
		v := *o.ControlMessageOffload

		if err = d.Set("control_message_offload", v); err != nil {
			return diag.Errorf("error reading control_message_offload: %v", err)
		}
	}

	if o.DataEthernetII != nil {
		v := *o.DataEthernetII

		if err = d.Set("data_ethernet_ii", v); err != nil {
			return diag.Errorf("error reading data_ethernet_ii: %v", err)
		}
	}

	if o.DiscoveryMcAddr != nil {
		v := *o.DiscoveryMcAddr

		if err = d.Set("discovery_mc_addr", v); err != nil {
			return diag.Errorf("error reading discovery_mc_addr: %v", err)
		}
	}

	if o.FiappEthType != nil {
		v := *o.FiappEthType

		if err = d.Set("fiapp_eth_type", v); err != nil {
			return diag.Errorf("error reading fiapp_eth_type: %v", err)
		}
	}

	if o.ImageDownload != nil {
		v := *o.ImageDownload

		if err = d.Set("image_download", v); err != nil {
			return diag.Errorf("error reading image_download: %v", err)
		}
	}

	if o.IpsecBaseIp != nil {
		v := *o.IpsecBaseIp

		if err = d.Set("ipsec_base_ip", v); err != nil {
			return diag.Errorf("error reading ipsec_base_ip: %v", err)
		}
	}

	if o.LinkAggregation != nil {
		v := *o.LinkAggregation

		if err = d.Set("link_aggregation", v); err != nil {
			return diag.Errorf("error reading link_aggregation: %v", err)
		}
	}

	if o.Location != nil {
		v := *o.Location

		if err = d.Set("location", v); err != nil {
			return diag.Errorf("error reading location: %v", err)
		}
	}

	if o.MaxClients != nil {
		v := *o.MaxClients

		if err = d.Set("max_clients", v); err != nil {
			return diag.Errorf("error reading max_clients: %v", err)
		}
	}

	if o.MaxRetransmit != nil {
		v := *o.MaxRetransmit

		if err = d.Set("max_retransmit", v); err != nil {
			return diag.Errorf("error reading max_retransmit: %v", err)
		}
	}

	if o.MeshEthType != nil {
		v := *o.MeshEthType

		if err = d.Set("mesh_eth_type", v); err != nil {
			return diag.Errorf("error reading mesh_eth_type: %v", err)
		}
	}

	if o.NacInterval != nil {
		v := *o.NacInterval

		if err = d.Set("nac_interval", v); err != nil {
			return diag.Errorf("error reading nac_interval: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.RogueScanMacAdjacency != nil {
		v := *o.RogueScanMacAdjacency

		if err = d.Set("rogue_scan_mac_adjacency", v); err != nil {
			return diag.Errorf("error reading rogue_scan_mac_adjacency: %v", err)
		}
	}

	if o.TunnelMode != nil {
		v := *o.TunnelMode

		if err = d.Set("tunnel_mode", v); err != nil {
			return diag.Errorf("error reading tunnel_mode: %v", err)
		}
	}

	if o.WtpShare != nil {
		v := *o.WtpShare

		if err = d.Set("wtp_share", v); err != nil {
			return diag.Errorf("error reading wtp_share: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerGlobal(d *schema.ResourceData, sv string) (*models.WirelessControllerGlobal, diag.Diagnostics) {
	obj := models.WirelessControllerGlobal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ap_log_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_log_server", sv)
				diags = append(diags, e)
			}
			obj.ApLogServer = &v2
		}
	}
	if v1, ok := d.GetOk("ap_log_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_log_server_ip", sv)
				diags = append(diags, e)
			}
			obj.ApLogServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("ap_log_server_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_log_server_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ApLogServerPort = &tmp
		}
	}
	if v1, ok := d.GetOk("control_message_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("control_message_offload", sv)
				diags = append(diags, e)
			}
			obj.ControlMessageOffload = &v2
		}
	}
	if v1, ok := d.GetOk("data_ethernet_ii"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("data_ethernet_ii", sv)
				diags = append(diags, e)
			}
			obj.DataEthernetII = &v2
		}
	}
	if v1, ok := d.GetOk("discovery_mc_addr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("discovery_mc_addr", sv)
				diags = append(diags, e)
			}
			obj.DiscoveryMcAddr = &v2
		}
	}
	if v1, ok := d.GetOk("fiapp_eth_type"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fiapp_eth_type", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FiappEthType = &tmp
		}
	}
	if v1, ok := d.GetOk("image_download"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("image_download", sv)
				diags = append(diags, e)
			}
			obj.ImageDownload = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_base_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_base_ip", sv)
				diags = append(diags, e)
			}
			obj.IpsecBaseIp = &v2
		}
	}
	if v1, ok := d.GetOk("link_aggregation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_aggregation", sv)
				diags = append(diags, e)
			}
			obj.LinkAggregation = &v2
		}
	}
	if v1, ok := d.GetOk("location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("location", sv)
				diags = append(diags, e)
			}
			obj.Location = &v2
		}
	}
	if v1, ok := d.GetOk("max_clients"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_clients", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxClients = &tmp
		}
	}
	if v1, ok := d.GetOk("max_retransmit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_retransmit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxRetransmit = &tmp
		}
	}
	if v1, ok := d.GetOk("mesh_eth_type"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mesh_eth_type", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MeshEthType = &tmp
		}
	}
	if v1, ok := d.GetOk("nac_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("nac_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NacInterval = &tmp
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
	if v1, ok := d.GetOk("rogue_scan_mac_adjacency"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rogue_scan_mac_adjacency", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RogueScanMacAdjacency = &tmp
		}
	}
	if v1, ok := d.GetOk("tunnel_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("tunnel_mode", sv)
				diags = append(diags, e)
			}
			obj.TunnelMode = &v2
		}
	}
	if v1, ok := d.GetOk("wtp_share"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wtp_share", sv)
				diags = append(diags, e)
			}
			obj.WtpShare = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWirelessControllerGlobal(d *schema.ResourceData, sv string) (*models.WirelessControllerGlobal, diag.Diagnostics) {
	obj := models.WirelessControllerGlobal{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
