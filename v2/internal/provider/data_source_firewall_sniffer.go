// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure sniffer.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallSniffer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure sniffer.",

		ReadContext: dataSourceFirewallSnifferRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"anomaly": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Denial of Service (DoS) anomaly settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action taken when the threshold is reached.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable anomaly logging.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Anomaly name.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Quarantine method.",
							Computed:    true,
						},
						"quarantine_expiry": {
							Type:        schema.TypeString,
							Description: "Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.",
							Computed:    true,
						},
						"quarantine_log": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine logging.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this anomaly.",
							Computed:    true,
						},
						"threshold": {
							Type:        schema.TypeInt,
							Description: "Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.",
							Computed:    true,
						},
						"thresholddefault": {
							Type:        schema.TypeInt,
							Description: "Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.",
							Computed:    true,
						},
					},
				},
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Name of an existing application list.",
				Computed:    true,
			},
			"application_list_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable application control profile.",
				Computed:    true,
			},
			"av_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing antivirus profile.",
				Computed:    true,
			},
			"av_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable antivirus profile.",
				Computed:    true,
			},
			"dlp_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing DLP profile.",
				Computed:    true,
			},
			"dlp_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable DLP profile.",
				Computed:    true,
			},
			"dlp_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing DLP sensor.",
				Computed:    true,
			},
			"dlp_sensor_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable DLP sensor.",
				Computed:    true,
			},
			"dsri": {
				Type:        schema.TypeString,
				Description: "Enable/disable DSRI.",
				Computed:    true,
			},
			"emailfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing email filter profile.",
				Computed:    true,
			},
			"emailfilter_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable emailfilter.",
				Computed:    true,
			},
			"file_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing file-filter profile.",
				Computed:    true,
			},
			"file_filter_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable file filter.",
				Computed:    true,
			},
			"host": {
				Type:        schema.TypeString,
				Description: "Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Sniffer ID (0 - 9999).",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name that traffic sniffing will take place on.",
				Computed:    true,
			},
			"ip_threatfeed": {
				Type:        schema.TypeList,
				Description: "Name of an existing IP threat feed.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Threat feed name.",
							Computed:    true,
						},
					},
				},
			},
			"ip_threatfeed_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable IP threat feed.",
				Computed:    true,
			},
			"ips_dos_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS DoS anomaly detection.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing IPS sensor.",
				Computed:    true,
			},
			"ips_sensor_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS sensor.",
				Computed:    true,
			},
			"ipv6": {
				Type:        schema.TypeString,
				Description: "Enable/disable sniffing IPv6 packets.",
				Computed:    true,
			},
			"logtraffic": {
				Type:        schema.TypeString,
				Description: "Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy.",
				Computed:    true,
			},
			"max_packet_count": {
				Type:        schema.TypeInt,
				Description: "Maximum packet count (1 - 1000000, default = 4000).",
				Computed:    true,
			},
			"non_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable sniffing non-IP packets.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeString,
				Description: "Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Integer value for the protocol type as defined by IANA (0 - 255).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable the active status of the sniffer.",
				Computed:    true,
			},
			"vlan": {
				Type:        schema.TypeString,
				Description: "List of VLANs to sniff.",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing web filter profile.",
				Computed:    true,
			},
			"webfilter_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable web filter profile.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallSnifferRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("fosid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadFirewallSniffer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSniffer dataSource: %v", err)
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

	diags := refreshObjectFirewallSniffer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
