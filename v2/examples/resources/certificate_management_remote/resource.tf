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
  name        = "Global_Remote_Example"
  certificate = local.cert
  scope       = "global"

  lifecycle {
    create_before_destroy = true
  }
}

output "debug" {
  value = fortios_certificate_management_remote.example.certificate_details
}