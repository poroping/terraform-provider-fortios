// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS sensor.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceIpsSensor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS sensor.",

		CreateContext: resourceIpsSensorCreate,
		ReadContext:   resourceIpsSensorRead,
		UpdateContext: resourceIpsSensorUpdate,
		DeleteContext: resourceIpsSensorDelete,

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
			"block_malicious_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable malicious URL blocking.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "IPS sensor filter.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "block", "reset", "default"}, false),

							Description: "Action taken with traffic in which signatures are detected.",
							Optional:    true,
							Computed:    true,
						},
						"application": {
							Type: schema.TypeString,

							Description: "Applications to be protected. set application ? lists available applications. all includes all applications. other includes all unlisted applications.",
							Optional:    true,
							Computed:    true,
						},
						"cve": {
							Type:        schema.TypeList,
							Description: "List of CVE IDs of the signatures to add to the sensor",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cve_entry": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 19),

										Description: "CVE IDs or CVE wildcards.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"exempt_ip": {
							Type:        schema.TypeList,
							Description: "Traffic from selected source or destination IP addresses is exempt from this signature.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4Classnet,

										Description: "Destination IP address and netmask (applies to packet matching the signature).",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Exempt IP ID.",
										Optional:    true,
										Computed:    true,
									},
									"src_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4Classnet,

										Description: "Source IP address and netmask (applies to packet matching the signature).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Rule ID in IPS database (0 - 4294967295).",
							Optional:    true,
							Computed:    true,
						},
						"location": {
							Type: schema.TypeString,

							Description: "Protect client or server traffic.",
							Optional:    true,
							Computed:    true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of signatures included in filter.",
							Optional:    true,
							Computed:    true,
						},
						"log_attack_context": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of attack context: URL buffer, header buffer, body buffer, packet buffer.",
							Optional:    true,
							Computed:    true,
						},
						"log_packet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable packet logging. Enable to save the packet that triggers the filter. You can download the packets in pcap format for diagnostic use.",
							Optional:    true,
							Computed:    true,
						},
						"os": {
							Type: schema.TypeString,

							Description: "Operating systems to be protected.  all includes all operating systems. other includes all unlisted operating systems.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type: schema.TypeString,

							Description: "Protocols to be examined. set protocol ? lists available protocols. all includes all protocols. other includes all unlisted protocols.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "attacker"}, false),

							Description: "Quarantine method.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_expiry": {
							Type: schema.TypeString,

							Description: "Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine logging.",
							Optional:    true,
							Computed:    true,
						},
						"rate_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Count of the rate.",
							Optional:    true,
							Computed:    true,
						},
						"rate_duration": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Duration (sec) of the rate.",
							Optional:    true,
							Computed:    true,
						},
						"rate_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"periodical", "continuous"}, false),

							Description: "Rate limit mode.",
							Optional:    true,
							Computed:    true,
						},
						"rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip", "dhcp-client-mac", "dns-domain"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"rule": {
							Type:        schema.TypeList,
							Description: "Identifies the predefined or custom IPS signatures to add to the sensor.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Rule IPS.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"severity": {
							Type: schema.TypeString,

							Description: "Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "default"}, false),

							Description: "Status of the signatures included in filter. default enables the filter and only use filters with default status of enable. Filters with default status of disable will not be used.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"extended_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable extended logging.",
				Optional:    true,
				Computed:    true,
			},
			"filter": {
				Type:        schema.TypeList,
				Description: "IPS sensor filter.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "block", "reset", "default"}, false),

							Description: "Action of selected rules.",
							Optional:    true,
							Computed:    true,
						},
						"application": {
							Type: schema.TypeString,

							Description: "Vulnerable application filter.",
							Optional:    true,
							Computed:    true,
						},
						"location": {
							Type: schema.TypeString,

							Description: "Vulnerability location filter.",
							Optional:    true,
							Computed:    true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of selected rules.",
							Optional:    true,
							Computed:    true,
						},
						"log_packet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable packet logging of selected rules.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Filter name.",
							Optional:    true,
							Computed:    true,
						},
						"os": {
							Type: schema.TypeString,

							Description: "Vulnerable OS filter.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type: schema.TypeString,

							Description: "Vulnerable protocol filter.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "attacker"}, false),

							Description: "Quarantine IP or interface.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_expiry": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 2147483647),

							Description: "Duration of quarantine in minute.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of selected quarantine.",
							Optional:    true,
							Computed:    true,
						},
						"severity": {
							Type: schema.TypeString,

							Description: "Vulnerability severity filter.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "default"}, false),

							Description: "Selected rules status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Sensor name.",
				ForceNew:    true,
				Required:    true,
			},
			"override": {
				Type:        schema.TypeList,
				Description: "IPS override rule.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "block", "reset"}, false),

							Description: "Action of override rule.",
							Optional:    true,
							Computed:    true,
						},
						"exempt_ip": {
							Type:        schema.TypeList,
							Description: "Exempted IP.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4Classnet,

										Description: "Destination IP address and netmask.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Exempt IP ID.",
										Optional:    true,
										Computed:    true,
									},
									"src_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4Classnet,

										Description: "Source IP address and netmask.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"log_packet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable packet logging.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "attacker"}, false),

							Description: "Quarantine IP or interface.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_expiry": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 2147483647),

							Description: "Duration of quarantine in minute.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of selected quarantine.",
							Optional:    true,
							Computed:    true,
						},
						"rule_id": {
							Type: schema.TypeInt,

							Description: "Override rule ID.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable status of override rule.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group.",
				Optional:    true,
				Computed:    true,
			},
			"scan_botnet_connections": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "block", "monitor"}, false),

				Description: "Block or monitor connections to Botnet servers, or disable Botnet scanning.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceIpsSensorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating IpsSensor resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectIpsSensor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIpsSensor(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsSensor")
	}

	return resourceIpsSensorRead(ctx, d, meta)
}

func resourceIpsSensorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIpsSensor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIpsSensor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IpsSensor resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsSensor")
	}

	return resourceIpsSensorRead(ctx, d, meta)
}

func resourceIpsSensorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteIpsSensor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IpsSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSensorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsSensor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsSensor resource: %v", err)
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

	diags := refreshObjectIpsSensor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenIpsSensorEntries(v *[]models.IpsSensorEntries, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Application; tmp != nil {
				v["application"] = *tmp
			}

			if tmp := cfg.Cve; tmp != nil {
				v["cve"] = flattenIpsSensorEntriesCve(tmp, sort)
			}

			if tmp := cfg.ExemptIp; tmp != nil {
				v["exempt_ip"] = flattenIpsSensorEntriesExemptIp(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Location; tmp != nil {
				v["location"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogAttackContext; tmp != nil {
				v["log_attack_context"] = *tmp
			}

			if tmp := cfg.LogPacket; tmp != nil {
				v["log_packet"] = *tmp
			}

			if tmp := cfg.Os; tmp != nil {
				v["os"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			if tmp := cfg.QuarantineExpiry; tmp != nil {
				v["quarantine_expiry"] = *tmp
			}

			if tmp := cfg.QuarantineLog; tmp != nil {
				v["quarantine_log"] = *tmp
			}

			if tmp := cfg.RateCount; tmp != nil {
				v["rate_count"] = *tmp
			}

			if tmp := cfg.RateDuration; tmp != nil {
				v["rate_duration"] = *tmp
			}

			if tmp := cfg.RateMode; tmp != nil {
				v["rate_mode"] = *tmp
			}

			if tmp := cfg.RateTrack; tmp != nil {
				v["rate_track"] = *tmp
			}

			if tmp := cfg.Rule; tmp != nil {
				v["rule"] = flattenIpsSensorEntriesRule(tmp, sort)
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenIpsSensorEntriesCve(v *[]models.IpsSensorEntriesCve, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CveEntry; tmp != nil {
				v["cve_entry"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "cve_entry")
	}

	return flat
}

func flattenIpsSensorEntriesExemptIp(v *[]models.IpsSensorEntriesExemptIp, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.DstIp; tmp != nil {
				v["dst_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.SrcIp; tmp != nil {
				v["src_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenIpsSensorEntriesRule(v *[]models.IpsSensorEntriesRule, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenIpsSensorFilter(v *[]models.IpsSensorFilter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Application; tmp != nil {
				v["application"] = *tmp
			}

			if tmp := cfg.Location; tmp != nil {
				v["location"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogPacket; tmp != nil {
				v["log_packet"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Os; tmp != nil {
				v["os"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			if tmp := cfg.QuarantineExpiry; tmp != nil {
				v["quarantine_expiry"] = *tmp
			}

			if tmp := cfg.QuarantineLog; tmp != nil {
				v["quarantine_log"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenIpsSensorOverride(v *[]models.IpsSensorOverride, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.ExemptIp; tmp != nil {
				v["exempt_ip"] = flattenIpsSensorOverrideExemptIp(tmp, sort)
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogPacket; tmp != nil {
				v["log_packet"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			if tmp := cfg.QuarantineExpiry; tmp != nil {
				v["quarantine_expiry"] = *tmp
			}

			if tmp := cfg.QuarantineLog; tmp != nil {
				v["quarantine_log"] = *tmp
			}

			if tmp := cfg.RuleId; tmp != nil {
				v["rule_id"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "rule_id")
	}

	return flat
}

func flattenIpsSensorOverrideExemptIp(v *[]models.IpsSensorOverrideExemptIp, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.DstIp; tmp != nil {
				v["dst_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.SrcIp; tmp != nil {
				v["src_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectIpsSensor(d *schema.ResourceData, o *models.IpsSensor, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BlockMaliciousUrl != nil {
		v := *o.BlockMaliciousUrl

		if err = d.Set("block_malicious_url", v); err != nil {
			return diag.Errorf("error reading block_malicious_url: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Entries != nil {
		if err = d.Set("entries", flattenIpsSensorEntries(o.Entries, sort)); err != nil {
			return diag.Errorf("error reading entries: %v", err)
		}
	}

	if o.ExtendedLog != nil {
		v := *o.ExtendedLog

		if err = d.Set("extended_log", v); err != nil {
			return diag.Errorf("error reading extended_log: %v", err)
		}
	}

	if o.Filter != nil {
		if err = d.Set("filter", flattenIpsSensorFilter(o.Filter, sort)); err != nil {
			return diag.Errorf("error reading filter: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Override != nil {
		if err = d.Set("override", flattenIpsSensorOverride(o.Override, sort)); err != nil {
			return diag.Errorf("error reading override: %v", err)
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.ScanBotnetConnections != nil {
		v := *o.ScanBotnetConnections

		if err = d.Set("scan_botnet_connections", v); err != nil {
			return diag.Errorf("error reading scan_botnet_connections: %v", err)
		}
	}

	return nil
}

func expandIpsSensorEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsSensorEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsSensorEntries

	for i := range l {
		tmp := models.IpsSensorEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.application", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Application = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cve", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandIpsSensorEntriesCve(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.IpsSensorEntriesCve
			// 	}
			tmp.Cve = v2

		}

		pre_append = fmt.Sprintf("%s.%d.exempt_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandIpsSensorEntriesExemptIp(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.IpsSensorEntriesExemptIp
			// 	}
			tmp.ExemptIp = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.location", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Location = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_attack_context", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogAttackContext = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_packet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogPacket = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.os", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Os = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineExpiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RateCount = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_duration", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RateDuration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rule", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandIpsSensorEntriesRule(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.IpsSensorEntriesRule
			// 	}
			tmp.Rule = v2

		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIpsSensorEntriesCve(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsSensorEntriesCve, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsSensorEntriesCve

	for i := range l {
		tmp := models.IpsSensorEntriesCve{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cve_entry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CveEntry = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIpsSensorEntriesExemptIp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsSensorEntriesExemptIp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsSensorEntriesExemptIp

	for i := range l {
		tmp := models.IpsSensorEntriesExemptIp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DstIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SrcIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIpsSensorEntriesRule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsSensorEntriesRule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsSensorEntriesRule

	for i := range l {
		tmp := models.IpsSensorEntriesRule{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIpsSensorFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsSensorFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsSensorFilter

	for i := range l {
		tmp := models.IpsSensorFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.application", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Application = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.location", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Location = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_packet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogPacket = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.os", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Os = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.QuarantineExpiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIpsSensorOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsSensorOverride, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsSensorOverride

	for i := range l {
		tmp := models.IpsSensorOverride{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.exempt_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandIpsSensorOverrideExemptIp(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.IpsSensorOverrideExemptIp
			// 	}
			tmp.ExemptIp = v2

		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_packet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogPacket = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.QuarantineExpiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rule_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RuleId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandIpsSensorOverrideExemptIp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsSensorOverrideExemptIp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsSensorOverrideExemptIp

	for i := range l {
		tmp := models.IpsSensorOverrideExemptIp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DstIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SrcIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectIpsSensor(d *schema.ResourceData, sv string) (*models.IpsSensor, diag.Diagnostics) {
	obj := models.IpsSensor{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("block_malicious_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("block_malicious_url", sv)
				diags = append(diags, e)
			}
			obj.BlockMaliciousUrl = &v2
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v, ok := d.GetOk("entries"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("entries", sv)
			diags = append(diags, e)
		}
		t, err := expandIpsSensorEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entries = t
		}
	} else if d.HasChange("entries") {
		old, new := d.GetChange("entries")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entries = &[]models.IpsSensorEntries{}
		}
	}
	if v1, ok := d.GetOk("extended_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extended_log", sv)
				diags = append(diags, e)
			}
			obj.ExtendedLog = &v2
		}
	}
	if v, ok := d.GetOk("filter"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("filter", sv)
			diags = append(diags, e)
		}
		t, err := expandIpsSensorFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Filter = t
		}
	} else if d.HasChange("filter") {
		old, new := d.GetChange("filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Filter = &[]models.IpsSensorFilter{}
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
	if v, ok := d.GetOk("override"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("override", sv)
			diags = append(diags, e)
		}
		t, err := expandIpsSensorOverride(d, v, "override", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Override = t
		}
	} else if d.HasChange("override") {
		old, new := d.GetChange("override")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Override = &[]models.IpsSensorOverride{}
		}
	}
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("scan_botnet_connections"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scan_botnet_connections", sv)
				diags = append(diags, e)
			}
			obj.ScanBotnetConnections = &v2
		}
	}
	return &obj, diags
}
