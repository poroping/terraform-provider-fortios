// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemsdwanHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemsdwanHealthCheckCreate,
		Read:   resourceSystemsdwanHealthCheckRead,
		Update: resourceSystemsdwanHealthCheckUpdate,
		Delete: resourceSystemsdwanHealthCheckDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
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
			"addr_mode": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"ipv4", "ipv6"}),

				Description: "Address mode (IPv4 or IPv6).",
				Optional:    true,
				Computed:    true,
			},
			"detect_mode": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"active", "passive", "prefer-passive"}),

				Description: "The mode determining how to detect the server.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode": {
				Type: schema.TypeString,

				Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
				Optional:    true,
				Computed:    true,
			},
			"dns_match_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Response IP expected from DNS server if the protocol is DNS.",
				Optional:    true,
				Computed:    true,
			},
			"dns_request_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fully qualified domain name to resolve for the DNS probe.",
				Optional:    true,
				Computed:    true,
			},
			"failtime": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Number of failures before server is considered lost (1 - 3600, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"ftp_file": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 254),

				Description: "Full path and file name on the FTP server to download for FTP health-check to probe.",
				Optional:    true,
				Computed:    true,
			},
			"ftp_mode": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"passive", "port"}),

				Description: "FTP mode.",
				Optional:    true,
				Computed:    true,
			},
			"ha_priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),

				Description: "HA election priority (1 - 50).",
				Optional:    true,
				Computed:    true,
			},
			"http_agent": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "String in the http-agent field in the HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"http_get": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "URL used to communicate with the server if the protocol if the protocol is HTTP.",
				Optional:    true,
				Computed:    true,
			},
			"http_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "Response string expected from the server if the protocol is HTTP.",
				Optional:    true,
				Computed:    true,
			},
			"interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 3600000),

				Description: "Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"members": {
				Type:        schema.TypeList,
				Description: "Member sequence number list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt,

							Description: "Member sequence number.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Status check or health check name.",
				ForceNew:    true,
				Required:    true,
			},
			"packet_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 1024),

				Description: "Packet size of a twamp test session,",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Twamp controller password in authentication mode",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port number used to communicate with the server over the selected protocol (0-65535, default = 0, auto select. http, twamp: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21).",
				Optional:    true,
				Computed:    true,
			},
			"probe_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 30),

				Description: "Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"probe_packets": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable transmission of probe packets.",
				Optional:    true,
				Computed:    true,
			},
			"probe_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 3600000),

				Description: "Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"ping", "tcp-echo", "udp-echo", "http", "twamp", "dns", "tcp-connect", "ftp"}),

				Description: "Protocol used to determine if the FortiGate can communicate with the server.",
				Optional:    true,
				Computed:    true,
			},
			"quality_measured_method": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"half-open", "half-close"}),

				Description: "Method to measure the quality of tcp-connect.",
				Optional:    true,
				Computed:    true,
			},
			"recoverytime": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Number of successful responses received before server is considered recovered (1 - 3600, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"security_mode": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"none", "authentication"}),

				Description: "Twamp controller security mode.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IP address or FQDN name of the server.",
				Optional:    true,
				Computed:    true,
			},
			"sla": {
				Type:        schema.TypeList,
				Description: "Service level agreement (SLA).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "SLA ID.",
							Optional:    true,
							Computed:    true,
						},
						"jitter_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
							Optional:    true,
							Computed:    true,
						},
						"latency_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
							Optional:    true,
							Computed:    true,
						},
						"link_cost_factor": {
							Type: schema.TypeString,

							DiffSuppressFunc: diffFakeListEqual,
							Description:      "Criteria on which to base link selection.",
							Optional:         true,
							Computed:         true,
						},
						"packetloss_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"sla_fail_log_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"sla_pass_log_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"system_dns": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable system DNS as the probe server.",
				Optional:    true,
				Computed:    true,
			},
			"threshold_alert_jitter": {
				Type: schema.TypeInt,

				Description: "Alert threshold for jitter (ms, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_alert_latency": {
				Type: schema.TypeInt,

				Description: "Alert threshold for latency (ms, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_alert_packetloss": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Alert threshold for packet loss (percentage, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_warning_jitter": {
				Type: schema.TypeInt,

				Description: "Warning threshold for jitter (ms, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_warning_latency": {
				Type: schema.TypeInt,

				Description: "Warning threshold for latency (ms, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_warning_packetloss": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Warning threshold for packet loss (percentage, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"update_cascade_interface": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable update cascade interface.",
				Optional:    true,
				Computed:    true,
			},
			"update_static_route": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable updating the static route.",
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "The user name to access probe server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemsdwanHealthCheckCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectSystemsdwanHealthCheck(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating SystemsdwanHealthCheck resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating SystemsdwanHealthCheck resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateSystemsdwanHealthCheck(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateSystemsdwanHealthCheck(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating SystemsdwanHealthCheck resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemsdwanHealthCheck")
	}

	return resourceSystemsdwanHealthCheckRead(d, m)
}

func resourceSystemsdwanHealthCheckUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectSystemsdwanHealthCheck(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanHealthCheck resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemsdwanHealthCheck(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanHealthCheck resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemsdwanHealthCheck")
	}

	return resourceSystemsdwanHealthCheckRead(d, m)
}

func resourceSystemsdwanHealthCheckDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	err := c.DeleteSystemsdwanHealthCheck(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemsdwanHealthCheck resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemsdwanHealthCheckRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	o, err := c.ReadSystemsdwanHealthCheck(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanHealthCheck resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemsdwanHealthCheck(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanHealthCheck resource from API: %v", err)
	}
	return nil
}

func flattenSystemsdwanHealthCheckAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckDetectMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckDiffservcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckDnsMatchIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckFailtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckFtpFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckFtpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckHaPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckHttpAgent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckHttpGet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckHttpMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {

			tmp["seq_num"] = flattenSystemsdwanHealthCheckMembersSeqNum(i["seq-num"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seq-num", d)
	return result
}

func flattenSystemsdwanHealthCheckMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckPacketSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckProbeCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckProbePackets(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckProbeTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckQualityMeasuredMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckRecoverytime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSecurityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSla(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemsdwanHealthCheckSlaId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {

			tmp["jitter_threshold"] = flattenSystemsdwanHealthCheckSlaJitterThreshold(i["jitter-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {

			tmp["latency_threshold"] = flattenSystemsdwanHealthCheckSlaLatencyThreshold(i["latency-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {

			tmp["link_cost_factor"] = flattenSystemsdwanHealthCheckSlaLinkCostFactor(i["link-cost-factor"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {

			tmp["packetloss_threshold"] = flattenSystemsdwanHealthCheckSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemsdwanHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckSystemDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanHealthCheckUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemsdwanHealthCheck(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("addr_mode", flattenSystemsdwanHealthCheckAddrMode(o["addr-mode"], d, "addr_mode", sv)); err != nil {
		if !fortiAPIPatch(o["addr-mode"]) {
			return fmt.Errorf("error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("detect_mode", flattenSystemsdwanHealthCheckDetectMode(o["detect-mode"], d, "detect_mode", sv)); err != nil {
		if !fortiAPIPatch(o["detect-mode"]) {
			return fmt.Errorf("error reading detect_mode: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenSystemsdwanHealthCheckDiffservcode(o["diffservcode"], d, "diffservcode", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dns_match_ip", flattenSystemsdwanHealthCheckDnsMatchIp(o["dns-match-ip"], d, "dns_match_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dns-match-ip"]) {
			return fmt.Errorf("error reading dns_match_ip: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", flattenSystemsdwanHealthCheckDnsRequestDomain(o["dns-request-domain"], d, "dns_request_domain", sv)); err != nil {
		if !fortiAPIPatch(o["dns-request-domain"]) {
			return fmt.Errorf("error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("failtime", flattenSystemsdwanHealthCheckFailtime(o["failtime"], d, "failtime", sv)); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("error reading failtime: %v", err)
		}
	}

	if err = d.Set("ftp_file", flattenSystemsdwanHealthCheckFtpFile(o["ftp-file"], d, "ftp_file", sv)); err != nil {
		if !fortiAPIPatch(o["ftp-file"]) {
			return fmt.Errorf("error reading ftp_file: %v", err)
		}
	}

	if err = d.Set("ftp_mode", flattenSystemsdwanHealthCheckFtpMode(o["ftp-mode"], d, "ftp_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ftp-mode"]) {
			return fmt.Errorf("error reading ftp_mode: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenSystemsdwanHealthCheckHaPriority(o["ha-priority"], d, "ha_priority", sv)); err != nil {
		if !fortiAPIPatch(o["ha-priority"]) {
			return fmt.Errorf("error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("http_agent", flattenSystemsdwanHealthCheckHttpAgent(o["http-agent"], d, "http_agent", sv)); err != nil {
		if !fortiAPIPatch(o["http-agent"]) {
			return fmt.Errorf("error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_get", flattenSystemsdwanHealthCheckHttpGet(o["http-get"], d, "http_get", sv)); err != nil {
		if !fortiAPIPatch(o["http-get"]) {
			return fmt.Errorf("error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenSystemsdwanHealthCheckHttpMatch(o["http-match"], d, "http_match", sv)); err != nil {
		if !fortiAPIPatch(o["http-match"]) {
			return fmt.Errorf("error reading http_match: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemsdwanHealthCheckInterval(o["interval"], d, "interval", sv)); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("error reading interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenSystemsdwanHealthCheckMembers(o["members"], d, "members", sv)); err != nil {
			if !fortiAPIPatch(o["members"]) {
				return fmt.Errorf("error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenSystemsdwanHealthCheckMembers(o["members"], d, "members", sv)); err != nil {
				if !fortiAPIPatch(o["members"]) {
					return fmt.Errorf("error reading members: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemsdwanHealthCheckName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("packet_size", flattenSystemsdwanHealthCheckPacketSize(o["packet-size"], d, "packet_size", sv)); err != nil {
		if !fortiAPIPatch(o["packet-size"]) {
			return fmt.Errorf("error reading packet_size: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemsdwanHealthCheckPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("error reading port: %v", err)
		}
	}

	if err = d.Set("probe_count", flattenSystemsdwanHealthCheckProbeCount(o["probe-count"], d, "probe_count", sv)); err != nil {
		if !fortiAPIPatch(o["probe-count"]) {
			return fmt.Errorf("error reading probe_count: %v", err)
		}
	}

	if err = d.Set("probe_packets", flattenSystemsdwanHealthCheckProbePackets(o["probe-packets"], d, "probe_packets", sv)); err != nil {
		if !fortiAPIPatch(o["probe-packets"]) {
			return fmt.Errorf("error reading probe_packets: %v", err)
		}
	}

	if err = d.Set("probe_timeout", flattenSystemsdwanHealthCheckProbeTimeout(o["probe-timeout"], d, "probe_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["probe-timeout"]) {
			return fmt.Errorf("error reading probe_timeout: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemsdwanHealthCheckProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_measured_method", flattenSystemsdwanHealthCheckQualityMeasuredMethod(o["quality-measured-method"], d, "quality_measured_method", sv)); err != nil {
		if !fortiAPIPatch(o["quality-measured-method"]) {
			return fmt.Errorf("error reading quality_measured_method: %v", err)
		}
	}

	if err = d.Set("recoverytime", flattenSystemsdwanHealthCheckRecoverytime(o["recoverytime"], d, "recoverytime", sv)); err != nil {
		if !fortiAPIPatch(o["recoverytime"]) {
			return fmt.Errorf("error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemsdwanHealthCheckSecurityMode(o["security-mode"], d, "security_mode", sv)); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("error reading security_mode: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemsdwanHealthCheckServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("error reading server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenSystemsdwanHealthCheckSla(o["sla"], d, "sla", sv)); err != nil {
			if !fortiAPIPatch(o["sla"]) {
				return fmt.Errorf("error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenSystemsdwanHealthCheckSla(o["sla"], d, "sla", sv)); err != nil {
				if !fortiAPIPatch(o["sla"]) {
					return fmt.Errorf("error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_fail_log_period", flattenSystemsdwanHealthCheckSlaFailLogPeriod(o["sla-fail-log-period"], d, "sla_fail_log_period", sv)); err != nil {
		if !fortiAPIPatch(o["sla-fail-log-period"]) {
			return fmt.Errorf("error reading sla_fail_log_period: %v", err)
		}
	}

	if err = d.Set("sla_pass_log_period", flattenSystemsdwanHealthCheckSlaPassLogPeriod(o["sla-pass-log-period"], d, "sla_pass_log_period", sv)); err != nil {
		if !fortiAPIPatch(o["sla-pass-log-period"]) {
			return fmt.Errorf("error reading sla_pass_log_period: %v", err)
		}
	}

	if err = d.Set("system_dns", flattenSystemsdwanHealthCheckSystemDns(o["system-dns"], d, "system_dns", sv)); err != nil {
		if !fortiAPIPatch(o["system-dns"]) {
			return fmt.Errorf("error reading system_dns: %v", err)
		}
	}

	if err = d.Set("threshold_alert_jitter", flattenSystemsdwanHealthCheckThresholdAlertJitter(o["threshold-alert-jitter"], d, "threshold_alert_jitter", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-alert-jitter"]) {
			return fmt.Errorf("error reading threshold_alert_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_alert_latency", flattenSystemsdwanHealthCheckThresholdAlertLatency(o["threshold-alert-latency"], d, "threshold_alert_latency", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-alert-latency"]) {
			return fmt.Errorf("error reading threshold_alert_latency: %v", err)
		}
	}

	if err = d.Set("threshold_alert_packetloss", flattenSystemsdwanHealthCheckThresholdAlertPacketloss(o["threshold-alert-packetloss"], d, "threshold_alert_packetloss", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-alert-packetloss"]) {
			return fmt.Errorf("error reading threshold_alert_packetloss: %v", err)
		}
	}

	if err = d.Set("threshold_warning_jitter", flattenSystemsdwanHealthCheckThresholdWarningJitter(o["threshold-warning-jitter"], d, "threshold_warning_jitter", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-warning-jitter"]) {
			return fmt.Errorf("error reading threshold_warning_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_warning_latency", flattenSystemsdwanHealthCheckThresholdWarningLatency(o["threshold-warning-latency"], d, "threshold_warning_latency", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-warning-latency"]) {
			return fmt.Errorf("error reading threshold_warning_latency: %v", err)
		}
	}

	if err = d.Set("threshold_warning_packetloss", flattenSystemsdwanHealthCheckThresholdWarningPacketloss(o["threshold-warning-packetloss"], d, "threshold_warning_packetloss", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-warning-packetloss"]) {
			return fmt.Errorf("error reading threshold_warning_packetloss: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", flattenSystemsdwanHealthCheckUpdateCascadeInterface(o["update-cascade-interface"], d, "update_cascade_interface", sv)); err != nil {
		if !fortiAPIPatch(o["update-cascade-interface"]) {
			return fmt.Errorf("error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_static_route", flattenSystemsdwanHealthCheckUpdateStaticRoute(o["update-static-route"], d, "update_static_route", sv)); err != nil {
		if !fortiAPIPatch(o["update-static-route"]) {
			return fmt.Errorf("error reading update_static_route: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemsdwanHealthCheckUser(o["user"], d, "user", sv)); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("error reading user: %v", err)
		}
	}

	return nil
}

func expandSystemsdwanHealthCheckAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckDetectMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckDiffservcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckDnsMatchIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckFailtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckFtpFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckFtpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckHaPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckHttpAgent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckHttpGet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckHttpMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["seq-num"], _ = expandSystemsdwanHealthCheckMembersSeqNum(d, i["seq_num"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanHealthCheckMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckPacketSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckProbeCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckProbePackets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckProbeTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckQualityMeasuredMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckRecoverytime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSecurityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemsdwanHealthCheckSlaId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["jitter-threshold"], _ = expandSystemsdwanHealthCheckSlaJitterThreshold(d, i["jitter_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["latency-threshold"], _ = expandSystemsdwanHealthCheckSlaLatencyThreshold(d, i["latency_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["link-cost-factor"], _ = expandSystemsdwanHealthCheckSlaLinkCostFactor(d, i["link_cost_factor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["packetloss-threshold"], _ = expandSystemsdwanHealthCheckSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanHealthCheckSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckSystemDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanHealthCheckUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemsdwanHealthCheck(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok {
		t, err := expandSystemsdwanHealthCheckAddrMode(d, v, "addr_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("detect_mode"); ok {
		t, err := expandSystemsdwanHealthCheckDetectMode(d, v, "detect_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-mode"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok {
		t, err := expandSystemsdwanHealthCheckDiffservcode(d, v, "diffservcode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dns_match_ip"); ok {
		t, err := expandSystemsdwanHealthCheckDnsMatchIp(d, v, "dns_match_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-match-ip"] = t
		}
	}

	if v, ok := d.GetOk("dns_request_domain"); ok {
		t, err := expandSystemsdwanHealthCheckDnsRequestDomain(d, v, "dns_request_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-request-domain"] = t
		}
	}

	if v, ok := d.GetOk("failtime"); ok {
		t, err := expandSystemsdwanHealthCheckFailtime(d, v, "failtime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("ftp_file"); ok {
		t, err := expandSystemsdwanHealthCheckFtpFile(d, v, "ftp_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-file"] = t
		}
	}

	if v, ok := d.GetOk("ftp_mode"); ok {
		t, err := expandSystemsdwanHealthCheckFtpMode(d, v, "ftp_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-mode"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok {
		t, err := expandSystemsdwanHealthCheckHaPriority(d, v, "ha_priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("http_agent"); ok {
		t, err := expandSystemsdwanHealthCheckHttpAgent(d, v, "http_agent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-agent"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok {
		t, err := expandSystemsdwanHealthCheckHttpGet(d, v, "http_get", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok {
		t, err := expandSystemsdwanHealthCheckHttpMatch(d, v, "http_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		t, err := expandSystemsdwanHealthCheckInterval(d, v, "interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok {
		t, err := expandSystemsdwanHealthCheckMembers(d, v, "members", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	} else if d.HasChange("members") {
		old, new := d.GetChange("members")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["members"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemsdwanHealthCheckName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_size"); ok {
		t, err := expandSystemsdwanHealthCheckPacketSize(d, v, "packet_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemsdwanHealthCheckPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemsdwanHealthCheckPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("probe_count"); ok {
		t, err := expandSystemsdwanHealthCheckProbeCount(d, v, "probe_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-count"] = t
		}
	}

	if v, ok := d.GetOk("probe_packets"); ok {
		t, err := expandSystemsdwanHealthCheckProbePackets(d, v, "probe_packets", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-packets"] = t
		}
	}

	if v, ok := d.GetOk("probe_timeout"); ok {
		t, err := expandSystemsdwanHealthCheckProbeTimeout(d, v, "probe_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-timeout"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSystemsdwanHealthCheckProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_measured_method"); ok {
		t, err := expandSystemsdwanHealthCheckQualityMeasuredMethod(d, v, "quality_measured_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-measured-method"] = t
		}
	}

	if v, ok := d.GetOk("recoverytime"); ok {
		t, err := expandSystemsdwanHealthCheckRecoverytime(d, v, "recoverytime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok {
		t, err := expandSystemsdwanHealthCheckSecurityMode(d, v, "security_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemsdwanHealthCheckServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok {
		t, err := expandSystemsdwanHealthCheckSla(d, v, "sla", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	} else if d.HasChange("sla") {
		old, new := d.GetChange("sla")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["sla"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("sla_fail_log_period"); ok {
		t, err := expandSystemsdwanHealthCheckSlaFailLogPeriod(d, v, "sla_fail_log_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-fail-log-period"] = t
		}
	}

	if v, ok := d.GetOk("sla_pass_log_period"); ok {
		t, err := expandSystemsdwanHealthCheckSlaPassLogPeriod(d, v, "sla_pass_log_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-pass-log-period"] = t
		}
	}

	if v, ok := d.GetOk("system_dns"); ok {
		t, err := expandSystemsdwanHealthCheckSystemDns(d, v, "system_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-dns"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_jitter"); ok {
		t, err := expandSystemsdwanHealthCheckThresholdAlertJitter(d, v, "threshold_alert_jitter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_latency"); ok {
		t, err := expandSystemsdwanHealthCheckThresholdAlertLatency(d, v, "threshold_alert_latency", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_packetloss"); ok {
		t, err := expandSystemsdwanHealthCheckThresholdAlertPacketloss(d, v, "threshold_alert_packetloss", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_jitter"); ok {
		t, err := expandSystemsdwanHealthCheckThresholdWarningJitter(d, v, "threshold_warning_jitter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_latency"); ok {
		t, err := expandSystemsdwanHealthCheckThresholdWarningLatency(d, v, "threshold_warning_latency", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_packetloss"); ok {
		t, err := expandSystemsdwanHealthCheckThresholdWarningPacketloss(d, v, "threshold_warning_packetloss", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("update_cascade_interface"); ok {
		t, err := expandSystemsdwanHealthCheckUpdateCascadeInterface(d, v, "update_cascade_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-cascade-interface"] = t
		}
	}

	if v, ok := d.GetOk("update_static_route"); ok {
		t, err := expandSystemsdwanHealthCheckUpdateStaticRoute(d, v, "update_static_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-static-route"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {
		t, err := expandSystemsdwanHealthCheckUser(d, v, "user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
