// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure central management.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Description: "Configure central management.",

		ReadContext: dataSourceSystemCentralManagementRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_monitor": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the central management server to remotely monitor this FortiGate",
				Computed:    true,
			},
			"allow_push_configuration": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the central management server to push configuration changes to this FortiGate.",
				Computed:    true,
			},
			"allow_push_firmware": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the central management server to push firmware updates to this FortiGate.",
				Computed:    true,
			},
			"allow_remote_firmware_upgrade": {
				Type:        schema.TypeString,
				Description: "Enable/disable remotely upgrading the firmware on this FortiGate from the central management server.",
				Computed:    true,
			},
			"ca_cert": {
				Type:        schema.TypeString,
				Description: "CA certificate to be used by FGFM protocol.",
				Computed:    true,
			},
			"enc_algorithm": {
				Type:        schema.TypeString,
				Description: "Encryption strength for communications between the FortiGate and central management.",
				Computed:    true,
			},
			"fmg": {
				Type:        schema.TypeString,
				Description: "IP address or FQDN of the FortiManager.",
				Computed:    true,
			},
			"fmg_source_ip": {
				Type:        schema.TypeString,
				Description: "IPv4 source address that this FortiGate uses when communicating with FortiManager.",
				Computed:    true,
			},
			"fmg_source_ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 source address that this FortiGate uses when communicating with FortiManager.",
				Computed:    true,
			},
			"fmg_update_port": {
				Type:        schema.TypeString,
				Description: "Port used to communicate with FortiManager that is acting as a FortiGuard update server.",
				Computed:    true,
			},
			"include_default_servers": {
				Type:        schema.TypeString,
				Description: "Enable/disable inclusion of public FortiGuard servers in the override server list.",
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
			"local_cert": {
				Type:        schema.TypeString,
				Description: "Certificate to be used by FGFM protocol.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Central management mode.",
				Computed:    true,
			},
			"schedule_config_restore": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the central management server to restore the configuration of this FortiGate.",
				Computed:    true,
			},
			"schedule_script_restore": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the central management server to restore the scripts stored on this FortiGate.",
				Computed:    true,
			},
			"serial_number": {
				Type:        schema.TypeString,
				Description: "Serial number.",
				Computed:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": {
							Type:        schema.TypeString,
							Description: "Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN.",
							Computed:    true,
						},
						"fqdn": {
							Type:        schema.TypeString,
							Description: "FQDN address of override server.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"server_address": {
							Type:        schema.TypeString,
							Description: "IPv4 address of override server.",
							Computed:    true,
						},
						"server_address6": {
							Type:        schema.TypeString,
							Description: "IPv6 address of override server.",
							Computed:    true,
						},
						"server_type": {
							Type:        schema.TypeString,
							Description: "FortiGuard service type.",
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Central management type.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeString,
				Description: "Virtual domain (VDOM) name to use when communicating with FortiManager.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemCentralManagementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemCentralManagement(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemCentralManagement dataSource: %v", err)
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

	diags := refreshObjectSystemCentralManagement(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
