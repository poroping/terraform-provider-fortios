---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profileprotocoloptions"
description: |-
  Configure protocol options.
---

## fortios_firewall_profileprotocoloptions
Configure protocol options.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comment` - Optional comments.
* `feature_set` - Flow/proxy feature set. Valid values: `flow` `proxy` .
* `name` - Name.
* `oversize_log` - Enable/disable logging for antivirus oversize file blocking. Valid values: `disable` `enable` .
* `replacemsg_group` - Name of the replacement message group to be used. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `rpc_over_http` - Enable/disable inspection of RPC over HTTP. Valid values: `enable` `disable` .
* `switching_protocols_log` - Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable` `enable` .
* `cifs` - Configure CIFS protocol options. The structure of `cifs` block is documented below.

The `cifs` block contains:

* `domain_controller` - Domain for which to decrypt CIFS traffic. This attribute must reference one of the following datasources: `user.domain-controller.name` `credential-store.domain-controller.server-name` .
* `options` - One or more options that can be applied to the session. Valid values: `oversize` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 445).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `server_credential_type` - CIFS server credential type. Valid values: `none` `credential-replication` `credential-keytab` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `auto-tuning` `system` `static` `dynamic` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `server_keytab` - Server keytab. The structure of `server_keytab` block is documented below.

The `server_keytab` block contains:

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `principal` - Service principal. For example, host/cifsserver.example.com@example.com.
* `dns` - Configure DNS protocol options. The structure of `dns` block is documented below.

The `dns` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 53).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `ftp` - Configure FTP protocol options. The structure of `ftp` block is documented below.

The `ftp` block contains:

* `comfort_amount` - Number of bytes to send in each transmission for client comforting (bytes).
* `comfort_interval` - Interval between successive transmissions of data for client comforting (seconds).
* `explicit_ftp_tls` - Enable/disable FTP redirection for explicit FTPS. Valid values: `enable` `disable` .
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable` `disable` .
* `options` - One or more options that can be applied to the session. Valid values: `clientcomfort` `oversize` `splice` `bypass-rest-command` `bypass-mode-command` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 21).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no` `yes` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `auto-tuning` `system` `static` `dynamic` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `http` - Configure HTTP protocol options. The structure of `http` block is documented below.

The `http` block contains:

* `address_ip_rating` - Enable/disable IP based URL rating. Valid values: `enable` `disable` .
* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `comfort_amount` - Number of bytes to send in each transmission for client comforting (bytes).
* `comfort_interval` - Interval between successive transmissions of data for client comforting (seconds).
* `fortinet_bar` - Enable/disable Fortinet bar on HTML content. Valid values: `enable` `disable` .
* `fortinet_bar_port` - Port for use by Fortinet Bar (1 - 65535, default = 8011).
* `h2c` - Enable/disable h2c HTTP connection upgrade. Valid values: `enable` `disable` .
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable` `disable` .
* `options` - One or more options that can be applied to the session. Valid values: `clientcomfort` `servercomfort` `oversize` `chunkedbypass` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 80).
* `post_lang` - ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets). Valid values: `jisx0201` `jisx0208` `jisx0212` `gb2312` `ksc5601-ex` `euc-jp` `sjis` `iso2022-jp` `iso2022-jp-1` `iso2022-jp-2` `euc-cn` `ces-gbk` `hz` `ces-big5` `euc-kr` `iso2022-jp-3` `iso8859-1` `tis620` `cp874` `cp1252` `cp1251` .
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `range_block` - Enable/disable blocking of partial downloads. Valid values: `disable` `enable` .
* `retry_count` - Number of attempts to retry HTTP connection (0 - 100, default = 0).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no` `yes` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `streaming_content_bypass` - Enable/disable bypassing of streaming content from buffering. Valid values: `enable` `disable` .
* `strip_x_forwarded_for` - Enable/disable stripping of HTTP X-Forwarded-For header. Valid values: `disable` `enable` .
* `switching_protocols` - Bypass from scanning, or block a connection that attempts to switch protocol. Valid values: `bypass` `block` .
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `auto-tuning` `system` `static` `dynamic` .
* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `enable` `disable` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `unknown_content_encoding` - Configure the action the FortiGate unit will take on unknown content-encoding. Valid values: `block` `inspect` `bypass` .
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `reject` `tunnel` `best-effort` .
* `verify_dns_for_policy_matching` - Enable/disable verification of DNS for policy matching. Valid values: `enable` `disable` .
* `imap` - Configure IMAP protocol options. The structure of `imap` block is documented below.

The `imap` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable` `disable` .
* `options` - One or more options that can be applied to the session. Valid values: `fragmail` `oversize` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 143).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no` `yes` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `mail_signature` - Configure Mail signature. The structure of `mail_signature` block is documented below.

The `mail_signature` block contains:

* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).
* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate. Valid values: `disable` `enable` .
* `mapi` - Configure MAPI protocol options. The structure of `mapi` block is documented below.

The `mapi` block contains:

* `options` - One or more options that can be applied to the session. Valid values: `fragmail` `oversize` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `nntp` - Configure NNTP protocol options. The structure of `nntp` block is documented below.

The `nntp` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable` `disable` .
* `options` - One or more options that can be applied to the session. Valid values: `oversize` `splice` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `pop3` - Configure POP3 protocol options. The structure of `pop3` block is documented below.

The `pop3` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable` `disable` .
* `options` - One or more options that can be applied to the session. Valid values: `fragmail` `oversize` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 110).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no` `yes` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `smtp` - Configure SMTP protocol options. The structure of `smtp` block is documented below.

The `smtp` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable` `disable` .
* `options` - One or more options that can be applied to the session. Valid values: `fragmail` `oversize` `splice` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `ports` - Ports to scan for content (1 - 65535, default = 25).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `server_busy` - Enable/disable SMTP server busy when server not available. Valid values: `enable` `disable` .
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no` `yes` .
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable` `disable` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.
* `ssh` - Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.

The `ssh` block contains:

* `comfort_amount` - Number of bytes to send in each transmission for client comforting (bytes).
* `comfort_interval` - Interval between successive transmissions of data for client comforting (seconds).
* `options` - One or more options that can be applied to the session. Valid values: `oversize` `clientcomfort` `servercomfort` .
* `oversize_limit` - Maximum in-memory file size that can be scanned (MB).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable` `disable` .
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no` `yes` .
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `auto-tuning` `system` `static` `dynamic` .
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_profileprotocoloptions can be imported using:
```sh
terraform import fortios_firewall_profileprotocoloptions.labelname {{mkey}}
```
