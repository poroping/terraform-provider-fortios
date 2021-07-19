package fortios

import (
	"crypto/x509"
	"encoding/pem"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func diffSuppCertificates(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	}
	if old == "" || new == "" {
		return isCertEqual(old, new)
	}
	return isCertEqual(old, new)
}

func isCertEqual(c1, c2 string) bool {
	der1, _ := pem.Decode([]byte(c1))
	der2, _ := pem.Decode([]byte(c2))

	if der1 == nil || der2 == nil {
		log.Printf("Unable to parse certificate")
		return false
	}

	cert1, err := x509.ParseCertificate(der1.Bytes)
	if err != nil {
		log.Printf("Unable to parse certificate")
		return false
	}

	cert2, err := x509.ParseCertificate(der2.Bytes)
	if err != nil {
		log.Printf("Unable to parse certificate")
		return false
	}

	return cert1.Equal(cert2)
}
