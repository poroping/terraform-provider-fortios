---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_global"
description: |-
  Configure global attributes.
---

## fortios_system_global
Configure global attributes.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `admin_concurrent` - Enable/disable concurrent administrator logins. Use policy-auth-concurrent for firewall authenticated users. Valid values: `enable` `disable` .
* `admin_console_timeout` - Console login timeout that overrides the admin timeout value (15 - 300 seconds, default = 0, which disables the timeout).
* `admin_forticloud_sso_login` - Enable/disable FortiCloud admin login via SSO. Valid values: `enable` `disable` .
* `admin_hsts_max_age` - HTTPS Strict-Transport-Security header max-age in seconds. A value of 0 will reset any HSTS records in the browser.When admin-https-redirect is disabled the header max-age will be 0.
* `admin_https_pki_required` - Enable/disable admin login method. Enable to force administrators to provide a valid certificate to log in if PKI is enabled. Disable to allow administrators to log in with a certificate or password. Valid values: `enable` `disable` .
* `admin_https_redirect` - Enable/disable redirection of HTTP administration access to HTTPS. Valid values: `enable` `disable` .
* `admin_https_ssl_banned_ciphers` - Select one or more cipher technologies that cannot be used in GUI HTTPS negotiations. Only applies to TLS 1.2 and below. Valid values: `RSA` `DHE` `ECDHE` `DSS` `ECDSA` `AES` `AESGCM` `CAMELLIA` `3DES` `SHA1` `SHA256` `SHA384` `STATIC` `CHACHA20` `ARIA` `AESCCM` .
* `admin_https_ssl_ciphersuites` - Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, remove TLS1.3 from admin-https-ssl-versions. Valid values: `TLS-AES-128-GCM-SHA256` `TLS-AES-256-GCM-SHA384` `TLS-CHACHA20-POLY1305-SHA256` `TLS-AES-128-CCM-SHA256` `TLS-AES-128-CCM-8-SHA256` .
* `admin_https_ssl_versions` - Allowed TLS versions for web administration. Valid values: `tlsv1-1` `tlsv1-2` `tlsv1-3` .
* `admin_lockout_duration` - Amount of time in seconds that an administrator account is locked out after reaching the admin-lockout-threshold for repeated failed login attempts.
* `admin_lockout_threshold` - Number of failed login attempts before an administrator account is locked out for the admin-lockout-duration.
* `admin_login_max` - Maximum number of administrators who can be logged in at the same time (1 - 100, default = 100).
* `admin_maintainer` - Enable/disable maintainer administrator login. When enabled, the maintainer account can be used to log in from the console after a hard reboot. The password is "bcpb" followed by the FortiGate unit serial number. You have limited time to complete this login. Valid values: `enable` `disable` .
* `admin_port` - Administrative access port for HTTP. (1 - 65535, default = 80).
* `admin_restrict_local` - Enable/disable local admin authentication restriction when remote authenticator is up and running (default = disable). Valid values: `enable` `disable` .
* `admin_scp` - Enable/disable using SCP to download the system configuration. You can use SCP as an alternative method for backing up the configuration. Valid values: `enable` `disable` .
* `admin_server_cert` - Server certificate that the FortiGate uses for HTTPS administrative connections. This attribute must reference one of the following datasources: `certificate.local.name` .
* `admin_sport` - Administrative access port for HTTPS. (1 - 65535, default = 443).
* `admin_ssh_grace_time` - Maximum time in seconds permitted between making an SSH connection to the FortiGate unit and authenticating (10 - 3600 sec (1 hour), default 120).
* `admin_ssh_password` - Enable/disable password authentication for SSH admin access. Valid values: `enable` `disable` .
* `admin_ssh_port` - Administrative access port for SSH. (1 - 65535, default = 22).
* `admin_ssh_v1` - Enable/disable SSH v1 compatibility. Valid values: `enable` `disable` .
* `admin_telnet` - Enable/disable TELNET service. Valid values: `enable` `disable` .
* `admin_telnet_port` - Administrative access port for TELNET. (1 - 65535, default = 23).
* `admintimeout` - Number of minutes before an idle administrator session times out (1 - 480 minutes (8 hours), default = 5). A shorter idle timeout is more secure.
* `alias` - Alias for your FortiGate unit.
* `allow_traffic_redirect` - Disable to prevent traffic with same local ingress and egress interface from being forwarded without policy check. Valid values: `enable` `disable` .
* `anti_replay` - Level of checking for packet replay and TCP sequence checking. Valid values: `disable` `loose` `strict` .
* `arp_max_entry` - Maximum number of dynamically learned MAC addresses that can be added to the ARP table (131072 - 2147483647, default = 131072).
* `auth_cert` - Server certificate that the FortiGate uses for HTTPS firewall authentication connections. This attribute must reference one of the following datasources: `certificate.local.name` .
* `auth_http_port` - User authentication HTTP port. (1 - 65535, default = 1000).
* `auth_https_port` - User authentication HTTPS port. (1 - 65535, default = 1003).
* `auth_keepalive` - Enable to prevent user authentication sessions from timing out when idle. Valid values: `enable` `disable` .
* `auth_session_limit` - Action to take when the number of allowed user authenticated sessions is reached. Valid values: `block-new` `logout-inactive` .
* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension devices. Valid values: `enable` `disable` .
* `autorun_log_fsck` - Enable/disable automatic log partition check after ungraceful shutdown. Valid values: `enable` `disable` .
* `av_affinity` - Affinity setting for AV scanning (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `av_failopen` - Set the action to take if the FortiGate is running low on memory or the proxy connection limit has been reached. Valid values: `pass` `off` `one-shot` .
* `av_failopen_session` - When enabled and a proxy for a protocol runs out of room in its session table, that protocol goes into failopen mode and enacts the action specified by av-failopen. Valid values: `enable` `disable` .
* `batch_cmdb` - Enable/disable batch mode, allowing you to enter a series of CLI commands that will execute as a group once they are loaded. Valid values: `enable` `disable` .
* `block_session_timer` - Duration in seconds for blocked sessions (1 - 300 sec  (5 minutes), default = 30).
* `br_fdb_max_entry` - Maximum number of bridge forwarding database (FDB) entries.
* `cert_chain_max` - Maximum number of certificates that can be traversed in a certificate chain.
* `cfg_revert_timeout` - Time-out for reverting to the last saved configuration. (10 - 4294967295 seconds, default = 600).
* `cfg_save` - Configuration file save mode for CLI changes. Valid values: `automatic` `manual` `revert` .
* `check_protocol_header` - Level of checking performed on protocol headers. Strict checking is more thorough but may affect performance. Loose checking is OK in most cases. Valid values: `loose` `strict` .
* `check_reset_range` - Configure ICMP error message verification. You can either apply strict RST range checking or disable it. Valid values: `strict` `disable` .
* `cli_audit_log` - Enable/disable CLI audit log. Valid values: `enable` `disable` .
* `cloud_communication` - Enable/disable all cloud communication. Valid values: `enable` `disable` .
* `clt_cert_req` - Enable/disable requiring administrators to have a client certificate to log into the GUI using HTTPS. Valid values: `enable` `disable` .
* `cmdbsvr_affinity` - Affinity setting for cmdbsvr (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `cpu_use_threshold` - Threshold at which CPU usage is reported (% of total CPU, default = 90).
* `csr_ca_attribute` - Enable/disable the CA attribute in certificates. Some CA servers reject CSRs that have the CA attribute. Valid values: `enable` `disable` .
* `daily_restart` - Enable/disable daily restart of FortiGate unit. Use the restart-time option to set the time of day for the restart. Valid values: `enable` `disable` .
* `default_service_source_port` - Default service source port range (default = 1 - 65535).
* `device_identification_active_scan_delay` - Number of seconds to passively scan a device before performing an active scan. (20 - 3600 sec, (20 sec to 1 hour), default = 90).
* `device_idle_timeout` - Time in seconds that a device must be idle to automatically log the device user out. (30 - 31536000 sec (30 sec to 1 year), default = 300).
* `dh_params` - Number of bits to use in the Diffie-Hellman exchange for HTTPS/SSH protocols. Valid values: `1024` `1536` `2048` `3072` `4096` `6144` `8192` .
* `dnsproxy_worker_count` - DNS proxy worker count. For a FortiGate with multiple logical CPUs, you can set the DNS process number from 1 to the number of logical CPUs.
* `dst` - Enable/disable daylight saving time. Valid values: `enable` `disable` .
* `extender_controller_reserved_network` - Configure reserved network subnet for managed LAN extension FortiExtender units. This is available when the FortiExtender daemon is running.
* `failtime` - Fail-time for server lost.
* `faz_disk_buffer_size` - Maximum disk buffer size to temporarily store logs destined for FortiAnalyzer. To be used in the event that FortiAnalyzer is unavailable.
* `fds_statistics` - Enable/disable sending IPS, Application Control, and AntiVirus data to FortiGuard. This data is used to improve FortiGuard services and is not shared with external parties and is protected by Fortinet's privacy policy. Valid values: `enable` `disable` .
* `fds_statistics_period` - FortiGuard statistics collection period in minutes. (1 - 1440 min (1 min to 24 hours), default = 60).
* `fec_port` - Local UDP port for Forward Error Correction (49152 - 65535).
* `fgd_alert_subscription` - Type of alert to retrieve from FortiGuard. Valid values: `advisory` `latest-threat` `latest-virus` `latest-attack` `new-antivirus-db` `new-attack-db` .
* `fortiextender` - Enable/disable FortiExtender. Valid values: `disable` `enable` .
* `fortiextender_data_port` - FortiExtender data port (1024 - 49150, default = 25246).
* `fortiextender_discovery_lockdown` - Enable/disable FortiExtender CAPWAP lockdown. Valid values: `disable` `enable` .
* `fortiextender_vlan_mode` - Enable/disable FortiExtender VLAN mode. Valid values: `enable` `disable` .
* `fortiipam_integration` - Enable/disable integration with the FortiIPAM cloud service. Valid values: `enable` `disable` .
* `fortiservice_port` - FortiService port (1 - 65535, default = 8013). Used by FortiClient endpoint compliance. Older versions of FortiClient used a different port.
* `fortitoken_cloud` - Enable/disable FortiToken Cloud service. Valid values: `enable` `disable` .
* `gui_allow_default_hostname` - Enable/disable the factory default hostname warning on the GUI setup wizard. Valid values: `enable` `disable` .
* `gui_cdn_usage` - Enable/disable Load GUI static files from a CDN. Valid values: `enable` `disable` .
* `gui_certificates` - Enable/disable the System > Certificate GUI page, allowing you to add and configure certificates from the GUI. Valid values: `enable` `disable` .
* `gui_custom_language` - Enable/disable custom languages in GUI. Valid values: `enable` `disable` .
* `gui_date_format` - Default date format used throughout GUI. Valid values: `yyyy/MM/dd` `dd/MM/yyyy` `MM/dd/yyyy` `yyyy-MM-dd` `dd-MM-yyyy` `MM-dd-yyyy` .
* `gui_date_time_source` - Source from which the FortiGate GUI uses to display date and time entries. Valid values: `system` `browser` .
* `gui_device_latitude` - Add the latitude of the location of this FortiGate to position it on the Threat Map.
* `gui_device_longitude` - Add the longitude of the location of this FortiGate to position it on the Threat Map.
* `gui_display_hostname` - Enable/disable displaying the FortiGate's hostname on the GUI login page. Valid values: `enable` `disable` .
* `gui_firmware_upgrade_setup_warning` - Enable/disable the firmware upgrade warning on GUI setup wizard. Valid values: `enable` `disable` .
* `gui_firmware_upgrade_warning` - Enable/disable the firmware upgrade warning on the GUI. Valid values: `enable` `disable` .
* `gui_forticare_registration_setup_warning` - Enable/disable the FortiCare registration setup warning on the GUI. Valid values: `enable` `disable` .
* `gui_fortigate_cloud_sandbox` - Enable/disable displaying FortiGate Cloud Sandbox on the GUI. Valid values: `enable` `disable` .
* `gui_fortisandbox_cloud` - Enable/disable displaying FortiSandbox Cloud on the GUI. Valid values: `enable` `disable` .
* `gui_ipv6` - Enable/disable IPv6 settings on the GUI. Valid values: `enable` `disable` .
* `gui_lines_per_page` - Number of lines to display per page for web administration.
* `gui_local_out` - Enable/disable Local-out traffic on the GUI. Valid values: `enable` `disable` .
* `gui_replacement_message_groups` - Enable/disable replacement message groups on the GUI. Valid values: `enable` `disable` .
* `gui_rest_api_cache` - Enable/disable REST API result caching on FortiGate. Valid values: `enable` `disable` .
* `gui_theme` - Color scheme for the administration GUI. Valid values: `jade` `neutrino` `mariner` `graphite` `melongene` `retro` `dark-matter` `onyx` `eclipse` .
* `gui_wireless_opensecurity` - Enable/disable wireless open security option on the GUI. Valid values: `enable` `disable` .
* `ha_affinity` - Affinity setting for HA daemons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `honor_df` - Enable/disable honoring of Don't-Fragment (DF) flag. Valid values: `enable` `disable` .
* `hostname` - FortiGate unit's hostname. Most models will truncate names longer than 24 characters. Some models support hostnames up to 35 characters.
* `igmp_state_limit` - Maximum number of IGMP memberships (96 - 64000, default = 3200).
* `internet_service_database` - Configure which Internet Service database size to download from FortiGuard and use. Valid values: `mini` `standard` `full` .
* `interval` - Dead gateway detection interval.
* `ip_src_port_range` - IP source port range used for traffic originating from the FortiGate unit.
* `ips_affinity` - Affinity setting for IPS (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx; allowed CPUs must be less than total number of IPS engine daemons).
* `ipsec_asic_offload` - Enable/disable ASIC offloading (hardware acceleration) for IPsec VPN traffic. Hardware acceleration can offload IPsec VPN sessions and accelerate encryption and decryption. Valid values: `enable` `disable` .
* `ipsec_ha_seqjump_rate` - ESP jump ahead rate (1G - 10G pps equivalent).
* `ipsec_hmac_offload` - Enable/disable offloading (hardware acceleration) of HMAC processing for IPsec VPN. Valid values: `enable` `disable` .
* `ipsec_soft_dec_async` - Enable/disable software decryption asynchronization (using multiple CPUs to do decryption) for IPsec VPN traffic. Valid values: `enable` `disable` .
* `ipv6_accept_dad` - Enable/disable acceptance of IPv6 Duplicate Address Detection (DAD).
* `ipv6_allow_anycast_probe` - Enable/disable IPv6 address probe through Anycast. Valid values: `enable` `disable` .
* `ipv6_allow_traffic_redirect` - Disable to prevent IPv6 traffic with same local ingress and egress interface from being forwarded without policy check. Valid values: `enable` `disable` .
* `irq_time_accounting` - Configure CPU IRQ time accounting mode. Valid values: `auto` `force` .
* `language` - GUI display language. Valid values: `english` `french` `spanish` `portuguese` `japanese` `trach` `simch` `korean` .
* `ldapconntimeout` - Global timeout for connections with remote LDAP servers in milliseconds (1 - 300000, default 500).
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception. Valid values: `enable` `disable` .
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission. Valid values: `enable` `disable` .
* `log_ssl_connection` - Enable/disable logging of SSL connection events. Valid values: `enable` `disable` .
* `log_uuid_address` - Enable/disable insertion of address UUIDs to traffic logs. Valid values: `enable` `disable` .
* `log_uuid_policy` - Enable/disable insertion of policy UUIDs to traffic logs. Valid values: `enable` `disable` .
* `login_timestamp` - Enable/disable login time recording. Valid values: `enable` `disable` .
* `long_vdom_name` - Enable/disable long VDOM name support. Valid values: `enable` `disable` .
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `management_port_use_admin_sport` - Enable/disable use of the admin-sport setting for the management port. If disabled, FortiGate will allow user to specify management-port. Valid values: `enable` `disable` .
* `management_vdom` - Management virtual domain name. This attribute must reference one of the following datasources: `system.vdom.name` .
* `max_dlpstat_memory` - Maximum DLP stat memory (0 - 4294967295).
* `max_route_cache_size` - Maximum number of IP route cache entries (0 - 2147483647).
* `memory_use_threshold_extreme` - Threshold at which memory usage is considered extreme (new sessions are dropped) (% of total RAM, default = 95).
* `memory_use_threshold_green` - Threshold at which memory usage forces the FortiGate to exit conserve mode (% of total RAM, default = 82).
* `memory_use_threshold_red` - Threshold at which memory usage forces the FortiGate to enter conserve mode (% of total RAM, default = 88).
* `miglog_affinity` - Affinity setting for logging (64-bit hexadecimal value in the format of xxxxxxxxxxxxxxxx).
* `miglogd_children` - Number of logging (miglogd) processes to be allowed to run. Higher number can reduce performance; lower number can slow log processing time. No logs will be dropped or lost if the number is changed.
* `multi_factor_authentication` - Enforce all login methods to require an additional authentication factor (default = optional). Valid values: `optional` `mandatory` .
* `ndp_max_entry` - Maximum number of NDP table entries (set to 65,536 or higher; if set to 0, kernel holds 65,536 entries).
* `per_user_bal` - Enable/disable per-user block/allow list filter. Valid values: `enable` `disable` .
* `per_user_bwl` - Enable/disable per-user black/white list filter. Valid values: `enable` `disable` .
* `pmtu_discovery` - Enable/disable path MTU discovery. Valid values: `enable` `disable` .
* `policy_auth_concurrent` - Number of concurrent firewall use logins from the same user (1 - 100, default = 0 means no limit).
* `post_login_banner` - Enable/disable displaying the administrator access disclaimer message after an administrator successfully logs in. Valid values: `disable` `enable` .
* `pre_login_banner` - Enable/disable displaying the administrator access disclaimer message on the login page before an administrator logs in. Valid values: `enable` `disable` .
* `private_data_encryption` - Enable/disable private data encryption using an AES 128-bit key or passpharse. Valid values: `disable` `enable` .
* `proxy_auth_lifetime` - Enable/disable authenticated users lifetime control. This is a cap on the total time a proxy user can be authenticated for after which re-authentication will take place. Valid values: `enable` `disable` .
* `proxy_auth_lifetime_timeout` - Lifetime timeout in minutes for authenticated users (5  - 65535 min, default=480 (8 hours)).
* `proxy_auth_timeout` - Authentication timeout in minutes for authenticated users (1 - 300 min, default = 10).
* `proxy_cert_use_mgmt_vdom` - Enable/disable using management VDOM to send requests. Valid values: `enable` `disable` .
* `proxy_cipher_hardware_acceleration` - Enable/disable using content processor (CP8 or CP9) hardware acceleration to encrypt and decrypt IPsec and SSL traffic. Valid values: `disable` `enable` .
* `proxy_kxp_hardware_acceleration` - Enable/disable using the content processor to accelerate KXP traffic. Valid values: `disable` `enable` .
* `proxy_re_authentication_mode` - Control if users must re-authenticate after a session is closed, traffic has been idle, or from the point at which the user was first created. Valid values: `session` `traffic` `absolute` .
* `proxy_resource_mode` - Enable/disable use of the maximum memory usage on the FortiGate unit's proxy processing of resources, such as block lists, allow lists, and external resources. Valid values: `enable` `disable` .
* `proxy_worker_count` - Proxy worker count.
* `radius_port` - RADIUS service port number.
* `reboot_upon_config_restore` - Enable/disable reboot of system upon restoring configuration. Valid values: `enable` `disable` .
* `refresh` - Statistics refresh interval second(s) in GUI.
* `remoteauthtimeout` - Number of seconds that the FortiGate waits for responses from remote RADIUS, LDAP, or TACACS+ authentication servers. (1-300 sec, default = 5).
* `reset_sessionless_tcp` - Action to perform if the FortiGate receives a TCP packet but cannot find a corresponding session in its session table. NAT/Route mode only. Valid values: `enable` `disable` .
* `restart_time` - Daily restart time (hh:mm).
* `revision_backup_on_logout` - Enable/disable back-up of the latest configuration revision when an administrator logs out of the CLI or GUI. Valid values: `enable` `disable` .
* `revision_image_auto_backup` - Enable/disable back-up of the latest image revision after the firmware is upgraded. Valid values: `enable` `disable` .
* `scanunit_count` - Number of scanunits. The range and the default depend on the number of CPUs. Only available on FortiGate units with multiple CPUs.
* `security_rating_result_submission` - Enable/disable the submission of Security Rating results to FortiGuard. Valid values: `enable` `disable` .
* `security_rating_run_on_schedule` - Enable/disable scheduled runs of Security Rating. Valid values: `enable` `disable` .
* `send_pmtu_icmp` - Enable/disable sending of path maximum transmission unit (PMTU) - ICMP destination unreachable packet and to support PMTUD protocol on your network to reduce fragmentation of packets. Valid values: `enable` `disable` .
* `snat_route_change` - Enable/disable the ability to change the static NAT route. Valid values: `enable` `disable` .
* `special_file_23_support` - Enable/disable detection of those special format files when using Data Leak Protection. Valid values: `disable` `enable` .
* `speedtest_server` - Enable/disable speed test server. Valid values: `enable` `disable` .
* `ssd_trim_date` - Date within a month to run ssd trim.
* `ssd_trim_freq` - How often to run SSD Trim (default = weekly). SSD Trim prevents SSD drive data loss by finding and isolating errors. Valid values: `never` `hourly` `daily` `weekly` `monthly` .
* `ssd_trim_hour` - Hour of the day on which to run SSD Trim (0 - 23, default = 1).
* `ssd_trim_min` - Minute of the hour on which to run SSD Trim (0 - 59, 60 for random).
* `ssd_trim_weekday` - Day of week to run SSD Trim. Valid values: `sunday` `monday` `tuesday` `wednesday` `thursday` `friday` `saturday` .
* `ssh_cbc_cipher` - Enable/disable CBC cipher for SSH access. Valid values: `enable` `disable` .
* `ssh_enc_algo` - Select one or more SSH ciphers. Valid values: `chacha20-poly1305@openssh.com` `aes128-ctr` `aes192-ctr` `aes256-ctr` `arcfour256` `arcfour128` `aes128-cbc` `3des-cbc` `blowfish-cbc` `cast128-cbc` `aes192-cbc` `aes256-cbc` `arcfour` `rijndael-cbc@lysator.liu.se` `aes128-gcm@openssh.com` `aes256-gcm@openssh.com` .
* `ssh_hmac_md5` - Enable/disable HMAC-MD5 for SSH access. Valid values: `enable` `disable` .
* `ssh_kex_algo` - Select one or more SSH kex algorithms. Valid values: `diffie-hellman-group1-sha1` `diffie-hellman-group14-sha1` `diffie-hellman-group-exchange-sha1` `diffie-hellman-group-exchange-sha256` `curve25519-sha256@libssh.org` `ecdh-sha2-nistp256` `ecdh-sha2-nistp384` `ecdh-sha2-nistp521` .
* `ssh_kex_sha1` - Enable/disable SHA1 key exchange for SSH access. Valid values: `enable` `disable` .
* `ssh_mac_algo` - Select one or more SSH MAC algorithms. Valid values: `hmac-md5` `hmac-md5-etm@openssh.com` `hmac-md5-96` `hmac-md5-96-etm@openssh.com` `hmac-sha1` `hmac-sha1-etm@openssh.com` `hmac-sha2-256` `hmac-sha2-256-etm@openssh.com` `hmac-sha2-512` `hmac-sha2-512-etm@openssh.com` `hmac-ripemd160` `hmac-ripemd160@openssh.com` `hmac-ripemd160-etm@openssh.com` `umac-64@openssh.com` `umac-128@openssh.com` `umac-64-etm@openssh.com` `umac-128-etm@openssh.com` .
* `ssh_mac_weak` - Enable/disable HMAC-SHA1 and UMAC-64-ETM for SSH access. Valid values: `enable` `disable` .
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default = TLSv1.2). Valid values: `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` `TLSv1-3` .
* `ssl_static_key_ciphers` - Enable/disable static key ciphers in SSL/TLS connections (e.g. AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256). Valid values: `enable` `disable` .
* `sslvpn_cipher_hardware_acceleration` - Enable/disable SSL VPN hardware acceleration. Valid values: `enable` `disable` .
* `sslvpn_ems_sn_check` - Enable/disable verification of EMS serial number in SSL-VPN connection. Valid values: `enable` `disable` .
* `sslvpn_kxp_hardware_acceleration` - Enable/disable SSL VPN KXP hardware acceleration. Valid values: `enable` `disable` .
* `sslvpn_max_worker_count` - Maximum number of SSL-VPN processes. Upper limit for this value is the number of CPUs and depends on the model. Default value of zero means the SSLVPN daemon decides the number of worker processes.
* `sslvpn_plugin_version_check` - Enable/disable checking browser's plugin version by SSL-VPN. Valid values: `enable` `disable` .
* `strict_dirty_session_check` - Enable to check the session against the original policy when revalidating. This can prevent dropping of redirected sessions when web-filtering and authentication are enabled together. If this option is enabled, the FortiGate unit deletes a session if a routing or policy change causes the session to no longer match the policy that originally allowed the session. Valid values: `enable` `disable` .
* `strong_crypto` - Enable to use strong encryption and only allow strong ciphers and digest for HTTPS/SSH/TLS/SSL functions. Valid values: `enable` `disable` .
* `switch_controller` - Enable/disable switch controller feature. Switch controller allows you to manage FortiSwitch from the FortiGate itself. Valid values: `disable` `enable` .
* `switch_controller_reserved_network` - Configure reserved network subnet for managed switches. This is available when the switch controller is enabled.
* `sys_perf_log_interval` - Time in minutes between updates of performance statistics logging. (1 - 15 min, default = 5, 0 = disabled).
* `tcp_halfclose_timer` - Number of seconds the FortiGate unit should wait to close a session after one peer has sent a FIN packet but the other has not responded (1 - 86400 sec (1 day), default = 120).
* `tcp_halfopen_timer` - Number of seconds the FortiGate unit should wait to close a session after one peer has sent an open session packet but the other has not responded (1 - 86400 sec (1 day), default = 10).
* `tcp_option` - Enable SACK, timestamp and MSS TCP options. Valid values: `enable` `disable` .
* `tcp_rst_timer` - Length of the TCP CLOSE state in seconds (5 - 300 sec, default = 5).
* `tcp_timewait_timer` - Length of the TCP TIME-WAIT state in seconds (1 - 300 sec, default = 1).
* `tftp` - Enable/disable TFTP. Valid values: `enable` `disable` .
* `timezone` - Number corresponding to your time zone from 00 to 86. Enter set timezone ? to view the list of time zones and the numbers that represent them. Valid values: `01` `02` `03` `04` `05` `81` `06` `07` `08` `09` `10` `11` `12` `13` `74` `14` `77` `15` `87` `16` `17` `18` `19` `20` `75` `21` `22` `23` `24` `80` `79` `25` `26` `27` `28` `78` `29` `30` `31` `32` `33` `34` `35` `36` `37` `38` `83` `84` `40` `85` `41` `42` `43` `39` `44` `46` `47` `51` `48` `45` `49` `50` `52` `53` `54` `55` `56` `57` `58` `59` `60` `62` `63` `61` `64` `65` `66` `67` `68` `69` `70` `71` `72` `00` `82` `73` `86` `76` .
* `traffic_priority` - Choose Type of Service (ToS) or Differentiated Services Code Point (DSCP) for traffic prioritization in traffic shaping. Valid values: `tos` `dscp` .
* `traffic_priority_level` - Default system-wide level of priority for traffic prioritization. Valid values: `low` `medium` `high` .
* `two_factor_email_expiry` - Email-based two-factor authentication session timeout (30 - 300 seconds (5 minutes), default = 60).
* `two_factor_fac_expiry` - FortiAuthenticator token authentication session timeout (10 - 3600 seconds (1 hour), default = 60).
* `two_factor_ftk_expiry` - FortiToken authentication session timeout (60 - 600 sec (10 minutes), default = 60).
* `two_factor_ftm_expiry` - FortiToken Mobile session timeout (1 - 168 hours (7 days), default = 72).
* `two_factor_sms_expiry` - SMS-based two-factor authentication session timeout (30 - 300 sec, default = 60).
* `udp_idle_timer` - UDP connection session timeout. This command can be useful in managing CPU and memory resources (1 - 86400 seconds (1 day), default = 60).
* `url_filter_affinity` - URL filter CPU affinity.
* `url_filter_count` - URL filter daemon count.
* `user_device_store_max_devices` - Maximum number of devices allowed in user device store.
* `user_device_store_max_unified_mem` - Maximum unified memory allowed in user device store.
* `user_device_store_max_users` - Maximum number of users allowed in user device store.
* `user_server_cert` - Certificate to use for https user authentication. This attribute must reference one of the following datasources: `certificate.local.name` .
* `vdom_mode` - Enable/disable support for split/multiple virtual domains (VDOMs). Valid values: `no-vdom` `split-vdom` `multi-vdom` .
* `vip_arp_range` - Controls the number of ARPs that the FortiGate sends for a Virtual IP (VIP) address range. Valid values: `unlimited` `restricted` .
* `wad_affinity` - Affinity setting for wad (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `wad_csvc_cs_count` - Number of concurrent WAD-cache-service object-cache processes.
* `wad_csvc_db_count` - Number of concurrent WAD-cache-service byte-cache processes.
* `wad_memory_change_granularity` - Minimum percentage change in system memory usage detected by the wad daemon prior to adjusting TCP window size for any active connection.
* `wad_source_affinity` - Enable/disable dispatching traffic to WAD workers based on source affinity. Valid values: `disable` `enable` .
* `wad_worker_count` - Number of explicit proxy WAN optimization daemon (WAD) processes. By default WAN optimization, explicit proxy, and web caching is handled by all of the CPU cores in a FortiGate unit.
* `wifi_ca_certificate` - CA certificate that verifies the WiFi certificate. This attribute must reference one of the following datasources: `certificate.ca.name` .
* `wifi_certificate` - Certificate to use for WiFi authentication. This attribute must reference one of the following datasources: `certificate.local.name` .
* `wimax_4g_usb` - Enable/disable comparability with WiMAX 4G USB devices. Valid values: `enable` `disable` .
* `wireless_controller` - Enable/disable the wireless controller feature to use the FortiGate unit to manage FortiAPs. Valid values: `enable` `disable` .
* `wireless_controller_port` - Port used for the control channel in wireless controller mode (wireless-mode is ac). The data channel port is the control channel port number plus one (1024 - 49150, default = 5246).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_global can be imported using:
```sh
terraform import fortios_system_global.labelname {{mkey}}
```
