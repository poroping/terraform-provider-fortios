package utils

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestValidateConvIPMask2CIDR(t *testing.T) {
	var tests = []struct {
		a      string
		b      string
		result string
	}{
		{
			"10.0.0.1 255.255.255.255", "10.0.0.1/32", "10.0.0.1/32",
		},
		{
			"10.0.0.1/32", "10.0.0.1/32", "10.0.0.1/32",
		},
		{
			"10.0.0.1 255.255.255.0", "10.0.0.1 255.255.255.0", "10.0.0.1 255.255.255.0",
		},
		{
			"10.0.0.1 255.255.0.0", "10.0.0.1/16", "10.0.0.1/16",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v - %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := ValidateConvIPMask2CIDR(tt.a, tt.b)
			log.Print(ans)
			if ans != tt.result {
				t.Errorf("got %v, want %v", ans, tt.result)
			}
		})
	}
}

// func ValidateConvIPMask2CIDR(oNewIP, oOldIP string) string {
// 	if oNewIP != oOldIP && strings.Contains(oNewIP, "/") && strings.Contains(oOldIP, " ") {
// 		line := strings.Split(oOldIP, " ")
// 		if len(line) >= 2 {
// 			ip := line[0]
// 			mask := line[1]
// 			prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
// 			return ip + "/" + strconv.Itoa(prefixSize)
// 		}
// 	}
// 	return oOldIP
// }

// func escapeFilter(filter string) string {
// 	filter = strings.ReplaceAll(filter, "_", "-")
// 	filter = strings.ReplaceAll(filter, "fosid", "id")
// 	filter = strings.ReplaceAll(filter, "&", "&filter=")
// 	filter = strings.ReplaceAll(filter, ".", "\\.")
// 	filter = strings.ReplaceAll(filter, "\\", "\\\\")
// 	filter = "filter=" + filter
// 	return filter
// }

func TestSortSubtable(t *testing.T) {
	var tests = []struct {
		a      []map[string]interface{}
		b      string
		result []map[string]interface{}
	}{
		{
			a: []map[string]interface{}{
				{
					"name": "zzz",
					"foo":  "bar1",
				},
				{
					"name": "ccc",
					"foo":  "bar",
				},
				{
					"name": "123",
					"foo":  "bar2",
				},
				{
					"name": "DDD",
					"foo":  "bar",
				},
			},
			b: "name",
			result: []map[string]interface{}{
				{
					"name": "123",
					"foo":  "bar2",
				},
				{
					"name": "DDD",
					"foo":  "bar",
				},
				{
					"name": "ccc",
					"foo":  "bar",
				},
				{
					"name": "zzz",
					"foo":  "bar1",
				},
			},
		},
		{
			a: []map[string]interface{}{
				{
					"id":  69,
					"foo": "bar",
				},
				{
					"id":  1,
					"foo": "bar",
				},
				{
					"id":  3,
					"foo": "bar",
				},
				{
					"id":  420,
					"foo": "bar",
				},
			},
			b: "id",
			result: []map[string]interface{}{
				{
					"id":  1,
					"foo": "bar",
				},
				{
					"id":  3,
					"foo": "bar",
				},
				{
					"id":  69,
					"foo": "bar",
				},
				{
					"id":  420,
					"foo": "bar",
				},
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("SortSubtable, %v", tt.b)
		t.Run(testname, func(t *testing.T) {
			SortSubtable(tt.a, tt.b)
			if !reflect.DeepEqual(tt.a, tt.result) {
				t.Errorf("got %v, want %v", tt.a, tt.result)
			}
		})
	}
}

func TestParseMkey(t *testing.T) {
	var tests = []struct {
		a      interface{}
		b      string
		result bool
	}{
		{
			"foo", "foo", true,
		},
		{
			"foo", "bar", false,
		},
		{
			123, "123", true,
		},
		{
			123, "bar", false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v - %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			if !((ParseMkey(tt.a) == tt.b) == tt.result) {
				t.Errorf("got %v, want %v", tt.a, tt.b)
			}
		})
	}
}

func TestCheckVer(t *testing.T) {
	var tests = []struct {
		a      string
		b      string
		c      string
		result bool
	}{
		{
			"", "v5.6.0", "", true,
		},
		{
			"", "", "v7.0.5", true,
		},
		{
			"v5.6.13", "", "", true,
		},
		{
			"v6.0.0", "v5.6.0", "", true,
		},
		{
			"v6.0.0", "", "v6.2.1", true,
		},
		{
			"v6.4.4", "v6.2.1", "v7.0.1", true,
		},
		{
			"v6.4.4", "", "v6.2.1", false,
		},
		{
			"v6.4.4", "v7.0.1", "", false,
		},
		{
			"v6.4.4", "v6.2.8", "v6.4.4", false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("cur: %v added: %v removed: %v", tt.a, tt.b, tt.c)
		t.Run(testname, func(t *testing.T) {
			ans := CheckVer(tt.a, tt.b, tt.c)
			if ans != tt.result {
				t.Errorf("got %v, want %v", ans, tt.result)
			}
		})
	}
}
