---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logdisk_setting"
description: |-
  Get information on a fortios Settings for local disk logging.
---

# Data Source: fortios_logdisk_setting
Use this data source to get information on a fortios Settings for local disk logging.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `diskfull` - Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite).
* `dlp_archive_quota` - DLP archive quota (MB).
* `full_final_warning_threshold` - Log full final warning threshold as a percent (3 - 100, default = 95).
* `full_first_warning_threshold` - Log full first warning threshold as a percent (1 - 98, default = 75).
* `full_second_warning_threshold` - Log full second warning threshold as a percent (2 - 99, default = 90).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `ips_archive` - Enable/disable IPS packet archiving to the local disk.
* `log_quota` - Disk log quota (MB).
* `max_log_file_size` - Maximum log file size before rolling (1 - 100 Mbytes).
* `max_policy_packet_capture_size` - Maximum size of policy sniffer in MB (0 means unlimited).
* `maximum_log_age` - Delete log files older than (days).
* `report_quota` - Report db quota (MB).
* `roll_day` - Day of week on which to roll log file.
* `roll_schedule` - Frequency to check log file for rolling.
* `roll_time` - Time of day to roll the log file (hh:mm).
* `source_ip` - Source IP address to use for uploading disk log files.
* `status` - Enable/disable local disk logging.
* `upload` - Enable/disable uploading log files when they are rolled.
* `upload_delete_files` - Delete log files after uploading (default = enable).
* `upload_destination` - The type of server to upload log files to. Only FTP is currently supported.
* `upload_ssl_conn` - Enable/disable encrypted FTPS communication to upload log files.
* `uploaddir` - The remote directory on the FTP server to upload log files to.
* `uploadip` - IP address of the FTP server to upload log files to.
* `uploadpass` - Password required to log into the FTP server to upload disk log files.
* `uploadport` - TCP port to use for communicating with the FTP server (default = 21).
* `uploadsched` - Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling).
* `uploadtime` - Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
* `uploadtype` - Types of log files to upload. Separate multiple entries with a space.
* `uploaduser` - Username required to log into the FTP server to upload disk log files.
