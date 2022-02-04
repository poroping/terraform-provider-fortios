// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system NTP information.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system NTP information.",

		ReadContext: dataSourceSystemNtpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"key": {
				Type:        schema.TypeString,
				Description: "Key for authentication.",
				Computed:    true,
				Sensitive:   true,
			},
			"key_id": {
				Type:        schema.TypeInt,
				Description: "Key ID for authentication.",
				Computed:    true,
			},
			"key_type": {
				Type:        schema.TypeString,
				Description: "Key type for authentication (MD5, SHA1).",
				Computed:    true,
			},
			"ntpserver": {
				Type:        schema.TypeList,
				Description: "Configure the FortiGate to connect to any available third-party NTP server.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Enable/disable MD5(NTPv3)/SHA1(NTPv4) authentication.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "NTP server ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Specify outgoing interface to reach server.",
							Computed:    true,
						},
						"interface_select_method": {
							Type:        schema.TypeString,
							Description: "Specify how to select outgoing interface to reach server.",
							Computed:    true,
						},
						"key": {
							Type:        schema.TypeString,
							Description: "Key for MD5(NTPv3)/SHA1(NTPv4) authentication.",
							Computed:    true,
							Sensitive:   true,
						},
						"key_id": {
							Type:        schema.TypeInt,
							Description: "Key ID for authentication.",
							Computed:    true,
						},
						"ntpv3": {
							Type:        schema.TypeString,
							Description: "Enable to use NTPv3 instead of NTPv4.",
							Computed:    true,
						},
						"server": {
							Type:        schema.TypeString,
							Description: "IP address or hostname of the NTP Server.",
							Computed:    true,
						},
					},
				},
			},
			"ntpsync": {
				Type:        schema.TypeString,
				Description: "Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.",
				Computed:    true,
			},
			"server_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address for communication to the NTP server.",
				Computed:    true,
			},
			"source_ip6": {
				Type:        schema.TypeString,
				Description: "Source IPv6 address for communication to the NTP server.",
				Computed:    true,
			},
			"syncinterval": {
				Type:        schema.TypeInt,
				Description: "NTP synchronization interval (1 - 1440 min).",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Use the FortiGuard NTP server or any other available NTP Server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemNtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemNtp"

	o, err := c.Cmdb.ReadSystemNtp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNtp dataSource: %v", err)
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

	diags := refreshObjectSystemNtp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
