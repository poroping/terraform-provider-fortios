// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemMobileTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.",

		ReadContext: dataSourceSystemMobileTunnelRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"hash_algorithm": {
				Type:        schema.TypeString,
				Description: "Hash Algorithm (Keyed MD5).",
				Computed:    true,
			},
			"home_address": {
				Type:        schema.TypeString,
				Description: "Home IP address (Format: xxx.xxx.xxx.xxx).",
				Computed:    true,
			},
			"home_agent": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).",
				Computed:    true,
			},
			"lifetime": {
				Type:        schema.TypeInt,
				Description: "NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).",
				Computed:    true,
			},
			"n_mhae_key": {
				Type:        schema.TypeString,
				Description: "NEMO authentication key.",
				Computed:    true,
			},
			"n_mhae_key_type": {
				Type:        schema.TypeString,
				Description: "NEMO authentication key type (ASCII or base64).",
				Computed:    true,
			},
			"n_mhae_spi": {
				Type:        schema.TypeInt,
				Description: "NEMO authentication SPI (default: 256).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Tunnel name.",
				Required:    true,
			},
			"network": {
				Type:        schema.TypeList,
				Description: "NEMO network configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Network entry ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Select the associated interface name from available options.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).",
							Computed:    true,
						},
					},
				},
			},
			"reg_interval": {
				Type:        schema.TypeInt,
				Description: "NMMO HA registration interval (5 - 300, default = 5).",
				Computed:    true,
			},
			"reg_retry": {
				Type:        schema.TypeInt,
				Description: "Maximum number of NMMO HA registration retries (1 to 30, default = 3).",
				Computed:    true,
			},
			"renew_interval": {
				Type:        schema.TypeInt,
				Description: "Time before lifetime expiration to send NMMO HA re-registration (5 - 60, default = 60).",
				Computed:    true,
			},
			"roaming_interface": {
				Type:        schema.TypeString,
				Description: "Select the associated interface name from available options.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this mobile tunnel.",
				Computed:    true,
			},
			"tunnel_mode": {
				Type:        schema.TypeString,
				Description: "NEMO tunnel mode (GRE tunnel).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemMobileTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemMobileTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemMobileTunnel dataSource: %v", err)
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

	diags := refreshObjectSystemMobileTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
