// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Overlay Controller VPN settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVpnOcvpn() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Overlay Controller VPN settings.",

		ReadContext: dataSourceVpnOcvpnRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auto_discovery": {
				Type:        schema.TypeString,
				Description: "Enable/disable auto-discovery shortcuts.",
				Computed:    true,
			},
			"auto_discovery_shortcut_mode": {
				Type:        schema.TypeString,
				Description: "Control deletion of child short-cut tunnels when the parent tunnel goes down.",
				Computed:    true,
			},
			"eap": {
				Type:        schema.TypeString,
				Description: "Enable/disable EAP client authentication.",
				Computed:    true,
			},
			"eap_users": {
				Type:        schema.TypeString,
				Description: "EAP authentication user group.",
				Computed:    true,
			},
			"forticlient_access": {
				Type:        schema.TypeList,
				Description: "Configure FortiClient settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_groups": {
							Type:        schema.TypeList,
							Description: "FortiClient user authentication groups.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_group": {
										Type:        schema.TypeString,
										Description: "Authentication user group for FortiClient access.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Group name.",
										Computed:    true,
									},
									"overlays": {
										Type:        schema.TypeList,
										Description: "OCVPN overlays to allow access to.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"overlay_name": {
													Type:        schema.TypeString,
													Description: "Overlay name.",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"psksecret": {
							Type:        schema.TypeString,
							Description: "Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).",
							Computed:    true,
							Sensitive:   true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable FortiClient to access OCVPN networks.",
							Computed:    true,
						},
					},
				},
			},
			"ip_allocation_block": {
				Type:        schema.TypeString,
				Description: "Class B subnet reserved for private IP address assignment.",
				Computed:    true,
			},
			"multipath": {
				Type:        schema.TypeString,
				Description: "Enable/disable multipath redundancy.",
				Computed:    true,
			},
			"nat": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT support.",
				Computed:    true,
			},
			"overlays": {
				Type:        schema.TypeList,
				Description: "Network overlays to register with Overlay Controller VPN service.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_ip": {
							Type:        schema.TypeString,
							Description: "Enable/disable mode-cfg address assignment.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"inter_overlay": {
							Type:        schema.TypeString,
							Description: "Allow or deny traffic from other overlays.",
							Computed:    true,
						},
						"ipv4_end_ip": {
							Type:        schema.TypeString,
							Description: "End of IPv4 range.",
							Computed:    true,
						},
						"ipv4_start_ip": {
							Type:        schema.TypeString,
							Description: "Start of IPv4 range.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Overlay name.",
							Computed:    true,
						},
						"overlay_name": {
							Type:        schema.TypeString,
							Description: "Overlay name.",
							Computed:    true,
						},
						"subnets": {
							Type:        schema.TypeList,
							Description: "Internal subnets to register with OCVPN service.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "ID.",
										Computed:    true,
									},
									"interface": {
										Type:        schema.TypeString,
										Description: "LAN interface.",
										Computed:    true,
									},
									"subnet": {
										Type:        schema.TypeString,
										Description: "IPv4 address and subnet mask.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Subnet type.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"poll_interval": {
				Type:        schema.TypeInt,
				Description: "Overlay Controller VPN polling interval.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "Set device role.",
				Computed:    true,
			},
			"sdwan": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding OCVPN tunnels to SD-WAN.",
				Computed:    true,
			},
			"sdwan_zone": {
				Type:        schema.TypeString,
				Description: "Set SD-WAN zone.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable Overlay Controller cloud assisted VPN.",
				Computed:    true,
			},
			"wan_interface": {
				Type:        schema.TypeList,
				Description: "FortiGate WAN interfaces to use with OCVPN.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVpnOcvpnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnOcvpn(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnOcvpn dataSource: %v", err)
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

	diags := refreshObjectVpnOcvpn(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
