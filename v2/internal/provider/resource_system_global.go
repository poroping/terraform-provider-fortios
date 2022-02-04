// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure global attributes.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure global attributes.",

		CreateContext: resourceSystemGlobalCreate,
		ReadContext:   resourceSystemGlobalRead,
		UpdateContext: resourceSystemGlobalUpdate,
		DeleteContext: resourceSystemGlobalDelete,

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
			"admin_concurrent": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable concurrent administrator logins. Use policy-auth-concurrent for firewall authenticated users.",
				Optional:    true,
				Computed:    true,
			},
			"admin_console_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 300),

				Description: "Console login timeout that overrides the admin timeout value (15 - 300 seconds, default = 0, which disables the timeout).",
				Optional:    true,
				Computed:    true,
			},
			"admin_forticloud_sso_login": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiCloud admin login via SSO.",
				Optional:    true,
				Computed:    true,
			},
			"admin_hsts_max_age": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),

				Description: "HTTPS Strict-Transport-Security header max-age in seconds. A value of 0 will reset any HSTS records in the browser.When admin-https-redirect is disabled the header max-age will be 0.",
				Optional:    true,
				Computed:    true,
			},
			"admin_https_pki_required": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable admin login method. Enable to force administrators to provide a valid certificate to log in if PKI is enabled. Disable to allow administrators to log in with a certificate or password.",
				Optional:    true,
				Computed:    true,
			},
			"admin_https_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redirection of HTTP administration access to HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"admin_https_ssl_banned_ciphers": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Select one or more cipher technologies that cannot be used in GUI HTTPS negotiations. Only applies to TLS 1.2 and below.",
				Optional:         true,
				Computed:         true,
			},
			"admin_https_ssl_ciphersuites": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, remove TLS1.3 from admin-https-ssl-versions.",
				Optional:         true,
				Computed:         true,
			},
			"admin_https_ssl_versions": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed TLS versions for web administration.",
				Optional:         true,
				Computed:         true,
			},
			"admin_lockout_duration": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483647),

				Description: "Amount of time in seconds that an administrator account is locked out after reaching the admin-lockout-threshold for repeated failed login attempts.",
				Optional:    true,
				Computed:    true,
			},
			"admin_lockout_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "Number of failed login attempts before an administrator account is locked out for the admin-lockout-duration.",
				Optional:    true,
				Computed:    true,
			},
			"admin_login_max": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Maximum number of administrators who can be logged in at the same time (1 - 100, default = 100).",
				Optional:    true,
				Computed:    true,
			},
			"admin_maintainer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable maintainer administrator login. When enabled, the maintainer account can be used to log in from the console after a hard reboot. The password is \"bcpb\" followed by the FortiGate unit serial number. You have limited time to complete this login.",
				Optional:    true,
				Computed:    true,
			},
			"admin_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Administrative access port for HTTP. (1 - 65535, default = 80).",
				Optional:    true,
				Computed:    true,
			},
			"admin_restrict_local": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local admin authentication restriction when remote authenticator is up and running (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"admin_scp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable using SCP to download the system configuration. You can use SCP as an alternative method for backing up the configuration.",
				Optional:    true,
				Computed:    true,
			},
			"admin_server_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Server certificate that the FortiGate uses for HTTPS administrative connections.",
				Optional:    true,
				Computed:    true,
			},
			"admin_sport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Administrative access port for HTTPS. (1 - 65535, default = 443).",
				Optional:    true,
				Computed:    true,
			},
			"admin_ssh_grace_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),

				Description: "Maximum time in seconds permitted between making an SSH connection to the FortiGate unit and authenticating (10 - 3600 sec (1 hour), default 120).",
				Optional:    true,
				Computed:    true,
			},
			"admin_ssh_password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable password authentication for SSH admin access.",
				Optional:    true,
				Computed:    true,
			},
			"admin_ssh_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Administrative access port for SSH. (1 - 65535, default = 22).",
				Optional:    true,
				Computed:    true,
			},
			"admin_ssh_v1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSH v1 compatibility.",
				Optional:    true,
				Computed:    true,
			},
			"admin_telnet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TELNET service.",
				Optional:    true,
				Computed:    true,
			},
			"admin_telnet_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Administrative access port for TELNET. (1 - 65535, default = 23).",
				Optional:    true,
				Computed:    true,
			},
			"admintimeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 480),

				Description: "Number of minutes before an idle administrator session times out (1 - 480 minutes (8 hours), default = 5). A shorter idle timeout is more secure.",
				Optional:    true,
				Computed:    true,
			},
			"alias": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Alias for your FortiGate unit.",
				Optional:    true,
				Computed:    true,
			},
			"allow_traffic_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Disable to prevent traffic with same local ingress and egress interface from being forwarded without policy check.",
				Optional:    true,
				Computed:    true,
			},
			"anti_replay": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "loose", "strict"}, false),

				Description: "Level of checking for packet replay and TCP sequence checking.",
				Optional:    true,
				Computed:    true,
			},
			"arp_max_entry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(131072, 2147483647),

				Description: "Maximum number of dynamically learned MAC addresses that can be added to the ARP table (131072 - 2147483647, default = 131072).",
				Optional:    true,
				Computed:    true,
			},
			"auth_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Server certificate that the FortiGate uses for HTTPS firewall authentication connections.",
				Optional:    true,
				Computed:    true,
			},
			"auth_http_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "User authentication HTTP port. (1 - 65535, default = 1000).",
				Optional:    true,
				Computed:    true,
			},
			"auth_https_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "User authentication HTTPS port. (1 - 65535, default = 1003).",
				Optional:    true,
				Computed:    true,
			},
			"auth_keepalive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to prevent user authentication sessions from timing out when idle.",
				Optional:    true,
				Computed:    true,
			},
			"auth_session_limit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"block-new", "logout-inactive"}, false),

				Description: "Action to take when the number of allowed user authenticated sessions is reached.",
				Optional:    true,
				Computed:    true,
			},
			"auto_auth_extension_device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic authorization of dedicated Fortinet extension devices.",
				Optional:    true,
				Computed:    true,
			},
			"autorun_log_fsck": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic log partition check after ungraceful shutdown.",
				Optional:    true,
				Computed:    true,
			},
			"av_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Affinity setting for AV scanning (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Optional:    true,
				Computed:    true,
			},
			"av_failopen": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "off", "one-shot"}, false),

				Description: "Set the action to take if the FortiGate is running low on memory or the proxy connection limit has been reached.",
				Optional:    true,
				Computed:    true,
			},
			"av_failopen_session": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled and a proxy for a protocol runs out of room in its session table, that protocol goes into failopen mode and enacts the action specified by av-failopen.",
				Optional:    true,
				Computed:    true,
			},
			"batch_cmdb": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable batch mode, allowing you to enter a series of CLI commands that will execute as a group once they are loaded.",
				Optional:    true,
				Computed:    true,
			},
			"block_session_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "Duration in seconds for blocked sessions (1 - 300 sec  (5 minutes), default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"br_fdb_max_entry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(8192, 2147483647),

				Description: "Maximum number of bridge forwarding database (FDB) entries.",
				Optional:    true,
				Computed:    true,
			},
			"cert_chain_max": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483647),

				Description: "Maximum number of certificates that can be traversed in a certificate chain.",
				Optional:    true,
				Computed:    true,
			},
			"cfg_revert_timeout": {
				Type: schema.TypeInt,

				Description: "Time-out for reverting to the last saved configuration. (10 - 4294967295 seconds, default = 600).",
				Optional:    true,
				Computed:    true,
			},
			"cfg_save": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"automatic", "manual", "revert"}, false),

				Description: "Configuration file save mode for CLI changes.",
				Optional:    true,
				Computed:    true,
			},
			"check_protocol_header": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"loose", "strict"}, false),

				Description: "Level of checking performed on protocol headers. Strict checking is more thorough but may affect performance. Loose checking is OK in most cases.",
				Optional:    true,
				Computed:    true,
			},
			"check_reset_range": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"strict", "disable"}, false),

				Description: "Configure ICMP error message verification. You can either apply strict RST range checking or disable it.",
				Optional:    true,
				Computed:    true,
			},
			"cli_audit_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable CLI audit log.",
				Optional:    true,
				Computed:    true,
			},
			"cloud_communication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable all cloud communication.",
				Optional:    true,
				Computed:    true,
			},
			"clt_cert_req": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable requiring administrators to have a client certificate to log into the GUI using HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"cmdbsvr_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Affinity setting for cmdbsvr (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Optional:    true,
				Computed:    true,
			},
			"cpu_use_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 99),

				Description: "Threshold at which CPU usage is reported (% of total CPU, default = 90).",
				Optional:    true,
				Computed:    true,
			},
			"csr_ca_attribute": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the CA attribute in certificates. Some CA servers reject CSRs that have the CA attribute.",
				Optional:    true,
				Computed:    true,
			},
			"daily_restart": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable daily restart of FortiGate unit. Use the restart-time option to set the time of day for the restart.",
				Optional:    true,
				Computed:    true,
			},
			"default_service_source_port": {
				Type: schema.TypeString,

				Description: "Default service source port range (default = 1 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"device_identification_active_scan_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 3600),

				Description: "Number of seconds to passively scan a device before performing an active scan. (20 - 3600 sec, (20 sec to 1 hour), default = 90).",
				Optional:    true,
				Computed:    true,
			},
			"device_idle_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 31536000),

				Description: "Time in seconds that a device must be idle to automatically log the device user out. (30 - 31536000 sec (30 sec to 1 year), default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"dh_params": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1024", "1536", "2048", "3072", "4096", "6144", "8192"}, false),

				Description: "Number of bits to use in the Diffie-Hellman exchange for HTTPS/SSH protocols.",
				Optional:    true,
				Computed:    true,
			},
			"dnsproxy_worker_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "DNS proxy worker count. For a FortiGate with multiple logical CPUs, you can set the DNS process number from 1 to the number of logical CPUs.",
				Optional:    true,
				Computed:    true,
			},
			"dst": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable daylight saving time.",
				Optional:    true,
				Computed:    true,
			},
			"extender_controller_reserved_network": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "Configure reserved network subnet for managed LAN extension FortiExtender units. This is available when the FortiExtender daemon is running.",
				Optional:    true,
				Computed:    true,
			},
			"failtime": {
				Type: schema.TypeInt,

				Description: "Fail-time for server lost.",
				Optional:    true,
				Computed:    true,
			},
			"faz_disk_buffer_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 214748364),

				Description: "Maximum disk buffer size to temporarily store logs destined for FortiAnalyzer. To be used in the event that FortiAnalyzer is unavailable.",
				Optional:    true,
				Computed:    true,
			},
			"fds_statistics": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending IPS, Application Control, and AntiVirus data to FortiGuard. This data is used to improve FortiGuard services and is not shared with external parties and is protected by Fortinet's privacy policy.",
				Optional:    true,
				Computed:    true,
			},
			"fds_statistics_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),

				Description: "FortiGuard statistics collection period in minutes. (1 - 1440 min (1 min to 24 hours), default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"fec_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(49152, 65535),

				Description: "Local UDP port for Forward Error Correction (49152 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"fgd_alert_subscription": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Type of alert to retrieve from FortiGuard.",
				Optional:         true,
				Computed:         true,
			},
			"fortiextender": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable FortiExtender.",
				Optional:    true,
				Computed:    true,
			},
			"fortiextender_data_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),

				Description: "FortiExtender data port (1024 - 49150, default = 25246).",
				Optional:    true,
				Computed:    true,
			},
			"fortiextender_discovery_lockdown": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable FortiExtender CAPWAP lockdown.",
				Optional:    true,
				Computed:    true,
			},
			"fortiextender_vlan_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiExtender VLAN mode.",
				Optional:    true,
				Computed:    true,
			},
			"fortiipam_integration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable integration with the FortiIPAM cloud service.",
				Optional:    true,
				Computed:    true,
			},
			"fortiservice_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "FortiService port (1 - 65535, default = 8013). Used by FortiClient endpoint compliance. Older versions of FortiClient used a different port.",
				Optional:    true,
				Computed:    true,
			},
			"fortitoken_cloud": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiToken Cloud service.",
				Optional:    true,
				Computed:    true,
			},
			"gui_allow_default_hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the factory default hostname warning on the GUI setup wizard.",
				Optional:    true,
				Computed:    true,
			},
			"gui_cdn_usage": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Load GUI static files from a CDN.",
				Optional:    true,
				Computed:    true,
			},
			"gui_certificates": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the System > Certificate GUI page, allowing you to add and configure certificates from the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_custom_language": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable custom languages in GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_date_format": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"yyyy/MM/dd", "dd/MM/yyyy", "MM/dd/yyyy", "yyyy-MM-dd", "dd-MM-yyyy", "MM-dd-yyyy"}, false),

				Description: "Default date format used throughout GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_date_time_source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"system", "browser"}, false),

				Description: "Source from which the FortiGate GUI uses to display date and time entries.",
				Optional:    true,
				Computed:    true,
			},
			"gui_device_latitude": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "Add the latitude of the location of this FortiGate to position it on the Threat Map.",
				Optional:    true,
				Computed:    true,
			},
			"gui_device_longitude": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "Add the longitude of the location of this FortiGate to position it on the Threat Map.",
				Optional:    true,
				Computed:    true,
			},
			"gui_display_hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable displaying the FortiGate's hostname on the GUI login page.",
				Optional:    true,
				Computed:    true,
			},
			"gui_firmware_upgrade_setup_warning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the firmware upgrade warning on GUI setup wizard.",
				Optional:    true,
				Computed:    true,
			},
			"gui_firmware_upgrade_warning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the firmware upgrade warning on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_forticare_registration_setup_warning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the FortiCare registration setup warning on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_fortigate_cloud_sandbox": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable displaying FortiGate Cloud Sandbox on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_fortisandbox_cloud": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable displaying FortiSandbox Cloud on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_ipv6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 settings on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_lines_per_page": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 1000),

				Description: "Number of lines to display per page for web administration.",
				Optional:    true,
				Computed:    true,
			},
			"gui_local_out": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Local-out traffic on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_replacement_message_groups": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable replacement message groups on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_rest_api_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable REST API result caching on FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"gui_theme": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"jade", "neutrino", "mariner", "graphite", "melongene", "retro", "dark-matter", "onyx", "eclipse"}, false),

				Description: "Color scheme for the administration GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_wireless_opensecurity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable wireless open security option on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"ha_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Affinity setting for HA daemons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Optional:    true,
				Computed:    true,
			},
			"honor_df": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable honoring of Don't-Fragment (DF) flag.",
				Optional:    true,
				Computed:    true,
			},
			"hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FortiGate unit's hostname. Most models will truncate names longer than 24 characters. Some models support hostnames up to 35 characters.",
				Optional:    true,
				Computed:    true,
			},
			"igmp_state_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(96, 128000),

				Description: "Maximum number of IGMP memberships (96 - 64000, default = 3200).",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_database": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mini", "standard", "full"}, false),

				Description: "Configure which Internet Service database size to download from FortiGuard and use.",
				Optional:    true,
				Computed:    true,
			},
			"interval": {
				Type: schema.TypeInt,

				Description: "Dead gateway detection interval.",
				Optional:    true,
				Computed:    true,
			},
			"ip_src_port_range": {
				Type: schema.TypeString,

				Description: "IP source port range used for traffic originating from the FortiGate unit.",
				Optional:    true,
				Computed:    true,
			},
			"ips_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Affinity setting for IPS (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx; allowed CPUs must be less than total number of IPS engine daemons).",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_asic_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ASIC offloading (hardware acceleration) for IPsec VPN traffic. Hardware acceleration can offload IPsec VPN sessions and accelerate encryption and decryption.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_ha_seqjump_rate": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "ESP jump ahead rate (1G - 10G pps equivalent).",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_hmac_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable offloading (hardware acceleration) of HMAC processing for IPsec VPN.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_soft_dec_async": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable software decryption asynchronization (using multiple CPUs to do decryption) for IPsec VPN traffic.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_accept_dad": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2),

				Description: "Enable/disable acceptance of IPv6 Duplicate Address Detection (DAD).",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_allow_anycast_probe": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 address probe through Anycast.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_allow_traffic_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Disable to prevent IPv6 traffic with same local ingress and egress interface from being forwarded without policy check.",
				Optional:    true,
				Computed:    true,
			},
			"irq_time_accounting": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "force"}, false),

				Description: "Configure CPU IRQ time accounting mode.",
				Optional:    true,
				Computed:    true,
			},
			"language": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"english", "french", "spanish", "portuguese", "japanese", "trach", "simch", "korean"}, false),

				Description: "GUI display language.",
				Optional:    true,
				Computed:    true,
			},
			"ldapconntimeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300000),

				Description: "Global timeout for connections with remote LDAP servers in milliseconds (1 - 300000, default 500).",
				Optional:    true,
				Computed:    true,
			},
			"lldp_reception": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) reception.",
				Optional:    true,
				Computed:    true,
			},
			"lldp_transmission": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) transmission.",
				Optional:    true,
				Computed:    true,
			},
			"log_ssl_connection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging of SSL connection events.",
				Optional:    true,
				Computed:    true,
			},
			"log_uuid_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable insertion of address UUIDs to traffic logs.",
				Optional:    true,
				Computed:    true,
			},
			"log_uuid_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable insertion of policy UUIDs to traffic logs.",
				Optional:    true,
				Computed:    true,
			},
			"login_timestamp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable login time recording.",
				Optional:    true,
				Computed:    true,
			},
			"long_vdom_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable long VDOM name support.",
				Optional:    true,
				Computed:    true,
			},
			"management_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.",
				Optional:    true,
				Computed:    true,
			},
			"management_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Overriding port for management connection (Overrides admin port).",
				Optional:    true,
				Computed:    true,
			},
			"management_port_use_admin_sport": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of the admin-sport setting for the management port. If disabled, FortiGate will allow user to specify management-port.",
				Optional:    true,
				Computed:    true,
			},
			"management_vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Management virtual domain name.",
				Optional:    true,
				Computed:    true,
			},
			"max_dlpstat_memory": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 0),

				Description: "Maximum DLP stat memory (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"max_route_cache_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),

				Description: "Maximum number of IP route cache entries (0 - 2147483647).",
				Optional:    true,
				Computed:    true,
			},
			"memory_use_threshold_extreme": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(70, 97),

				Description: "Threshold at which memory usage is considered extreme (new sessions are dropped) (% of total RAM, default = 95).",
				Optional:    true,
				Computed:    true,
			},
			"memory_use_threshold_green": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(70, 97),

				Description: "Threshold at which memory usage forces the FortiGate to exit conserve mode (% of total RAM, default = 82).",
				Optional:    true,
				Computed:    true,
			},
			"memory_use_threshold_red": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(70, 97),

				Description: "Threshold at which memory usage forces the FortiGate to enter conserve mode (% of total RAM, default = 88).",
				Optional:    true,
				Computed:    true,
			},
			"miglog_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "Affinity setting for logging (64-bit hexadecimal value in the format of xxxxxxxxxxxxxxxx).",
				Optional:    true,
				Computed:    true,
			},
			"miglogd_children": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),

				Description: "Number of logging (miglogd) processes to be allowed to run. Higher number can reduce performance; lower number can slow log processing time. No logs will be dropped or lost if the number is changed.",
				Optional:    true,
				Computed:    true,
			},
			"multi_factor_authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"optional", "mandatory"}, false),

				Description: "Enforce all login methods to require an additional authentication factor (default = optional).",
				Optional:    true,
				Computed:    true,
			},
			"ndp_max_entry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(65536, 2147483647),

				Description: "Maximum number of NDP table entries (set to 65,536 or higher; if set to 0, kernel holds 65,536 entries).",
				Optional:    true,
				Computed:    true,
			},
			"per_user_bal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable per-user block/allow list filter.",
				Optional:    true,
				Computed:    true,
			},
			"per_user_bwl": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable per-user black/white list filter.",
				Optional:    true,
				Computed:    true,
			},
			"pmtu_discovery": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable path MTU discovery.",
				Optional:    true,
				Computed:    true,
			},
			"policy_auth_concurrent": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Number of concurrent firewall use logins from the same user (1 - 100, default = 0 means no limit).",
				Optional:    true,
				Computed:    true,
			},
			"post_login_banner": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable displaying the administrator access disclaimer message after an administrator successfully logs in.",
				Optional:    true,
				Computed:    true,
			},
			"pre_login_banner": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable displaying the administrator access disclaimer message on the login page before an administrator logs in.",
				Optional:    true,
				Computed:    true,
			},
			"private_data_encryption": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable private data encryption using an AES 128-bit key or passpharse.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_auth_lifetime": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authenticated users lifetime control. This is a cap on the total time a proxy user can be authenticated for after which re-authentication will take place.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_auth_lifetime_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 65535),

				Description: "Lifetime timeout in minutes for authenticated users (5  - 65535 min, default=480 (8 hours)).",
				Optional:    true,
				Computed:    true,
			},
			"proxy_auth_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "Authentication timeout in minutes for authenticated users (1 - 300 min, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"proxy_cert_use_mgmt_vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable using management VDOM to send requests.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_cipher_hardware_acceleration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable using content processor (CP8 or CP9) hardware acceleration to encrypt and decrypt IPsec and SSL traffic.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_kxp_hardware_acceleration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable using the content processor to accelerate KXP traffic.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_re_authentication_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"session", "traffic", "absolute"}, false),

				Description: "Control if users must re-authenticate after a session is closed, traffic has been idle, or from the point at which the user was first created.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_resource_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of the maximum memory usage on the FortiGate unit's proxy processing of resources, such as block lists, allow lists, and external resources.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_worker_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Proxy worker count.",
				Optional:    true,
				Computed:    true,
			},
			"radius_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "RADIUS service port number.",
				Optional:    true,
				Computed:    true,
			},
			"reboot_upon_config_restore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable reboot of system upon restoring configuration.",
				Optional:    true,
				Computed:    true,
			},
			"refresh": {
				Type: schema.TypeInt,

				Description: "Statistics refresh interval second(s) in GUI.",
				Optional:    true,
				Computed:    true,
			},
			"remoteauthtimeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "Number of seconds that the FortiGate waits for responses from remote RADIUS, LDAP, or TACACS+ authentication servers. (1-300 sec, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"reset_sessionless_tcp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Action to perform if the FortiGate receives a TCP packet but cannot find a corresponding session in its session table. NAT/Route mode only.",
				Optional:    true,
				Computed:    true,
			},
			"restart_time": {
				Type: schema.TypeString,

				Description: "Daily restart time (hh:mm).",
				Optional:    true,
				Computed:    true,
			},
			"revision_backup_on_logout": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable back-up of the latest configuration revision when an administrator logs out of the CLI or GUI.",
				Optional:    true,
				Computed:    true,
			},
			"revision_image_auto_backup": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable back-up of the latest image revision after the firmware is upgraded.",
				Optional:    true,
				Computed:    true,
			},
			"scanunit_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),

				Description: "Number of scanunits. The range and the default depend on the number of CPUs. Only available on FortiGate units with multiple CPUs.",
				Optional:    true,
				Computed:    true,
			},
			"security_rating_result_submission": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the submission of Security Rating results to FortiGuard.",
				Optional:    true,
				Computed:    true,
			},
			"security_rating_run_on_schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable scheduled runs of Security Rating.",
				Optional:    true,
				Computed:    true,
			},
			"send_pmtu_icmp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending of path maximum transmission unit (PMTU) - ICMP destination unreachable packet and to support PMTUD protocol on your network to reduce fragmentation of packets.",
				Optional:    true,
				Computed:    true,
			},
			"snat_route_change": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the ability to change the static NAT route.",
				Optional:    true,
				Computed:    true,
			},
			"special_file_23_support": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable detection of those special format files when using Data Leak Protection.",
				Optional:    true,
				Computed:    true,
			},
			"speedtest_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable speed test server.",
				Optional:    true,
				Computed:    true,
			},
			"ssd_trim_date": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),

				Description: "Date within a month to run ssd trim.",
				Optional:    true,
				Computed:    true,
			},
			"ssd_trim_freq": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"never", "hourly", "daily", "weekly", "monthly"}, false),

				Description: "How often to run SSD Trim (default = weekly). SSD Trim prevents SSD drive data loss by finding and isolating errors.",
				Optional:    true,
				Computed:    true,
			},
			"ssd_trim_hour": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),

				Description: "Hour of the day on which to run SSD Trim (0 - 23, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"ssd_trim_min": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),

				Description: "Minute of the hour on which to run SSD Trim (0 - 59, 60 for random).",
				Optional:    true,
				Computed:    true,
			},
			"ssd_trim_weekday": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}, false),

				Description: "Day of week to run SSD Trim.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_cbc_cipher": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable CBC cipher for SSH access.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_enc_algo": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Select one or more SSH ciphers.",
				Optional:         true,
				Computed:         true,
			},
			"ssh_hmac_md5": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HMAC-MD5 for SSH access.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_kex_algo": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Select one or more SSH kex algorithms.",
				Optional:         true,
				Computed:         true,
			},
			"ssh_kex_sha1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SHA1 key exchange for SSH access.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_mac_algo": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Select one or more SSH MAC algorithms.",
				Optional:         true,
				Computed:         true,
			},
			"ssh_mac_weak": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HMAC-SHA1 and UMAC-64-ETM for SSH access.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"SSLv3", "TLSv1", "TLSv1-1", "TLSv1-2", "TLSv1-3"}, false),

				Description: "Minimum supported protocol version for SSL/TLS connections (default = TLSv1.2).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_static_key_ciphers": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable static key ciphers in SSL/TLS connections (e.g. AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256).",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_cipher_hardware_acceleration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL VPN hardware acceleration.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_ems_sn_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable verification of EMS serial number in SSL-VPN connection.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_kxp_hardware_acceleration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL VPN KXP hardware acceleration.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_max_worker_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "Maximum number of SSL-VPN processes. Upper limit for this value is the number of CPUs and depends on the model. Default value of zero means the SSLVPN daemon decides the number of worker processes.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_plugin_version_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable checking browser's plugin version by SSL-VPN.",
				Optional:    true,
				Computed:    true,
			},
			"strict_dirty_session_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to check the session against the original policy when revalidating. This can prevent dropping of redirected sessions when web-filtering and authentication are enabled together. If this option is enabled, the FortiGate unit deletes a session if a routing or policy change causes the session to no longer match the policy that originally allowed the session.",
				Optional:    true,
				Computed:    true,
			},
			"strong_crypto": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to use strong encryption and only allow strong ciphers and digest for HTTPS/SSH/TLS/SSL functions.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable switch controller feature. Switch controller allows you to manage FortiSwitch from the FortiGate itself.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_reserved_network": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "Configure reserved network subnet for managed switches. This is available when the switch controller is enabled.",
				Optional:    true,
				Computed:    true,
			},
			"sys_perf_log_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),

				Description: "Time in minutes between updates of performance statistics logging. (1 - 15 min, default = 5, 0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_halfclose_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),

				Description: "Number of seconds the FortiGate unit should wait to close a session after one peer has sent a FIN packet but the other has not responded (1 - 86400 sec (1 day), default = 120).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_halfopen_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),

				Description: "Number of seconds the FortiGate unit should wait to close a session after one peer has sent an open session packet but the other has not responded (1 - 86400 sec (1 day), default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable SACK, timestamp and MSS TCP options.",
				Optional:    true,
				Computed:    true,
			},
			"tcp_rst_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 300),

				Description: "Length of the TCP CLOSE state in seconds (5 - 300 sec, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_timewait_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),

				Description: "Length of the TCP TIME-WAIT state in seconds (1 - 300 sec, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"tftp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TFTP.",
				Optional:    true,
				Computed:    true,
			},
			"timezone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"01", "02", "03", "04", "05", "81", "06", "07", "08", "09", "10", "11", "12", "13", "74", "14", "77", "15", "87", "16", "17", "18", "19", "20", "75", "21", "22", "23", "24", "80", "79", "25", "26", "27", "28", "78", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "83", "84", "40", "85", "41", "42", "43", "39", "44", "46", "47", "51", "48", "45", "49", "50", "52", "53", "54", "55", "56", "57", "58", "59", "60", "62", "63", "61", "64", "65", "66", "67", "68", "69", "70", "71", "72", "00", "82", "73", "86", "76"}, false),

				Description: "Number corresponding to your time zone from 00 to 86. Enter set timezone ? to view the list of time zones and the numbers that represent them.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_priority": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tos", "dscp"}, false),

				Description: "Choose Type of Service (ToS) or Differentiated Services Code Point (DSCP) for traffic prioritization in traffic shaping.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_priority_level": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high"}, false),

				Description: "Default system-wide level of priority for traffic prioritization.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_email_expiry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 300),

				Description: "Email-based two-factor authentication session timeout (30 - 300 seconds (5 minutes), default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_fac_expiry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),

				Description: "FortiAuthenticator token authentication session timeout (10 - 3600 seconds (1 hour), default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_ftk_expiry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 600),

				Description: "FortiToken authentication session timeout (60 - 600 sec (10 minutes), default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_ftm_expiry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 168),

				Description: "FortiToken Mobile session timeout (1 - 168 hours (7 days), default = 72).",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_sms_expiry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 300),

				Description: "SMS-based two-factor authentication session timeout (30 - 300 sec, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"udp_idle_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),

				Description: "UDP connection session timeout. This command can be useful in managing CPU and memory resources (1 - 86400 seconds (1 day), default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"url_filter_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "URL filter CPU affinity.",
				Optional:    true,
				Computed:    true,
			},
			"url_filter_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "URL filter daemon count.",
				Optional:    true,
				Computed:    true,
			},
			"user_device_store_max_devices": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(36677, 104794),

				Description: "Maximum number of devices allowed in user device store.",
				Optional:    true,
				Computed:    true,
			},
			"user_device_store_max_unified_mem": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(73355960, 733559603),

				Description: "Maximum unified memory allowed in user device store.",
				Optional:    true,
				Computed:    true,
			},
			"user_device_store_max_users": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(36677, 104794),

				Description: "Maximum number of users allowed in user device store.",
				Optional:    true,
				Computed:    true,
			},
			"user_server_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Certificate to use for https user authentication.",
				Optional:    true,
				Computed:    true,
			},
			"vdom_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"no-vdom", "split-vdom", "multi-vdom"}, false),

				Description: "Enable/disable support for split/multiple virtual domains (VDOMs).",
				Optional:    true,
				Computed:    true,
			},
			"vip_arp_range": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unlimited", "restricted"}, false),

				Description: "Controls the number of ARPs that the FortiGate sends for a Virtual IP (VIP) address range.",
				Optional:    true,
				Computed:    true,
			},
			"wad_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Affinity setting for wad (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Optional:    true,
				Computed:    true,
			},
			"wad_csvc_cs_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1),

				Description: "Number of concurrent WAD-cache-service object-cache processes.",
				Optional:    true,
				Computed:    true,
			},
			"wad_csvc_db_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Number of concurrent WAD-cache-service byte-cache processes.",
				Optional:    true,
				Computed:    true,
			},
			"wad_memory_change_granularity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 25),

				Description: "Minimum percentage change in system memory usage detected by the wad daemon prior to adjusting TCP window size for any active connection.",
				Optional:    true,
				Computed:    true,
			},
			"wad_source_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable dispatching traffic to WAD workers based on source affinity.",
				Optional:    true,
				Computed:    true,
			},
			"wad_worker_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Number of explicit proxy WAN optimization daemon (WAD) processes. By default WAN optimization, explicit proxy, and web caching is handled by all of the CPU cores in a FortiGate unit.",
				Optional:    true,
				Computed:    true,
			},
			"wifi_ca_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "CA certificate that verifies the WiFi certificate.",
				Optional:    true,
				Computed:    true,
			},
			"wifi_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Certificate to use for WiFi authentication.",
				Optional:    true,
				Computed:    true,
			},
			"wimax_4g_usb": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable comparability with WiMAX 4G USB devices.",
				Optional:    true,
				Computed:    true,
			},
			"wireless_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the wireless controller feature to use the FortiGate unit to manage FortiAPs.",
				Optional:    true,
				Computed:    true,
			},
			"wireless_controller_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),

				Description: "Port used for the control channel in wireless controller mode (wireless-mode is ac). The data channel port is the control channel port number plus one (1024 - 49150, default = 5246).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemGlobal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemGlobal")
	}

	return resourceSystemGlobalRead(ctx, d, meta)
}

func resourceSystemGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemGlobal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemGlobal")
	}

	return resourceSystemGlobalRead(ctx, d, meta)
}

func resourceSystemGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemGlobal resource: %v", err)
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

	diags := refreshObjectSystemGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemGlobal(d *schema.ResourceData, o *models.SystemGlobal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AdminConcurrent != nil {
		v := *o.AdminConcurrent

		if err = d.Set("admin_concurrent", v); err != nil {
			return diag.Errorf("error reading admin_concurrent: %v", err)
		}
	}

	if o.AdminConsoleTimeout != nil {
		v := *o.AdminConsoleTimeout

		if err = d.Set("admin_console_timeout", v); err != nil {
			return diag.Errorf("error reading admin_console_timeout: %v", err)
		}
	}

	if o.AdminForticloudSsoLogin != nil {
		v := *o.AdminForticloudSsoLogin

		if err = d.Set("admin_forticloud_sso_login", v); err != nil {
			return diag.Errorf("error reading admin_forticloud_sso_login: %v", err)
		}
	}

	if o.AdminHstsMaxAge != nil {
		v := *o.AdminHstsMaxAge

		if err = d.Set("admin_hsts_max_age", v); err != nil {
			return diag.Errorf("error reading admin_hsts_max_age: %v", err)
		}
	}

	if o.AdminHttpsPkiRequired != nil {
		v := *o.AdminHttpsPkiRequired

		if err = d.Set("admin_https_pki_required", v); err != nil {
			return diag.Errorf("error reading admin_https_pki_required: %v", err)
		}
	}

	if o.AdminHttpsRedirect != nil {
		v := *o.AdminHttpsRedirect

		if err = d.Set("admin_https_redirect", v); err != nil {
			return diag.Errorf("error reading admin_https_redirect: %v", err)
		}
	}

	if o.AdminHttpsSslBannedCiphers != nil {
		v := *o.AdminHttpsSslBannedCiphers

		if err = d.Set("admin_https_ssl_banned_ciphers", v); err != nil {
			return diag.Errorf("error reading admin_https_ssl_banned_ciphers: %v", err)
		}
	}

	if o.AdminHttpsSslCiphersuites != nil {
		v := *o.AdminHttpsSslCiphersuites

		if err = d.Set("admin_https_ssl_ciphersuites", v); err != nil {
			return diag.Errorf("error reading admin_https_ssl_ciphersuites: %v", err)
		}
	}

	if o.AdminHttpsSslVersions != nil {
		v := *o.AdminHttpsSslVersions

		if err = d.Set("admin_https_ssl_versions", v); err != nil {
			return diag.Errorf("error reading admin_https_ssl_versions: %v", err)
		}
	}

	if o.AdminLockoutDuration != nil {
		v := *o.AdminLockoutDuration

		if err = d.Set("admin_lockout_duration", v); err != nil {
			return diag.Errorf("error reading admin_lockout_duration: %v", err)
		}
	}

	if o.AdminLockoutThreshold != nil {
		v := *o.AdminLockoutThreshold

		if err = d.Set("admin_lockout_threshold", v); err != nil {
			return diag.Errorf("error reading admin_lockout_threshold: %v", err)
		}
	}

	if o.AdminLoginMax != nil {
		v := *o.AdminLoginMax

		if err = d.Set("admin_login_max", v); err != nil {
			return diag.Errorf("error reading admin_login_max: %v", err)
		}
	}

	if o.AdminMaintainer != nil {
		v := *o.AdminMaintainer

		if err = d.Set("admin_maintainer", v); err != nil {
			return diag.Errorf("error reading admin_maintainer: %v", err)
		}
	}

	if o.AdminPort != nil {
		v := *o.AdminPort

		if err = d.Set("admin_port", v); err != nil {
			return diag.Errorf("error reading admin_port: %v", err)
		}
	}

	if o.AdminRestrictLocal != nil {
		v := *o.AdminRestrictLocal

		if err = d.Set("admin_restrict_local", v); err != nil {
			return diag.Errorf("error reading admin_restrict_local: %v", err)
		}
	}

	if o.AdminScp != nil {
		v := *o.AdminScp

		if err = d.Set("admin_scp", v); err != nil {
			return diag.Errorf("error reading admin_scp: %v", err)
		}
	}

	if o.AdminServerCert != nil {
		v := *o.AdminServerCert

		if err = d.Set("admin_server_cert", v); err != nil {
			return diag.Errorf("error reading admin_server_cert: %v", err)
		}
	}

	if o.AdminSport != nil {
		v := *o.AdminSport

		if err = d.Set("admin_sport", v); err != nil {
			return diag.Errorf("error reading admin_sport: %v", err)
		}
	}

	if o.AdminSshGraceTime != nil {
		v := *o.AdminSshGraceTime

		if err = d.Set("admin_ssh_grace_time", v); err != nil {
			return diag.Errorf("error reading admin_ssh_grace_time: %v", err)
		}
	}

	if o.AdminSshPassword != nil {
		v := *o.AdminSshPassword

		if err = d.Set("admin_ssh_password", v); err != nil {
			return diag.Errorf("error reading admin_ssh_password: %v", err)
		}
	}

	if o.AdminSshPort != nil {
		v := *o.AdminSshPort

		if err = d.Set("admin_ssh_port", v); err != nil {
			return diag.Errorf("error reading admin_ssh_port: %v", err)
		}
	}

	if o.AdminSshV1 != nil {
		v := *o.AdminSshV1

		if err = d.Set("admin_ssh_v1", v); err != nil {
			return diag.Errorf("error reading admin_ssh_v1: %v", err)
		}
	}

	if o.AdminTelnet != nil {
		v := *o.AdminTelnet

		if err = d.Set("admin_telnet", v); err != nil {
			return diag.Errorf("error reading admin_telnet: %v", err)
		}
	}

	if o.AdminTelnetPort != nil {
		v := *o.AdminTelnetPort

		if err = d.Set("admin_telnet_port", v); err != nil {
			return diag.Errorf("error reading admin_telnet_port: %v", err)
		}
	}

	if o.Admintimeout != nil {
		v := *o.Admintimeout

		if err = d.Set("admintimeout", v); err != nil {
			return diag.Errorf("error reading admintimeout: %v", err)
		}
	}

	if o.Alias != nil {
		v := *o.Alias

		if err = d.Set("alias", v); err != nil {
			return diag.Errorf("error reading alias: %v", err)
		}
	}

	if o.AllowTrafficRedirect != nil {
		v := *o.AllowTrafficRedirect

		if err = d.Set("allow_traffic_redirect", v); err != nil {
			return diag.Errorf("error reading allow_traffic_redirect: %v", err)
		}
	}

	if o.AntiReplay != nil {
		v := *o.AntiReplay

		if err = d.Set("anti_replay", v); err != nil {
			return diag.Errorf("error reading anti_replay: %v", err)
		}
	}

	if o.ArpMaxEntry != nil {
		v := *o.ArpMaxEntry

		if err = d.Set("arp_max_entry", v); err != nil {
			return diag.Errorf("error reading arp_max_entry: %v", err)
		}
	}

	if o.AuthCert != nil {
		v := *o.AuthCert

		if err = d.Set("auth_cert", v); err != nil {
			return diag.Errorf("error reading auth_cert: %v", err)
		}
	}

	if o.AuthHttpPort != nil {
		v := *o.AuthHttpPort

		if err = d.Set("auth_http_port", v); err != nil {
			return diag.Errorf("error reading auth_http_port: %v", err)
		}
	}

	if o.AuthHttpsPort != nil {
		v := *o.AuthHttpsPort

		if err = d.Set("auth_https_port", v); err != nil {
			return diag.Errorf("error reading auth_https_port: %v", err)
		}
	}

	if o.AuthKeepalive != nil {
		v := *o.AuthKeepalive

		if err = d.Set("auth_keepalive", v); err != nil {
			return diag.Errorf("error reading auth_keepalive: %v", err)
		}
	}

	if o.AuthSessionLimit != nil {
		v := *o.AuthSessionLimit

		if err = d.Set("auth_session_limit", v); err != nil {
			return diag.Errorf("error reading auth_session_limit: %v", err)
		}
	}

	if o.AutoAuthExtensionDevice != nil {
		v := *o.AutoAuthExtensionDevice

		if err = d.Set("auto_auth_extension_device", v); err != nil {
			return diag.Errorf("error reading auto_auth_extension_device: %v", err)
		}
	}

	if o.AutorunLogFsck != nil {
		v := *o.AutorunLogFsck

		if err = d.Set("autorun_log_fsck", v); err != nil {
			return diag.Errorf("error reading autorun_log_fsck: %v", err)
		}
	}

	if o.AvAffinity != nil {
		v := *o.AvAffinity

		if err = d.Set("av_affinity", v); err != nil {
			return diag.Errorf("error reading av_affinity: %v", err)
		}
	}

	if o.AvFailopen != nil {
		v := *o.AvFailopen

		if err = d.Set("av_failopen", v); err != nil {
			return diag.Errorf("error reading av_failopen: %v", err)
		}
	}

	if o.AvFailopenSession != nil {
		v := *o.AvFailopenSession

		if err = d.Set("av_failopen_session", v); err != nil {
			return diag.Errorf("error reading av_failopen_session: %v", err)
		}
	}

	if o.BatchCmdb != nil {
		v := *o.BatchCmdb

		if err = d.Set("batch_cmdb", v); err != nil {
			return diag.Errorf("error reading batch_cmdb: %v", err)
		}
	}

	if o.BlockSessionTimer != nil {
		v := *o.BlockSessionTimer

		if err = d.Set("block_session_timer", v); err != nil {
			return diag.Errorf("error reading block_session_timer: %v", err)
		}
	}

	if o.BrFdbMaxEntry != nil {
		v := *o.BrFdbMaxEntry

		if err = d.Set("br_fdb_max_entry", v); err != nil {
			return diag.Errorf("error reading br_fdb_max_entry: %v", err)
		}
	}

	if o.CertChainMax != nil {
		v := *o.CertChainMax

		if err = d.Set("cert_chain_max", v); err != nil {
			return diag.Errorf("error reading cert_chain_max: %v", err)
		}
	}

	if o.CfgRevertTimeout != nil {
		v := *o.CfgRevertTimeout

		if err = d.Set("cfg_revert_timeout", v); err != nil {
			return diag.Errorf("error reading cfg_revert_timeout: %v", err)
		}
	}

	if o.CfgSave != nil {
		v := *o.CfgSave

		if err = d.Set("cfg_save", v); err != nil {
			return diag.Errorf("error reading cfg_save: %v", err)
		}
	}

	if o.CheckProtocolHeader != nil {
		v := *o.CheckProtocolHeader

		if err = d.Set("check_protocol_header", v); err != nil {
			return diag.Errorf("error reading check_protocol_header: %v", err)
		}
	}

	if o.CheckResetRange != nil {
		v := *o.CheckResetRange

		if err = d.Set("check_reset_range", v); err != nil {
			return diag.Errorf("error reading check_reset_range: %v", err)
		}
	}

	if o.CliAuditLog != nil {
		v := *o.CliAuditLog

		if err = d.Set("cli_audit_log", v); err != nil {
			return diag.Errorf("error reading cli_audit_log: %v", err)
		}
	}

	if o.CloudCommunication != nil {
		v := *o.CloudCommunication

		if err = d.Set("cloud_communication", v); err != nil {
			return diag.Errorf("error reading cloud_communication: %v", err)
		}
	}

	if o.CltCertReq != nil {
		v := *o.CltCertReq

		if err = d.Set("clt_cert_req", v); err != nil {
			return diag.Errorf("error reading clt_cert_req: %v", err)
		}
	}

	if o.CmdbsvrAffinity != nil {
		v := *o.CmdbsvrAffinity

		if err = d.Set("cmdbsvr_affinity", v); err != nil {
			return diag.Errorf("error reading cmdbsvr_affinity: %v", err)
		}
	}

	if o.CpuUseThreshold != nil {
		v := *o.CpuUseThreshold

		if err = d.Set("cpu_use_threshold", v); err != nil {
			return diag.Errorf("error reading cpu_use_threshold: %v", err)
		}
	}

	if o.CsrCaAttribute != nil {
		v := *o.CsrCaAttribute

		if err = d.Set("csr_ca_attribute", v); err != nil {
			return diag.Errorf("error reading csr_ca_attribute: %v", err)
		}
	}

	if o.DailyRestart != nil {
		v := *o.DailyRestart

		if err = d.Set("daily_restart", v); err != nil {
			return diag.Errorf("error reading daily_restart: %v", err)
		}
	}

	if o.DefaultServiceSourcePort != nil {
		v := *o.DefaultServiceSourcePort

		if err = d.Set("default_service_source_port", v); err != nil {
			return diag.Errorf("error reading default_service_source_port: %v", err)
		}
	}

	if o.DeviceIdentificationActiveScanDelay != nil {
		v := *o.DeviceIdentificationActiveScanDelay

		if err = d.Set("device_identification_active_scan_delay", v); err != nil {
			return diag.Errorf("error reading device_identification_active_scan_delay: %v", err)
		}
	}

	if o.DeviceIdleTimeout != nil {
		v := *o.DeviceIdleTimeout

		if err = d.Set("device_idle_timeout", v); err != nil {
			return diag.Errorf("error reading device_idle_timeout: %v", err)
		}
	}

	if o.DhParams != nil {
		v := *o.DhParams

		if err = d.Set("dh_params", v); err != nil {
			return diag.Errorf("error reading dh_params: %v", err)
		}
	}

	if o.DnsproxyWorkerCount != nil {
		v := *o.DnsproxyWorkerCount

		if err = d.Set("dnsproxy_worker_count", v); err != nil {
			return diag.Errorf("error reading dnsproxy_worker_count: %v", err)
		}
	}

	if o.Dst != nil {
		v := *o.Dst

		if err = d.Set("dst", v); err != nil {
			return diag.Errorf("error reading dst: %v", err)
		}
	}

	if o.ExtenderControllerReservedNetwork != nil {
		v := *o.ExtenderControllerReservedNetwork
		if current, ok := d.GetOk("extender_controller_reserved_network"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("extender_controller_reserved_network", v); err != nil {
			return diag.Errorf("error reading extender_controller_reserved_network: %v", err)
		}
	}

	if o.Failtime != nil {
		v := *o.Failtime

		if err = d.Set("failtime", v); err != nil {
			return diag.Errorf("error reading failtime: %v", err)
		}
	}

	if o.FazDiskBufferSize != nil {
		v := *o.FazDiskBufferSize

		if err = d.Set("faz_disk_buffer_size", v); err != nil {
			return diag.Errorf("error reading faz_disk_buffer_size: %v", err)
		}
	}

	if o.FdsStatistics != nil {
		v := *o.FdsStatistics

		if err = d.Set("fds_statistics", v); err != nil {
			return diag.Errorf("error reading fds_statistics: %v", err)
		}
	}

	if o.FdsStatisticsPeriod != nil {
		v := *o.FdsStatisticsPeriod

		if err = d.Set("fds_statistics_period", v); err != nil {
			return diag.Errorf("error reading fds_statistics_period: %v", err)
		}
	}

	if o.FecPort != nil {
		v := *o.FecPort

		if err = d.Set("fec_port", v); err != nil {
			return diag.Errorf("error reading fec_port: %v", err)
		}
	}

	if o.FgdAlertSubscription != nil {
		v := *o.FgdAlertSubscription

		if err = d.Set("fgd_alert_subscription", v); err != nil {
			return diag.Errorf("error reading fgd_alert_subscription: %v", err)
		}
	}

	if o.Fortiextender != nil {
		v := *o.Fortiextender

		if err = d.Set("fortiextender", v); err != nil {
			return diag.Errorf("error reading fortiextender: %v", err)
		}
	}

	if o.FortiextenderDataPort != nil {
		v := *o.FortiextenderDataPort

		if err = d.Set("fortiextender_data_port", v); err != nil {
			return diag.Errorf("error reading fortiextender_data_port: %v", err)
		}
	}

	if o.FortiextenderDiscoveryLockdown != nil {
		v := *o.FortiextenderDiscoveryLockdown

		if err = d.Set("fortiextender_discovery_lockdown", v); err != nil {
			return diag.Errorf("error reading fortiextender_discovery_lockdown: %v", err)
		}
	}

	if o.FortiextenderVlanMode != nil {
		v := *o.FortiextenderVlanMode

		if err = d.Set("fortiextender_vlan_mode", v); err != nil {
			return diag.Errorf("error reading fortiextender_vlan_mode: %v", err)
		}
	}

	if o.FortiipamIntegration != nil {
		v := *o.FortiipamIntegration

		if err = d.Set("fortiipam_integration", v); err != nil {
			return diag.Errorf("error reading fortiipam_integration: %v", err)
		}
	}

	if o.FortiservicePort != nil {
		v := *o.FortiservicePort

		if err = d.Set("fortiservice_port", v); err != nil {
			return diag.Errorf("error reading fortiservice_port: %v", err)
		}
	}

	if o.FortitokenCloud != nil {
		v := *o.FortitokenCloud

		if err = d.Set("fortitoken_cloud", v); err != nil {
			return diag.Errorf("error reading fortitoken_cloud: %v", err)
		}
	}

	if o.GuiAllowDefaultHostname != nil {
		v := *o.GuiAllowDefaultHostname

		if err = d.Set("gui_allow_default_hostname", v); err != nil {
			return diag.Errorf("error reading gui_allow_default_hostname: %v", err)
		}
	}

	if o.GuiCdnUsage != nil {
		v := *o.GuiCdnUsage

		if err = d.Set("gui_cdn_usage", v); err != nil {
			return diag.Errorf("error reading gui_cdn_usage: %v", err)
		}
	}

	if o.GuiCertificates != nil {
		v := *o.GuiCertificates

		if err = d.Set("gui_certificates", v); err != nil {
			return diag.Errorf("error reading gui_certificates: %v", err)
		}
	}

	if o.GuiCustomLanguage != nil {
		v := *o.GuiCustomLanguage

		if err = d.Set("gui_custom_language", v); err != nil {
			return diag.Errorf("error reading gui_custom_language: %v", err)
		}
	}

	if o.GuiDateFormat != nil {
		v := *o.GuiDateFormat

		if err = d.Set("gui_date_format", v); err != nil {
			return diag.Errorf("error reading gui_date_format: %v", err)
		}
	}

	if o.GuiDateTimeSource != nil {
		v := *o.GuiDateTimeSource

		if err = d.Set("gui_date_time_source", v); err != nil {
			return diag.Errorf("error reading gui_date_time_source: %v", err)
		}
	}

	if o.GuiDeviceLatitude != nil {
		v := *o.GuiDeviceLatitude

		if err = d.Set("gui_device_latitude", v); err != nil {
			return diag.Errorf("error reading gui_device_latitude: %v", err)
		}
	}

	if o.GuiDeviceLongitude != nil {
		v := *o.GuiDeviceLongitude

		if err = d.Set("gui_device_longitude", v); err != nil {
			return diag.Errorf("error reading gui_device_longitude: %v", err)
		}
	}

	if o.GuiDisplayHostname != nil {
		v := *o.GuiDisplayHostname

		if err = d.Set("gui_display_hostname", v); err != nil {
			return diag.Errorf("error reading gui_display_hostname: %v", err)
		}
	}

	if o.GuiFirmwareUpgradeSetupWarning != nil {
		v := *o.GuiFirmwareUpgradeSetupWarning

		if err = d.Set("gui_firmware_upgrade_setup_warning", v); err != nil {
			return diag.Errorf("error reading gui_firmware_upgrade_setup_warning: %v", err)
		}
	}

	if o.GuiFirmwareUpgradeWarning != nil {
		v := *o.GuiFirmwareUpgradeWarning

		if err = d.Set("gui_firmware_upgrade_warning", v); err != nil {
			return diag.Errorf("error reading gui_firmware_upgrade_warning: %v", err)
		}
	}

	if o.GuiForticareRegistrationSetupWarning != nil {
		v := *o.GuiForticareRegistrationSetupWarning

		if err = d.Set("gui_forticare_registration_setup_warning", v); err != nil {
			return diag.Errorf("error reading gui_forticare_registration_setup_warning: %v", err)
		}
	}

	if o.GuiFortigateCloudSandbox != nil {
		v := *o.GuiFortigateCloudSandbox

		if err = d.Set("gui_fortigate_cloud_sandbox", v); err != nil {
			return diag.Errorf("error reading gui_fortigate_cloud_sandbox: %v", err)
		}
	}

	if o.GuiFortisandboxCloud != nil {
		v := *o.GuiFortisandboxCloud

		if err = d.Set("gui_fortisandbox_cloud", v); err != nil {
			return diag.Errorf("error reading gui_fortisandbox_cloud: %v", err)
		}
	}

	if o.GuiIpv6 != nil {
		v := *o.GuiIpv6

		if err = d.Set("gui_ipv6", v); err != nil {
			return diag.Errorf("error reading gui_ipv6: %v", err)
		}
	}

	if o.GuiLinesPerPage != nil {
		v := *o.GuiLinesPerPage

		if err = d.Set("gui_lines_per_page", v); err != nil {
			return diag.Errorf("error reading gui_lines_per_page: %v", err)
		}
	}

	if o.GuiLocalOut != nil {
		v := *o.GuiLocalOut

		if err = d.Set("gui_local_out", v); err != nil {
			return diag.Errorf("error reading gui_local_out: %v", err)
		}
	}

	if o.GuiReplacementMessageGroups != nil {
		v := *o.GuiReplacementMessageGroups

		if err = d.Set("gui_replacement_message_groups", v); err != nil {
			return diag.Errorf("error reading gui_replacement_message_groups: %v", err)
		}
	}

	if o.GuiRestApiCache != nil {
		v := *o.GuiRestApiCache

		if err = d.Set("gui_rest_api_cache", v); err != nil {
			return diag.Errorf("error reading gui_rest_api_cache: %v", err)
		}
	}

	if o.GuiTheme != nil {
		v := *o.GuiTheme

		if err = d.Set("gui_theme", v); err != nil {
			return diag.Errorf("error reading gui_theme: %v", err)
		}
	}

	if o.GuiWirelessOpensecurity != nil {
		v := *o.GuiWirelessOpensecurity

		if err = d.Set("gui_wireless_opensecurity", v); err != nil {
			return diag.Errorf("error reading gui_wireless_opensecurity: %v", err)
		}
	}

	if o.HaAffinity != nil {
		v := *o.HaAffinity

		if err = d.Set("ha_affinity", v); err != nil {
			return diag.Errorf("error reading ha_affinity: %v", err)
		}
	}

	if o.HonorDf != nil {
		v := *o.HonorDf

		if err = d.Set("honor_df", v); err != nil {
			return diag.Errorf("error reading honor_df: %v", err)
		}
	}

	if o.Hostname != nil {
		v := *o.Hostname

		if err = d.Set("hostname", v); err != nil {
			return diag.Errorf("error reading hostname: %v", err)
		}
	}

	if o.IgmpStateLimit != nil {
		v := *o.IgmpStateLimit

		if err = d.Set("igmp_state_limit", v); err != nil {
			return diag.Errorf("error reading igmp_state_limit: %v", err)
		}
	}

	if o.InternetServiceDatabase != nil {
		v := *o.InternetServiceDatabase

		if err = d.Set("internet_service_database", v); err != nil {
			return diag.Errorf("error reading internet_service_database: %v", err)
		}
	}

	if o.Interval != nil {
		v := *o.Interval

		if err = d.Set("interval", v); err != nil {
			return diag.Errorf("error reading interval: %v", err)
		}
	}

	if o.IpSrcPortRange != nil {
		v := *o.IpSrcPortRange

		if err = d.Set("ip_src_port_range", v); err != nil {
			return diag.Errorf("error reading ip_src_port_range: %v", err)
		}
	}

	if o.IpsAffinity != nil {
		v := *o.IpsAffinity

		if err = d.Set("ips_affinity", v); err != nil {
			return diag.Errorf("error reading ips_affinity: %v", err)
		}
	}

	if o.IpsecAsicOffload != nil {
		v := *o.IpsecAsicOffload

		if err = d.Set("ipsec_asic_offload", v); err != nil {
			return diag.Errorf("error reading ipsec_asic_offload: %v", err)
		}
	}

	if o.IpsecHaSeqjumpRate != nil {
		v := *o.IpsecHaSeqjumpRate

		if err = d.Set("ipsec_ha_seqjump_rate", v); err != nil {
			return diag.Errorf("error reading ipsec_ha_seqjump_rate: %v", err)
		}
	}

	if o.IpsecHmacOffload != nil {
		v := *o.IpsecHmacOffload

		if err = d.Set("ipsec_hmac_offload", v); err != nil {
			return diag.Errorf("error reading ipsec_hmac_offload: %v", err)
		}
	}

	if o.IpsecSoftDecAsync != nil {
		v := *o.IpsecSoftDecAsync

		if err = d.Set("ipsec_soft_dec_async", v); err != nil {
			return diag.Errorf("error reading ipsec_soft_dec_async: %v", err)
		}
	}

	if o.Ipv6AcceptDad != nil {
		v := *o.Ipv6AcceptDad

		if err = d.Set("ipv6_accept_dad", v); err != nil {
			return diag.Errorf("error reading ipv6_accept_dad: %v", err)
		}
	}

	if o.Ipv6AllowAnycastProbe != nil {
		v := *o.Ipv6AllowAnycastProbe

		if err = d.Set("ipv6_allow_anycast_probe", v); err != nil {
			return diag.Errorf("error reading ipv6_allow_anycast_probe: %v", err)
		}
	}

	if o.Ipv6AllowTrafficRedirect != nil {
		v := *o.Ipv6AllowTrafficRedirect

		if err = d.Set("ipv6_allow_traffic_redirect", v); err != nil {
			return diag.Errorf("error reading ipv6_allow_traffic_redirect: %v", err)
		}
	}

	if o.IrqTimeAccounting != nil {
		v := *o.IrqTimeAccounting

		if err = d.Set("irq_time_accounting", v); err != nil {
			return diag.Errorf("error reading irq_time_accounting: %v", err)
		}
	}

	if o.Language != nil {
		v := *o.Language

		if err = d.Set("language", v); err != nil {
			return diag.Errorf("error reading language: %v", err)
		}
	}

	if o.Ldapconntimeout != nil {
		v := *o.Ldapconntimeout

		if err = d.Set("ldapconntimeout", v); err != nil {
			return diag.Errorf("error reading ldapconntimeout: %v", err)
		}
	}

	if o.LldpReception != nil {
		v := *o.LldpReception

		if err = d.Set("lldp_reception", v); err != nil {
			return diag.Errorf("error reading lldp_reception: %v", err)
		}
	}

	if o.LldpTransmission != nil {
		v := *o.LldpTransmission

		if err = d.Set("lldp_transmission", v); err != nil {
			return diag.Errorf("error reading lldp_transmission: %v", err)
		}
	}

	if o.LogSslConnection != nil {
		v := *o.LogSslConnection

		if err = d.Set("log_ssl_connection", v); err != nil {
			return diag.Errorf("error reading log_ssl_connection: %v", err)
		}
	}

	if o.LogUuidAddress != nil {
		v := *o.LogUuidAddress

		if err = d.Set("log_uuid_address", v); err != nil {
			return diag.Errorf("error reading log_uuid_address: %v", err)
		}
	}

	if o.LogUuidPolicy != nil {
		v := *o.LogUuidPolicy

		if err = d.Set("log_uuid_policy", v); err != nil {
			return diag.Errorf("error reading log_uuid_policy: %v", err)
		}
	}

	if o.LoginTimestamp != nil {
		v := *o.LoginTimestamp

		if err = d.Set("login_timestamp", v); err != nil {
			return diag.Errorf("error reading login_timestamp: %v", err)
		}
	}

	if o.LongVdomName != nil {
		v := *o.LongVdomName

		if err = d.Set("long_vdom_name", v); err != nil {
			return diag.Errorf("error reading long_vdom_name: %v", err)
		}
	}

	if o.ManagementIp != nil {
		v := *o.ManagementIp

		if err = d.Set("management_ip", v); err != nil {
			return diag.Errorf("error reading management_ip: %v", err)
		}
	}

	if o.ManagementPort != nil {
		v := *o.ManagementPort

		if err = d.Set("management_port", v); err != nil {
			return diag.Errorf("error reading management_port: %v", err)
		}
	}

	if o.ManagementPortUseAdminSport != nil {
		v := *o.ManagementPortUseAdminSport

		if err = d.Set("management_port_use_admin_sport", v); err != nil {
			return diag.Errorf("error reading management_port_use_admin_sport: %v", err)
		}
	}

	if o.ManagementVdom != nil {
		v := *o.ManagementVdom

		if err = d.Set("management_vdom", v); err != nil {
			return diag.Errorf("error reading management_vdom: %v", err)
		}
	}

	if o.MaxDlpstatMemory != nil {
		v := *o.MaxDlpstatMemory

		if err = d.Set("max_dlpstat_memory", v); err != nil {
			return diag.Errorf("error reading max_dlpstat_memory: %v", err)
		}
	}

	if o.MaxRouteCacheSize != nil {
		v := *o.MaxRouteCacheSize

		if err = d.Set("max_route_cache_size", v); err != nil {
			return diag.Errorf("error reading max_route_cache_size: %v", err)
		}
	}

	if o.MemoryUseThresholdExtreme != nil {
		v := *o.MemoryUseThresholdExtreme

		if err = d.Set("memory_use_threshold_extreme", v); err != nil {
			return diag.Errorf("error reading memory_use_threshold_extreme: %v", err)
		}
	}

	if o.MemoryUseThresholdGreen != nil {
		v := *o.MemoryUseThresholdGreen

		if err = d.Set("memory_use_threshold_green", v); err != nil {
			return diag.Errorf("error reading memory_use_threshold_green: %v", err)
		}
	}

	if o.MemoryUseThresholdRed != nil {
		v := *o.MemoryUseThresholdRed

		if err = d.Set("memory_use_threshold_red", v); err != nil {
			return diag.Errorf("error reading memory_use_threshold_red: %v", err)
		}
	}

	if o.MiglogAffinity != nil {
		v := *o.MiglogAffinity

		if err = d.Set("miglog_affinity", v); err != nil {
			return diag.Errorf("error reading miglog_affinity: %v", err)
		}
	}

	if o.MiglogdChildren != nil {
		v := *o.MiglogdChildren

		if err = d.Set("miglogd_children", v); err != nil {
			return diag.Errorf("error reading miglogd_children: %v", err)
		}
	}

	if o.MultiFactorAuthentication != nil {
		v := *o.MultiFactorAuthentication

		if err = d.Set("multi_factor_authentication", v); err != nil {
			return diag.Errorf("error reading multi_factor_authentication: %v", err)
		}
	}

	if o.NdpMaxEntry != nil {
		v := *o.NdpMaxEntry

		if err = d.Set("ndp_max_entry", v); err != nil {
			return diag.Errorf("error reading ndp_max_entry: %v", err)
		}
	}

	if o.PerUserBal != nil {
		v := *o.PerUserBal

		if err = d.Set("per_user_bal", v); err != nil {
			return diag.Errorf("error reading per_user_bal: %v", err)
		}
	}

	if o.PerUserBwl != nil {
		v := *o.PerUserBwl

		if err = d.Set("per_user_bwl", v); err != nil {
			return diag.Errorf("error reading per_user_bwl: %v", err)
		}
	}

	if o.PmtuDiscovery != nil {
		v := *o.PmtuDiscovery

		if err = d.Set("pmtu_discovery", v); err != nil {
			return diag.Errorf("error reading pmtu_discovery: %v", err)
		}
	}

	if o.PolicyAuthConcurrent != nil {
		v := *o.PolicyAuthConcurrent

		if err = d.Set("policy_auth_concurrent", v); err != nil {
			return diag.Errorf("error reading policy_auth_concurrent: %v", err)
		}
	}

	if o.PostLoginBanner != nil {
		v := *o.PostLoginBanner

		if err = d.Set("post_login_banner", v); err != nil {
			return diag.Errorf("error reading post_login_banner: %v", err)
		}
	}

	if o.PreLoginBanner != nil {
		v := *o.PreLoginBanner

		if err = d.Set("pre_login_banner", v); err != nil {
			return diag.Errorf("error reading pre_login_banner: %v", err)
		}
	}

	if o.PrivateDataEncryption != nil {
		v := *o.PrivateDataEncryption

		if err = d.Set("private_data_encryption", v); err != nil {
			return diag.Errorf("error reading private_data_encryption: %v", err)
		}
	}

	if o.ProxyAuthLifetime != nil {
		v := *o.ProxyAuthLifetime

		if err = d.Set("proxy_auth_lifetime", v); err != nil {
			return diag.Errorf("error reading proxy_auth_lifetime: %v", err)
		}
	}

	if o.ProxyAuthLifetimeTimeout != nil {
		v := *o.ProxyAuthLifetimeTimeout

		if err = d.Set("proxy_auth_lifetime_timeout", v); err != nil {
			return diag.Errorf("error reading proxy_auth_lifetime_timeout: %v", err)
		}
	}

	if o.ProxyAuthTimeout != nil {
		v := *o.ProxyAuthTimeout

		if err = d.Set("proxy_auth_timeout", v); err != nil {
			return diag.Errorf("error reading proxy_auth_timeout: %v", err)
		}
	}

	if o.ProxyCertUseMgmtVdom != nil {
		v := *o.ProxyCertUseMgmtVdom

		if err = d.Set("proxy_cert_use_mgmt_vdom", v); err != nil {
			return diag.Errorf("error reading proxy_cert_use_mgmt_vdom: %v", err)
		}
	}

	if o.ProxyCipherHardwareAcceleration != nil {
		v := *o.ProxyCipherHardwareAcceleration

		if err = d.Set("proxy_cipher_hardware_acceleration", v); err != nil {
			return diag.Errorf("error reading proxy_cipher_hardware_acceleration: %v", err)
		}
	}

	if o.ProxyKxpHardwareAcceleration != nil {
		v := *o.ProxyKxpHardwareAcceleration

		if err = d.Set("proxy_kxp_hardware_acceleration", v); err != nil {
			return diag.Errorf("error reading proxy_kxp_hardware_acceleration: %v", err)
		}
	}

	if o.ProxyReAuthenticationMode != nil {
		v := *o.ProxyReAuthenticationMode

		if err = d.Set("proxy_re_authentication_mode", v); err != nil {
			return diag.Errorf("error reading proxy_re_authentication_mode: %v", err)
		}
	}

	if o.ProxyResourceMode != nil {
		v := *o.ProxyResourceMode

		if err = d.Set("proxy_resource_mode", v); err != nil {
			return diag.Errorf("error reading proxy_resource_mode: %v", err)
		}
	}

	if o.ProxyWorkerCount != nil {
		v := *o.ProxyWorkerCount

		if err = d.Set("proxy_worker_count", v); err != nil {
			return diag.Errorf("error reading proxy_worker_count: %v", err)
		}
	}

	if o.RadiusPort != nil {
		v := *o.RadiusPort

		if err = d.Set("radius_port", v); err != nil {
			return diag.Errorf("error reading radius_port: %v", err)
		}
	}

	if o.RebootUponConfigRestore != nil {
		v := *o.RebootUponConfigRestore

		if err = d.Set("reboot_upon_config_restore", v); err != nil {
			return diag.Errorf("error reading reboot_upon_config_restore: %v", err)
		}
	}

	if o.Refresh != nil {
		v := *o.Refresh

		if err = d.Set("refresh", v); err != nil {
			return diag.Errorf("error reading refresh: %v", err)
		}
	}

	if o.Remoteauthtimeout != nil {
		v := *o.Remoteauthtimeout

		if err = d.Set("remoteauthtimeout", v); err != nil {
			return diag.Errorf("error reading remoteauthtimeout: %v", err)
		}
	}

	if o.ResetSessionlessTcp != nil {
		v := *o.ResetSessionlessTcp

		if err = d.Set("reset_sessionless_tcp", v); err != nil {
			return diag.Errorf("error reading reset_sessionless_tcp: %v", err)
		}
	}

	if o.RestartTime != nil {
		v := *o.RestartTime

		if err = d.Set("restart_time", v); err != nil {
			return diag.Errorf("error reading restart_time: %v", err)
		}
	}

	if o.RevisionBackupOnLogout != nil {
		v := *o.RevisionBackupOnLogout

		if err = d.Set("revision_backup_on_logout", v); err != nil {
			return diag.Errorf("error reading revision_backup_on_logout: %v", err)
		}
	}

	if o.RevisionImageAutoBackup != nil {
		v := *o.RevisionImageAutoBackup

		if err = d.Set("revision_image_auto_backup", v); err != nil {
			return diag.Errorf("error reading revision_image_auto_backup: %v", err)
		}
	}

	if o.ScanunitCount != nil {
		v := *o.ScanunitCount

		if err = d.Set("scanunit_count", v); err != nil {
			return diag.Errorf("error reading scanunit_count: %v", err)
		}
	}

	if o.SecurityRatingResultSubmission != nil {
		v := *o.SecurityRatingResultSubmission

		if err = d.Set("security_rating_result_submission", v); err != nil {
			return diag.Errorf("error reading security_rating_result_submission: %v", err)
		}
	}

	if o.SecurityRatingRunOnSchedule != nil {
		v := *o.SecurityRatingRunOnSchedule

		if err = d.Set("security_rating_run_on_schedule", v); err != nil {
			return diag.Errorf("error reading security_rating_run_on_schedule: %v", err)
		}
	}

	if o.SendPmtuIcmp != nil {
		v := *o.SendPmtuIcmp

		if err = d.Set("send_pmtu_icmp", v); err != nil {
			return diag.Errorf("error reading send_pmtu_icmp: %v", err)
		}
	}

	if o.SnatRouteChange != nil {
		v := *o.SnatRouteChange

		if err = d.Set("snat_route_change", v); err != nil {
			return diag.Errorf("error reading snat_route_change: %v", err)
		}
	}

	if o.SpecialFile23Support != nil {
		v := *o.SpecialFile23Support

		if err = d.Set("special_file_23_support", v); err != nil {
			return diag.Errorf("error reading special_file_23_support: %v", err)
		}
	}

	if o.SpeedtestServer != nil {
		v := *o.SpeedtestServer

		if err = d.Set("speedtest_server", v); err != nil {
			return diag.Errorf("error reading speedtest_server: %v", err)
		}
	}

	if o.SsdTrimDate != nil {
		v := *o.SsdTrimDate

		if err = d.Set("ssd_trim_date", v); err != nil {
			return diag.Errorf("error reading ssd_trim_date: %v", err)
		}
	}

	if o.SsdTrimFreq != nil {
		v := *o.SsdTrimFreq

		if err = d.Set("ssd_trim_freq", v); err != nil {
			return diag.Errorf("error reading ssd_trim_freq: %v", err)
		}
	}

	if o.SsdTrimHour != nil {
		v := *o.SsdTrimHour

		if err = d.Set("ssd_trim_hour", v); err != nil {
			return diag.Errorf("error reading ssd_trim_hour: %v", err)
		}
	}

	if o.SsdTrimMin != nil {
		v := *o.SsdTrimMin

		if err = d.Set("ssd_trim_min", v); err != nil {
			return diag.Errorf("error reading ssd_trim_min: %v", err)
		}
	}

	if o.SsdTrimWeekday != nil {
		v := *o.SsdTrimWeekday

		if err = d.Set("ssd_trim_weekday", v); err != nil {
			return diag.Errorf("error reading ssd_trim_weekday: %v", err)
		}
	}

	if o.SshCbcCipher != nil {
		v := *o.SshCbcCipher

		if err = d.Set("ssh_cbc_cipher", v); err != nil {
			return diag.Errorf("error reading ssh_cbc_cipher: %v", err)
		}
	}

	if o.SshEncAlgo != nil {
		v := *o.SshEncAlgo

		if err = d.Set("ssh_enc_algo", v); err != nil {
			return diag.Errorf("error reading ssh_enc_algo: %v", err)
		}
	}

	if o.SshHmacMd5 != nil {
		v := *o.SshHmacMd5

		if err = d.Set("ssh_hmac_md5", v); err != nil {
			return diag.Errorf("error reading ssh_hmac_md5: %v", err)
		}
	}

	if o.SshKexAlgo != nil {
		v := *o.SshKexAlgo

		if err = d.Set("ssh_kex_algo", v); err != nil {
			return diag.Errorf("error reading ssh_kex_algo: %v", err)
		}
	}

	if o.SshKexSha1 != nil {
		v := *o.SshKexSha1

		if err = d.Set("ssh_kex_sha1", v); err != nil {
			return diag.Errorf("error reading ssh_kex_sha1: %v", err)
		}
	}

	if o.SshMacAlgo != nil {
		v := *o.SshMacAlgo

		if err = d.Set("ssh_mac_algo", v); err != nil {
			return diag.Errorf("error reading ssh_mac_algo: %v", err)
		}
	}

	if o.SshMacWeak != nil {
		v := *o.SshMacWeak

		if err = d.Set("ssh_mac_weak", v); err != nil {
			return diag.Errorf("error reading ssh_mac_weak: %v", err)
		}
	}

	if o.SslMinProtoVersion != nil {
		v := *o.SslMinProtoVersion

		if err = d.Set("ssl_min_proto_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_proto_version: %v", err)
		}
	}

	if o.SslStaticKeyCiphers != nil {
		v := *o.SslStaticKeyCiphers

		if err = d.Set("ssl_static_key_ciphers", v); err != nil {
			return diag.Errorf("error reading ssl_static_key_ciphers: %v", err)
		}
	}

	if o.SslvpnCipherHardwareAcceleration != nil {
		v := *o.SslvpnCipherHardwareAcceleration

		if err = d.Set("sslvpn_cipher_hardware_acceleration", v); err != nil {
			return diag.Errorf("error reading sslvpn_cipher_hardware_acceleration: %v", err)
		}
	}

	if o.SslvpnEmsSnCheck != nil {
		v := *o.SslvpnEmsSnCheck

		if err = d.Set("sslvpn_ems_sn_check", v); err != nil {
			return diag.Errorf("error reading sslvpn_ems_sn_check: %v", err)
		}
	}

	if o.SslvpnKxpHardwareAcceleration != nil {
		v := *o.SslvpnKxpHardwareAcceleration

		if err = d.Set("sslvpn_kxp_hardware_acceleration", v); err != nil {
			return diag.Errorf("error reading sslvpn_kxp_hardware_acceleration: %v", err)
		}
	}

	if o.SslvpnMaxWorkerCount != nil {
		v := *o.SslvpnMaxWorkerCount

		if err = d.Set("sslvpn_max_worker_count", v); err != nil {
			return diag.Errorf("error reading sslvpn_max_worker_count: %v", err)
		}
	}

	if o.SslvpnPluginVersionCheck != nil {
		v := *o.SslvpnPluginVersionCheck

		if err = d.Set("sslvpn_plugin_version_check", v); err != nil {
			return diag.Errorf("error reading sslvpn_plugin_version_check: %v", err)
		}
	}

	if o.StrictDirtySessionCheck != nil {
		v := *o.StrictDirtySessionCheck

		if err = d.Set("strict_dirty_session_check", v); err != nil {
			return diag.Errorf("error reading strict_dirty_session_check: %v", err)
		}
	}

	if o.StrongCrypto != nil {
		v := *o.StrongCrypto

		if err = d.Set("strong_crypto", v); err != nil {
			return diag.Errorf("error reading strong_crypto: %v", err)
		}
	}

	if o.SwitchController != nil {
		v := *o.SwitchController

		if err = d.Set("switch_controller", v); err != nil {
			return diag.Errorf("error reading switch_controller: %v", err)
		}
	}

	if o.SwitchControllerReservedNetwork != nil {
		v := *o.SwitchControllerReservedNetwork
		if current, ok := d.GetOk("switch_controller_reserved_network"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("switch_controller_reserved_network", v); err != nil {
			return diag.Errorf("error reading switch_controller_reserved_network: %v", err)
		}
	}

	if o.SysPerfLogInterval != nil {
		v := *o.SysPerfLogInterval

		if err = d.Set("sys_perf_log_interval", v); err != nil {
			return diag.Errorf("error reading sys_perf_log_interval: %v", err)
		}
	}

	if o.TcpHalfcloseTimer != nil {
		v := *o.TcpHalfcloseTimer

		if err = d.Set("tcp_halfclose_timer", v); err != nil {
			return diag.Errorf("error reading tcp_halfclose_timer: %v", err)
		}
	}

	if o.TcpHalfopenTimer != nil {
		v := *o.TcpHalfopenTimer

		if err = d.Set("tcp_halfopen_timer", v); err != nil {
			return diag.Errorf("error reading tcp_halfopen_timer: %v", err)
		}
	}

	if o.TcpOption != nil {
		v := *o.TcpOption

		if err = d.Set("tcp_option", v); err != nil {
			return diag.Errorf("error reading tcp_option: %v", err)
		}
	}

	if o.TcpRstTimer != nil {
		v := *o.TcpRstTimer

		if err = d.Set("tcp_rst_timer", v); err != nil {
			return diag.Errorf("error reading tcp_rst_timer: %v", err)
		}
	}

	if o.TcpTimewaitTimer != nil {
		v := *o.TcpTimewaitTimer

		if err = d.Set("tcp_timewait_timer", v); err != nil {
			return diag.Errorf("error reading tcp_timewait_timer: %v", err)
		}
	}

	if o.Tftp != nil {
		v := *o.Tftp

		if err = d.Set("tftp", v); err != nil {
			return diag.Errorf("error reading tftp: %v", err)
		}
	}

	if o.Timezone != nil {
		v := *o.Timezone

		if err = d.Set("timezone", v); err != nil {
			return diag.Errorf("error reading timezone: %v", err)
		}
	}

	if o.TrafficPriority != nil {
		v := *o.TrafficPriority

		if err = d.Set("traffic_priority", v); err != nil {
			return diag.Errorf("error reading traffic_priority: %v", err)
		}
	}

	if o.TrafficPriorityLevel != nil {
		v := *o.TrafficPriorityLevel

		if err = d.Set("traffic_priority_level", v); err != nil {
			return diag.Errorf("error reading traffic_priority_level: %v", err)
		}
	}

	if o.TwoFactorEmailExpiry != nil {
		v := *o.TwoFactorEmailExpiry

		if err = d.Set("two_factor_email_expiry", v); err != nil {
			return diag.Errorf("error reading two_factor_email_expiry: %v", err)
		}
	}

	if o.TwoFactorFacExpiry != nil {
		v := *o.TwoFactorFacExpiry

		if err = d.Set("two_factor_fac_expiry", v); err != nil {
			return diag.Errorf("error reading two_factor_fac_expiry: %v", err)
		}
	}

	if o.TwoFactorFtkExpiry != nil {
		v := *o.TwoFactorFtkExpiry

		if err = d.Set("two_factor_ftk_expiry", v); err != nil {
			return diag.Errorf("error reading two_factor_ftk_expiry: %v", err)
		}
	}

	if o.TwoFactorFtmExpiry != nil {
		v := *o.TwoFactorFtmExpiry

		if err = d.Set("two_factor_ftm_expiry", v); err != nil {
			return diag.Errorf("error reading two_factor_ftm_expiry: %v", err)
		}
	}

	if o.TwoFactorSmsExpiry != nil {
		v := *o.TwoFactorSmsExpiry

		if err = d.Set("two_factor_sms_expiry", v); err != nil {
			return diag.Errorf("error reading two_factor_sms_expiry: %v", err)
		}
	}

	if o.UdpIdleTimer != nil {
		v := *o.UdpIdleTimer

		if err = d.Set("udp_idle_timer", v); err != nil {
			return diag.Errorf("error reading udp_idle_timer: %v", err)
		}
	}

	if o.UrlFilterAffinity != nil {
		v := *o.UrlFilterAffinity

		if err = d.Set("url_filter_affinity", v); err != nil {
			return diag.Errorf("error reading url_filter_affinity: %v", err)
		}
	}

	if o.UrlFilterCount != nil {
		v := *o.UrlFilterCount

		if err = d.Set("url_filter_count", v); err != nil {
			return diag.Errorf("error reading url_filter_count: %v", err)
		}
	}

	if o.UserDeviceStoreMaxDevices != nil {
		v := *o.UserDeviceStoreMaxDevices

		if err = d.Set("user_device_store_max_devices", v); err != nil {
			return diag.Errorf("error reading user_device_store_max_devices: %v", err)
		}
	}

	if o.UserDeviceStoreMaxUnifiedMem != nil {
		v := *o.UserDeviceStoreMaxUnifiedMem

		if err = d.Set("user_device_store_max_unified_mem", v); err != nil {
			return diag.Errorf("error reading user_device_store_max_unified_mem: %v", err)
		}
	}

	if o.UserDeviceStoreMaxUsers != nil {
		v := *o.UserDeviceStoreMaxUsers

		if err = d.Set("user_device_store_max_users", v); err != nil {
			return diag.Errorf("error reading user_device_store_max_users: %v", err)
		}
	}

	if o.UserServerCert != nil {
		v := *o.UserServerCert

		if err = d.Set("user_server_cert", v); err != nil {
			return diag.Errorf("error reading user_server_cert: %v", err)
		}
	}

	if o.VdomMode != nil {
		v := *o.VdomMode

		if err = d.Set("vdom_mode", v); err != nil {
			return diag.Errorf("error reading vdom_mode: %v", err)
		}
	}

	if o.VipArpRange != nil {
		v := *o.VipArpRange

		if err = d.Set("vip_arp_range", v); err != nil {
			return diag.Errorf("error reading vip_arp_range: %v", err)
		}
	}

	if o.WadAffinity != nil {
		v := *o.WadAffinity

		if err = d.Set("wad_affinity", v); err != nil {
			return diag.Errorf("error reading wad_affinity: %v", err)
		}
	}

	if o.WadCsvcCsCount != nil {
		v := *o.WadCsvcCsCount

		if err = d.Set("wad_csvc_cs_count", v); err != nil {
			return diag.Errorf("error reading wad_csvc_cs_count: %v", err)
		}
	}

	if o.WadCsvcDbCount != nil {
		v := *o.WadCsvcDbCount

		if err = d.Set("wad_csvc_db_count", v); err != nil {
			return diag.Errorf("error reading wad_csvc_db_count: %v", err)
		}
	}

	if o.WadMemoryChangeGranularity != nil {
		v := *o.WadMemoryChangeGranularity

		if err = d.Set("wad_memory_change_granularity", v); err != nil {
			return diag.Errorf("error reading wad_memory_change_granularity: %v", err)
		}
	}

	if o.WadSourceAffinity != nil {
		v := *o.WadSourceAffinity

		if err = d.Set("wad_source_affinity", v); err != nil {
			return diag.Errorf("error reading wad_source_affinity: %v", err)
		}
	}

	if o.WadWorkerCount != nil {
		v := *o.WadWorkerCount

		if err = d.Set("wad_worker_count", v); err != nil {
			return diag.Errorf("error reading wad_worker_count: %v", err)
		}
	}

	if o.WifiCaCertificate != nil {
		v := *o.WifiCaCertificate

		if err = d.Set("wifi_ca_certificate", v); err != nil {
			return diag.Errorf("error reading wifi_ca_certificate: %v", err)
		}
	}

	if o.WifiCertificate != nil {
		v := *o.WifiCertificate

		if err = d.Set("wifi_certificate", v); err != nil {
			return diag.Errorf("error reading wifi_certificate: %v", err)
		}
	}

	if o.Wimax4gUsb != nil {
		v := *o.Wimax4gUsb

		if err = d.Set("wimax_4g_usb", v); err != nil {
			return diag.Errorf("error reading wimax_4g_usb: %v", err)
		}
	}

	if o.WirelessController != nil {
		v := *o.WirelessController

		if err = d.Set("wireless_controller", v); err != nil {
			return diag.Errorf("error reading wireless_controller: %v", err)
		}
	}

	if o.WirelessControllerPort != nil {
		v := *o.WirelessControllerPort

		if err = d.Set("wireless_controller_port", v); err != nil {
			return diag.Errorf("error reading wireless_controller_port: %v", err)
		}
	}

	return nil
}

func getObjectSystemGlobal(d *schema.ResourceData, sv string) (*models.SystemGlobal, diag.Diagnostics) {
	obj := models.SystemGlobal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("admin_concurrent"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_concurrent", sv)
				diags = append(diags, e)
			}
			obj.AdminConcurrent = &v2
		}
	}
	if v1, ok := d.GetOk("admin_console_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_console_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminConsoleTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_forticloud_sso_login"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("admin_forticloud_sso_login", sv)
				diags = append(diags, e)
			}
			obj.AdminForticloudSsoLogin = &v2
		}
	}
	if v1, ok := d.GetOk("admin_hsts_max_age"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_hsts_max_age", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminHstsMaxAge = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_https_pki_required"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_https_pki_required", sv)
				diags = append(diags, e)
			}
			obj.AdminHttpsPkiRequired = &v2
		}
	}
	if v1, ok := d.GetOk("admin_https_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_https_redirect", sv)
				diags = append(diags, e)
			}
			obj.AdminHttpsRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("admin_https_ssl_banned_ciphers"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("admin_https_ssl_banned_ciphers", sv)
				diags = append(diags, e)
			}
			obj.AdminHttpsSslBannedCiphers = &v2
		}
	}
	if v1, ok := d.GetOk("admin_https_ssl_ciphersuites"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("admin_https_ssl_ciphersuites", sv)
				diags = append(diags, e)
			}
			obj.AdminHttpsSslCiphersuites = &v2
		}
	}
	if v1, ok := d.GetOk("admin_https_ssl_versions"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_https_ssl_versions", sv)
				diags = append(diags, e)
			}
			obj.AdminHttpsSslVersions = &v2
		}
	}
	if v1, ok := d.GetOk("admin_lockout_duration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_lockout_duration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminLockoutDuration = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_lockout_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_lockout_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminLockoutThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_login_max"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_login_max", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminLoginMax = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_maintainer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_maintainer", sv)
				diags = append(diags, e)
			}
			obj.AdminMaintainer = &v2
		}
	}
	if v1, ok := d.GetOk("admin_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminPort = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_restrict_local"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_restrict_local", sv)
				diags = append(diags, e)
			}
			obj.AdminRestrictLocal = &v2
		}
	}
	if v1, ok := d.GetOk("admin_scp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_scp", sv)
				diags = append(diags, e)
			}
			obj.AdminScp = &v2
		}
	}
	if v1, ok := d.GetOk("admin_server_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_server_cert", sv)
				diags = append(diags, e)
			}
			obj.AdminServerCert = &v2
		}
	}
	if v1, ok := d.GetOk("admin_sport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_sport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminSport = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_ssh_grace_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_ssh_grace_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminSshGraceTime = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_ssh_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_ssh_password", sv)
				diags = append(diags, e)
			}
			obj.AdminSshPassword = &v2
		}
	}
	if v1, ok := d.GetOk("admin_ssh_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_ssh_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminSshPort = &tmp
		}
	}
	if v1, ok := d.GetOk("admin_ssh_v1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_ssh_v1", sv)
				diags = append(diags, e)
			}
			obj.AdminSshV1 = &v2
		}
	}
	if v1, ok := d.GetOk("admin_telnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_telnet", sv)
				diags = append(diags, e)
			}
			obj.AdminTelnet = &v2
		}
	}
	if v1, ok := d.GetOk("admin_telnet_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_telnet_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdminTelnetPort = &tmp
		}
	}
	if v1, ok := d.GetOk("admintimeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admintimeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Admintimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("alias"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("alias", sv)
				diags = append(diags, e)
			}
			obj.Alias = &v2
		}
	}
	if v1, ok := d.GetOk("allow_traffic_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_traffic_redirect", sv)
				diags = append(diags, e)
			}
			obj.AllowTrafficRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("anti_replay"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("anti_replay", sv)
				diags = append(diags, e)
			}
			obj.AntiReplay = &v2
		}
	}
	if v1, ok := d.GetOk("arp_max_entry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arp_max_entry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ArpMaxEntry = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_cert", sv)
				diags = append(diags, e)
			}
			obj.AuthCert = &v2
		}
	}
	if v1, ok := d.GetOk("auth_http_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_http_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthHttpPort = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_https_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_https_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthHttpsPort = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_keepalive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_keepalive", sv)
				diags = append(diags, e)
			}
			obj.AuthKeepalive = &v2
		}
	}
	if v1, ok := d.GetOk("auth_session_limit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_session_limit", sv)
				diags = append(diags, e)
			}
			obj.AuthSessionLimit = &v2
		}
	}
	if v1, ok := d.GetOk("auto_auth_extension_device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_auth_extension_device", sv)
				diags = append(diags, e)
			}
			obj.AutoAuthExtensionDevice = &v2
		}
	}
	if v1, ok := d.GetOk("autorun_log_fsck"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("autorun_log_fsck", sv)
				diags = append(diags, e)
			}
			obj.AutorunLogFsck = &v2
		}
	}
	if v1, ok := d.GetOk("av_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_affinity", sv)
				diags = append(diags, e)
			}
			obj.AvAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("av_failopen"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_failopen", sv)
				diags = append(diags, e)
			}
			obj.AvFailopen = &v2
		}
	}
	if v1, ok := d.GetOk("av_failopen_session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_failopen_session", sv)
				diags = append(diags, e)
			}
			obj.AvFailopenSession = &v2
		}
	}
	if v1, ok := d.GetOk("batch_cmdb"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("batch_cmdb", sv)
				diags = append(diags, e)
			}
			obj.BatchCmdb = &v2
		}
	}
	if v1, ok := d.GetOk("block_session_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("block_session_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BlockSessionTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("br_fdb_max_entry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("br_fdb_max_entry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BrFdbMaxEntry = &tmp
		}
	}
	if v1, ok := d.GetOk("cert_chain_max"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert_chain_max", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CertChainMax = &tmp
		}
	}
	if v1, ok := d.GetOk("cfg_revert_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cfg_revert_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CfgRevertTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("cfg_save"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cfg_save", sv)
				diags = append(diags, e)
			}
			obj.CfgSave = &v2
		}
	}
	if v1, ok := d.GetOk("check_protocol_header"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("check_protocol_header", sv)
				diags = append(diags, e)
			}
			obj.CheckProtocolHeader = &v2
		}
	}
	if v1, ok := d.GetOk("check_reset_range"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("check_reset_range", sv)
				diags = append(diags, e)
			}
			obj.CheckResetRange = &v2
		}
	}
	if v1, ok := d.GetOk("cli_audit_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cli_audit_log", sv)
				diags = append(diags, e)
			}
			obj.CliAuditLog = &v2
		}
	}
	if v1, ok := d.GetOk("cloud_communication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cloud_communication", sv)
				diags = append(diags, e)
			}
			obj.CloudCommunication = &v2
		}
	}
	if v1, ok := d.GetOk("clt_cert_req"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("clt_cert_req", sv)
				diags = append(diags, e)
			}
			obj.CltCertReq = &v2
		}
	}
	if v1, ok := d.GetOk("cmdbsvr_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("cmdbsvr_affinity", sv)
				diags = append(diags, e)
			}
			obj.CmdbsvrAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("cpu_use_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cpu_use_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CpuUseThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("csr_ca_attribute"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("csr_ca_attribute", sv)
				diags = append(diags, e)
			}
			obj.CsrCaAttribute = &v2
		}
	}
	if v1, ok := d.GetOk("daily_restart"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("daily_restart", sv)
				diags = append(diags, e)
			}
			obj.DailyRestart = &v2
		}
	}
	if v1, ok := d.GetOk("default_service_source_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_service_source_port", sv)
				diags = append(diags, e)
			}
			obj.DefaultServiceSourcePort = &v2
		}
	}
	if v1, ok := d.GetOk("device_identification_active_scan_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("device_identification_active_scan_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeviceIdentificationActiveScanDelay = &tmp
		}
	}
	if v1, ok := d.GetOk("device_idle_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device_idle_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeviceIdleTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("dh_params"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dh_params", sv)
				diags = append(diags, e)
			}
			obj.DhParams = &v2
		}
	}
	if v1, ok := d.GetOk("dnsproxy_worker_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dnsproxy_worker_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DnsproxyWorkerCount = &tmp
		}
	}
	if v1, ok := d.GetOk("dst"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst", sv)
				diags = append(diags, e)
			}
			obj.Dst = &v2
		}
	}
	if v1, ok := d.GetOk("extender_controller_reserved_network"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("extender_controller_reserved_network", sv)
				diags = append(diags, e)
			}
			obj.ExtenderControllerReservedNetwork = &v2
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
	if v1, ok := d.GetOk("faz_disk_buffer_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("faz_disk_buffer_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FazDiskBufferSize = &tmp
		}
	}
	if v1, ok := d.GetOk("fds_statistics"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fds_statistics", sv)
				diags = append(diags, e)
			}
			obj.FdsStatistics = &v2
		}
	}
	if v1, ok := d.GetOk("fds_statistics_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fds_statistics_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FdsStatisticsPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("fec_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("fec_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FecPort = &tmp
		}
	}
	if v1, ok := d.GetOk("fgd_alert_subscription"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fgd_alert_subscription", sv)
				diags = append(diags, e)
			}
			obj.FgdAlertSubscription = &v2
		}
	}
	if v1, ok := d.GetOk("fortiextender"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("fortiextender", sv)
				diags = append(diags, e)
			}
			obj.Fortiextender = &v2
		}
	}
	if v1, ok := d.GetOk("fortiextender_data_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("fortiextender_data_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FortiextenderDataPort = &tmp
		}
	}
	if v1, ok := d.GetOk("fortiextender_discovery_lockdown"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("fortiextender_discovery_lockdown", sv)
				diags = append(diags, e)
			}
			obj.FortiextenderDiscoveryLockdown = &v2
		}
	}
	if v1, ok := d.GetOk("fortiextender_vlan_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("fortiextender_vlan_mode", sv)
				diags = append(diags, e)
			}
			obj.FortiextenderVlanMode = &v2
		}
	}
	if v1, ok := d.GetOk("fortiipam_integration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "v7.0.2") {
				e := utils.AttributeVersionWarning("fortiipam_integration", sv)
				diags = append(diags, e)
			}
			obj.FortiipamIntegration = &v2
		}
	}
	if v1, ok := d.GetOk("fortiservice_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortiservice_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FortiservicePort = &tmp
		}
	}
	if v1, ok := d.GetOk("fortitoken_cloud"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortitoken_cloud", sv)
				diags = append(diags, e)
			}
			obj.FortitokenCloud = &v2
		}
	}
	if v1, ok := d.GetOk("gui_allow_default_hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_allow_default_hostname", sv)
				diags = append(diags, e)
			}
			obj.GuiAllowDefaultHostname = &v2
		}
	}
	if v1, ok := d.GetOk("gui_cdn_usage"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("gui_cdn_usage", sv)
				diags = append(diags, e)
			}
			obj.GuiCdnUsage = &v2
		}
	}
	if v1, ok := d.GetOk("gui_certificates"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_certificates", sv)
				diags = append(diags, e)
			}
			obj.GuiCertificates = &v2
		}
	}
	if v1, ok := d.GetOk("gui_custom_language"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_custom_language", sv)
				diags = append(diags, e)
			}
			obj.GuiCustomLanguage = &v2
		}
	}
	if v1, ok := d.GetOk("gui_date_format"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_date_format", sv)
				diags = append(diags, e)
			}
			obj.GuiDateFormat = &v2
		}
	}
	if v1, ok := d.GetOk("gui_date_time_source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_date_time_source", sv)
				diags = append(diags, e)
			}
			obj.GuiDateTimeSource = &v2
		}
	}
	if v1, ok := d.GetOk("gui_device_latitude"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_device_latitude", sv)
				diags = append(diags, e)
			}
			obj.GuiDeviceLatitude = &v2
		}
	}
	if v1, ok := d.GetOk("gui_device_longitude"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_device_longitude", sv)
				diags = append(diags, e)
			}
			obj.GuiDeviceLongitude = &v2
		}
	}
	if v1, ok := d.GetOk("gui_display_hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_display_hostname", sv)
				diags = append(diags, e)
			}
			obj.GuiDisplayHostname = &v2
		}
	}
	if v1, ok := d.GetOk("gui_firmware_upgrade_setup_warning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "v6.4.2") {
				e := utils.AttributeVersionWarning("gui_firmware_upgrade_setup_warning", sv)
				diags = append(diags, e)
			}
			obj.GuiFirmwareUpgradeSetupWarning = &v2
		}
	}
	if v1, ok := d.GetOk("gui_firmware_upgrade_warning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("gui_firmware_upgrade_warning", sv)
				diags = append(diags, e)
			}
			obj.GuiFirmwareUpgradeWarning = &v2
		}
	}
	if v1, ok := d.GetOk("gui_forticare_registration_setup_warning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("gui_forticare_registration_setup_warning", sv)
				diags = append(diags, e)
			}
			obj.GuiForticareRegistrationSetupWarning = &v2
		}
	}
	if v1, ok := d.GetOk("gui_fortigate_cloud_sandbox"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gui_fortigate_cloud_sandbox", sv)
				diags = append(diags, e)
			}
			obj.GuiFortigateCloudSandbox = &v2
		}
	}
	if v1, ok := d.GetOk("gui_fortisandbox_cloud"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("gui_fortisandbox_cloud", sv)
				diags = append(diags, e)
			}
			obj.GuiFortisandboxCloud = &v2
		}
	}
	if v1, ok := d.GetOk("gui_ipv6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_ipv6", sv)
				diags = append(diags, e)
			}
			obj.GuiIpv6 = &v2
		}
	}
	if v1, ok := d.GetOk("gui_lines_per_page"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("gui_lines_per_page", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GuiLinesPerPage = &tmp
		}
	}
	if v1, ok := d.GetOk("gui_local_out"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gui_local_out", sv)
				diags = append(diags, e)
			}
			obj.GuiLocalOut = &v2
		}
	}
	if v1, ok := d.GetOk("gui_replacement_message_groups"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gui_replacement_message_groups", sv)
				diags = append(diags, e)
			}
			obj.GuiReplacementMessageGroups = &v2
		}
	}
	if v1, ok := d.GetOk("gui_rest_api_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gui_rest_api_cache", sv)
				diags = append(diags, e)
			}
			obj.GuiRestApiCache = &v2
		}
	}
	if v1, ok := d.GetOk("gui_theme"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_theme", sv)
				diags = append(diags, e)
			}
			obj.GuiTheme = &v2
		}
	}
	if v1, ok := d.GetOk("gui_wireless_opensecurity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_wireless_opensecurity", sv)
				diags = append(diags, e)
			}
			obj.GuiWirelessOpensecurity = &v2
		}
	}
	if v1, ok := d.GetOk("ha_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ha_affinity", sv)
				diags = append(diags, e)
			}
			obj.HaAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("honor_df"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("honor_df", sv)
				diags = append(diags, e)
			}
			obj.HonorDf = &v2
		}
	}
	if v1, ok := d.GetOk("hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostname", sv)
				diags = append(diags, e)
			}
			obj.Hostname = &v2
		}
	}
	if v1, ok := d.GetOk("igmp_state_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("igmp_state_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IgmpStateLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("internet_service_database"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("internet_service_database", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceDatabase = &v2
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
	if v1, ok := d.GetOk("ip_src_port_range"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_src_port_range", sv)
				diags = append(diags, e)
			}
			obj.IpSrcPortRange = &v2
		}
	}
	if v1, ok := d.GetOk("ips_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_affinity", sv)
				diags = append(diags, e)
			}
			obj.IpsAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_asic_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("ipsec_asic_offload", sv)
				diags = append(diags, e)
			}
			obj.IpsecAsicOffload = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_ha_seqjump_rate"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ipsec_ha_seqjump_rate", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IpsecHaSeqjumpRate = &tmp
		}
	}
	if v1, ok := d.GetOk("ipsec_hmac_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("ipsec_hmac_offload", sv)
				diags = append(diags, e)
			}
			obj.IpsecHmacOffload = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_soft_dec_async"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_soft_dec_async", sv)
				diags = append(diags, e)
			}
			obj.IpsecSoftDecAsync = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_accept_dad"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_accept_dad", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Ipv6AcceptDad = &tmp
		}
	}
	if v1, ok := d.GetOk("ipv6_allow_anycast_probe"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_allow_anycast_probe", sv)
				diags = append(diags, e)
			}
			obj.Ipv6AllowAnycastProbe = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_allow_traffic_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ipv6_allow_traffic_redirect", sv)
				diags = append(diags, e)
			}
			obj.Ipv6AllowTrafficRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("irq_time_accounting"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("irq_time_accounting", sv)
				diags = append(diags, e)
			}
			obj.IrqTimeAccounting = &v2
		}
	}
	if v1, ok := d.GetOk("language"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("language", sv)
				diags = append(diags, e)
			}
			obj.Language = &v2
		}
	}
	if v1, ok := d.GetOk("ldapconntimeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldapconntimeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Ldapconntimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("lldp_reception"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp_reception", sv)
				diags = append(diags, e)
			}
			obj.LldpReception = &v2
		}
	}
	if v1, ok := d.GetOk("lldp_transmission"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp_transmission", sv)
				diags = append(diags, e)
			}
			obj.LldpTransmission = &v2
		}
	}
	if v1, ok := d.GetOk("log_ssl_connection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_ssl_connection", sv)
				diags = append(diags, e)
			}
			obj.LogSslConnection = &v2
		}
	}
	if v1, ok := d.GetOk("log_uuid_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_uuid_address", sv)
				diags = append(diags, e)
			}
			obj.LogUuidAddress = &v2
		}
	}
	if v1, ok := d.GetOk("log_uuid_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("log_uuid_policy", sv)
				diags = append(diags, e)
			}
			obj.LogUuidPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("login_timestamp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_timestamp", sv)
				diags = append(diags, e)
			}
			obj.LoginTimestamp = &v2
		}
	}
	if v1, ok := d.GetOk("long_vdom_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("long_vdom_name", sv)
				diags = append(diags, e)
			}
			obj.LongVdomName = &v2
		}
	}
	if v1, ok := d.GetOk("management_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("management_ip", sv)
				diags = append(diags, e)
			}
			obj.ManagementIp = &v2
		}
	}
	if v1, ok := d.GetOk("management_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("management_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ManagementPort = &tmp
		}
	}
	if v1, ok := d.GetOk("management_port_use_admin_sport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("management_port_use_admin_sport", sv)
				diags = append(diags, e)
			}
			obj.ManagementPortUseAdminSport = &v2
		}
	}
	if v1, ok := d.GetOk("management_vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("management_vdom", sv)
				diags = append(diags, e)
			}
			obj.ManagementVdom = &v2
		}
	}
	if v1, ok := d.GetOk("max_dlpstat_memory"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("max_dlpstat_memory", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxDlpstatMemory = &tmp
		}
	}
	if v1, ok := d.GetOk("max_route_cache_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_route_cache_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxRouteCacheSize = &tmp
		}
	}
	if v1, ok := d.GetOk("memory_use_threshold_extreme"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("memory_use_threshold_extreme", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MemoryUseThresholdExtreme = &tmp
		}
	}
	if v1, ok := d.GetOk("memory_use_threshold_green"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("memory_use_threshold_green", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MemoryUseThresholdGreen = &tmp
		}
	}
	if v1, ok := d.GetOk("memory_use_threshold_red"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("memory_use_threshold_red", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MemoryUseThresholdRed = &tmp
		}
	}
	if v1, ok := d.GetOk("miglog_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("miglog_affinity", sv)
				diags = append(diags, e)
			}
			obj.MiglogAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("miglogd_children"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("miglogd_children", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MiglogdChildren = &tmp
		}
	}
	if v1, ok := d.GetOk("multi_factor_authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multi_factor_authentication", sv)
				diags = append(diags, e)
			}
			obj.MultiFactorAuthentication = &v2
		}
	}
	if v1, ok := d.GetOk("ndp_max_entry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ndp_max_entry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NdpMaxEntry = &tmp
		}
	}
	if v1, ok := d.GetOk("per_user_bal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("per_user_bal", sv)
				diags = append(diags, e)
			}
			obj.PerUserBal = &v2
		}
	}
	if v1, ok := d.GetOk("per_user_bwl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("per_user_bwl", sv)
				diags = append(diags, e)
			}
			obj.PerUserBwl = &v2
		}
	}
	if v1, ok := d.GetOk("pmtu_discovery"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.0") {
				e := utils.AttributeVersionWarning("pmtu_discovery", sv)
				diags = append(diags, e)
			}
			obj.PmtuDiscovery = &v2
		}
	}
	if v1, ok := d.GetOk("policy_auth_concurrent"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policy_auth_concurrent", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PolicyAuthConcurrent = &tmp
		}
	}
	if v1, ok := d.GetOk("post_login_banner"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("post_login_banner", sv)
				diags = append(diags, e)
			}
			obj.PostLoginBanner = &v2
		}
	}
	if v1, ok := d.GetOk("pre_login_banner"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pre_login_banner", sv)
				diags = append(diags, e)
			}
			obj.PreLoginBanner = &v2
		}
	}
	if v1, ok := d.GetOk("private_data_encryption"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("private_data_encryption", sv)
				diags = append(diags, e)
			}
			obj.PrivateDataEncryption = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_auth_lifetime"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_auth_lifetime", sv)
				diags = append(diags, e)
			}
			obj.ProxyAuthLifetime = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_auth_lifetime_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_auth_lifetime_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProxyAuthLifetimeTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("proxy_auth_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_auth_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProxyAuthTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("proxy_cert_use_mgmt_vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("proxy_cert_use_mgmt_vdom", sv)
				diags = append(diags, e)
			}
			obj.ProxyCertUseMgmtVdom = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_cipher_hardware_acceleration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("proxy_cipher_hardware_acceleration", sv)
				diags = append(diags, e)
			}
			obj.ProxyCipherHardwareAcceleration = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_kxp_hardware_acceleration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("proxy_kxp_hardware_acceleration", sv)
				diags = append(diags, e)
			}
			obj.ProxyKxpHardwareAcceleration = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_re_authentication_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_re_authentication_mode", sv)
				diags = append(diags, e)
			}
			obj.ProxyReAuthenticationMode = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_resource_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("proxy_resource_mode", sv)
				diags = append(diags, e)
			}
			obj.ProxyResourceMode = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_worker_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_worker_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProxyWorkerCount = &tmp
		}
	}
	if v1, ok := d.GetOk("radius_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RadiusPort = &tmp
		}
	}
	if v1, ok := d.GetOk("reboot_upon_config_restore"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reboot_upon_config_restore", sv)
				diags = append(diags, e)
			}
			obj.RebootUponConfigRestore = &v2
		}
	}
	if v1, ok := d.GetOk("refresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("refresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Refresh = &tmp
		}
	}
	if v1, ok := d.GetOk("remoteauthtimeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remoteauthtimeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Remoteauthtimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("reset_sessionless_tcp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reset_sessionless_tcp", sv)
				diags = append(diags, e)
			}
			obj.ResetSessionlessTcp = &v2
		}
	}
	if v1, ok := d.GetOk("restart_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("restart_time", sv)
				diags = append(diags, e)
			}
			obj.RestartTime = &v2
		}
	}
	if v1, ok := d.GetOk("revision_backup_on_logout"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("revision_backup_on_logout", sv)
				diags = append(diags, e)
			}
			obj.RevisionBackupOnLogout = &v2
		}
	}
	if v1, ok := d.GetOk("revision_image_auto_backup"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("revision_image_auto_backup", sv)
				diags = append(diags, e)
			}
			obj.RevisionImageAutoBackup = &v2
		}
	}
	if v1, ok := d.GetOk("scanunit_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scanunit_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ScanunitCount = &tmp
		}
	}
	if v1, ok := d.GetOk("security_rating_result_submission"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_rating_result_submission", sv)
				diags = append(diags, e)
			}
			obj.SecurityRatingResultSubmission = &v2
		}
	}
	if v1, ok := d.GetOk("security_rating_run_on_schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_rating_run_on_schedule", sv)
				diags = append(diags, e)
			}
			obj.SecurityRatingRunOnSchedule = &v2
		}
	}
	if v1, ok := d.GetOk("send_pmtu_icmp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_pmtu_icmp", sv)
				diags = append(diags, e)
			}
			obj.SendPmtuIcmp = &v2
		}
	}
	if v1, ok := d.GetOk("snat_route_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("snat_route_change", sv)
				diags = append(diags, e)
			}
			obj.SnatRouteChange = &v2
		}
	}
	if v1, ok := d.GetOk("special_file_23_support"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("special_file_23_support", sv)
				diags = append(diags, e)
			}
			obj.SpecialFile23Support = &v2
		}
	}
	if v1, ok := d.GetOk("speedtest_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("speedtest_server", sv)
				diags = append(diags, e)
			}
			obj.SpeedtestServer = &v2
		}
	}
	if v1, ok := d.GetOk("ssd_trim_date"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssd_trim_date", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SsdTrimDate = &tmp
		}
	}
	if v1, ok := d.GetOk("ssd_trim_freq"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssd_trim_freq", sv)
				diags = append(diags, e)
			}
			obj.SsdTrimFreq = &v2
		}
	}
	if v1, ok := d.GetOk("ssd_trim_hour"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssd_trim_hour", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SsdTrimHour = &tmp
		}
	}
	if v1, ok := d.GetOk("ssd_trim_min"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssd_trim_min", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SsdTrimMin = &tmp
		}
	}
	if v1, ok := d.GetOk("ssd_trim_weekday"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssd_trim_weekday", sv)
				diags = append(diags, e)
			}
			obj.SsdTrimWeekday = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_cbc_cipher"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("ssh_cbc_cipher", sv)
				diags = append(diags, e)
			}
			obj.SshCbcCipher = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_enc_algo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ssh_enc_algo", sv)
				diags = append(diags, e)
			}
			obj.SshEncAlgo = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_hmac_md5"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("ssh_hmac_md5", sv)
				diags = append(diags, e)
			}
			obj.SshHmacMd5 = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_kex_algo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ssh_kex_algo", sv)
				diags = append(diags, e)
			}
			obj.SshKexAlgo = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_kex_sha1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("ssh_kex_sha1", sv)
				diags = append(diags, e)
			}
			obj.SshKexSha1 = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_mac_algo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ssh_mac_algo", sv)
				diags = append(diags, e)
			}
			obj.SshMacAlgo = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_mac_weak"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("ssh_mac_weak", sv)
				diags = append(diags, e)
			}
			obj.SshMacWeak = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_proto_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_proto_version", sv)
				diags = append(diags, e)
			}
			obj.SslMinProtoVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_static_key_ciphers"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_static_key_ciphers", sv)
				diags = append(diags, e)
			}
			obj.SslStaticKeyCiphers = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_cipher_hardware_acceleration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("sslvpn_cipher_hardware_acceleration", sv)
				diags = append(diags, e)
			}
			obj.SslvpnCipherHardwareAcceleration = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_ems_sn_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("sslvpn_ems_sn_check", sv)
				diags = append(diags, e)
			}
			obj.SslvpnEmsSnCheck = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_kxp_hardware_acceleration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("sslvpn_kxp_hardware_acceleration", sv)
				diags = append(diags, e)
			}
			obj.SslvpnKxpHardwareAcceleration = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_max_worker_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sslvpn_max_worker_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslvpnMaxWorkerCount = &tmp
		}
	}
	if v1, ok := d.GetOk("sslvpn_plugin_version_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sslvpn_plugin_version_check", sv)
				diags = append(diags, e)
			}
			obj.SslvpnPluginVersionCheck = &v2
		}
	}
	if v1, ok := d.GetOk("strict_dirty_session_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strict_dirty_session_check", sv)
				diags = append(diags, e)
			}
			obj.StrictDirtySessionCheck = &v2
		}
	}
	if v1, ok := d.GetOk("strong_crypto"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strong_crypto", sv)
				diags = append(diags, e)
			}
			obj.StrongCrypto = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller", sv)
				diags = append(diags, e)
			}
			obj.SwitchController = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_reserved_network"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_reserved_network", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerReservedNetwork = &v2
		}
	}
	if v1, ok := d.GetOk("sys_perf_log_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sys_perf_log_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SysPerfLogInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_halfclose_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_halfclose_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpHalfcloseTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_halfopen_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_halfopen_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpHalfopenTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_option", sv)
				diags = append(diags, e)
			}
			obj.TcpOption = &v2
		}
	}
	if v1, ok := d.GetOk("tcp_rst_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("tcp_rst_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpRstTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_timewait_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_timewait_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpTimewaitTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("tftp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tftp", sv)
				diags = append(diags, e)
			}
			obj.Tftp = &v2
		}
	}
	if v1, ok := d.GetOk("timezone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timezone", sv)
				diags = append(diags, e)
			}
			obj.Timezone = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_priority"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_priority", sv)
				diags = append(diags, e)
			}
			obj.TrafficPriority = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_priority_level"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_priority_level", sv)
				diags = append(diags, e)
			}
			obj.TrafficPriorityLevel = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_email_expiry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_email_expiry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TwoFactorEmailExpiry = &tmp
		}
	}
	if v1, ok := d.GetOk("two_factor_fac_expiry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_fac_expiry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TwoFactorFacExpiry = &tmp
		}
	}
	if v1, ok := d.GetOk("two_factor_ftk_expiry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_ftk_expiry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TwoFactorFtkExpiry = &tmp
		}
	}
	if v1, ok := d.GetOk("two_factor_ftm_expiry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_ftm_expiry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TwoFactorFtmExpiry = &tmp
		}
	}
	if v1, ok := d.GetOk("two_factor_sms_expiry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_sms_expiry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TwoFactorSmsExpiry = &tmp
		}
	}
	if v1, ok := d.GetOk("udp_idle_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("udp_idle_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UdpIdleTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("url_filter_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url_filter_affinity", sv)
				diags = append(diags, e)
			}
			obj.UrlFilterAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("url_filter_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url_filter_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UrlFilterCount = &tmp
		}
	}
	if v1, ok := d.GetOk("user_device_store_max_devices"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("user_device_store_max_devices", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UserDeviceStoreMaxDevices = &tmp
		}
	}
	if v1, ok := d.GetOk("user_device_store_max_unified_mem"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("user_device_store_max_unified_mem", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UserDeviceStoreMaxUnifiedMem = &tmp
		}
	}
	if v1, ok := d.GetOk("user_device_store_max_users"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("user_device_store_max_users", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UserDeviceStoreMaxUsers = &tmp
		}
	}
	if v1, ok := d.GetOk("user_server_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_server_cert", sv)
				diags = append(diags, e)
			}
			obj.UserServerCert = &v2
		}
	}
	if v1, ok := d.GetOk("vdom_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("vdom_mode", sv)
				diags = append(diags, e)
			}
			obj.VdomMode = &v2
		}
	}
	if v1, ok := d.GetOk("vip_arp_range"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vip_arp_range", sv)
				diags = append(diags, e)
			}
			obj.VipArpRange = &v2
		}
	}
	if v1, ok := d.GetOk("wad_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wad_affinity", sv)
				diags = append(diags, e)
			}
			obj.WadAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("wad_csvc_cs_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wad_csvc_cs_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WadCsvcCsCount = &tmp
		}
	}
	if v1, ok := d.GetOk("wad_csvc_db_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wad_csvc_db_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WadCsvcDbCount = &tmp
		}
	}
	if v1, ok := d.GetOk("wad_memory_change_granularity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wad_memory_change_granularity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WadMemoryChangeGranularity = &tmp
		}
	}
	if v1, ok := d.GetOk("wad_source_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wad_source_affinity", sv)
				diags = append(diags, e)
			}
			obj.WadSourceAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("wad_worker_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wad_worker_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WadWorkerCount = &tmp
		}
	}
	if v1, ok := d.GetOk("wifi_ca_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wifi_ca_certificate", sv)
				diags = append(diags, e)
			}
			obj.WifiCaCertificate = &v2
		}
	}
	if v1, ok := d.GetOk("wifi_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wifi_certificate", sv)
				diags = append(diags, e)
			}
			obj.WifiCertificate = &v2
		}
	}
	if v1, ok := d.GetOk("wimax_4g_usb"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wimax_4g_usb", sv)
				diags = append(diags, e)
			}
			obj.Wimax4gUsb = &v2
		}
	}
	if v1, ok := d.GetOk("wireless_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wireless_controller", sv)
				diags = append(diags, e)
			}
			obj.WirelessController = &v2
		}
	}
	if v1, ok := d.GetOk("wireless_controller_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wireless_controller_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WirelessControllerPort = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemGlobal(d *schema.ResourceData, sv string) (*models.SystemGlobal, diag.Diagnostics) {
	obj := models.SystemGlobal{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
