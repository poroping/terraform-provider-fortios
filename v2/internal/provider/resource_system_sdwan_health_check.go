// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemSdwanHealthCheck() *schema.Resource {
	return &schema.Resource{
		Description: "SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.",

		CreateContext: resourceSystemSdwanHealthCheckCreate,
		ReadContext:   resourceSystemSdwanHealthCheckRead,
		UpdateContext: resourceSystemSdwanHealthCheckUpdate,
		DeleteContext: resourceSystemSdwanHealthCheckDelete,

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
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"addr_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

				Description: "Address mode (IPv4 or IPv6).",
				Optional:    true,
				Computed:    true,
			},
			"detect_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"active", "passive", "prefer-passive"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"passive", "port"}, false),

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

				Description: "Packet size of a TWAMP test session.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "TWAMP controller password in authentication mode.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Port number used to communicate with the server over the selected protocol (0 - 65535, default = 0, auto select. http, tcp-connect: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21, twamp: 862).",
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
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"ping", "tcp-echo", "udp-echo", "http", "twamp", "dns", "tcp-connect", "ftp"}, false),

				Description: "Protocol used to determine if the FortiGate can communicate with the server.",
				Optional:    true,
				Computed:    true,
			},
			"quality_measured_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"half-open", "half-close"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"none", "authentication"}, false),

				Description: "Twamp controller security mode.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.StringLenBetween(0, 79),
				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "IP address or FQDN name of the server.",
				Optional:         true,
				Computed:         true,
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

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
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
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable update cascade interface.",
				Optional:    true,
				Computed:    true,
			},
			"update_static_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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

func resourceSystemSdwanHealthCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemSdwanHealthCheck resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSdwanHealthCheck(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSdwanHealthCheck(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanHealthCheck")
	}

	return resourceSystemSdwanHealthCheckRead(ctx, d, meta)
}

func resourceSystemSdwanHealthCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSdwanHealthCheck(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSdwanHealthCheck(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSdwanHealthCheck resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwanHealthCheck")
	}

	return resourceSystemSdwanHealthCheckRead(ctx, d, meta)
}

func resourceSystemSdwanHealthCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSdwanHealthCheck(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSdwanHealthCheck resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanHealthCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadSystemSdwanHealthCheck(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanHealthCheck resource: %v", err)
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

	diags := refreshObjectSystemSdwanHealthCheck(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemSdwanHealthCheck(d *schema.ResourceData, o *models.SystemSdwanHealthCheck, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrMode != nil {
		v := *o.AddrMode

		if err = d.Set("addr_mode", v); err != nil {
			return diag.Errorf("error reading addr_mode: %v", err)
		}
	}

	if o.DetectMode != nil {
		v := *o.DetectMode

		if err = d.Set("detect_mode", v); err != nil {
			return diag.Errorf("error reading detect_mode: %v", err)
		}
	}

	if o.Diffservcode != nil {
		v := *o.Diffservcode

		if err = d.Set("diffservcode", v); err != nil {
			return diag.Errorf("error reading diffservcode: %v", err)
		}
	}

	if o.DnsMatchIp != nil {
		v := *o.DnsMatchIp

		if err = d.Set("dns_match_ip", v); err != nil {
			return diag.Errorf("error reading dns_match_ip: %v", err)
		}
	}

	if o.DnsRequestDomain != nil {
		v := *o.DnsRequestDomain

		if err = d.Set("dns_request_domain", v); err != nil {
			return diag.Errorf("error reading dns_request_domain: %v", err)
		}
	}

	if o.Failtime != nil {
		v := *o.Failtime

		if err = d.Set("failtime", v); err != nil {
			return diag.Errorf("error reading failtime: %v", err)
		}
	}

	if o.FtpFile != nil {
		v := *o.FtpFile

		if err = d.Set("ftp_file", v); err != nil {
			return diag.Errorf("error reading ftp_file: %v", err)
		}
	}

	if o.FtpMode != nil {
		v := *o.FtpMode

		if err = d.Set("ftp_mode", v); err != nil {
			return diag.Errorf("error reading ftp_mode: %v", err)
		}
	}

	if o.HaPriority != nil {
		v := *o.HaPriority

		if err = d.Set("ha_priority", v); err != nil {
			return diag.Errorf("error reading ha_priority: %v", err)
		}
	}

	if o.HttpAgent != nil {
		v := *o.HttpAgent

		if err = d.Set("http_agent", v); err != nil {
			return diag.Errorf("error reading http_agent: %v", err)
		}
	}

	if o.HttpGet != nil {
		v := *o.HttpGet

		if err = d.Set("http_get", v); err != nil {
			return diag.Errorf("error reading http_get: %v", err)
		}
	}

	if o.HttpMatch != nil {
		v := *o.HttpMatch

		if err = d.Set("http_match", v); err != nil {
			return diag.Errorf("error reading http_match: %v", err)
		}
	}

	if o.Interval != nil {
		v := *o.Interval

		if err = d.Set("interval", v); err != nil {
			return diag.Errorf("error reading interval: %v", err)
		}
	}

	if o.Members != nil {
		if err = d.Set("members", flattenSystemSdwanHealthCheckMembers(d, o.Members, "members", sort)); err != nil {
			return diag.Errorf("error reading members: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PacketSize != nil {
		v := *o.PacketSize

		if err = d.Set("packet_size", v); err != nil {
			return diag.Errorf("error reading packet_size: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.ProbeCount != nil {
		v := *o.ProbeCount

		if err = d.Set("probe_count", v); err != nil {
			return diag.Errorf("error reading probe_count: %v", err)
		}
	}

	if o.ProbePackets != nil {
		v := *o.ProbePackets

		if err = d.Set("probe_packets", v); err != nil {
			return diag.Errorf("error reading probe_packets: %v", err)
		}
	}

	if o.ProbeTimeout != nil {
		v := *o.ProbeTimeout

		if err = d.Set("probe_timeout", v); err != nil {
			return diag.Errorf("error reading probe_timeout: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.QualityMeasuredMethod != nil {
		v := *o.QualityMeasuredMethod

		if err = d.Set("quality_measured_method", v); err != nil {
			return diag.Errorf("error reading quality_measured_method: %v", err)
		}
	}

	if o.Recoverytime != nil {
		v := *o.Recoverytime

		if err = d.Set("recoverytime", v); err != nil {
			return diag.Errorf("error reading recoverytime: %v", err)
		}
	}

	if o.SecurityMode != nil {
		v := *o.SecurityMode

		if err = d.Set("security_mode", v); err != nil {
			return diag.Errorf("error reading security_mode: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.Sla != nil {
		if err = d.Set("sla", flattenSystemSdwanHealthCheckSla(d, o.Sla, "sla", sort)); err != nil {
			return diag.Errorf("error reading sla: %v", err)
		}
	}

	if o.SlaFailLogPeriod != nil {
		v := *o.SlaFailLogPeriod

		if err = d.Set("sla_fail_log_period", v); err != nil {
			return diag.Errorf("error reading sla_fail_log_period: %v", err)
		}
	}

	if o.SlaPassLogPeriod != nil {
		v := *o.SlaPassLogPeriod

		if err = d.Set("sla_pass_log_period", v); err != nil {
			return diag.Errorf("error reading sla_pass_log_period: %v", err)
		}
	}

	if o.SystemDns != nil {
		v := *o.SystemDns

		if err = d.Set("system_dns", v); err != nil {
			return diag.Errorf("error reading system_dns: %v", err)
		}
	}

	if o.ThresholdAlertJitter != nil {
		v := *o.ThresholdAlertJitter

		if err = d.Set("threshold_alert_jitter", v); err != nil {
			return diag.Errorf("error reading threshold_alert_jitter: %v", err)
		}
	}

	if o.ThresholdAlertLatency != nil {
		v := *o.ThresholdAlertLatency

		if err = d.Set("threshold_alert_latency", v); err != nil {
			return diag.Errorf("error reading threshold_alert_latency: %v", err)
		}
	}

	if o.ThresholdAlertPacketloss != nil {
		v := *o.ThresholdAlertPacketloss

		if err = d.Set("threshold_alert_packetloss", v); err != nil {
			return diag.Errorf("error reading threshold_alert_packetloss: %v", err)
		}
	}

	if o.ThresholdWarningJitter != nil {
		v := *o.ThresholdWarningJitter

		if err = d.Set("threshold_warning_jitter", v); err != nil {
			return diag.Errorf("error reading threshold_warning_jitter: %v", err)
		}
	}

	if o.ThresholdWarningLatency != nil {
		v := *o.ThresholdWarningLatency

		if err = d.Set("threshold_warning_latency", v); err != nil {
			return diag.Errorf("error reading threshold_warning_latency: %v", err)
		}
	}

	if o.ThresholdWarningPacketloss != nil {
		v := *o.ThresholdWarningPacketloss

		if err = d.Set("threshold_warning_packetloss", v); err != nil {
			return diag.Errorf("error reading threshold_warning_packetloss: %v", err)
		}
	}

	if o.UpdateCascadeInterface != nil {
		v := *o.UpdateCascadeInterface

		if err = d.Set("update_cascade_interface", v); err != nil {
			return diag.Errorf("error reading update_cascade_interface: %v", err)
		}
	}

	if o.UpdateStaticRoute != nil {
		v := *o.UpdateStaticRoute

		if err = d.Set("update_static_route", v); err != nil {
			return diag.Errorf("error reading update_static_route: %v", err)
		}
	}

	if o.User != nil {
		v := *o.User

		if err = d.Set("user", v); err != nil {
			return diag.Errorf("error reading user: %v", err)
		}
	}

	return nil
}

func getObjectSystemSdwanHealthCheck(d *schema.ResourceData, sv string) (*models.SystemSdwanHealthCheck, diag.Diagnostics) {
	obj := models.SystemSdwanHealthCheck{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addr_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("addr_mode", sv)
				diags = append(diags, e)
			}
			obj.AddrMode = &v2
		}
	}
	if v1, ok := d.GetOk("detect_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("detect_mode", sv)
				diags = append(diags, e)
			}
			obj.DetectMode = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode", sv)
				diags = append(diags, e)
			}
			obj.Diffservcode = &v2
		}
	}
	if v1, ok := d.GetOk("dns_match_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_match_ip", sv)
				diags = append(diags, e)
			}
			obj.DnsMatchIp = &v2
		}
	}
	if v1, ok := d.GetOk("dns_request_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_request_domain", sv)
				diags = append(diags, e)
			}
			obj.DnsRequestDomain = &v2
		}
	}
	if v1, ok := d.GetOk("failtime"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("failtime", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Failtime = &tmp
		}
	}
	if v1, ok := d.GetOk("ftp_file"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ftp_file", sv)
				diags = append(diags, e)
			}
			obj.FtpFile = &v2
		}
	}
	if v1, ok := d.GetOk("ftp_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ftp_mode", sv)
				diags = append(diags, e)
			}
			obj.FtpMode = &v2
		}
	}
	if v1, ok := d.GetOk("ha_priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HaPriority = &tmp
		}
	}
	if v1, ok := d.GetOk("http_agent"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_agent", sv)
				diags = append(diags, e)
			}
			obj.HttpAgent = &v2
		}
	}
	if v1, ok := d.GetOk("http_get"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_get", sv)
				diags = append(diags, e)
			}
			obj.HttpGet = &v2
		}
	}
	if v1, ok := d.GetOk("http_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_match", sv)
				diags = append(diags, e)
			}
			obj.HttpMatch = &v2
		}
	}
	if v1, ok := d.GetOk("interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Interval = &tmp
		}
	}
	if v, ok := d.GetOk("members"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("members", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdwanHealthCheckMembers(d, v, "members", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Members = t
		}
	} else if d.HasChange("members") {
		old, new := d.GetChange("members")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Members = &[]models.SystemSdwanHealthCheckMembers{}
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
	if v1, ok := d.GetOk("packet_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PacketSize = &tmp
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("probe_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("probe_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProbeCount = &tmp
		}
	}
	if v1, ok := d.GetOk("probe_packets"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("probe_packets", sv)
				diags = append(diags, e)
			}
			obj.ProbePackets = &v2
		}
	}
	if v1, ok := d.GetOk("probe_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("probe_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProbeTimeout = &tmp
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
	if v1, ok := d.GetOk("quality_measured_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("quality_measured_method", sv)
				diags = append(diags, e)
			}
			obj.QualityMeasuredMethod = &v2
		}
	}
	if v1, ok := d.GetOk("recoverytime"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("recoverytime", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Recoverytime = &tmp
		}
	}
	if v1, ok := d.GetOk("security_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_mode", sv)
				diags = append(diags, e)
			}
			obj.SecurityMode = &v2
		}
	}
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v, ok := d.GetOk("sla"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("sla", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdwanHealthCheckSla(d, v, "sla", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Sla = t
		}
	} else if d.HasChange("sla") {
		old, new := d.GetChange("sla")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Sla = &[]models.SystemSdwanHealthCheckSla{}
		}
	}
	if v1, ok := d.GetOk("sla_fail_log_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sla_fail_log_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SlaFailLogPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("sla_pass_log_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sla_pass_log_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SlaPassLogPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("system_dns"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("system_dns", sv)
				diags = append(diags, e)
			}
			obj.SystemDns = &v2
		}
	}
	if v1, ok := d.GetOk("threshold_alert_jitter"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_alert_jitter", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdAlertJitter = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_alert_latency"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_alert_latency", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdAlertLatency = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_alert_packetloss"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_alert_packetloss", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdAlertPacketloss = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_warning_jitter"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_warning_jitter", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdWarningJitter = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_warning_latency"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_warning_latency", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdWarningLatency = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_warning_packetloss"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_warning_packetloss", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdWarningPacketloss = &tmp
		}
	}
	if v1, ok := d.GetOk("update_cascade_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_cascade_interface", sv)
				diags = append(diags, e)
			}
			obj.UpdateCascadeInterface = &v2
		}
	}
	if v1, ok := d.GetOk("update_static_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_static_route", sv)
				diags = append(diags, e)
			}
			obj.UpdateStaticRoute = &v2
		}
	}
	if v1, ok := d.GetOk("user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user", sv)
				diags = append(diags, e)
			}
			obj.User = &v2
		}
	}
	return &obj, diags
}
