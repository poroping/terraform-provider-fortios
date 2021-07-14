---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_management_remote"
description: |-
  Manage fortios remote certificates.
---

# fortios_certificate_management_remote
Use this resource to manage a fortios remote certificate

## Example Usage

```hcl
locals {
  cert = <<EOF
-----BEGIN CERTIFICATE-----
MIIDQjCCAiqgAwIBAgIDAYahMA0GCSqGSIb3DQEBCwUAMBoxGDAWBgNVBAMMD2Zh
Yy5leGFtcGxlLmNvbTAeFw0yMTA3MDgwODAyMDdaFw0yNjA3MDcwODAyMDdaMBox
GDAWBgNVBAMMD2ZhYy5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALm6PEQFsJM9p/kuC980P1fq7suSmIafP5r96cTbjlIh+63Y/y2g
6piw/klyfVPyhIOzYZJ4zLmggAJ2JKy6+bernCZ+pAMwtvTeE+ly22Z9LG1bWjrB
ixb9CCJD7ylXnMyovSneUdKpQv1mSMt5nnRvMTD+bXUi0EMWiRgllUNEeAjyON4A
9MBaPBGo/AaShAMyzaYfV3bjcnftXRTl1vTD7srQCydFi5LTIGxun3m8USyRZgOZ
vWWaiAjBJJEqtY/CnrM3G9Mc1AAIx55O5mNVj7p/O5XvMSbikKuBabmAKGtelskY
tjzl/TQ53gfGavEpVwSarswDL65kF/4J93sCAwEAAaOBkDCBjTAMBgNVHRMBAf8E
AjAAMB0GA1UdDgQWBBQs3pdfZ22XvV93u/mknCDcLD+0fjBJBgNVHSMEQjBAgBRG
kgxtUDeVZ6nAMws2qbpfQLZ/YKEepBwwGjEYMBYGA1UEAwwPZmFjLmV4YW1wbGUu
Y29tgghLhCcPSljtZjATBgNVHSUEDDAKBggrBgEFBQcDATANBgkqhkiG9w0BAQsF
AAOCAQEADDxyPJKyXeHteKrWs9BZv15XbjCtzwSaiugJPlBc4jYiRQZvZvnsd/5N
8wcnxtmyDrb7PV6S4pXyYlTucfyv0RGunp1XJnTaOhskOnGVgzFrbLPRgeCZZyJj
b/I40UUWJoC/3/sCEPIkkxII8CX/188Cp5lsjZ/rLJ7DEaozTXt9fKRaCcalV1dd
/jhAPMEitI9i19iiDreJHAXxtgK6MY0zqwwQ7bO0yz5PeW7g/O+73JTPnxCVyOSP
rSmZ2d/sZ5ltmOCt9pywSomXKMP26VZcuFgGbM/yF2aJQ4bI7uaTXg9Zck0PbGbc
pcq522T2hkRqn6mpHYwq4hShjk9L2A==
-----END CERTIFICATE-----
EOF
}

resource "fortios_certificate_management_remote" "example" {
  name        = "Global_Example"
  certificate = local.cert
  scope       = "global"

  lifecycle {
    create_before_destroy = true
  }
}

output "debug" {
  value = fortios_certificate_management_remote.example.certificate_details
}
```

## Argument Reference

* `name` - (Required) Name of certificate. Changing the name forces a new resource to be created.
* `scope` - (Required) Scope of remote certificate (vdom|global). Global scope is only accessible for global administrators. Changing the scope forces a new resource to be created.
* `certificate` - PEM format certificate. Changing the certificate forces a new resource to be created.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:

* `certificate_details` - Certificate details. The structure of `certificate_details` block is documented below.

The `certificate_details` block contains:

* `signature_algorithm` - The algorithm used to sign the certificate.
* `public_key_algorithm` - Tag category.
* `serial_number` - Number that uniquely identifies the certificate with the CA's system. The `format`
    function can be used to convert this base 10 number into other bases, such as hex.
* `is_ca` - `true` if this certificate is a ca certificate.
* `is_valid` - `true` if the certificate is within valid time period.
* `version` - The version the certificate is in.
* `issuer` - Who verified and signed the certificate, roughly following RFC2253.
* `subject` - The entity the certificate belongs to, roughly following RFC2253.
* `not_before` - The time after which the certificate is valid, as an RFC3339 timestamp.
* `not_after` - The time until which the certificate is invalid, as an RFC3339 timestamp.
* `sha1_fingerprint` - The SHA1 fingerprint of the public key of the certificate.