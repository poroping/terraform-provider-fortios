package suppressors

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DiffSuppCertificates(k, old, new string, d *schema.ResourceData) bool {
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

func DiffIPEqual(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	}
	return isIPEqual(old, new)
}

func isIPEqual(s1, s2 string) bool {
	ip1 := net.ParseIP(s1)
	ip2 := net.ParseIP(s2)

	if ip1 == nil || ip2 == nil {
		log.Printf("Error parsing IP %s %s", s1, s2)
		return false
	}

	return ip1.Equal(ip2)
}

func DiffCidrEqual(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	}
	old = toCidr(old)
	new = toCidr(new)
	return isCidrEqual(old, new)
}

func isCidrEqual(s1, s2 string) bool {
	ip1, _, err := net.ParseCIDR(s1)
	if err != nil {
		log.Printf("Error parsing CIDR %s", s1)
		return false
	}
	ip2, _, err := net.ParseCIDR(s2)
	if err != nil {
		log.Printf("Error parsing CIDR %s", s2)
		return false
	}

	return ip1.Equal(ip2)
}

func DiffFakeListEqual(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	}
	return isFakeListEqual(old, new)
}

func isFakeListEqual(s1, s2 string) bool {
	l1 := strings.Split(s1, " ")
	l2 := strings.Split(s2, " ")
	sort.Strings(l1)
	sort.Strings(l2)

	return reflect.DeepEqual(l1, l2)
}

func DiffMultiStringEqual(k, old, new string, d *schema.ResourceData) bool {
	return isMultiStringEqual(old, new)
}

func isMultiStringEqual(s1, s2 string) bool {
	s1 = strings.Trim(s1, "\"")
	s2 = strings.Trim(s2, "\"")
	return s1 == s2
}

func toCidr(in string) (cidr string) {
	_, _, err := net.ParseCIDR(in)
	if err == nil {
		return in
	}
	s := strings.Split(in, " ")
	if len(s) == 2 {
		ip := s[0]
		mask := s[1]
		prefix_length, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
		if mask != "0.0.0.0" && prefix_length == 0 {
			return in
		}
		cidr = ip + "/" + strconv.Itoa(prefix_length)
		return
	}
	return in
}

// func isSubtableEqual(v1, v2 []map[string]interface{}) bool {
// 	l1 := strings.Split(s1, " ")
// 	l2 := strings.Split(s2, " ")
// 	sort.Strings(l1)
// 	sort.Strings(l2)

// 	return reflect.DeepEqual(l1, l2)
// }
