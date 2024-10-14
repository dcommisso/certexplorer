package certformatter

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToColonNotation(t *testing.T) {
	cases := map[string]struct {
		input    []byte
		expected string
	}{
		"byte": {
			input:    []byte{93, 147, 141, 48, 103, 54, 200, 6, 29, 26, 199, 84, 132, 105, 7},
			expected: "5d:93:8d:30:67:36:c8:06:1d:1a:c7:54:84:69:07",
		},
		"big int": {
			input:    big.NewInt(758990378568).Bytes(),
			expected: "b0:b7:5a:16:48",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := ToColonNotation(tc.input)
			if got != tc.expected {
				t.Errorf("expected: %v - got: %v", tc.expected, got)
			}
		})
	}
}

func TestGetFormattedCertificate(t *testing.T) {
	cs := NewCertstore()
	cs.Load(getSampleCert("multiple certificates with comments"), "test")
	certformatter := cs.NewFormatter()

	cases := map[string]struct {
		inputCertIndex      int
		inputSelectedFields []Outputfield
		expectedOutput      FormattedCertificate
		expectedError       string
	}{
		"two fields selected": {
			inputCertIndex:      0,
			inputSelectedFields: []Outputfield{OutputFieldSourceFile, OutputFieldSubject},
			expectedOutput: FormattedCertificate{
				OutputFieldSubject:          "Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldSourceFile:       "From file: test",
				OutputFieldCertificateIndex: "0",
			},
		},
		"all fields selected - some empty": {
			inputCertIndex: 0,
			expectedOutput: FormattedCertificate{
				OutputFieldSubject: "Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldIssuer:  "Issuer: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldSerialNumber: `Serial Number:
    5d:93:8d:30:67:36:c8:06:1d:1a:c7:54:84:69:07`,
				OutputFieldValidity: `Validity
    Not Before: 2008-10-29 15:59:56 +0000 UTC
    Not After : 2030-01-01 00:00:00 +0000 UTC`,
				OutputFieldNotBefore: "Not Before: 2008-10-29 15:59:56 +0000 UTC",
				OutputFieldNotAfter:  "Not After : 2030-01-01 00:00:00 +0000 UTC",
				OutputFieldSKID: `Subject Key Identifier:
    F7:7D:C5:FD:C4:E8:9A:1B:77:64:A7:F5:1D:A0:CC:BF:87:60:9A:6D`,
				OutputFieldAKID: `Authority Key Identifier:
    -`,
				OutputFieldSANs: `Subject Alternative Name:
    -`,
				OutputFieldRawCert: `Raw Certificate:
    -----BEGIN CERTIFICATE-----
    MIIFgzCCA2ugAwIBAgIPXZONMGc2yAYdGsdUhGkHMA0GCSqGSIb3DQEBCwUAMDsx
    CzAJBgNVBAYTAkVTMREwDwYDVQQKDAhGTk1ULVJDTTEZMBcGA1UECwwQQUMgUkFJ
    WiBGTk1ULVJDTTAeFw0wODEwMjkxNTU5NTZaFw0zMDAxMDEwMDAwMDBaMDsxCzAJ
    BgNVBAYTAkVTMREwDwYDVQQKDAhGTk1ULVJDTTEZMBcGA1UECwwQQUMgUkFJWiBG
    Tk1ULVJDTTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBALpxgHpMhm5/
    yBNtwMZ9HACXjywMI7sQmkCpGreHiPibVmr75nuOi5KOpyVdWRHbNi63URcfqQgf
    BBckWKo3Shjf5TnUV/3XwSyRAZHiItQDwFj8d0fsjz50Q7qsNI1NOHZnjrDIbzAz
    WHFctPVrbtQBULgTfmxKo0nRIBnuvMApGGWn3v7v3QqQIecaZ5JCEJhfTzC8PhxF
    tBDXaEAUwED653cXeuYLj2VbPNmaUtu1vZ5Gzz3rkQUCwJaydkxNEJY7kvqcfw+Z
    374jNUUeAlz+taibmSXaXvMiwzn15Cou08YfxGyqxRxqAQVKL9LFwag0Jl1mpdIC
    IfkYtwb1TplvqKtMUejPUBjFd8g5CSxJkjKZqLsXF3mwWsXmo8RZZUc1g16p6DUL
    mbvkzSDGm0oGObVo/CK67lWMK07q87Hj/LaZmtVC+nFNCM+HHmpxffnTtOmlcYF7
    wk5HlqX2doWjKI/pgG6BU6VtX7hI+cL5NqYuSf+4lsKMB7ObiFj86xsc3i1w4peS
    MKGJ47xVqCfWS+2QrYv6YyVZLag13cqXM7zlzced0ezvXg5KkAYmY6252TUtB7p2
    ZSysV4999AeU14ECll2jB0nVetBX+RvnU0Z1qrB5QstocQjpYL05ac70r8NWQMet
    UqIJ5G+GR4of6ygnXYMgrwTJbFaai0b1AgMBAAGjgYMwgYAwDwYDVR0TAQH/BAUw
    AwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYEFPd9xf3E6Jobd2Sn9R2gzL+H
    YJptMD4GA1UdIAQ3MDUwMwYEVR0gADArMCkGCCsGAQUFBwIBFh1odHRwOi8vd3d3
    LmNlcnQuZm5tdC5lcy9kcGNzLzANBgkqhkiG9w0BAQsFAAOCAgEAB5BK3/MjTvDD
    nFFlm5wioooMhfNzKWtN/gHiqQxjAb8EZ6WdmF/9ARP67Jpi6Yb+tmLSbkyU+8B1
    RXxlDPiyN8+sD8+Nb/kZ94/sHvJwnvDKuO+3/3Y3dlv2bojzr2IyIpMNOmqOFGYM
    LVN0V2Ue1bLdI4E7pWYjJ2cJj+F3qkPNZVEI7VFY/uY5+ctHhKQV8Xa7pO6kO8Rf
    77IzlhEYt8llvhjho6Tc+hj507wTmzl6NLrTQfv6MooqtyuGC2mDOL7Nii4LcK2N
    JpLuHvUBKwrZ1pebbuCoGRw6IYsMHkCtA+fdZn71uSANA+iW+YJF1DngoABd15jm
    fZ5nc8OaKveri6E6FO80vFIOiZiaBECEHX5FaZNXzuvO+FB8TxxuBEOb+dY7Ixjp
    6o7RTUaN8Tvkasq6+yO3m/qZASlaWFot4/nUbQ4mrcFuNLwy+AwF+mWj2zs3gyLp
    1txyM/1d8iC9djwj2ij3+RvrWWTV3F9yfiD8zYm1kGdNYno/Tq0dwzn+evQoFt9B
    9kiABdcPUXmsEKvU7ANm5mqwujGSQkBqvjrTcuFqN1W8rB2Vt2lh8kORdOag0wok
    RqEIr9baRRmW1FMdW4R58MD3R++Lj8UGrp1MYp3/RgT408m2ECVAdf4WqslKYIYv
    uu8wd+RU4riEmViAqhOLUTpPSPaLtrM=
    -----END CERTIFICATE-----`,
				OutputFieldSourceFile:       "From file: test",
				OutputFieldCertificateIndex: "0",
			},
		},
		"OutputFieldSourceFile as invalid selected field": {
			inputCertIndex:      0,
			inputSelectedFields: []Outputfield{OutputFieldCertificateIndex},
			expectedError:       "invalid OutputField",
		},
		"mixed valid and invalid fields": {
			inputCertIndex:      0,
			inputSelectedFields: []Outputfield{OutputFieldSubject, OutputFieldCertificateIndex},
			expectedError:       "invalid OutputField",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := certformatter.GetFormattedCertificate(tc.inputCertIndex, tc.inputSelectedFields...)

			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError, "error expected: %v", tc.expectedError)
				return
			}

			assert.Equal(t, tc.expectedOutput, got)
		})
	}
}
