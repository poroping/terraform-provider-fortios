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
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure global attributes.",

		ReadContext: dataSourceSystemGlobalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"admin_concurrent": {
				Type:        schema.TypeString,
				Description: "Enable/disable concurrent administrator logins. Use policy-auth-concurrent for firewall authenticated users.",
				Computed:    true,
			},
			"admin_console_timeout": {
				Type:        schema.TypeInt,
				Description: "Console login timeout that overrides the admin timeout value (15 - 300 seconds, default = 0, which disables the timeout).",
				Computed:    true,
			},
			"admin_forticloud_sso_login": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiCloud admin login via SSO.",
				Computed:    true,
			},
			"admin_hsts_max_age": {
				Type:        schema.TypeInt,
				Description: "HTTPS Strict-Transport-Security header max-age in seconds. A value of 0 will reset any HSTS records in the browser.When admin-https-redirect is disabled the header max-age will be 0.",
				Computed:    true,
			},
			"admin_https_pki_required": {
				Type:        schema.TypeString,
				Description: "Enable/disable admin login method. Enable to force administrators to provide a valid certificate to log in if PKI is enabled. Disable to allow administrators to log in with a certificate or password.",
				Computed:    true,
			},
			"admin_https_redirect": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirection of HTTP administration access to HTTPS.",
				Computed:    true,
			},
			"admin_https_ssl_banned_ciphers": {
				Type:        schema.TypeString,
				Description: "Select one or more cipher technologies that cannot be used in GUI HTTPS negotiations. Only applies to TLS 1.2 and below.",
				Computed:    true,
			},
			"admin_https_ssl_ciphersuites": {
				Type:        schema.TypeString,
				Description: "Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, remove TLS1.3 from admin-https-ssl-versions.",
				Computed:    true,
			},
			"admin_https_ssl_versions": {
				Type:        schema.TypeString,
				Description: "Allowed TLS versions for web administration.",
				Computed:    true,
			},
			"admin_lockout_duration": {
				Type:        schema.TypeInt,
				Description: "Amount of time in seconds that an administrator account is locked out after reaching the admin-lockout-threshold for repeated failed login attempts.",
				Computed:    true,
			},
			"admin_lockout_threshold": {
				Type:        schema.TypeInt,
				Description: "Number of failed login attempts before an administrator account is locked out for the admin-lockout-duration.",
				Computed:    true,
			},
			"admin_login_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of administrators who can be logged in at the same time (1 - 100, default = 100).",
				Computed:    true,
			},
			"admin_maintainer": {
				Type:        schema.TypeString,
				Description: "Enable/disable maintainer administrator login. When enabled, the maintainer account can be used to log in from the console after a hard reboot. The password is \"bcpb\" followed by the FortiGate unit serial number. You have limited time to complete this login.",
				Computed:    true,
			},
			"admin_port": {
				Type:        schema.TypeInt,
				Description: "Administrative access port for HTTP. (1 - 65535, default = 80).",
				Computed:    true,
			},
			"admin_restrict_local": {
				Type:        schema.TypeString,
				Description: "Enable/disable local admin authentication restriction when remote authenticator is up and running (default = disable).",
				Computed:    true,
			},
			"admin_scp": {
				Type:        schema.TypeString,
				Description: "Enable/disable using SCP to download the system configuration. You can use SCP as an alternative method for backing up the configuration.",
				Computed:    true,
			},
			"admin_server_cert": {
				Type:        schema.TypeString,
				Description: "Server certificate that the FortiGate uses for HTTPS administrative connections.",
				Computed:    true,
			},
			"admin_sport": {
				Type:        schema.TypeInt,
				Description: "Administrative access port for HTTPS. (1 - 65535, default = 443).",
				Computed:    true,
			},
			"admin_ssh_grace_time": {
				Type:        schema.TypeInt,
				Description: "Maximum time in seconds permitted between making an SSH connection to the FortiGate unit and authenticating (10 - 3600 sec (1 hour), default 120).",
				Computed:    true,
			},
			"admin_ssh_password": {
				Type:        schema.TypeString,
				Description: "Enable/disable password authentication for SSH admin access.",
				Computed:    true,
			},
			"admin_ssh_port": {
				Type:        schema.TypeInt,
				Description: "Administrative access port for SSH. (1 - 65535, default = 22).",
				Computed:    true,
			},
			"admin_ssh_v1": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSH v1 compatibility.",
				Computed:    true,
			},
			"admin_telnet": {
				Type:        schema.TypeString,
				Description: "Enable/disable TELNET service.",
				Computed:    true,
			},
			"admin_telnet_port": {
				Type:        schema.TypeInt,
				Description: "Administrative access port for TELNET. (1 - 65535, default = 23).",
				Computed:    true,
			},
			"admintimeout": {
				Type:        schema.TypeInt,
				Description: "Number of minutes before an idle administrator session times out (1 - 480 minutes (8 hours), default = 5). A shorter idle timeout is more secure.",
				Computed:    true,
			},
			"alias": {
				Type:        schema.TypeString,
				Description: "Alias for your FortiGate unit.",
				Computed:    true,
			},
			"allow_traffic_redirect": {
				Type:        schema.TypeString,
				Description: "Disable to prevent traffic with same local ingress and egress interface from being forwarded without policy check.",
				Computed:    true,
			},
			"anti_replay": {
				Type:        schema.TypeString,
				Description: "Level of checking for packet replay and TCP sequence checking.",
				Computed:    true,
			},
			"arp_max_entry": {
				Type:        schema.TypeInt,
				Description: "Maximum number of dynamically learned MAC addresses that can be added to the ARP table (131072 - 2147483647, default = 131072).",
				Computed:    true,
			},
			"auth_cert": {
				Type:        schema.TypeString,
				Description: "Server certificate that the FortiGate uses for HTTPS firewall authentication connections.",
				Computed:    true,
			},
			"auth_http_port": {
				Type:        schema.TypeInt,
				Description: "User authentication HTTP port. (1 - 65535, default = 1000).",
				Computed:    true,
			},
			"auth_https_port": {
				Type:        schema.TypeInt,
				Description: "User authentication HTTPS port. (1 - 65535, default = 1003).",
				Computed:    true,
			},
			"auth_keepalive": {
				Type:        schema.TypeString,
				Description: "Enable to prevent user authentication sessions from timing out when idle.",
				Computed:    true,
			},
			"auth_session_limit": {
				Type:        schema.TypeString,
				Description: "Action to take when the number of allowed user authenticated sessions is reached.",
				Computed:    true,
			},
			"auto_auth_extension_device": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic authorization of dedicated Fortinet extension devices.",
				Computed:    true,
			},
			"autorun_log_fsck": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic log partition check after ungraceful shutdown.",
				Computed:    true,
			},
			"av_affinity": {
				Type:        schema.TypeString,
				Description: "Affinity setting for AV scanning (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Computed:    true,
			},
			"av_failopen": {
				Type:        schema.TypeString,
				Description: "Set the action to take if the FortiGate is running low on memory or the proxy connection limit has been reached.",
				Computed:    true,
			},
			"av_failopen_session": {
				Type:        schema.TypeString,
				Description: "When enabled and a proxy for a protocol runs out of room in its session table, that protocol goes into failopen mode and enacts the action specified by av-failopen.",
				Computed:    true,
			},
			"batch_cmdb": {
				Type:        schema.TypeString,
				Description: "Enable/disable batch mode, allowing you to enter a series of CLI commands that will execute as a group once they are loaded.",
				Computed:    true,
			},
			"block_session_timer": {
				Type:        schema.TypeInt,
				Description: "Duration in seconds for blocked sessions (1 - 300 sec  (5 minutes), default = 30).",
				Computed:    true,
			},
			"br_fdb_max_entry": {
				Type:        schema.TypeInt,
				Description: "Maximum number of bridge forwarding database (FDB) entries.",
				Computed:    true,
			},
			"cert_chain_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of certificates that can be traversed in a certificate chain.",
				Computed:    true,
			},
			"cfg_revert_timeout": {
				Type:        schema.TypeInt,
				Description: "Time-out for reverting to the last saved configuration. (10 - 4294967295 seconds, default = 600).",
				Computed:    true,
			},
			"cfg_save": {
				Type:        schema.TypeString,
				Description: "Configuration file save mode for CLI changes.",
				Computed:    true,
			},
			"check_protocol_header": {
				Type:        schema.TypeString,
				Description: "Level of checking performed on protocol headers. Strict checking is more thorough but may affect performance. Loose checking is OK in most cases.",
				Computed:    true,
			},
			"check_reset_range": {
				Type:        schema.TypeString,
				Description: "Configure ICMP error message verification. You can either apply strict RST range checking or disable it.",
				Computed:    true,
			},
			"cli_audit_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable CLI audit log.",
				Computed:    true,
			},
			"cloud_communication": {
				Type:        schema.TypeString,
				Description: "Enable/disable all cloud communication.",
				Computed:    true,
			},
			"clt_cert_req": {
				Type:        schema.TypeString,
				Description: "Enable/disable requiring administrators to have a client certificate to log into the GUI using HTTPS.",
				Computed:    true,
			},
			"cmdbsvr_affinity": {
				Type:        schema.TypeString,
				Description: "Affinity setting for cmdbsvr (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Computed:    true,
			},
			"cpu_use_threshold": {
				Type:        schema.TypeInt,
				Description: "Threshold at which CPU usage is reported (% of total CPU, default = 90).",
				Computed:    true,
			},
			"csr_ca_attribute": {
				Type:        schema.TypeString,
				Description: "Enable/disable the CA attribute in certificates. Some CA servers reject CSRs that have the CA attribute.",
				Computed:    true,
			},
			"daily_restart": {
				Type:        schema.TypeString,
				Description: "Enable/disable daily restart of FortiGate unit. Use the restart-time option to set the time of day for the restart.",
				Computed:    true,
			},
			"default_service_source_port": {
				Type:        schema.TypeString,
				Description: "Default service source port range (default = 1 - 65535).",
				Computed:    true,
			},
			"device_identification_active_scan_delay": {
				Type:        schema.TypeInt,
				Description: "Number of seconds to passively scan a device before performing an active scan. (20 - 3600 sec, (20 sec to 1 hour), default = 90).",
				Computed:    true,
			},
			"device_idle_timeout": {
				Type:        schema.TypeInt,
				Description: "Time in seconds that a device must be idle to automatically log the device user out. (30 - 31536000 sec (30 sec to 1 year), default = 300).",
				Computed:    true,
			},
			"dh_params": {
				Type:        schema.TypeString,
				Description: "Number of bits to use in the Diffie-Hellman exchange for HTTPS/SSH protocols.",
				Computed:    true,
			},
			"dnsproxy_worker_count": {
				Type:        schema.TypeInt,
				Description: "DNS proxy worker count. For a FortiGate with multiple logical CPUs, you can set the DNS process number from 1 to the number of logical CPUs.",
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeString,
				Description: "Enable/disable daylight saving time.",
				Computed:    true,
			},
			"extender_controller_reserved_network": {
				Type:        schema.TypeString,
				Description: "Configure reserved network subnet for managed LAN extension FortiExtender units. This is available when the FortiExtender daemon is running.",
				Computed:    true,
			},
			"failtime": {
				Type:        schema.TypeInt,
				Description: "Fail-time for server lost.",
				Computed:    true,
			},
			"faz_disk_buffer_size": {
				Type:        schema.TypeInt,
				Description: "Maximum disk buffer size to temporarily store logs destined for FortiAnalyzer. To be used in the event that FortiAnalyzer is unavailable.",
				Computed:    true,
			},
			"fds_statistics": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending IPS, Application Control, and AntiVirus data to FortiGuard. This data is used to improve FortiGuard services and is not shared with external parties and is protected by Fortinet's privacy policy.",
				Computed:    true,
			},
			"fds_statistics_period": {
				Type:        schema.TypeInt,
				Description: "FortiGuard statistics collection period in minutes. (1 - 1440 min (1 min to 24 hours), default = 60).",
				Computed:    true,
			},
			"fec_port": {
				Type:        schema.TypeInt,
				Description: "Local UDP port for Forward Error Correction (49152 - 65535).",
				Computed:    true,
			},
			"fgd_alert_subscription": {
				Type:        schema.TypeString,
				Description: "Type of alert to retrieve from FortiGuard.",
				Computed:    true,
			},
			"fortiextender": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiExtender.",
				Computed:    true,
			},
			"fortiextender_data_port": {
				Type:        schema.TypeInt,
				Description: "FortiExtender data port (1024 - 49150, default = 25246).",
				Computed:    true,
			},
			"fortiextender_discovery_lockdown": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiExtender CAPWAP lockdown.",
				Computed:    true,
			},
			"fortiextender_vlan_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiExtender VLAN mode.",
				Computed:    true,
			},
			"fortiipam_integration": {
				Type:        schema.TypeString,
				Description: "Enable/disable integration with the FortiIPAM cloud service.",
				Computed:    true,
			},
			"fortiservice_port": {
				Type:        schema.TypeInt,
				Description: "FortiService port (1 - 65535, default = 8013). Used by FortiClient endpoint compliance. Older versions of FortiClient used a different port.",
				Computed:    true,
			},
			"fortitoken_cloud": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiToken Cloud service.",
				Computed:    true,
			},
			"gui_allow_default_hostname": {
				Type:        schema.TypeString,
				Description: "Enable/disable the factory default hostname warning on the GUI setup wizard.",
				Computed:    true,
			},
			"gui_cdn_usage": {
				Type:        schema.TypeString,
				Description: "Enable/disable Load GUI static files from a CDN.",
				Computed:    true,
			},
			"gui_certificates": {
				Type:        schema.TypeString,
				Description: "Enable/disable the System > Certificate GUI page, allowing you to add and configure certificates from the GUI.",
				Computed:    true,
			},
			"gui_custom_language": {
				Type:        schema.TypeString,
				Description: "Enable/disable custom languages in GUI.",
				Computed:    true,
			},
			"gui_date_format": {
				Type:        schema.TypeString,
				Description: "Default date format used throughout GUI.",
				Computed:    true,
			},
			"gui_date_time_source": {
				Type:        schema.TypeString,
				Description: "Source from which the FortiGate GUI uses to display date and time entries.",
				Computed:    true,
			},
			"gui_device_latitude": {
				Type:        schema.TypeString,
				Description: "Add the latitude of the location of this FortiGate to position it on the Threat Map.",
				Computed:    true,
			},
			"gui_device_longitude": {
				Type:        schema.TypeString,
				Description: "Add the longitude of the location of this FortiGate to position it on the Threat Map.",
				Computed:    true,
			},
			"gui_display_hostname": {
				Type:        schema.TypeString,
				Description: "Enable/disable displaying the FortiGate's hostname on the GUI login page.",
				Computed:    true,
			},
			"gui_firmware_upgrade_setup_warning": {
				Type:        schema.TypeString,
				Description: "Enable/disable the firmware upgrade warning on GUI setup wizard.",
				Computed:    true,
			},
			"gui_firmware_upgrade_warning": {
				Type:        schema.TypeString,
				Description: "Enable/disable the firmware upgrade warning on the GUI.",
				Computed:    true,
			},
			"gui_forticare_registration_setup_warning": {
				Type:        schema.TypeString,
				Description: "Enable/disable the FortiCare registration setup warning on the GUI.",
				Computed:    true,
			},
			"gui_fortigate_cloud_sandbox": {
				Type:        schema.TypeString,
				Description: "Enable/disable displaying FortiGate Cloud Sandbox on the GUI.",
				Computed:    true,
			},
			"gui_fortisandbox_cloud": {
				Type:        schema.TypeString,
				Description: "Enable/disable displaying FortiSandbox Cloud on the GUI.",
				Computed:    true,
			},
			"gui_ipv6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 settings on the GUI.",
				Computed:    true,
			},
			"gui_lines_per_page": {
				Type:        schema.TypeInt,
				Description: "Number of lines to display per page for web administration.",
				Computed:    true,
			},
			"gui_local_out": {
				Type:        schema.TypeString,
				Description: "Enable/disable Local-out traffic on the GUI.",
				Computed:    true,
			},
			"gui_replacement_message_groups": {
				Type:        schema.TypeString,
				Description: "Enable/disable replacement message groups on the GUI.",
				Computed:    true,
			},
			"gui_rest_api_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable REST API result caching on FortiGate.",
				Computed:    true,
			},
			"gui_theme": {
				Type:        schema.TypeString,
				Description: "Color scheme for the administration GUI.",
				Computed:    true,
			},
			"gui_wireless_opensecurity": {
				Type:        schema.TypeString,
				Description: "Enable/disable wireless open security option on the GUI.",
				Computed:    true,
			},
			"ha_affinity": {
				Type:        schema.TypeString,
				Description: "Affinity setting for HA daemons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Computed:    true,
			},
			"honor_df": {
				Type:        schema.TypeString,
				Description: "Enable/disable honoring of Don't-Fragment (DF) flag.",
				Computed:    true,
			},
			"hostname": {
				Type:        schema.TypeString,
				Description: "FortiGate unit's hostname. Most models will truncate names longer than 24 characters. Some models support hostnames up to 35 characters.",
				Computed:    true,
			},
			"igmp_state_limit": {
				Type:        schema.TypeInt,
				Description: "Maximum number of IGMP memberships (96 - 64000, default = 3200).",
				Computed:    true,
			},
			"internet_service_database": {
				Type:        schema.TypeString,
				Description: "Configure which Internet Service database size to download from FortiGuard and use.",
				Computed:    true,
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "Dead gateway detection interval.",
				Computed:    true,
			},
			"ip_src_port_range": {
				Type:        schema.TypeString,
				Description: "IP source port range used for traffic originating from the FortiGate unit.",
				Computed:    true,
			},
			"ips_affinity": {
				Type:        schema.TypeString,
				Description: "Affinity setting for IPS (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx; allowed CPUs must be less than total number of IPS engine daemons).",
				Computed:    true,
			},
			"ipsec_asic_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable ASIC offloading (hardware acceleration) for IPsec VPN traffic. Hardware acceleration can offload IPsec VPN sessions and accelerate encryption and decryption.",
				Computed:    true,
			},
			"ipsec_ha_seqjump_rate": {
				Type:        schema.TypeInt,
				Description: "ESP jump ahead rate (1G - 10G pps equivalent).",
				Computed:    true,
			},
			"ipsec_hmac_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable offloading (hardware acceleration) of HMAC processing for IPsec VPN.",
				Computed:    true,
			},
			"ipsec_soft_dec_async": {
				Type:        schema.TypeString,
				Description: "Enable/disable software decryption asynchronization (using multiple CPUs to do decryption) for IPsec VPN traffic.",
				Computed:    true,
			},
			"ipv6_accept_dad": {
				Type:        schema.TypeInt,
				Description: "Enable/disable acceptance of IPv6 Duplicate Address Detection (DAD).",
				Computed:    true,
			},
			"ipv6_allow_anycast_probe": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 address probe through Anycast.",
				Computed:    true,
			},
			"ipv6_allow_traffic_redirect": {
				Type:        schema.TypeString,
				Description: "Disable to prevent IPv6 traffic with same local ingress and egress interface from being forwarded without policy check.",
				Computed:    true,
			},
			"irq_time_accounting": {
				Type:        schema.TypeString,
				Description: "Configure CPU IRQ time accounting mode.",
				Computed:    true,
			},
			"language": {
				Type:        schema.TypeString,
				Description: "GUI display language.",
				Computed:    true,
			},
			"ldapconntimeout": {
				Type:        schema.TypeInt,
				Description: "Global timeout for connections with remote LDAP servers in milliseconds (1 - 300000, default 500).",
				Computed:    true,
			},
			"lldp_reception": {
				Type:        schema.TypeString,
				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) reception.",
				Computed:    true,
			},
			"lldp_transmission": {
				Type:        schema.TypeString,
				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) transmission.",
				Computed:    true,
			},
			"log_ssl_connection": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of SSL connection events.",
				Computed:    true,
			},
			"log_uuid_address": {
				Type:        schema.TypeString,
				Description: "Enable/disable insertion of address UUIDs to traffic logs.",
				Computed:    true,
			},
			"log_uuid_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable insertion of policy UUIDs to traffic logs.",
				Computed:    true,
			},
			"login_timestamp": {
				Type:        schema.TypeString,
				Description: "Enable/disable login time recording.",
				Computed:    true,
			},
			"long_vdom_name": {
				Type:        schema.TypeString,
				Description: "Enable/disable long VDOM name support.",
				Computed:    true,
			},
			"management_ip": {
				Type:        schema.TypeString,
				Description: "Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.",
				Computed:    true,
			},
			"management_port": {
				Type:        schema.TypeInt,
				Description: "Overriding port for management connection (Overrides admin port).",
				Computed:    true,
			},
			"management_port_use_admin_sport": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of the admin-sport setting for the management port. If disabled, FortiGate will allow user to specify management-port.",
				Computed:    true,
			},
			"management_vdom": {
				Type:        schema.TypeString,
				Description: "Management virtual domain name.",
				Computed:    true,
			},
			"max_dlpstat_memory": {
				Type:        schema.TypeInt,
				Description: "Maximum DLP stat memory (0 - 4294967295).",
				Computed:    true,
			},
			"max_route_cache_size": {
				Type:        schema.TypeInt,
				Description: "Maximum number of IP route cache entries (0 - 2147483647).",
				Computed:    true,
			},
			"memory_use_threshold_extreme": {
				Type:        schema.TypeInt,
				Description: "Threshold at which memory usage is considered extreme (new sessions are dropped) (% of total RAM, default = 95).",
				Computed:    true,
			},
			"memory_use_threshold_green": {
				Type:        schema.TypeInt,
				Description: "Threshold at which memory usage forces the FortiGate to exit conserve mode (% of total RAM, default = 82).",
				Computed:    true,
			},
			"memory_use_threshold_red": {
				Type:        schema.TypeInt,
				Description: "Threshold at which memory usage forces the FortiGate to enter conserve mode (% of total RAM, default = 88).",
				Computed:    true,
			},
			"miglog_affinity": {
				Type:        schema.TypeString,
				Description: "Affinity setting for logging (64-bit hexadecimal value in the format of xxxxxxxxxxxxxxxx).",
				Computed:    true,
			},
			"miglogd_children": {
				Type:        schema.TypeInt,
				Description: "Number of logging (miglogd) processes to be allowed to run. Higher number can reduce performance; lower number can slow log processing time. No logs will be dropped or lost if the number is changed.",
				Computed:    true,
			},
			"multi_factor_authentication": {
				Type:        schema.TypeString,
				Description: "Enforce all login methods to require an additional authentication factor (default = optional).",
				Computed:    true,
			},
			"ndp_max_entry": {
				Type:        schema.TypeInt,
				Description: "Maximum number of NDP table entries (set to 65,536 or higher; if set to 0, kernel holds 65,536 entries).",
				Computed:    true,
			},
			"per_user_bal": {
				Type:        schema.TypeString,
				Description: "Enable/disable per-user block/allow list filter.",
				Computed:    true,
			},
			"per_user_bwl": {
				Type:        schema.TypeString,
				Description: "Enable/disable per-user black/white list filter.",
				Computed:    true,
			},
			"pmtu_discovery": {
				Type:        schema.TypeString,
				Description: "Enable/disable path MTU discovery.",
				Computed:    true,
			},
			"policy_auth_concurrent": {
				Type:        schema.TypeInt,
				Description: "Number of concurrent firewall use logins from the same user (1 - 100, default = 0 means no limit).",
				Computed:    true,
			},
			"post_login_banner": {
				Type:        schema.TypeString,
				Description: "Enable/disable displaying the administrator access disclaimer message after an administrator successfully logs in.",
				Computed:    true,
			},
			"pre_login_banner": {
				Type:        schema.TypeString,
				Description: "Enable/disable displaying the administrator access disclaimer message on the login page before an administrator logs in.",
				Computed:    true,
			},
			"private_data_encryption": {
				Type:        schema.TypeString,
				Description: "Enable/disable private data encryption using an AES 128-bit key or passpharse.",
				Computed:    true,
			},
			"proxy_auth_lifetime": {
				Type:        schema.TypeString,
				Description: "Enable/disable authenticated users lifetime control. This is a cap on the total time a proxy user can be authenticated for after which re-authentication will take place.",
				Computed:    true,
			},
			"proxy_auth_lifetime_timeout": {
				Type:        schema.TypeInt,
				Description: "Lifetime timeout in minutes for authenticated users (5  - 65535 min, default=480 (8 hours)).",
				Computed:    true,
			},
			"proxy_auth_timeout": {
				Type:        schema.TypeInt,
				Description: "Authentication timeout in minutes for authenticated users (1 - 300 min, default = 10).",
				Computed:    true,
			},
			"proxy_cert_use_mgmt_vdom": {
				Type:        schema.TypeString,
				Description: "Enable/disable using management VDOM to send requests.",
				Computed:    true,
			},
			"proxy_cipher_hardware_acceleration": {
				Type:        schema.TypeString,
				Description: "Enable/disable using content processor (CP8 or CP9) hardware acceleration to encrypt and decrypt IPsec and SSL traffic.",
				Computed:    true,
			},
			"proxy_kxp_hardware_acceleration": {
				Type:        schema.TypeString,
				Description: "Enable/disable using the content processor to accelerate KXP traffic.",
				Computed:    true,
			},
			"proxy_re_authentication_mode": {
				Type:        schema.TypeString,
				Description: "Control if users must re-authenticate after a session is closed, traffic has been idle, or from the point at which the user was first created.",
				Computed:    true,
			},
			"proxy_resource_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of the maximum memory usage on the FortiGate unit's proxy processing of resources, such as block lists, allow lists, and external resources.",
				Computed:    true,
			},
			"proxy_worker_count": {
				Type:        schema.TypeInt,
				Description: "Proxy worker count.",
				Computed:    true,
			},
			"radius_port": {
				Type:        schema.TypeInt,
				Description: "RADIUS service port number.",
				Computed:    true,
			},
			"reboot_upon_config_restore": {
				Type:        schema.TypeString,
				Description: "Enable/disable reboot of system upon restoring configuration.",
				Computed:    true,
			},
			"refresh": {
				Type:        schema.TypeInt,
				Description: "Statistics refresh interval second(s) in GUI.",
				Computed:    true,
			},
			"remoteauthtimeout": {
				Type:        schema.TypeInt,
				Description: "Number of seconds that the FortiGate waits for responses from remote RADIUS, LDAP, or TACACS+ authentication servers. (1-300 sec, default = 5).",
				Computed:    true,
			},
			"reset_sessionless_tcp": {
				Type:        schema.TypeString,
				Description: "Action to perform if the FortiGate receives a TCP packet but cannot find a corresponding session in its session table. NAT/Route mode only.",
				Computed:    true,
			},
			"restart_time": {
				Type:        schema.TypeString,
				Description: "Daily restart time (hh:mm).",
				Computed:    true,
			},
			"revision_backup_on_logout": {
				Type:        schema.TypeString,
				Description: "Enable/disable back-up of the latest configuration revision when an administrator logs out of the CLI or GUI.",
				Computed:    true,
			},
			"revision_image_auto_backup": {
				Type:        schema.TypeString,
				Description: "Enable/disable back-up of the latest image revision after the firmware is upgraded.",
				Computed:    true,
			},
			"scanunit_count": {
				Type:        schema.TypeInt,
				Description: "Number of scanunits. The range and the default depend on the number of CPUs. Only available on FortiGate units with multiple CPUs.",
				Computed:    true,
			},
			"security_rating_result_submission": {
				Type:        schema.TypeString,
				Description: "Enable/disable the submission of Security Rating results to FortiGuard.",
				Computed:    true,
			},
			"security_rating_run_on_schedule": {
				Type:        schema.TypeString,
				Description: "Enable/disable scheduled runs of Security Rating.",
				Computed:    true,
			},
			"send_pmtu_icmp": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending of path maximum transmission unit (PMTU) - ICMP destination unreachable packet and to support PMTUD protocol on your network to reduce fragmentation of packets.",
				Computed:    true,
			},
			"snat_route_change": {
				Type:        schema.TypeString,
				Description: "Enable/disable the ability to change the static NAT route.",
				Computed:    true,
			},
			"special_file_23_support": {
				Type:        schema.TypeString,
				Description: "Enable/disable detection of those special format files when using Data Leak Protection.",
				Computed:    true,
			},
			"speedtest_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable speed test server.",
				Computed:    true,
			},
			"ssd_trim_date": {
				Type:        schema.TypeInt,
				Description: "Date within a month to run ssd trim.",
				Computed:    true,
			},
			"ssd_trim_freq": {
				Type:        schema.TypeString,
				Description: "How often to run SSD Trim (default = weekly). SSD Trim prevents SSD drive data loss by finding and isolating errors.",
				Computed:    true,
			},
			"ssd_trim_hour": {
				Type:        schema.TypeInt,
				Description: "Hour of the day on which to run SSD Trim (0 - 23, default = 1).",
				Computed:    true,
			},
			"ssd_trim_min": {
				Type:        schema.TypeInt,
				Description: "Minute of the hour on which to run SSD Trim (0 - 59, 60 for random).",
				Computed:    true,
			},
			"ssd_trim_weekday": {
				Type:        schema.TypeString,
				Description: "Day of week to run SSD Trim.",
				Computed:    true,
			},
			"ssh_cbc_cipher": {
				Type:        schema.TypeString,
				Description: "Enable/disable CBC cipher for SSH access.",
				Computed:    true,
			},
			"ssh_enc_algo": {
				Type:        schema.TypeString,
				Description: "Select one or more SSH ciphers.",
				Computed:    true,
			},
			"ssh_hmac_md5": {
				Type:        schema.TypeString,
				Description: "Enable/disable HMAC-MD5 for SSH access.",
				Computed:    true,
			},
			"ssh_kex_algo": {
				Type:        schema.TypeString,
				Description: "Select one or more SSH kex algorithms.",
				Computed:    true,
			},
			"ssh_kex_sha1": {
				Type:        schema.TypeString,
				Description: "Enable/disable SHA1 key exchange for SSH access.",
				Computed:    true,
			},
			"ssh_mac_algo": {
				Type:        schema.TypeString,
				Description: "Select one or more SSH MAC algorithms.",
				Computed:    true,
			},
			"ssh_mac_weak": {
				Type:        schema.TypeString,
				Description: "Enable/disable HMAC-SHA1 and UMAC-64-ETM for SSH access.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default = TLSv1.2).",
				Computed:    true,
			},
			"ssl_static_key_ciphers": {
				Type:        schema.TypeString,
				Description: "Enable/disable static key ciphers in SSL/TLS connections (e.g. AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256).",
				Computed:    true,
			},
			"sslvpn_cipher_hardware_acceleration": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL VPN hardware acceleration.",
				Computed:    true,
			},
			"sslvpn_ems_sn_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable verification of EMS serial number in SSL-VPN connection.",
				Computed:    true,
			},
			"sslvpn_kxp_hardware_acceleration": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL VPN KXP hardware acceleration.",
				Computed:    true,
			},
			"sslvpn_max_worker_count": {
				Type:        schema.TypeInt,
				Description: "Maximum number of SSL-VPN processes. Upper limit for this value is the number of CPUs and depends on the model. Default value of zero means the SSLVPN daemon decides the number of worker processes.",
				Computed:    true,
			},
			"sslvpn_plugin_version_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable checking browser's plugin version by SSL-VPN.",
				Computed:    true,
			},
			"strict_dirty_session_check": {
				Type:        schema.TypeString,
				Description: "Enable to check the session against the original policy when revalidating. This can prevent dropping of redirected sessions when web-filtering and authentication are enabled together. If this option is enabled, the FortiGate unit deletes a session if a routing or policy change causes the session to no longer match the policy that originally allowed the session.",
				Computed:    true,
			},
			"strong_crypto": {
				Type:        schema.TypeString,
				Description: "Enable to use strong encryption and only allow strong ciphers and digest for HTTPS/SSH/TLS/SSL functions.",
				Computed:    true,
			},
			"switch_controller": {
				Type:        schema.TypeString,
				Description: "Enable/disable switch controller feature. Switch controller allows you to manage FortiSwitch from the FortiGate itself.",
				Computed:    true,
			},
			"switch_controller_reserved_network": {
				Type:        schema.TypeString,
				Description: "Configure reserved network subnet for managed switches. This is available when the switch controller is enabled.",
				Computed:    true,
			},
			"sys_perf_log_interval": {
				Type:        schema.TypeInt,
				Description: "Time in minutes between updates of performance statistics logging. (1 - 15 min, default = 5, 0 = disabled).",
				Computed:    true,
			},
			"tcp_halfclose_timer": {
				Type:        schema.TypeInt,
				Description: "Number of seconds the FortiGate unit should wait to close a session after one peer has sent a FIN packet but the other has not responded (1 - 86400 sec (1 day), default = 120).",
				Computed:    true,
			},
			"tcp_halfopen_timer": {
				Type:        schema.TypeInt,
				Description: "Number of seconds the FortiGate unit should wait to close a session after one peer has sent an open session packet but the other has not responded (1 - 86400 sec (1 day), default = 10).",
				Computed:    true,
			},
			"tcp_option": {
				Type:        schema.TypeString,
				Description: "Enable SACK, timestamp and MSS TCP options.",
				Computed:    true,
			},
			"tcp_rst_timer": {
				Type:        schema.TypeInt,
				Description: "Length of the TCP CLOSE state in seconds (5 - 300 sec, default = 5).",
				Computed:    true,
			},
			"tcp_timewait_timer": {
				Type:        schema.TypeInt,
				Description: "Length of the TCP TIME-WAIT state in seconds (1 - 300 sec, default = 1).",
				Computed:    true,
			},
			"tftp": {
				Type:        schema.TypeString,
				Description: "Enable/disable TFTP.",
				Computed:    true,
			},
			"timezone": {
				Type:        schema.TypeString,
				Description: "Number corresponding to your time zone from 00 to 86. Enter set timezone ? to view the list of time zones and the numbers that represent them.",
				Computed:    true,
			},
			"traffic_priority": {
				Type:        schema.TypeString,
				Description: "Choose Type of Service (ToS) or Differentiated Services Code Point (DSCP) for traffic prioritization in traffic shaping.",
				Computed:    true,
			},
			"traffic_priority_level": {
				Type:        schema.TypeString,
				Description: "Default system-wide level of priority for traffic prioritization.",
				Computed:    true,
			},
			"two_factor_email_expiry": {
				Type:        schema.TypeInt,
				Description: "Email-based two-factor authentication session timeout (30 - 300 seconds (5 minutes), default = 60).",
				Computed:    true,
			},
			"two_factor_fac_expiry": {
				Type:        schema.TypeInt,
				Description: "FortiAuthenticator token authentication session timeout (10 - 3600 seconds (1 hour), default = 60).",
				Computed:    true,
			},
			"two_factor_ftk_expiry": {
				Type:        schema.TypeInt,
				Description: "FortiToken authentication session timeout (60 - 600 sec (10 minutes), default = 60).",
				Computed:    true,
			},
			"two_factor_ftm_expiry": {
				Type:        schema.TypeInt,
				Description: "FortiToken Mobile session timeout (1 - 168 hours (7 days), default = 72).",
				Computed:    true,
			},
			"two_factor_sms_expiry": {
				Type:        schema.TypeInt,
				Description: "SMS-based two-factor authentication session timeout (30 - 300 sec, default = 60).",
				Computed:    true,
			},
			"udp_idle_timer": {
				Type:        schema.TypeInt,
				Description: "UDP connection session timeout. This command can be useful in managing CPU and memory resources (1 - 86400 seconds (1 day), default = 60).",
				Computed:    true,
			},
			"url_filter_affinity": {
				Type:        schema.TypeString,
				Description: "URL filter CPU affinity.",
				Computed:    true,
			},
			"url_filter_count": {
				Type:        schema.TypeInt,
				Description: "URL filter daemon count.",
				Computed:    true,
			},
			"user_device_store_max_devices": {
				Type:        schema.TypeInt,
				Description: "Maximum number of devices allowed in user device store.",
				Computed:    true,
			},
			"user_device_store_max_unified_mem": {
				Type:        schema.TypeInt,
				Description: "Maximum unified memory allowed in user device store.",
				Computed:    true,
			},
			"user_device_store_max_users": {
				Type:        schema.TypeInt,
				Description: "Maximum number of users allowed in user device store.",
				Computed:    true,
			},
			"user_server_cert": {
				Type:        schema.TypeString,
				Description: "Certificate to use for https user authentication.",
				Computed:    true,
			},
			"vdom_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable support for split/multiple virtual domains (VDOMs).",
				Computed:    true,
			},
			"vip_arp_range": {
				Type:        schema.TypeString,
				Description: "Controls the number of ARPs that the FortiGate sends for a Virtual IP (VIP) address range.",
				Computed:    true,
			},
			"wad_affinity": {
				Type:        schema.TypeString,
				Description: "Affinity setting for wad (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Computed:    true,
			},
			"wad_csvc_cs_count": {
				Type:        schema.TypeInt,
				Description: "Number of concurrent WAD-cache-service object-cache processes.",
				Computed:    true,
			},
			"wad_csvc_db_count": {
				Type:        schema.TypeInt,
				Description: "Number of concurrent WAD-cache-service byte-cache processes.",
				Computed:    true,
			},
			"wad_memory_change_granularity": {
				Type:        schema.TypeInt,
				Description: "Minimum percentage change in system memory usage detected by the wad daemon prior to adjusting TCP window size for any active connection.",
				Computed:    true,
			},
			"wad_source_affinity": {
				Type:        schema.TypeString,
				Description: "Enable/disable dispatching traffic to WAD workers based on source affinity.",
				Computed:    true,
			},
			"wad_worker_count": {
				Type:        schema.TypeInt,
				Description: "Number of explicit proxy WAN optimization daemon (WAD) processes. By default WAN optimization, explicit proxy, and web caching is handled by all of the CPU cores in a FortiGate unit.",
				Computed:    true,
			},
			"wifi_ca_certificate": {
				Type:        schema.TypeString,
				Description: "CA certificate that verifies the WiFi certificate.",
				Computed:    true,
			},
			"wifi_certificate": {
				Type:        schema.TypeString,
				Description: "Certificate to use for WiFi authentication.",
				Computed:    true,
			},
			"wimax_4g_usb": {
				Type:        schema.TypeString,
				Description: "Enable/disable comparability with WiMAX 4G USB devices.",
				Computed:    true,
			},
			"wireless_controller": {
				Type:        schema.TypeString,
				Description: "Enable/disable the wireless controller feature to use the FortiGate unit to manage FortiAPs.",
				Computed:    true,
			},
			"wireless_controller_port": {
				Type:        schema.TypeInt,
				Description: "Port used for the control channel in wireless controller mode (wireless-mode is ac). The data channel port is the control channel port number plus one (1024 - 49150, default = 5246).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemGlobal"

	o, err := c.Cmdb.ReadSystemGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemGlobal dataSource: %v", err)
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

	diags := refreshObjectSystemGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
