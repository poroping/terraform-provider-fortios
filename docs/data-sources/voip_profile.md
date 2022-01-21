---
subcategory: "FortiGate Voip"
layout: "fortios"
page_title: "FortiOS: fortios_voip_profile"
description: |-
  Get information on a fortios Configure VoIP profiles.
---

# Data Source: fortios_voip_profile
Use this data source to get information on a fortios Configure VoIP profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `feature_set` - Flow or proxy inspection feature set.
* `name` - Profile name.
* `msrp` - MSRP.The structure of `msrp` block is documented below.

The `msrp` block contains:

* `log_violations` - Enable/disable logging of MSRP violations.
* `max_msg_size` - Maximum allowable MSRP message size (1-65535).
* `max_msg_size_action` - Action for violation of max-msg-size.
* `status` - Enable/disable MSRP.
* `sccp` - SCCP.The structure of `sccp` block is documented below.

The `sccp` block contains:

* `block_mcast` - Enable/disable block multicast RTP connections.
* `log_call_summary` - Enable/disable log summary of SCCP calls.
* `log_violations` - Enable/disable logging of SCCP violations.
* `max_calls` - Maximum calls per minute per SCCP client (max 65535).
* `status` - Enable/disable SCCP.
* `verify_header` - Enable/disable verify SCCP header content.
* `sip` - SIP.The structure of `sip` block is documented below.

The `sip` block contains:

* `ack_rate` - ACK request rate limit (per second, per policy).
* `ack_rate_track` - Track the packet protocol field.
* `block_ack` - Enable/disable block ACK requests.
* `block_bye` - Enable/disable block BYE requests.
* `block_cancel` - Enable/disable block CANCEL requests.
* `block_geo_red_options` - Enable/disable block OPTIONS requests, but OPTIONS requests still notify for redundancy.
* `block_info` - Enable/disable block INFO requests.
* `block_invite` - Enable/disable block INVITE requests.
* `block_long_lines` - Enable/disable block requests with headers exceeding max-line-length.
* `block_message` - Enable/disable block MESSAGE requests.
* `block_notify` - Enable/disable block NOTIFY requests.
* `block_options` - Enable/disable block OPTIONS requests and no OPTIONS as notifying message for redundancy either.
* `block_prack` - Enable/disable block prack requests.
* `block_publish` - Enable/disable block PUBLISH requests.
* `block_refer` - Enable/disable block REFER requests.
* `block_register` - Enable/disable block REGISTER requests.
* `block_subscribe` - Enable/disable block SUBSCRIBE requests.
* `block_unknown` - Block unrecognized SIP requests (enabled by default).
* `block_update` - Enable/disable block UPDATE requests.
* `bye_rate` - BYE request rate limit (per second, per policy).
* `bye_rate_track` - Track the packet protocol field.
* `call_keepalive` - Continue tracking calls with no RTP for this many minutes.
* `cancel_rate` - CANCEL request rate limit (per second, per policy).
* `cancel_rate_track` - Track the packet protocol field.
* `contact_fixup` - Fixup contact anyway even if contact's IP:port doesn't match session's IP:port.
* `hnt_restrict_source_ip` - Enable/disable restrict RTP source IP to be the same as SIP source IP when HNT is enabled.
* `hosted_nat_traversal` - Hosted NAT Traversal (HNT).
* `info_rate` - INFO request rate limit (per second, per policy).
* `info_rate_track` - Track the packet protocol field.
* `invite_rate` - INVITE request rate limit (per second, per policy).
* `invite_rate_track` - Track the packet protocol field.
* `ips_rtp` - Enable/disable allow IPS on RTP.
* `log_call_summary` - Enable/disable logging of SIP call summary.
* `log_violations` - Enable/disable logging of SIP violations.
* `malformed_header_allow` - Action for malformed Allow header.
* `malformed_header_call_id` - Action for malformed Call-ID header.
* `malformed_header_contact` - Action for malformed Contact header.
* `malformed_header_content_length` - Action for malformed Content-Length header.
* `malformed_header_content_type` - Action for malformed Content-Type header.
* `malformed_header_cseq` - Action for malformed CSeq header.
* `malformed_header_expires` - Action for malformed Expires header.
* `malformed_header_from` - Action for malformed From header.
* `malformed_header_max_forwards` - Action for malformed Max-Forwards header.
* `malformed_header_no_proxy_require` - Action for malformed SIP messages without Proxy-Require header.
* `malformed_header_no_require` - Action for malformed SIP messages without Require header.
* `malformed_header_p_asserted_identity` - Action for malformed P-Asserted-Identity header.
* `malformed_header_rack` - Action for malformed RAck header.
* `malformed_header_record_route` - Action for malformed Record-Route header.
* `malformed_header_route` - Action for malformed Route header.
* `malformed_header_rseq` - Action for malformed RSeq header.
* `malformed_header_sdp_a` - Action for malformed SDP a line.
* `malformed_header_sdp_b` - Action for malformed SDP b line.
* `malformed_header_sdp_c` - Action for malformed SDP c line.
* `malformed_header_sdp_i` - Action for malformed SDP i line.
* `malformed_header_sdp_k` - Action for malformed SDP k line.
* `malformed_header_sdp_m` - Action for malformed SDP m line.
* `malformed_header_sdp_o` - Action for malformed SDP o line.
* `malformed_header_sdp_r` - Action for malformed SDP r line.
* `malformed_header_sdp_s` - Action for malformed SDP s line.
* `malformed_header_sdp_t` - Action for malformed SDP t line.
* `malformed_header_sdp_v` - Action for malformed SDP v line.
* `malformed_header_sdp_z` - Action for malformed SDP z line.
* `malformed_header_to` - Action for malformed To header.
* `malformed_header_via` - Action for malformed VIA header.
* `malformed_request_line` - Action for malformed request line.
* `max_body_length` - Maximum SIP message body length (0 meaning no limit).
* `max_dialogs` - Maximum number of concurrent calls/dialogs (per policy).
* `max_idle_dialogs` - Maximum number established but idle dialogs to retain (per policy).
* `max_line_length` - Maximum SIP header line length (78-4096).
* `message_rate` - MESSAGE request rate limit (per second, per policy).
* `message_rate_track` - Track the packet protocol field.
* `nat_port_range` - RTP NAT port range.
* `nat_trace` - Enable/disable preservation of original IP in SDP i line.
* `no_sdp_fixup` - Enable/disable no SDP fix-up.
* `notify_rate` - NOTIFY request rate limit (per second, per policy).
* `notify_rate_track` - Track the packet protocol field.
* `open_contact_pinhole` - Enable/disable open pinhole for non-REGISTER Contact port.
* `open_record_route_pinhole` - Enable/disable open pinhole for Record-Route port.
* `open_register_pinhole` - Enable/disable open pinhole for REGISTER Contact port.
* `open_via_pinhole` - Enable/disable open pinhole for Via port.
* `options_rate` - OPTIONS request rate limit (per second, per policy).
* `options_rate_track` - Track the packet protocol field.
* `prack_rate` - PRACK request rate limit (per second, per policy).
* `prack_rate_track` - Track the packet protocol field.
* `preserve_override` - Override i line to preserve original IPS (default: append).
* `provisional_invite_expiry_time` - Expiry time (10-3600, in seconds) for provisional INVITE.
* `publish_rate` - PUBLISH request rate limit (per second, per policy).
* `publish_rate_track` - Track the packet protocol field.
* `refer_rate` - REFER request rate limit (per second, per policy).
* `refer_rate_track` - Track the packet protocol field.
* `register_contact_trace` - Enable/disable trace original IP/port within the contact header of REGISTER requests.
* `register_rate` - REGISTER request rate limit (per second, per policy).
* `register_rate_track` - Track the packet protocol field.
* `rfc2543_branch` - Enable/disable support via branch compliant with RFC 2543.
* `rtp` - Enable/disable create pinholes for RTP traffic to traverse firewall.
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation.
* `ssl_auth_client` - Require a client certificate and authenticate it with the peer/peergrp.
* `ssl_auth_server` - Authenticate the server's certificate with the peer/peergrp.
* `ssl_client_certificate` - Name of Certificate to offer to server if requested.
* `ssl_client_renegotiation` - Allow/block client renegotiation by server.
* `ssl_max_version` - Highest SSL/TLS version to negotiate.
* `ssl_min_version` - Lowest SSL/TLS version to negotiate.
* `ssl_mode` - SSL/TLS mode for encryption & decryption of traffic.
* `ssl_pfs` - SSL Perfect Forward Secrecy.
* `ssl_send_empty_frags` - Send empty fragments to avoid attack on CBC IV (SSL 3.0 & TLS 1.0 only).
* `ssl_server_certificate` - Name of Certificate return to the client in every SSL connection.
* `status` - Enable/disable SIP.
* `strict_register` - Enable/disable only allow the registrar to connect.
* `subscribe_rate` - SUBSCRIBE request rate limit (per second, per policy).
* `subscribe_rate_track` - Track the packet protocol field.
* `unknown_header` - Action for unknown SIP header.
* `update_rate` - UPDATE request rate limit (per second, per policy).
* `update_rate_track` - Track the packet protocol field.
