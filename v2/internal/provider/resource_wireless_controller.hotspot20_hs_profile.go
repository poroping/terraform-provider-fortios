// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure hotspot profile.

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

func resourceWirelessControllerHotspot20HsProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure hotspot profile.",

		CreateContext: resourceWirelessControllerHotspot20HsProfileCreate,
		ReadContext:   resourceWirelessControllerHotspot20HsProfileRead,
		UpdateContext: resourceWirelessControllerHotspot20HsProfileUpdate,
		DeleteContext: resourceWirelessControllerHotspot20HsProfileDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"3gpp_plmn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "3GPP PLMN name.",
				Optional:    true,
				Computed:    true,
			},
			"access_network_asra": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable additional step required for access (ASRA).",
				Optional:    true,
				Computed:    true,
			},
			"access_network_esr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable emergency services reachable (ESR).",
				Optional:    true,
				Computed:    true,
			},
			"access_network_internet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable connectivity to the Internet.",
				Optional:    true,
				Computed:    true,
			},
			"access_network_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"private-network", "private-network-with-guest-access", "chargeable-public-network", "free-public-network", "personal-device-network", "emergency-services-only-network", "test-or-experimental", "wildcard"}, false),

				Description: "Access network type.",
				Optional:    true,
				Computed:    true,
			},
			"access_network_uesa": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable unauthenticated emergency service accessible (UESA).",
				Optional:    true,
				Computed:    true,
			},
			"advice_of_charge": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Advice of charge.",
				Optional:    true,
				Computed:    true,
			},
			"anqp_domain_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "ANQP Domain ID (0-65535).",
				Optional:    true,
				Computed:    true,
			},
			"bss_transition": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable basic service set (BSS) transition Support.",
				Optional:    true,
				Computed:    true,
			},
			"conn_cap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Connection capability name.",
				Optional:    true,
				Computed:    true,
			},
			"deauth_request_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 120),

				Description: "Deauthentication request timeout (in seconds).",
				Optional:    true,
				Computed:    true,
			},
			"dgaf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable downstream group-addressed forwarding (DGAF).",
				Optional:    true,
				Computed:    true,
			},
			"domain_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Domain name.",
				Optional:    true,
				Computed:    true,
			},
			"gas_comeback_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 10000),

				Description: "GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"gas_fragmentation_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(512, 4096),

				Description: "GAS fragmentation limit (512 - 4096, default = 1024).",
				Optional:    true,
				Computed:    true,
			},
			"hessid": {
				Type: schema.TypeString,

				Description: "Homogeneous extended service set identifier (HESSID).",
				Optional:    true,
				Computed:    true,
			},
			"ip_addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IP address type name.",
				Optional:    true,
				Computed:    true,
			},
			"l2tif": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Layer 2 traffic inspection and filtering.",
				Optional:    true,
				Computed:    true,
			},
			"nai_realm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "NAI realm list name.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Hotspot profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"network_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Network authentication name.",
				Optional:    true,
				Computed:    true,
			},
			"oper_friendly_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Operator friendly name.",
				Optional:    true,
				Computed:    true,
			},
			"oper_icon": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Operator icon.",
				Optional:    true,
				Computed:    true,
			},
			"osu_provider": {
				Type:        schema.TypeList,
				Description: "Manually selected list of OSU provider(s).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "OSU provider name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"osu_provider_nai": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "OSU Provider NAI.",
				Optional:    true,
				Computed:    true,
			},
			"osu_ssid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Online sign up (OSU) SSID.",
				Optional:    true,
				Computed:    true,
			},
			"pame_bi": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI).",
				Optional:    true,
				Computed:    true,
			},
			"proxy_arp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Proxy ARP.",
				Optional:    true,
				Computed:    true,
			},
			"qos_map": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "QoS MAP set ID.",
				Optional:    true,
				Computed:    true,
			},
			"release": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3),

				Description: "Hotspot 2.0 Release number (1, 2, 3, default = 2).",
				Optional:    true,
				Computed:    true,
			},
			"roaming_consortium": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Roaming consortium list name.",
				Optional:    true,
				Computed:    true,
			},
			"terms_and_conditions": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Terms and conditions.",
				Optional:    true,
				Computed:    true,
			},
			"venue_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unspecified", "assembly", "business", "educational", "factory", "institutional", "mercantile", "residential", "storage", "utility", "vehicular", "outdoor"}, false),

				Description: "Venue group.",
				Optional:    true,
				Computed:    true,
			},
			"venue_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Venue name.",
				Optional:    true,
				Computed:    true,
			},
			"venue_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unspecified", "arena", "stadium", "passenger-terminal", "amphitheater", "amusement-park", "place-of-worship", "convention-center", "library", "museum", "restaurant", "theater", "bar", "coffee-shop", "zoo-or-aquarium", "emergency-center", "doctor-office", "bank", "fire-station", "police-station", "post-office", "professional-office", "research-facility", "attorney-office", "primary-school", "secondary-school", "university-or-college", "factory", "hospital", "long-term-care-facility", "rehab-center", "group-home", "prison-or-jail", "retail-store", "grocery-market", "auto-service-station", "shopping-mall", "gas-station", "private", "hotel-or-motel", "dormitory", "boarding-house", "automobile", "airplane", "bus", "ferry", "ship-or-boat", "train", "motor-bike", "muni-mesh-network", "city-park", "rest-area", "traffic-control", "bus-stop", "kiosk"}, false),

				Description: "Venue type.",
				Optional:    true,
				Computed:    true,
			},
			"venue_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Venue name.",
				Optional:    true,
				Computed:    true,
			},
			"wan_metrics": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WAN metric name.",
				Optional:    true,
				Computed:    true,
			},
			"wnm_sleep_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable wireless network management (WNM) sleep mode.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20HsProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20HsProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20HsProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20HsProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20HsProfile")
	}

	return resourceWirelessControllerHotspot20HsProfileRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20HsProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20HsProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20HsProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20HsProfile")
	}

	return resourceWirelessControllerHotspot20HsProfileRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20HsProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20HsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20HsProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20HsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20HsProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20HsProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerHotspot20HsProfileOsuProvider(v *[]models.WirelessControllerHotspot20HsProfileOsuProvider, sort bool) interface{} {
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

func refreshObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData, o *models.WirelessControllerHotspot20HsProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.The3gppPlmn != nil {
		v := *o.The3gppPlmn

		if err = d.Set("3gpp_plmn", v); err != nil {
			return diag.Errorf("error reading 3gpp_plmn: %v", err)
		}
	}

	if o.AccessNetworkAsra != nil {
		v := *o.AccessNetworkAsra

		if err = d.Set("access_network_asra", v); err != nil {
			return diag.Errorf("error reading access_network_asra: %v", err)
		}
	}

	if o.AccessNetworkEsr != nil {
		v := *o.AccessNetworkEsr

		if err = d.Set("access_network_esr", v); err != nil {
			return diag.Errorf("error reading access_network_esr: %v", err)
		}
	}

	if o.AccessNetworkInternet != nil {
		v := *o.AccessNetworkInternet

		if err = d.Set("access_network_internet", v); err != nil {
			return diag.Errorf("error reading access_network_internet: %v", err)
		}
	}

	if o.AccessNetworkType != nil {
		v := *o.AccessNetworkType

		if err = d.Set("access_network_type", v); err != nil {
			return diag.Errorf("error reading access_network_type: %v", err)
		}
	}

	if o.AccessNetworkUesa != nil {
		v := *o.AccessNetworkUesa

		if err = d.Set("access_network_uesa", v); err != nil {
			return diag.Errorf("error reading access_network_uesa: %v", err)
		}
	}

	if o.AdviceOfCharge != nil {
		v := *o.AdviceOfCharge

		if err = d.Set("advice_of_charge", v); err != nil {
			return diag.Errorf("error reading advice_of_charge: %v", err)
		}
	}

	if o.AnqpDomainId != nil {
		v := *o.AnqpDomainId

		if err = d.Set("anqp_domain_id", v); err != nil {
			return diag.Errorf("error reading anqp_domain_id: %v", err)
		}
	}

	if o.BssTransition != nil {
		v := *o.BssTransition

		if err = d.Set("bss_transition", v); err != nil {
			return diag.Errorf("error reading bss_transition: %v", err)
		}
	}

	if o.ConnCap != nil {
		v := *o.ConnCap

		if err = d.Set("conn_cap", v); err != nil {
			return diag.Errorf("error reading conn_cap: %v", err)
		}
	}

	if o.DeauthRequestTimeout != nil {
		v := *o.DeauthRequestTimeout

		if err = d.Set("deauth_request_timeout", v); err != nil {
			return diag.Errorf("error reading deauth_request_timeout: %v", err)
		}
	}

	if o.Dgaf != nil {
		v := *o.Dgaf

		if err = d.Set("dgaf", v); err != nil {
			return diag.Errorf("error reading dgaf: %v", err)
		}
	}

	if o.DomainName != nil {
		v := *o.DomainName

		if err = d.Set("domain_name", v); err != nil {
			return diag.Errorf("error reading domain_name: %v", err)
		}
	}

	if o.GasComebackDelay != nil {
		v := *o.GasComebackDelay

		if err = d.Set("gas_comeback_delay", v); err != nil {
			return diag.Errorf("error reading gas_comeback_delay: %v", err)
		}
	}

	if o.GasFragmentationLimit != nil {
		v := *o.GasFragmentationLimit

		if err = d.Set("gas_fragmentation_limit", v); err != nil {
			return diag.Errorf("error reading gas_fragmentation_limit: %v", err)
		}
	}

	if o.Hessid != nil {
		v := *o.Hessid

		if err = d.Set("hessid", v); err != nil {
			return diag.Errorf("error reading hessid: %v", err)
		}
	}

	if o.IpAddrType != nil {
		v := *o.IpAddrType

		if err = d.Set("ip_addr_type", v); err != nil {
			return diag.Errorf("error reading ip_addr_type: %v", err)
		}
	}

	if o.L2tif != nil {
		v := *o.L2tif

		if err = d.Set("l2tif", v); err != nil {
			return diag.Errorf("error reading l2tif: %v", err)
		}
	}

	if o.NaiRealm != nil {
		v := *o.NaiRealm

		if err = d.Set("nai_realm", v); err != nil {
			return diag.Errorf("error reading nai_realm: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NetworkAuth != nil {
		v := *o.NetworkAuth

		if err = d.Set("network_auth", v); err != nil {
			return diag.Errorf("error reading network_auth: %v", err)
		}
	}

	if o.OperFriendlyName != nil {
		v := *o.OperFriendlyName

		if err = d.Set("oper_friendly_name", v); err != nil {
			return diag.Errorf("error reading oper_friendly_name: %v", err)
		}
	}

	if o.OperIcon != nil {
		v := *o.OperIcon

		if err = d.Set("oper_icon", v); err != nil {
			return diag.Errorf("error reading oper_icon: %v", err)
		}
	}

	if o.OsuProvider != nil {
		if err = d.Set("osu_provider", flattenWirelessControllerHotspot20HsProfileOsuProvider(o.OsuProvider, sort)); err != nil {
			return diag.Errorf("error reading osu_provider: %v", err)
		}
	}

	if o.OsuProviderNai != nil {
		v := *o.OsuProviderNai

		if err = d.Set("osu_provider_nai", v); err != nil {
			return diag.Errorf("error reading osu_provider_nai: %v", err)
		}
	}

	if o.OsuSsid != nil {
		v := *o.OsuSsid

		if err = d.Set("osu_ssid", v); err != nil {
			return diag.Errorf("error reading osu_ssid: %v", err)
		}
	}

	if o.PameBi != nil {
		v := *o.PameBi

		if err = d.Set("pame_bi", v); err != nil {
			return diag.Errorf("error reading pame_bi: %v", err)
		}
	}

	if o.ProxyArp != nil {
		v := *o.ProxyArp

		if err = d.Set("proxy_arp", v); err != nil {
			return diag.Errorf("error reading proxy_arp: %v", err)
		}
	}

	if o.QosMap != nil {
		v := *o.QosMap

		if err = d.Set("qos_map", v); err != nil {
			return diag.Errorf("error reading qos_map: %v", err)
		}
	}

	if o.Release != nil {
		v := *o.Release

		if err = d.Set("release", v); err != nil {
			return diag.Errorf("error reading release: %v", err)
		}
	}

	if o.RoamingConsortium != nil {
		v := *o.RoamingConsortium

		if err = d.Set("roaming_consortium", v); err != nil {
			return diag.Errorf("error reading roaming_consortium: %v", err)
		}
	}

	if o.TermsAndConditions != nil {
		v := *o.TermsAndConditions

		if err = d.Set("terms_and_conditions", v); err != nil {
			return diag.Errorf("error reading terms_and_conditions: %v", err)
		}
	}

	if o.VenueGroup != nil {
		v := *o.VenueGroup

		if err = d.Set("venue_group", v); err != nil {
			return diag.Errorf("error reading venue_group: %v", err)
		}
	}

	if o.VenueName != nil {
		v := *o.VenueName

		if err = d.Set("venue_name", v); err != nil {
			return diag.Errorf("error reading venue_name: %v", err)
		}
	}

	if o.VenueType != nil {
		v := *o.VenueType

		if err = d.Set("venue_type", v); err != nil {
			return diag.Errorf("error reading venue_type: %v", err)
		}
	}

	if o.VenueUrl != nil {
		v := *o.VenueUrl

		if err = d.Set("venue_url", v); err != nil {
			return diag.Errorf("error reading venue_url: %v", err)
		}
	}

	if o.WanMetrics != nil {
		v := *o.WanMetrics

		if err = d.Set("wan_metrics", v); err != nil {
			return diag.Errorf("error reading wan_metrics: %v", err)
		}
	}

	if o.WnmSleepMode != nil {
		v := *o.WnmSleepMode

		if err = d.Set("wnm_sleep_mode", v); err != nil {
			return diag.Errorf("error reading wnm_sleep_mode: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerHotspot20HsProfileOsuProvider(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20HsProfileOsuProvider, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20HsProfileOsuProvider

	for i := range l {
		tmp := models.WirelessControllerHotspot20HsProfileOsuProvider{}
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

func getObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20HsProfile, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20HsProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("3gpp_plmn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("3gpp_plmn", sv)
				diags = append(diags, e)
			}
			obj.The3gppPlmn = &v2
		}
	}
	if v1, ok := d.GetOk("access_network_asra"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_network_asra", sv)
				diags = append(diags, e)
			}
			obj.AccessNetworkAsra = &v2
		}
	}
	if v1, ok := d.GetOk("access_network_esr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_network_esr", sv)
				diags = append(diags, e)
			}
			obj.AccessNetworkEsr = &v2
		}
	}
	if v1, ok := d.GetOk("access_network_internet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_network_internet", sv)
				diags = append(diags, e)
			}
			obj.AccessNetworkInternet = &v2
		}
	}
	if v1, ok := d.GetOk("access_network_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_network_type", sv)
				diags = append(diags, e)
			}
			obj.AccessNetworkType = &v2
		}
	}
	if v1, ok := d.GetOk("access_network_uesa"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_network_uesa", sv)
				diags = append(diags, e)
			}
			obj.AccessNetworkUesa = &v2
		}
	}
	if v1, ok := d.GetOk("advice_of_charge"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("advice_of_charge", sv)
				diags = append(diags, e)
			}
			obj.AdviceOfCharge = &v2
		}
	}
	if v1, ok := d.GetOk("anqp_domain_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("anqp_domain_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AnqpDomainId = &tmp
		}
	}
	if v1, ok := d.GetOk("bss_transition"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bss_transition", sv)
				diags = append(diags, e)
			}
			obj.BssTransition = &v2
		}
	}
	if v1, ok := d.GetOk("conn_cap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("conn_cap", sv)
				diags = append(diags, e)
			}
			obj.ConnCap = &v2
		}
	}
	if v1, ok := d.GetOk("deauth_request_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deauth_request_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeauthRequestTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("dgaf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dgaf", sv)
				diags = append(diags, e)
			}
			obj.Dgaf = &v2
		}
	}
	if v1, ok := d.GetOk("domain_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain_name", sv)
				diags = append(diags, e)
			}
			obj.DomainName = &v2
		}
	}
	if v1, ok := d.GetOk("gas_comeback_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gas_comeback_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GasComebackDelay = &tmp
		}
	}
	if v1, ok := d.GetOk("gas_fragmentation_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gas_fragmentation_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GasFragmentationLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("hessid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hessid", sv)
				diags = append(diags, e)
			}
			obj.Hessid = &v2
		}
	}
	if v1, ok := d.GetOk("ip_addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_addr_type", sv)
				diags = append(diags, e)
			}
			obj.IpAddrType = &v2
		}
	}
	if v1, ok := d.GetOk("l2tif"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("l2tif", sv)
				diags = append(diags, e)
			}
			obj.L2tif = &v2
		}
	}
	if v1, ok := d.GetOk("nai_realm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nai_realm", sv)
				diags = append(diags, e)
			}
			obj.NaiRealm = &v2
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
	if v1, ok := d.GetOk("network_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("network_auth", sv)
				diags = append(diags, e)
			}
			obj.NetworkAuth = &v2
		}
	}
	if v1, ok := d.GetOk("oper_friendly_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("oper_friendly_name", sv)
				diags = append(diags, e)
			}
			obj.OperFriendlyName = &v2
		}
	}
	if v1, ok := d.GetOk("oper_icon"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("oper_icon", sv)
				diags = append(diags, e)
			}
			obj.OperIcon = &v2
		}
	}
	if v, ok := d.GetOk("osu_provider"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("osu_provider", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20HsProfileOsuProvider(d, v, "osu_provider", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OsuProvider = t
		}
	} else if d.HasChange("osu_provider") {
		old, new := d.GetChange("osu_provider")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OsuProvider = &[]models.WirelessControllerHotspot20HsProfileOsuProvider{}
		}
	}
	if v1, ok := d.GetOk("osu_provider_nai"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("osu_provider_nai", sv)
				diags = append(diags, e)
			}
			obj.OsuProviderNai = &v2
		}
	}
	if v1, ok := d.GetOk("osu_ssid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("osu_ssid", sv)
				diags = append(diags, e)
			}
			obj.OsuSsid = &v2
		}
	}
	if v1, ok := d.GetOk("pame_bi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pame_bi", sv)
				diags = append(diags, e)
			}
			obj.PameBi = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_arp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_arp", sv)
				diags = append(diags, e)
			}
			obj.ProxyArp = &v2
		}
	}
	if v1, ok := d.GetOk("qos_map"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("qos_map", sv)
				diags = append(diags, e)
			}
			obj.QosMap = &v2
		}
	}
	if v1, ok := d.GetOk("release"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("release", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Release = &tmp
		}
	}
	if v1, ok := d.GetOk("roaming_consortium"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("roaming_consortium", sv)
				diags = append(diags, e)
			}
			obj.RoamingConsortium = &v2
		}
	}
	if v1, ok := d.GetOk("terms_and_conditions"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("terms_and_conditions", sv)
				diags = append(diags, e)
			}
			obj.TermsAndConditions = &v2
		}
	}
	if v1, ok := d.GetOk("venue_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("venue_group", sv)
				diags = append(diags, e)
			}
			obj.VenueGroup = &v2
		}
	}
	if v1, ok := d.GetOk("venue_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("venue_name", sv)
				diags = append(diags, e)
			}
			obj.VenueName = &v2
		}
	}
	if v1, ok := d.GetOk("venue_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("venue_type", sv)
				diags = append(diags, e)
			}
			obj.VenueType = &v2
		}
	}
	if v1, ok := d.GetOk("venue_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("venue_url", sv)
				diags = append(diags, e)
			}
			obj.VenueUrl = &v2
		}
	}
	if v1, ok := d.GetOk("wan_metrics"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wan_metrics", sv)
				diags = append(diags, e)
			}
			obj.WanMetrics = &v2
		}
	}
	if v1, ok := d.GetOk("wnm_sleep_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wnm_sleep_mode", sv)
				diags = append(diags, e)
			}
			obj.WnmSleepMode = &v2
		}
	}
	return &obj, diags
}
