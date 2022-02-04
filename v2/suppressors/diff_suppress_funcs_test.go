package suppressors

import (
	"fmt"
	"testing"
)

func TestDiffSuppCertificates(t *testing.T) {
	var tests = []struct {
		c1, c2 string
		equal  bool
	}{
		{
			`-----BEGIN CERTIFICATE-----
MIIDpDCCAoygAwIBAgIGAXJ6CjorMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQG
EwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNj
bzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMM
CmRldi03NTIzMzgxHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjAw
NjAzMTE1NTMwWhcNMzAwNjAzMTE1NjMwWjCBkjELMAkGA1UEBhMCVVMxEzARBgNV
BAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoM
BE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNzUyMzM4
MRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAh2ffRazJFEIN9+DnXnGhEd6veltxrFRttY7bjANZVB2Y
MHd1NxfzVRrvpkbXy+18CGBMsepJRwmx/yyzwfZvxl9ma3l4q9Ir4Xor9eOIrYIb
Yvg/RQHsOg78Pon7ch31uoBJAX/CD7peQdpeGXxW2MIdzudgfXkOThD1DeHAgRcG
LpQFXp3t8YHgQYFgTcQOHtNiOEZkCcc7QvU7zE40SJYHcE/jSz0gGRZ4yVm7gHdt
VmkKDFIjSDfj+GNxAeGHl3iXun4PmWC2nuUpw6GVfpB6HROpWt2EYsKYG53VxfUm
nPygfO6KO9XcPWrP7V5KF9Ji46L7LwFoXksvRKFqYQIDAQABMA0GCSqGSIb3DQEB
CwUAA4IBAQAgbh24hd+cdRF9suTtjRovU0h3tbGY0LQw8FCMfkPc8dAH/J/zPB0S
VuYf74bSIDcTFLsNCIdEOMFoQjN2kvjQDguwj1tca5J7RbzEVYcNa6mf7ZQPOgNa
I1ZCRL8v5XRM0ApL4RYxAr+Pn4D+z8InqY9dljgyLpgQ6ME4/XQH3wdwfkGMUm7L
YJ04iMtokZy9Kd+9B7a9q+R9LRFC5/DICN9h0+ID87BE769flT0UhU+jHA0cS1oi
bH8GdR23jxB3xt4tM5fRIU+O/j7pDZZmXAMOJ/ijLvr9TmOYOxjke1wq8wd/ML5j
Ck+FLktutUqufpzBpZChVAgXUcTR9WzZ
-----END CERTIFICATE-----`,
			`-----BEGIN CERTIFICATE-----
MIIDpDCCAoygAwIBAgIGAXJ6CjorMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEG
A1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEU
MBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMMCmRldi03NTIzMzgxHDAaBgkqhkiG9w0BCQEW
DWluZm9Ab2t0YS5jb20wHhcNMjAwNjAzMTE1NTMwWhcNMzAwNjAzMTE1NjMwWjCBkjELMAkGA1UE
BhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNV
BAoMBE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNzUyMzM4MRwwGgYJ
KoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
h2ffRazJFEIN9+DnXnGhEd6veltxrFRttY7bjANZVB2YMHd1NxfzVRrvpkbXy+18CGBMsepJRwmx
/yyzwfZvxl9ma3l4q9Ir4Xor9eOIrYIbYvg/RQHsOg78Pon7ch31uoBJAX/CD7peQdpeGXxW2MId
zudgfXkOThD1DeHAgRcGLpQFXp3t8YHgQYFgTcQOHtNiOEZkCcc7QvU7zE40SJYHcE/jSz0gGRZ4
yVm7gHdtVmkKDFIjSDfj+GNxAeGHl3iXun4PmWC2nuUpw6GVfpB6HROpWt2EYsKYG53VxfUmnPyg
fO6KO9XcPWrP7V5KF9Ji46L7LwFoXksvRKFqYQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAgbh24
hd+cdRF9suTtjRovU0h3tbGY0LQw8FCMfkPc8dAH/J/zPB0SVuYf74bSIDcTFLsNCIdEOMFoQjN2
kvjQDguwj1tca5J7RbzEVYcNa6mf7ZQPOgNaI1ZCRL8v5XRM0ApL4RYxAr+Pn4D+z8InqY9dljgy
LpgQ6ME4/XQH3wdwfkGMUm7LYJ04iMtokZy9Kd+9B7a9q+R9LRFC5/DICN9h0+ID87BE769flT0U
hU+jHA0cS1oibH8GdR23jxB3xt4tM5fRIU+O/j7pDZZmXAMOJ/ijLvr9TmOYOxjke1wq8wd/ML5j
Ck+FLktutUqufpzBpZChVAgXUcTR9WzZ
-----END CERTIFICATE-----`, true},
		{
			`-----BEGIN CERTIFICATE-----
MIIDpDCCAoygAwIBAgIGAXJ6CjorMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQG
EwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNj
bzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMM
CmRldi03NTIzMzgxHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjAw
NjAzMTE1NTMwWhcNMzAwNjAzMTE1NjMwWjCBkjELMAkGA1UEBhMCVVMxEzARBgNV
BAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoM
BE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNzUyMzM4
MRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAh2ffRazJFEIN9+DnXnGhEd6veltxrFRttY7bjANZVB2Y
MHd1NxfzVRrvpkbXy+18CGBMsepJRwmx/yyzwfZvxl9ma3l4q9Ir4Xor9eOIrYIb
Yvg/RQHsOg78Pon7ch31uoBJAX/CD7peQdpeGXxW2MIdzudgfXkOThD1DeHAgRcG
LpQFXp3t8YHgQYFgTcQOHtNiOEZkCcc7QvU7zE40SJYHcE/jSz0gGRZ4yVm7gHdt
VmkKDFIjSDfj+GNxAeGHl3iXun4PmWC2nuUpw6GVfpB6HROpWt2EYsKYG53VxfUm
nPygfO6KO9XcPWrP7V5KF9Ji46L7LwFoXksvRKFqYQIDAQABMA0GCSqGSIb3DQEB
CwUAA4IBAQAgbh24hd+cdRF9suTtjRovU0h3tbGY0LQw8FCMfkPc8dAH/J/zPB0S
VuYf74bSIDcTFLsNCIdEOMFoQjN2kvjQDguwj1tca5J7RbzEVYcNa6mf7ZQPOgNa
I1ZCRL8v5XRM0ApL4RYxAr+Pn4D+z8InqY9dljgyLpgQ6ME4/XQH3wdwfkGMUm7L
YJ04iMtokZy9Kd+9B7a9q+R9LRFC5/DICN9h0+ID87BE769flT0UhU+jHA0cS1oi
bH8GdR23jxB3xt4tM5fRIU+O/j7pDZZmXAMOJ/ijLvr9TmOYOxjke1wq8wd/ML5j
Ck+FLktutUqufpzBpZChVAgXUcTR9WzZ
-----END CERTIFICATE-----`,
			`-----BEGIN CERTIFICATE-----
MIIDBTCCAe2gAwIBAgIQN33ROaIJ6bJBWDCxtmJEbjANBgkqhkiG9w0BAQsFADAtMSswKQYDVQQD
EyJhY2NvdW50cy5hY2Nlc3Njb250cm9sLndpbmRvd3MubmV0MB4XDTIwMTIyMTIwNTAxN1oXDTI1
MTIyMDIwNTAxN1owLTErMCkGA1UEAxMiYWNjb3VudHMuYWNjZXNzY29udHJvbC53aW5kb3dzLm5l
dDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKGiy0/YZHEo9rRn2bI27u189Sq7NKhI
nFz5hLCSjgUB2rmf5ETNR3RJIDiW1M51LKROsTrjkl45cxK6gcVwLuEgr3L1TgmBtr/Rt/riKyxe
XbLQ9LGBwaNVaJrSscxfdFbJa5J+qzUIFBiFoL7kE8ZtbkZJWBTxHEyEcNC52JJ8ydOhgvZYyket
e8AAVa2TZAbg4ECo9+6nMsaGsSBncRHJlRWVycq8Q4HV4faMEZmZ+iyCZRo2fZufXpn7sJwZ7CEB
uw4qycHvUl6y153sUUFqsswnZGGjqpKSq7I7sVI9vjB199RarHaSSbDgL2FxjmASiUY4RqxnTjVa
2XVHUwUCAwEAAaMhMB8wHQYDVR0OBBYEFI5mN5ftHloEDVNoIa8sQs7kJAeTMA0GCSqGSIb3DQEB
CwUAA4IBAQBnaGnojxNgnV4+TCPZ9br4ox1nRn9tzY8b5pwKTW2McJTe0yEvrHyaItK8KbmeKJOB
vASf+QwHkp+F2BAXzRiTl4Z+gNFQULPzsQWpmKlz6fIWhc7ksgpTkMK6AaTbwWYTfmpKnQw/KJm/
6rboLDWYyKFpQcStu67RZ+aRvQz68Ev2ga5JsXlcOJ3gP/lE5WC1S0rjfabzdMOGP8qZQhXk4wBO
gtFBaisDnbjV5pcIrjRPlhoCxvKgC/290nZ9/DLBH3TbHk8xwHXeBAnAjyAqOZij92uksAv7ZLq4
MODcnQshVINXwsYshG1pQqOLwMertNaY5WtrubMRku44Dw7R
-----END CERTIFICATE-----`, false},
		{
			`-----BEGIN CERTIFICATE-----
MIIDpDCCAoygAwIBAgIGAXJ6CjorMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQG
EwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNj
bzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMM
CmRldi03NTIzMzgxHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjAw
NjAzMTE1NTMwWhcNMzAwNjAzMTE1NjMwWjCBkjELMAkGA1UEBhMCVVMxEzARBgNV
BAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoM
BE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNzUyMzM4
MRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAh2ffRazJFEIN9+DnXnGhEd6veltxrFRttY7bjANZVB2Y
MHd1NxfzVRrvpkbXy+18CGBMsepJRwmx/yyzwfZvxl9ma3l4q9Ir4Xor9eOIrYIb
Yvg/RQHsOg78Pon7ch31uoBJAX/CD7peQdpeGXxW2MIdzudgfXkOThD1DeHAgRcG
LpQFXp3t8YHgQYFgTcQOHtNiOEZkCcc7QvU7zE40SJYHcE/jSz0gGRZ4yVm7gHdt
VmkKDFIjSDfj+GNxAeGHl3iXun4PmWC2nuUpw6GVfpB6HROpWt2EYsKYG53VxfUm
nPygfO6KO9XcPWrP7V5KF9Ji46L7LwFoXksvRKFqYQIDAQABMA0GCSqGSIb3DQEB
CwUAA4IBAQAgbh24hd+cdRF9suTtjRovU0h3tbGY0LQw8FCMfkPc8dAH/J/zPB0S
VuYf74bSIDcTFLsNCIdEOMFoQjN2kvjQDguwj1tca5J7RbzEVYcNa6mf7ZQPOgNa
I1ZCRL8v5XRM0ApL4RYxAr+Pn4D+z8InqY9dljgyLpgQ6ME4/XQH3wdwfkGMUm7L
YJ04iMtokZy9Kd+9B7a9q+R9LRFC5/DICN9h0+ID87BE769flT0UhU+jHA0cS1oi
bH8GdR23jxB3xt4tM5fRIU+O/j7pDZZmXAMOJ/ijLvr9TmOYOxjke1wq8wd/ML5j
Ck+FLktutUqufpzBpZChVAgXUcTR9WzZ
-----END CERTIFICATE-----`, "", false},
		{
			`-----BEGIN CERTIFICATE-----
MIIDpDCCAoygAwIBAgIGAXJ6CjorMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQG
EwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNj
bzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMM
CmRldi03NTIzMzgxHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjAw
NjAzMTE1NTMwWhcNMzAwNjAzMTE1NjMwWjCBkjELMAkGA1UEBhMCVVMxEzARBgNV
BAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoM
BE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNzUyMzM4
MRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAh2ffRazJFEIN9+DnXnGhEd6veltxrFRttY7bjANZVB2Y
MHd1NxfzVRrvpkbXy+18CGBMsepJRwmx/yyzwfZvxl9ma3l4q9Ir4Xor9eOIrYIb
Yvg/RQHsOg78Pon7ch31uoBJAX/CD7peQdpeGXxW2MIdzudgfXkOThD1DeHAgRcG
LpQFXp3t8YHgQYFgTcQOHtNiOEZkCcc7QvU7zE40SJYHcE/jSz0gGRZ4yVm7gHdt
VmkKDFIjSDfj+GNxAeGHl3iXun4PmWC2nuUpw6GVfpB6HROpWt2EYsKYG53VxfUm
nPygfO6KO9XcPWrP7V5KF9Ji46L7LwFoXksvRKFqYQIDAQABMA0GCSqGSIb3DQEB
CwUAA4IBAQAgbh24hd+cdRF9suTtjRovU0h3tbGY0LQw8FCMfkPc8dAH/J/zPB0S
VuYf74bSIDcTFLsNCIdEOMFoQjN2kvjQDguwj1tca5J7RbzEVYcNa6mf7ZQPOgNa
I1ZCRL8v5XRM0ApL4RYxAr+Pn4D+z8InqY9dljgyLpgQ6ME4/XQH3wdwfkGMUm7L
YJ04iMtokZy9Kd+9B7a9q+R9LRFC5/DICN9h0+ID87BE769flT0UhU+jHA0cS1oi
bH8GdR23jxB3xt4tM5fRIU+O/j7pDZZmXAMOJ/ijLvr9TmOYOxjke1wq8wd/ML5j
Ck+FLktutUqufpzBpZChVAgXUcTR9WzZ
-----END CERTIFICATE-----`,
			`-----BEGIN CERTIFICATE-----
INVALID
-----END CERTIFICATE-----`, false},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Testing cert equality, %v", i)
		t.Run(testname, func(t *testing.T) {
			ans := DiffSuppCertificates("", tt.c1, tt.c2, nil)
			if ans != tt.equal {
				t.Errorf("got %v, want %v", ans, tt.equal)
			}
		})
	}
}

func TestIsIPEqual(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		equal  bool
	}{
		{"2a02:03c8::0", "2a02:3c8::0", true},
		{"notip", "2a02:3c8::0", false},
		{"", "2a02:3c8::0", false},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Testing IP equality, %v", i)
		t.Run(testname, func(t *testing.T) {
			ans := isIPEqual(tt.s1, tt.s2)
			if ans != tt.equal {
				t.Errorf("got %v, want %v", ans, tt.equal)
			}
		})
	}
}

func TestIsCidrEqual(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		equal  bool
	}{
		{"2a02:03c8::0/128", "2a02:3c8::0/128", true},
		{"notip", "2a02:3c8::0/128", false},
		{"", "2a02:3c8::0/128", false},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Testing CIDR equality, %v", i)
		t.Run(testname, func(t *testing.T) {
			ans := isCidrEqual(tt.s1, tt.s2)
			if ans != tt.equal {
				t.Errorf("got %v, want %v", ans, tt.equal)
			}
		})
	}
}

func TestIsFakeListEqual(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		equal  bool
	}{
		{"21 14 5", "21 5 14", true},
		{"21 14 5", "5 14", false},
		{"21 14 5", "20 5 14", false},
		{"a b c", "a c b", true},
		{"a b c", "d e f", false},
		{"a b c", "", false},
		{"certificate_01 certificate_02 certificate_03", "certificate_03 certificate_01 certificate_02", true}}

	for i, tt := range tests {
		testname := fmt.Sprintf("Testing FortiOS shitty 'list' equality, %v", i)
		t.Run(testname, func(t *testing.T) {
			ans := isFakeListEqual(tt.s1, tt.s2)
			if ans != tt.equal {
				t.Errorf("got %v, want %v", ans, tt.equal)
			}
		})
	}
}

func TestIsMultiStringEqual(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		equal  bool
	}{
		{"\"169.254.255.1\"", "169.254.255.1", true},
		{`"169.254.255.1"`, "169.254.255.1", true},
		{"", "", true},
		{"\"169.254.255.2\"", "169.254.255.1", false},
		{"\"169.254.255.1\"", "\"169.254.255.1\"", true},
		{"\"169.254.255.1\" \"169.254.255.2\"", "\"169.254.255.1\" \"169.254.255.2\"", true},
		{"\"169.254.255.1\" \"169.254.255.2\"", "\"169.254.255.1\" \"169.254.255.3\"", false},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Testing FortiOS 'multi-string' equality, %v", i)
		t.Run(testname, func(t *testing.T) {
			ans := isMultiStringEqual(tt.s1, tt.s2)
			if ans != tt.equal {
				t.Errorf("got %v, want %v", ans, tt.equal)
			}
		})
	}
}

func TestToCIDR(t *testing.T) {
	var tests = []struct {
		s1, s2 string
	}{
		{"0.0.0.0 0.0.0.0", "0.0.0.0/0"},
		{"172.16.0.0 255.240.0.0", "172.16.0.0/12"},
		{"192.168.0.0/16", "192.168.0.0/16"},
		{"asdsadasdasd", "asdsadasdasd"},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Testing FortiOS 'multi-string' equality, %v", i)
		t.Run(testname, func(t *testing.T) {
			ans := toCidr(tt.s1)
			if ans != tt.s2 {
				t.Errorf("got %v, want %v", ans, tt.s2)
			}
		})
	}
}
