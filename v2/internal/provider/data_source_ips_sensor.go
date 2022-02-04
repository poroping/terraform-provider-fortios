// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS sensor.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceIpsSensor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS sensor.",

		ReadContext: dataSourceIpsSensorRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"block_malicious_url": {
				Type:        schema.TypeString,
				Description: "Enable/disable malicious URL blocking.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "IPS sensor filter.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action taken with traffic in which signatures are detected.",
							Computed:    true,
						},
						"application": {
							Type:        schema.TypeString,
							Description: "Operating systems to be protected. Use all for every application and other for unlisted application.",
							Computed:    true,
						},
						"cve": {
							Type:        schema.TypeList,
							Description: "List of CVE IDs of the signatures to add to the sensor.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cve_entry": {
										Type:        schema.TypeString,
										Description: "CVE IDs or CVE wildcards.",
										Computed:    true,
									},
								},
							},
						},
						"exempt_ip": {
							Type:        schema.TypeList,
							Description: "Traffic from selected source or destination IP addresses is exempt from this signature.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": {
										Type:        schema.TypeString,
										Description: "Destination IP address and netmask (applies to packet matching the signature).",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Exempt IP ID.",
										Computed:    true,
									},
									"src_ip": {
										Type:        schema.TypeString,
										Description: "Source IP address and netmask (applies to packet matching the signature).",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Rule ID in IPS database (0 - 4294967295).",
							Computed:    true,
						},
						"location": {
							Type:        schema.TypeString,
							Description: "Protect client or server traffic.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of signatures included in filter.",
							Computed:    true,
						},
						"log_attack_context": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of attack context: URL buffer, header buffer, body buffer, packet buffer.",
							Computed:    true,
						},
						"log_packet": {
							Type:        schema.TypeString,
							Description: "Enable/disable packet logging. Enable to save the packet that triggers the filter. You can download the packets in pcap format for diagnostic use.",
							Computed:    true,
						},
						"os": {
							Type:        schema.TypeString,
							Description: "Operating systems to be protected. Use all for every operating system and other for unlisted operating systems.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocols to be examined. Use all for every protocol and other for unlisted protocols.",
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
						"rate_count": {
							Type:        schema.TypeInt,
							Description: "Count of the rate.",
							Computed:    true,
						},
						"rate_duration": {
							Type:        schema.TypeInt,
							Description: "Duration (sec) of the rate.",
							Computed:    true,
						},
						"rate_mode": {
							Type:        schema.TypeString,
							Description: "Rate limit mode.",
							Computed:    true,
						},
						"rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"rule": {
							Type:        schema.TypeList,
							Description: "Identifies the predefined or custom IPS signatures to add to the sensor.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Rule IPS.",
										Computed:    true,
									},
								},
							},
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status of the signatures included in filter. Only those filters with a status to enable are used.",
							Computed:    true,
						},
					},
				},
			},
			"extended_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable extended logging.",
				Computed:    true,
			},
			"filter": {
				Type:        schema.TypeList,
				Description: "IPS sensor filter.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action of selected rules.",
							Computed:    true,
						},
						"application": {
							Type:        schema.TypeString,
							Description: "Vulnerable application filter.",
							Computed:    true,
						},
						"location": {
							Type:        schema.TypeString,
							Description: "Vulnerability location filter.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of selected rules.",
							Computed:    true,
						},
						"log_packet": {
							Type:        schema.TypeString,
							Description: "Enable/disable packet logging of selected rules.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Filter name.",
							Computed:    true,
						},
						"os": {
							Type:        schema.TypeString,
							Description: "Vulnerable OS filter.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Vulnerable protocol filter.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Quarantine IP or interface.",
							Computed:    true,
						},
						"quarantine_expiry": {
							Type:        schema.TypeInt,
							Description: "Duration of quarantine in minute.",
							Computed:    true,
						},
						"quarantine_log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of selected quarantine.",
							Computed:    true,
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Vulnerability severity filter.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Selected rules status.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Sensor name.",
				Required:    true,
			},
			"override": {
				Type:        schema.TypeList,
				Description: "IPS override rule.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action of override rule.",
							Computed:    true,
						},
						"exempt_ip": {
							Type:        schema.TypeList,
							Description: "Exempted IP.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": {
										Type:        schema.TypeString,
										Description: "Destination IP address and netmask.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Exempt IP ID.",
										Computed:    true,
									},
									"src_ip": {
										Type:        schema.TypeString,
										Description: "Source IP address and netmask.",
										Computed:    true,
									},
								},
							},
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"log_packet": {
							Type:        schema.TypeString,
							Description: "Enable/disable packet logging.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Quarantine IP or interface.",
							Computed:    true,
						},
						"quarantine_expiry": {
							Type:        schema.TypeInt,
							Description: "Duration of quarantine in minute.",
							Computed:    true,
						},
						"quarantine_log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of selected quarantine.",
							Computed:    true,
						},
						"rule_id": {
							Type:        schema.TypeInt,
							Description: "Override rule ID.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable status of override rule.",
							Computed:    true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"scan_botnet_connections": {
				Type:        schema.TypeString,
				Description: "Block or monitor connections to Botnet servers, or disable Botnet scanning.",
				Computed:    true,
			},
		},
	}
}

func dataSourceIpsSensorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsSensor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsSensor dataSource: %v", err)
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

	diags := refreshObjectIpsSensor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
