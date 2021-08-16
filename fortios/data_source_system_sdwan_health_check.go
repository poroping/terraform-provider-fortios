// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSystemsdwanHealthCheck() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemsdwanHealthCheckRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_mode": {
				Type:        schema.TypeString,
				Description: "Address mode (IPv4 or IPv6).",
				Computed:    true,
			},
			"detect_mode": {
				Type:        schema.TypeString,
				Description: "The mode determining how to detect the server.",
				Computed:    true,
			},
			"diffservcode": {
				Type:        schema.TypeString,
				Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
				Computed:    true,
			},
			"dns_match_ip": {
				Type:        schema.TypeString,
				Description: "Response IP expected from DNS server if the protocol is DNS.",
				Computed:    true,
			},
			"dns_request_domain": {
				Type:        schema.TypeString,
				Description: "Fully qualified domain name to resolve for the DNS probe.",
				Computed:    true,
			},
			"failtime": {
				Type:        schema.TypeInt,
				Description: "Number of failures before server is considered lost (1 - 3600, default = 5).",
				Computed:    true,
			},
			"ftp_file": {
				Type:        schema.TypeString,
				Description: "Full path and file name on the FTP server to download for FTP health-check to probe.",
				Computed:    true,
			},
			"ftp_mode": {
				Type:        schema.TypeString,
				Description: "FTP mode.",
				Computed:    true,
			},
			"ha_priority": {
				Type:        schema.TypeInt,
				Description: "HA election priority (1 - 50).",
				Computed:    true,
			},
			"http_agent": {
				Type:        schema.TypeString,
				Description: "String in the http-agent field in the HTTP header.",
				Computed:    true,
			},
			"http_get": {
				Type:        schema.TypeString,
				Description: "URL used to communicate with the server if the protocol if the protocol is HTTP.",
				Computed:    true,
			},
			"http_match": {
				Type:        schema.TypeString,
				Description: "Response string expected from the server if the protocol is HTTP.",
				Computed:    true,
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).",
				Computed:    true,
			},
			"members": {
				Type:        schema.TypeList,
				Description: "Member sequence number list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type:        schema.TypeInt,
							Description: "Member sequence number.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Status check or health check name.",
				Required:    true,
			},
			"packet_size": {
				Type:        schema.TypeInt,
				Description: "Packet size of a twamp test session,",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Twamp controller password in authentication mode",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port number used to communicate with the server over the selected protocol (0-65535, default = 0, auto select. http, twamp: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21).",
				Computed:    true,
			},
			"probe_count": {
				Type:        schema.TypeInt,
				Description: "Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).",
				Computed:    true,
			},
			"probe_packets": {
				Type:        schema.TypeString,
				Description: "Enable/disable transmission of probe packets.",
				Computed:    true,
			},
			"probe_timeout": {
				Type:        schema.TypeInt,
				Description: "Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol used to determine if the FortiGate can communicate with the server.",
				Computed:    true,
			},
			"quality_measured_method": {
				Type:        schema.TypeString,
				Description: "Method to measure the quality of tcp-connect.",
				Computed:    true,
			},
			"recoverytime": {
				Type:        schema.TypeInt,
				Description: "Number of successful responses received before server is considered recovered (1 - 3600, default = 5).",
				Computed:    true,
			},
			"security_mode": {
				Type:        schema.TypeString,
				Description: "Twamp controller security mode.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "IP address or FQDN name of the server.",
				Computed:    true,
			},
			"sla": {
				Type:        schema.TypeList,
				Description: "Service level agreement (SLA).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "SLA ID.",
							Computed:    true,
						},
						"jitter_threshold": {
							Type:        schema.TypeInt,
							Description: "Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
							Computed:    true,
						},
						"latency_threshold": {
							Type:        schema.TypeInt,
							Description: "Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
							Computed:    true,
						},
						"link_cost_factor": {
							Type:        schema.TypeString,
							Description: "Criteria on which to base link selection.",
							Computed:    true,
						},
						"packetloss_threshold": {
							Type:        schema.TypeInt,
							Description: "Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).",
							Computed:    true,
						},
					},
				},
			},
			"sla_fail_log_period": {
				Type:        schema.TypeInt,
				Description: "Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).",
				Computed:    true,
			},
			"sla_pass_log_period": {
				Type:        schema.TypeInt,
				Description: "Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).",
				Computed:    true,
			},
			"system_dns": {
				Type:        schema.TypeString,
				Description: "Enable/disable system DNS as the probe server.",
				Computed:    true,
			},
			"threshold_alert_jitter": {
				Type:        schema.TypeInt,
				Description: "Alert threshold for jitter (ms, default = 0).",
				Computed:    true,
			},
			"threshold_alert_latency": {
				Type:        schema.TypeInt,
				Description: "Alert threshold for latency (ms, default = 0).",
				Computed:    true,
			},
			"threshold_alert_packetloss": {
				Type:        schema.TypeInt,
				Description: "Alert threshold for packet loss (percentage, default = 0).",
				Computed:    true,
			},
			"threshold_warning_jitter": {
				Type:        schema.TypeInt,
				Description: "Warning threshold for jitter (ms, default = 0).",
				Computed:    true,
			},
			"threshold_warning_latency": {
				Type:        schema.TypeInt,
				Description: "Warning threshold for latency (ms, default = 0).",
				Computed:    true,
			},
			"threshold_warning_packetloss": {
				Type:        schema.TypeInt,
				Description: "Warning threshold for packet loss (percentage, default = 0).",
				Computed:    true,
			},
			"update_cascade_interface": {
				Type:        schema.TypeString,
				Description: "Enable/disable update cascade interface.",
				Computed:    true,
			},
			"update_static_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable updating the static route.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "The user name to access probe server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemsdwanHealthCheckRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""
	key := "name"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing SystemsdwanHealthCheck: type error")
	}

	o, err := c.ReadSystemsdwanHealthCheck(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanHealthCheck: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemsdwanHealthCheck(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanHealthCheck from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemsdwanHealthCheckAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckDetectMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckDnsMatchIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckFtpFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckFtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["seq_num"] = dataSourceFlattenSystemsdwanHealthCheckMembersSeqNum(i["seq-num"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanHealthCheckMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckProbeCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckProbePackets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckProbeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckQualityMeasuredMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenSystemsdwanHealthCheckSlaId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {

			tmp["jitter_threshold"] = dataSourceFlattenSystemsdwanHealthCheckSlaJitterThreshold(i["jitter-threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {

			tmp["latency_threshold"] = dataSourceFlattenSystemsdwanHealthCheckSlaLatencyThreshold(i["latency-threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {

			tmp["link_cost_factor"] = dataSourceFlattenSystemsdwanHealthCheckSlaLinkCostFactor(i["link-cost-factor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {

			tmp["packetloss_threshold"] = dataSourceFlattenSystemsdwanHealthCheckSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckSystemDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanHealthCheckUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemsdwanHealthCheck(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_mode", dataSourceFlattenSystemsdwanHealthCheckAddrMode(o["addr-mode"], d, "addr_mode")); err != nil {
		if !fortiAPIPatch(o["addr-mode"]) {
			return fmt.Errorf("error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("detect_mode", dataSourceFlattenSystemsdwanHealthCheckDetectMode(o["detect-mode"], d, "detect_mode")); err != nil {
		if !fortiAPIPatch(o["detect-mode"]) {
			return fmt.Errorf("error reading detect_mode: %v", err)
		}
	}

	if err = d.Set("diffservcode", dataSourceFlattenSystemsdwanHealthCheckDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dns_match_ip", dataSourceFlattenSystemsdwanHealthCheckDnsMatchIp(o["dns-match-ip"], d, "dns_match_ip")); err != nil {
		if !fortiAPIPatch(o["dns-match-ip"]) {
			return fmt.Errorf("error reading dns_match_ip: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", dataSourceFlattenSystemsdwanHealthCheckDnsRequestDomain(o["dns-request-domain"], d, "dns_request_domain")); err != nil {
		if !fortiAPIPatch(o["dns-request-domain"]) {
			return fmt.Errorf("error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("failtime", dataSourceFlattenSystemsdwanHealthCheckFailtime(o["failtime"], d, "failtime")); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("error reading failtime: %v", err)
		}
	}

	if err = d.Set("ftp_file", dataSourceFlattenSystemsdwanHealthCheckFtpFile(o["ftp-file"], d, "ftp_file")); err != nil {
		if !fortiAPIPatch(o["ftp-file"]) {
			return fmt.Errorf("error reading ftp_file: %v", err)
		}
	}

	if err = d.Set("ftp_mode", dataSourceFlattenSystemsdwanHealthCheckFtpMode(o["ftp-mode"], d, "ftp_mode")); err != nil {
		if !fortiAPIPatch(o["ftp-mode"]) {
			return fmt.Errorf("error reading ftp_mode: %v", err)
		}
	}

	if err = d.Set("ha_priority", dataSourceFlattenSystemsdwanHealthCheckHaPriority(o["ha-priority"], d, "ha_priority")); err != nil {
		if !fortiAPIPatch(o["ha-priority"]) {
			return fmt.Errorf("error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("http_agent", dataSourceFlattenSystemsdwanHealthCheckHttpAgent(o["http-agent"], d, "http_agent")); err != nil {
		if !fortiAPIPatch(o["http-agent"]) {
			return fmt.Errorf("error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_get", dataSourceFlattenSystemsdwanHealthCheckHttpGet(o["http-get"], d, "http_get")); err != nil {
		if !fortiAPIPatch(o["http-get"]) {
			return fmt.Errorf("error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", dataSourceFlattenSystemsdwanHealthCheckHttpMatch(o["http-match"], d, "http_match")); err != nil {
		if !fortiAPIPatch(o["http-match"]) {
			return fmt.Errorf("error reading http_match: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemsdwanHealthCheckInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("error reading interval: %v", err)
		}
	}

	if err = d.Set("members", dataSourceFlattenSystemsdwanHealthCheckMembers(o["members"], d, "members")); err != nil {
		if !fortiAPIPatch(o["members"]) {
			return fmt.Errorf("error reading members: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemsdwanHealthCheckName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("packet_size", dataSourceFlattenSystemsdwanHealthCheckPacketSize(o["packet-size"], d, "packet_size")); err != nil {
		if !fortiAPIPatch(o["packet-size"]) {
			return fmt.Errorf("error reading packet_size: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemsdwanHealthCheckPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("error reading port: %v", err)
		}
	}

	if err = d.Set("probe_count", dataSourceFlattenSystemsdwanHealthCheckProbeCount(o["probe-count"], d, "probe_count")); err != nil {
		if !fortiAPIPatch(o["probe-count"]) {
			return fmt.Errorf("error reading probe_count: %v", err)
		}
	}

	if err = d.Set("probe_packets", dataSourceFlattenSystemsdwanHealthCheckProbePackets(o["probe-packets"], d, "probe_packets")); err != nil {
		if !fortiAPIPatch(o["probe-packets"]) {
			return fmt.Errorf("error reading probe_packets: %v", err)
		}
	}

	if err = d.Set("probe_timeout", dataSourceFlattenSystemsdwanHealthCheckProbeTimeout(o["probe-timeout"], d, "probe_timeout")); err != nil {
		if !fortiAPIPatch(o["probe-timeout"]) {
			return fmt.Errorf("error reading probe_timeout: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenSystemsdwanHealthCheckProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_measured_method", dataSourceFlattenSystemsdwanHealthCheckQualityMeasuredMethod(o["quality-measured-method"], d, "quality_measured_method")); err != nil {
		if !fortiAPIPatch(o["quality-measured-method"]) {
			return fmt.Errorf("error reading quality_measured_method: %v", err)
		}
	}

	if err = d.Set("recoverytime", dataSourceFlattenSystemsdwanHealthCheckRecoverytime(o["recoverytime"], d, "recoverytime")); err != nil {
		if !fortiAPIPatch(o["recoverytime"]) {
			return fmt.Errorf("error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("security_mode", dataSourceFlattenSystemsdwanHealthCheckSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("error reading security_mode: %v", err)
		}
	}

	if err = d.Set("server", dataSourceFlattenSystemsdwanHealthCheckServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("error reading server: %v", err)
		}
	}

	if err = d.Set("sla", dataSourceFlattenSystemsdwanHealthCheckSla(o["sla"], d, "sla")); err != nil {
		if !fortiAPIPatch(o["sla"]) {
			return fmt.Errorf("error reading sla: %v", err)
		}
	}

	if err = d.Set("sla_fail_log_period", dataSourceFlattenSystemsdwanHealthCheckSlaFailLogPeriod(o["sla-fail-log-period"], d, "sla_fail_log_period")); err != nil {
		if !fortiAPIPatch(o["sla-fail-log-period"]) {
			return fmt.Errorf("error reading sla_fail_log_period: %v", err)
		}
	}

	if err = d.Set("sla_pass_log_period", dataSourceFlattenSystemsdwanHealthCheckSlaPassLogPeriod(o["sla-pass-log-period"], d, "sla_pass_log_period")); err != nil {
		if !fortiAPIPatch(o["sla-pass-log-period"]) {
			return fmt.Errorf("error reading sla_pass_log_period: %v", err)
		}
	}

	if err = d.Set("system_dns", dataSourceFlattenSystemsdwanHealthCheckSystemDns(o["system-dns"], d, "system_dns")); err != nil {
		if !fortiAPIPatch(o["system-dns"]) {
			return fmt.Errorf("error reading system_dns: %v", err)
		}
	}

	if err = d.Set("threshold_alert_jitter", dataSourceFlattenSystemsdwanHealthCheckThresholdAlertJitter(o["threshold-alert-jitter"], d, "threshold_alert_jitter")); err != nil {
		if !fortiAPIPatch(o["threshold-alert-jitter"]) {
			return fmt.Errorf("error reading threshold_alert_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_alert_latency", dataSourceFlattenSystemsdwanHealthCheckThresholdAlertLatency(o["threshold-alert-latency"], d, "threshold_alert_latency")); err != nil {
		if !fortiAPIPatch(o["threshold-alert-latency"]) {
			return fmt.Errorf("error reading threshold_alert_latency: %v", err)
		}
	}

	if err = d.Set("threshold_alert_packetloss", dataSourceFlattenSystemsdwanHealthCheckThresholdAlertPacketloss(o["threshold-alert-packetloss"], d, "threshold_alert_packetloss")); err != nil {
		if !fortiAPIPatch(o["threshold-alert-packetloss"]) {
			return fmt.Errorf("error reading threshold_alert_packetloss: %v", err)
		}
	}

	if err = d.Set("threshold_warning_jitter", dataSourceFlattenSystemsdwanHealthCheckThresholdWarningJitter(o["threshold-warning-jitter"], d, "threshold_warning_jitter")); err != nil {
		if !fortiAPIPatch(o["threshold-warning-jitter"]) {
			return fmt.Errorf("error reading threshold_warning_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_warning_latency", dataSourceFlattenSystemsdwanHealthCheckThresholdWarningLatency(o["threshold-warning-latency"], d, "threshold_warning_latency")); err != nil {
		if !fortiAPIPatch(o["threshold-warning-latency"]) {
			return fmt.Errorf("error reading threshold_warning_latency: %v", err)
		}
	}

	if err = d.Set("threshold_warning_packetloss", dataSourceFlattenSystemsdwanHealthCheckThresholdWarningPacketloss(o["threshold-warning-packetloss"], d, "threshold_warning_packetloss")); err != nil {
		if !fortiAPIPatch(o["threshold-warning-packetloss"]) {
			return fmt.Errorf("error reading threshold_warning_packetloss: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", dataSourceFlattenSystemsdwanHealthCheckUpdateCascadeInterface(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if !fortiAPIPatch(o["update-cascade-interface"]) {
			return fmt.Errorf("error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_static_route", dataSourceFlattenSystemsdwanHealthCheckUpdateStaticRoute(o["update-static-route"], d, "update_static_route")); err != nil {
		if !fortiAPIPatch(o["update-static-route"]) {
			return fmt.Errorf("error reading update_static_route: %v", err)
		}
	}

	if err = d.Set("user", dataSourceFlattenSystemsdwanHealthCheckUser(o["user"], d, "user")); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("error reading user: %v", err)
		}
	}

	return nil
}
