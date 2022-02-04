---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profileprotocoloptions"
description: |-
  Get information on a fortios Configure protocol options.
---

# Data Source: fortios_firewall_profileprotocoloptions
Use this data source to get information on a fortios Configure protocol options.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Optional comments.
* `feature_set` - Flow/proxy feature set.
* `name` - Name.
* `oversize_log` - Enable/disable logging for antivirus oversize file blocking.
* `replacemsg_group` - Name of the replacement message group to be used.
* `rpc_over_http` - Enable/disable inspection of RPC over HTTP.
* `switching_protocols_log` - Enable/disable logging for HTTP/HTTPS switching protocols.
* `cifs` - Configure CIFS protocol options.The structure of `cifs` block is documented below.

The `cifs` block contains:

* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 445).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `server_credential_type` - CIFS server credential type.
* `status` - Enable/disable the active status of scanning for this protocol.
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `server_keytab` - Server keytab.The structure of `server_keytab` block is documented below.

The `server_keytab` block contains:

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `principal` - Service principal. For example, host/cifsserver.example.com@example.com.
* `dns` - Configure DNS protocol options.The structure of `dns` block is documented below.

The `dns` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 53).
* `status` - Enable/disable the active status of scanning for this protocol.
* `ftp` - Configure FTP protocol options.The structure of `ftp` block is documented below.

The `ftp` block contains:

* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 21).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.
* `status` - Enable/disable the active status of scanning for this protocol.
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `http` - Configure HTTP protocol options.The structure of `http` block is documented below.

The `http` block contains:

* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `fortinet_bar` - Enable/disable Fortinet bar on HTML content.
* `fortinet_bar_port` - Port for use by Fortinet Bar (1 - 65535, default = 8011).
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 80).
* `post_lang` - ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `range_block` - Enable/disable blocking of partial downloads.
* `retry_count` - Number of attempts to retry HTTP connection (0 - 100, default = 0).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.
* `status` - Enable/disable the active status of scanning for this protocol.
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `streaming_content_bypass` - Enable/disable bypassing of streaming content from buffering.
* `strip_x_forwarded_for` - Enable/disable stripping of HTTP X-Forwarded-For header.
* `switching_protocols` - Bypass from scanning, or block a connection that attempts to switch protocol.
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol.
* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.
* `imap` - Configure IMAP protocol options.The structure of `imap` block is documented below.

The `imap` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 143).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.
* `status` - Enable/disable the active status of scanning for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `mail_signature` - Configure Mail signature.The structure of `mail_signature` block is documented below.

The `mail_signature` block contains:

* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).
* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.
* `mapi` - Configure MAPI protocol options.The structure of `mapi` block is documented below.

The `mapi` block contains:

* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `status` - Enable/disable the active status of scanning for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `nntp` - Configure NNTP protocol options.The structure of `nntp` block is documented below.

The `nntp` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `status` - Enable/disable the active status of scanning for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `pop3` - Configure POP3 protocol options.The structure of `pop3` block is documented below.

The `pop3` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 110).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.
* `status` - Enable/disable the active status of scanning for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `smtp` - Configure SMTP protocol options.The structure of `smtp` block is documented below.

The `smtp` block contains:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 25).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `server_busy` - Enable/disable SMTP server busy when server not available.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.
* `status` - Enable/disable the active status of scanning for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
* `ssh` - Configure SFTP and SCP protocol options.The structure of `ssh` block is documented below.

The `ssh` block contains:

* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol.
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).
