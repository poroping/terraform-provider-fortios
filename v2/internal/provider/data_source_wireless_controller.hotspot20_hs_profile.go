// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure hotspot profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerHotspot20HsProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure hotspot profile.",

		ReadContext: dataSourceWirelessControllerHotspot20HsProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"3gpp_plmn": {
				Type:        schema.TypeString,
				Description: "3GPP PLMN name.",
				Computed:    true,
			},
			"access_network_asra": {
				Type:        schema.TypeString,
				Description: "Enable/disable additional step required for access (ASRA).",
				Computed:    true,
			},
			"access_network_esr": {
				Type:        schema.TypeString,
				Description: "Enable/disable emergency services reachable (ESR).",
				Computed:    true,
			},
			"access_network_internet": {
				Type:        schema.TypeString,
				Description: "Enable/disable connectivity to the Internet.",
				Computed:    true,
			},
			"access_network_type": {
				Type:        schema.TypeString,
				Description: "Access network type.",
				Computed:    true,
			},
			"access_network_uesa": {
				Type:        schema.TypeString,
				Description: "Enable/disable unauthenticated emergency service accessible (UESA).",
				Computed:    true,
			},
			"advice_of_charge": {
				Type:        schema.TypeString,
				Description: "Advice of charge.",
				Computed:    true,
			},
			"anqp_domain_id": {
				Type:        schema.TypeInt,
				Description: "ANQP Domain ID (0-65535).",
				Computed:    true,
			},
			"bss_transition": {
				Type:        schema.TypeString,
				Description: "Enable/disable basic service set (BSS) transition Support.",
				Computed:    true,
			},
			"conn_cap": {
				Type:        schema.TypeString,
				Description: "Connection capability name.",
				Computed:    true,
			},
			"deauth_request_timeout": {
				Type:        schema.TypeInt,
				Description: "Deauthentication request timeout (in seconds).",
				Computed:    true,
			},
			"dgaf": {
				Type:        schema.TypeString,
				Description: "Enable/disable downstream group-addressed forwarding (DGAF).",
				Computed:    true,
			},
			"domain_name": {
				Type:        schema.TypeString,
				Description: "Domain name.",
				Computed:    true,
			},
			"gas_comeback_delay": {
				Type:        schema.TypeInt,
				Description: "GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).",
				Computed:    true,
			},
			"gas_fragmentation_limit": {
				Type:        schema.TypeInt,
				Description: "GAS fragmentation limit (512 - 4096, default = 1024).",
				Computed:    true,
			},
			"hessid": {
				Type:        schema.TypeString,
				Description: "Homogeneous extended service set identifier (HESSID).",
				Computed:    true,
			},
			"ip_addr_type": {
				Type:        schema.TypeString,
				Description: "IP address type name.",
				Computed:    true,
			},
			"l2tif": {
				Type:        schema.TypeString,
				Description: "Enable/disable Layer 2 traffic inspection and filtering.",
				Computed:    true,
			},
			"nai_realm": {
				Type:        schema.TypeString,
				Description: "NAI realm list name.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Hotspot profile name.",
				Required:    true,
			},
			"network_auth": {
				Type:        schema.TypeString,
				Description: "Network authentication name.",
				Computed:    true,
			},
			"oper_friendly_name": {
				Type:        schema.TypeString,
				Description: "Operator friendly name.",
				Computed:    true,
			},
			"oper_icon": {
				Type:        schema.TypeString,
				Description: "Operator icon.",
				Computed:    true,
			},
			"osu_provider": {
				Type:        schema.TypeList,
				Description: "Manually selected list of OSU provider(s).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "OSU provider name.",
							Computed:    true,
						},
					},
				},
			},
			"osu_provider_nai": {
				Type:        schema.TypeString,
				Description: "OSU Provider NAI.",
				Computed:    true,
			},
			"osu_ssid": {
				Type:        schema.TypeString,
				Description: "Online sign up (OSU) SSID.",
				Computed:    true,
			},
			"pame_bi": {
				Type:        schema.TypeString,
				Description: "Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI).",
				Computed:    true,
			},
			"proxy_arp": {
				Type:        schema.TypeString,
				Description: "Enable/disable Proxy ARP.",
				Computed:    true,
			},
			"qos_map": {
				Type:        schema.TypeString,
				Description: "QoS MAP set ID.",
				Computed:    true,
			},
			"release": {
				Type:        schema.TypeInt,
				Description: "Hotspot 2.0 Release number (1, 2, 3, default = 2).",
				Computed:    true,
			},
			"roaming_consortium": {
				Type:        schema.TypeString,
				Description: "Roaming consortium list name.",
				Computed:    true,
			},
			"terms_and_conditions": {
				Type:        schema.TypeString,
				Description: "Terms and conditions.",
				Computed:    true,
			},
			"venue_group": {
				Type:        schema.TypeString,
				Description: "Venue group.",
				Computed:    true,
			},
			"venue_name": {
				Type:        schema.TypeString,
				Description: "Venue name.",
				Computed:    true,
			},
			"venue_type": {
				Type:        schema.TypeString,
				Description: "Venue type.",
				Computed:    true,
			},
			"venue_url": {
				Type:        schema.TypeString,
				Description: "Venue name.",
				Computed:    true,
			},
			"wan_metrics": {
				Type:        schema.TypeString,
				Description: "WAN metric name.",
				Computed:    true,
			},
			"wnm_sleep_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable wireless network management (WNM) sleep mode.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerHotspot20HsProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadWirelessControllerHotspot20HsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20HsProfile dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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

	d.SetId(mkey)

	return nil
}
