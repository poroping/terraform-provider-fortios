// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure sniffer.

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

func resourceFirewallSniffer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure sniffer.",

		CreateContext: resourceFirewallSnifferCreate,
		ReadContext:   resourceFirewallSnifferRead,
		UpdateContext: resourceFirewallSnifferUpdate,
		DeleteContext: resourceFirewallSnifferDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"anomaly": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Denial of Service (DoS) anomaly settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "block"}, false),

							Description: "Action taken when the threshold is reached.",
							Optional:    true,
							Computed:    true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable anomaly logging.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Anomaly name.",
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
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable this anomaly.",
							Optional:    true,
							Computed:    true,
						},
						"threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 2147483647),

							Description: "Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.",
							Optional:    true,
							Computed:    true,
						},
						"thresholddefault": {
							Type: schema.TypeInt,

							Description: "Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"application_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing application list.",
				Optional:    true,
				Computed:    true,
			},
			"application_list_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable application control profile.",
				Optional:    true,
				Computed:    true,
			},
			"av_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing antivirus profile.",
				Optional:    true,
				Computed:    true,
			},
			"av_profile_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable antivirus profile.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DLP sensor.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_sensor_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DLP sensor.",
				Optional:    true,
				Computed:    true,
			},
			"dsri": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DSRI.",
				Optional:    true,
				Computed:    true,
			},
			"emailfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing email filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"emailfilter_profile_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable emailfilter.",
				Optional:    true,
				Computed:    true,
			},
			"file_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing file-filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"file_filter_profile_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable file filter.",
				Optional:    true,
				Computed:    true,
			},
			"host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9999),

				Description: "Sniffer ID (0 - 9999).",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface name that traffic sniffing will take place on.",
				Optional:    true,
				Computed:    true,
			},
			"ip_threatfeed": {
				Type:        schema.TypeList,
				Description: "Name of an existing IP threat feed.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Threat feed name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip_threatfeed_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IP threat feed.",
				Optional:    true,
				Computed:    true,
			},
			"ips_dos_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPS DoS anomaly detection.",
				Optional:    true,
				Computed:    true,
			},
			"ips_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing IPS sensor.",
				Optional:    true,
				Computed:    true,
			},
			"ips_sensor_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPS sensor.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sniffing IPv6 packets.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "utm", "disable"}, false),

				Description: "Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy.",
				Optional:    true,
				Computed:    true,
			},
			"max_packet_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000000),

				Description: "Maximum packet count (1 - 1000000, default = 4000).",
				Optional:    true,
				Computed:    true,
			},
			"non_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sniffing non-IP packets.",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Integer value for the protocol type as defined by IANA (0 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the active status of the sniffer.",
				Optional:    true,
				Computed:    true,
			},
			"vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "List of VLANs to sniff.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing web filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_profile_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable web filter profile.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallSnifferCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallSniffer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallSniffer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallSniffer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSniffer")
	}

	return resourceFirewallSnifferRead(ctx, d, meta)
}

func resourceFirewallSnifferUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSniffer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallSniffer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallSniffer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSniffer")
	}

	return resourceFirewallSnifferRead(ctx, d, meta)
}

func resourceFirewallSnifferDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallSniffer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallSniffer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSnifferRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSniffer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSniffer resource: %v", err)
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

	diags := refreshObjectFirewallSniffer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallSnifferAnomaly(v *[]models.FirewallSnifferAnomaly, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
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

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Threshold; tmp != nil {
				v["threshold"] = *tmp
			}

			if tmp := cfg.Thresholddefault; tmp != nil {
				v["thresholddefault"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallSnifferIpThreatfeed(v *[]models.FirewallSnifferIpThreatfeed, sort bool) interface{} {
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

func refreshObjectFirewallSniffer(d *schema.ResourceData, o *models.FirewallSniffer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Anomaly != nil {
		if err = d.Set("anomaly", flattenFirewallSnifferAnomaly(o.Anomaly, sort)); err != nil {
			return diag.Errorf("error reading anomaly: %v", err)
		}
	}

	if o.ApplicationList != nil {
		v := *o.ApplicationList

		if err = d.Set("application_list", v); err != nil {
			return diag.Errorf("error reading application_list: %v", err)
		}
	}

	if o.ApplicationListStatus != nil {
		v := *o.ApplicationListStatus

		if err = d.Set("application_list_status", v); err != nil {
			return diag.Errorf("error reading application_list_status: %v", err)
		}
	}

	if o.AvProfile != nil {
		v := *o.AvProfile

		if err = d.Set("av_profile", v); err != nil {
			return diag.Errorf("error reading av_profile: %v", err)
		}
	}

	if o.AvProfileStatus != nil {
		v := *o.AvProfileStatus

		if err = d.Set("av_profile_status", v); err != nil {
			return diag.Errorf("error reading av_profile_status: %v", err)
		}
	}

	if o.DlpSensor != nil {
		v := *o.DlpSensor

		if err = d.Set("dlp_sensor", v); err != nil {
			return diag.Errorf("error reading dlp_sensor: %v", err)
		}
	}

	if o.DlpSensorStatus != nil {
		v := *o.DlpSensorStatus

		if err = d.Set("dlp_sensor_status", v); err != nil {
			return diag.Errorf("error reading dlp_sensor_status: %v", err)
		}
	}

	if o.Dsri != nil {
		v := *o.Dsri

		if err = d.Set("dsri", v); err != nil {
			return diag.Errorf("error reading dsri: %v", err)
		}
	}

	if o.EmailfilterProfile != nil {
		v := *o.EmailfilterProfile

		if err = d.Set("emailfilter_profile", v); err != nil {
			return diag.Errorf("error reading emailfilter_profile: %v", err)
		}
	}

	if o.EmailfilterProfileStatus != nil {
		v := *o.EmailfilterProfileStatus

		if err = d.Set("emailfilter_profile_status", v); err != nil {
			return diag.Errorf("error reading emailfilter_profile_status: %v", err)
		}
	}

	if o.FileFilterProfile != nil {
		v := *o.FileFilterProfile

		if err = d.Set("file_filter_profile", v); err != nil {
			return diag.Errorf("error reading file_filter_profile: %v", err)
		}
	}

	if o.FileFilterProfileStatus != nil {
		v := *o.FileFilterProfileStatus

		if err = d.Set("file_filter_profile_status", v); err != nil {
			return diag.Errorf("error reading file_filter_profile_status: %v", err)
		}
	}

	if o.Host != nil {
		v := *o.Host

		if err = d.Set("host", v); err != nil {
			return diag.Errorf("error reading host: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.IpThreatfeed != nil {
		if err = d.Set("ip_threatfeed", flattenFirewallSnifferIpThreatfeed(o.IpThreatfeed, sort)); err != nil {
			return diag.Errorf("error reading ip_threatfeed: %v", err)
		}
	}

	if o.IpThreatfeedStatus != nil {
		v := *o.IpThreatfeedStatus

		if err = d.Set("ip_threatfeed_status", v); err != nil {
			return diag.Errorf("error reading ip_threatfeed_status: %v", err)
		}
	}

	if o.IpsDosStatus != nil {
		v := *o.IpsDosStatus

		if err = d.Set("ips_dos_status", v); err != nil {
			return diag.Errorf("error reading ips_dos_status: %v", err)
		}
	}

	if o.IpsSensor != nil {
		v := *o.IpsSensor

		if err = d.Set("ips_sensor", v); err != nil {
			return diag.Errorf("error reading ips_sensor: %v", err)
		}
	}

	if o.IpsSensorStatus != nil {
		v := *o.IpsSensorStatus

		if err = d.Set("ips_sensor_status", v); err != nil {
			return diag.Errorf("error reading ips_sensor_status: %v", err)
		}
	}

	if o.Ipv6 != nil {
		v := *o.Ipv6

		if err = d.Set("ipv6", v); err != nil {
			return diag.Errorf("error reading ipv6: %v", err)
		}
	}

	if o.Logtraffic != nil {
		v := *o.Logtraffic

		if err = d.Set("logtraffic", v); err != nil {
			return diag.Errorf("error reading logtraffic: %v", err)
		}
	}

	if o.MaxPacketCount != nil {
		v := *o.MaxPacketCount

		if err = d.Set("max_packet_count", v); err != nil {
			return diag.Errorf("error reading max_packet_count: %v", err)
		}
	}

	if o.NonIp != nil {
		v := *o.NonIp

		if err = d.Set("non_ip", v); err != nil {
			return diag.Errorf("error reading non_ip: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Vlan != nil {
		v := *o.Vlan

		if err = d.Set("vlan", v); err != nil {
			return diag.Errorf("error reading vlan: %v", err)
		}
	}

	if o.WebfilterProfile != nil {
		v := *o.WebfilterProfile

		if err = d.Set("webfilter_profile", v); err != nil {
			return diag.Errorf("error reading webfilter_profile: %v", err)
		}
	}

	if o.WebfilterProfileStatus != nil {
		v := *o.WebfilterProfileStatus

		if err = d.Set("webfilter_profile_status", v); err != nil {
			return diag.Errorf("error reading webfilter_profile_status: %v", err)
		}
	}

	return nil
}

func expandFirewallSnifferAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSnifferAnomaly, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSnifferAnomaly

	for i := range l {
		tmp := models.FirewallSnifferAnomaly{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
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

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Threshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.thresholddefault", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Thresholddefault = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSnifferIpThreatfeed(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSnifferIpThreatfeed, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSnifferIpThreatfeed

	for i := range l {
		tmp := models.FirewallSnifferIpThreatfeed{}
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

func getObjectFirewallSniffer(d *schema.ResourceData, sv string) (*models.FirewallSniffer, diag.Diagnostics) {
	obj := models.FirewallSniffer{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("anomaly"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("anomaly", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSnifferAnomaly(d, v, "anomaly", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Anomaly = t
		}
	} else if d.HasChange("anomaly") {
		old, new := d.GetChange("anomaly")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Anomaly = &[]models.FirewallSnifferAnomaly{}
		}
	}
	if v1, ok := d.GetOk("application_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application_list", sv)
				diags = append(diags, e)
			}
			obj.ApplicationList = &v2
		}
	}
	if v1, ok := d.GetOk("application_list_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application_list_status", sv)
				diags = append(diags, e)
			}
			obj.ApplicationListStatus = &v2
		}
	}
	if v1, ok := d.GetOk("av_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_profile", sv)
				diags = append(diags, e)
			}
			obj.AvProfile = &v2
		}
	}
	if v1, ok := d.GetOk("av_profile_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_profile_status", sv)
				diags = append(diags, e)
			}
			obj.AvProfileStatus = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dlp_sensor", sv)
				diags = append(diags, e)
			}
			obj.DlpSensor = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_sensor_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dlp_sensor_status", sv)
				diags = append(diags, e)
			}
			obj.DlpSensorStatus = &v2
		}
	}
	if v1, ok := d.GetOk("dsri"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dsri", sv)
				diags = append(diags, e)
			}
			obj.Dsri = &v2
		}
	}
	if v1, ok := d.GetOk("emailfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("emailfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.EmailfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("emailfilter_profile_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("emailfilter_profile_status", sv)
				diags = append(diags, e)
			}
			obj.EmailfilterProfileStatus = &v2
		}
	}
	if v1, ok := d.GetOk("file_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("file_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.FileFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("file_filter_profile_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("file_filter_profile_status", sv)
				diags = append(diags, e)
			}
			obj.FileFilterProfileStatus = &v2
		}
	}
	if v1, ok := d.GetOk("host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host", sv)
				diags = append(diags, e)
			}
			obj.Host = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v, ok := d.GetOk("ip_threatfeed"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("ip_threatfeed", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSnifferIpThreatfeed(d, v, "ip_threatfeed", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpThreatfeed = t
		}
	} else if d.HasChange("ip_threatfeed") {
		old, new := d.GetChange("ip_threatfeed")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpThreatfeed = &[]models.FirewallSnifferIpThreatfeed{}
		}
	}
	if v1, ok := d.GetOk("ip_threatfeed_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ip_threatfeed_status", sv)
				diags = append(diags, e)
			}
			obj.IpThreatfeedStatus = &v2
		}
	}
	if v1, ok := d.GetOk("ips_dos_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_dos_status", sv)
				diags = append(diags, e)
			}
			obj.IpsDosStatus = &v2
		}
	}
	if v1, ok := d.GetOk("ips_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_sensor", sv)
				diags = append(diags, e)
			}
			obj.IpsSensor = &v2
		}
	}
	if v1, ok := d.GetOk("ips_sensor_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_sensor_status", sv)
				diags = append(diags, e)
			}
			obj.IpsSensorStatus = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6", sv)
				diags = append(diags, e)
			}
			obj.Ipv6 = &v2
		}
	}
	if v1, ok := d.GetOk("logtraffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logtraffic", sv)
				diags = append(diags, e)
			}
			obj.Logtraffic = &v2
		}
	}
	if v1, ok := d.GetOk("max_packet_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_packet_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxPacketCount = &tmp
		}
	}
	if v1, ok := d.GetOk("non_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("non_ip", sv)
				diags = append(diags, e)
			}
			obj.NonIp = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			obj.Port = &v2
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan", sv)
				diags = append(diags, e)
			}
			obj.Vlan = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.WebfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_profile_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_profile_status", sv)
				diags = append(diags, e)
			}
			obj.WebfilterProfileStatus = &v2
		}
	}
	return &obj, diags
}
