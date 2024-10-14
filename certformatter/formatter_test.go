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
		"all fields selected - none empty": {
			inputCertIndex: 3,
			expectedOutput: FormattedCertificate{
				OutputFieldSubject: "Subject: CN=example.com",
				OutputFieldIssuer:  "Issuer: CN=example.com",
				OutputFieldSerialNumber: `Serial Number:
    39:28:ed:76:45:6f:84:d0:77:9a:cb:0c:0f:e2:f4:d3:87:e5:b3:64`,
				OutputFieldValidity: `Validity
    Not Before: 2024-10-07 15:44:12 +0000 UTC
    Not After : 2034-10-05 15:44:12 +0000 UTC`,
				OutputFieldNotBefore: "Not Before: 2024-10-07 15:44:12 +0000 UTC",
				OutputFieldNotAfter:  "Not After : 2034-10-05 15:44:12 +0000 UTC",
				OutputFieldSKID: `Subject Key Identifier:
    12:97:38:99:6E:64:A2:7E:CB:2F:57:7D:5B:E6:10:17:F7:2A:CA:55`,
				OutputFieldAKID: `Authority Key Identifier:
    12:97:38:99:6E:64:A2:7E:CB:2F:57:7D:5B:E6:10:17:F7:2A:CA:55`,
				OutputFieldSANs: `Subject Alternative Name:
    DNS:example.com, DNS:*.example.com, IP Address:10.0.0.1, IP Address:127.0.0.1`,
				OutputFieldRawCert: `Raw Certificate:
    -----BEGIN CERTIFICATE-----
    MIIFQjCCAyqgAwIBAgIUOSjtdkVvhNB3mssMD+L004fls2QwDQYJKoZIhvcNAQEL
    BQAwFjEUMBIGA1UEAwwLZXhhbXBsZS5jb20wHhcNMjQxMDA3MTU0NDEyWhcNMzQx
    MDA1MTU0NDEyWjAWMRQwEgYDVQQDDAtleGFtcGxlLmNvbTCCAiIwDQYJKoZIhvcN
    AQEBBQADggIPADCCAgoCggIBANuEMkqyS0UrlkAMKMzPsmvdwp0fefH564JXXOZ5
    v5TnkubqI7ijKBXwSvPCTCuebZFYqIpqN58dSyeMewkEjngAH99+NAuIlt5nyAg5
    +BxZiRNgugdnc6kAcBH+2C/3T4P0HSd0NPGVQOC4GxMAy6Shz5LiztunPkWPUxWy
    8OZYO86xilW4L6Dv6JuPCsX7vzv90a31rtfghPlFUz4frcfGpdmTUeXv38/aVZRz
    ToGgMvQUIaVvfxfmTqrWLCTy/4sTvIY8mVgCaGw5wNDv2oUiZ7C18w0orWeY0vc4
    b+F+ma2Mwa04pj0KluwJFx85qH3lEk50yfnLYWUr3ZYwu2TOTqpDr61aR+NZqbw+
    QZWrUImipTzbtVqezmzCLQqXI3zPsmhXb6+k/Ykj1xj9deE7vYAIc8MBVxhw1pYQ
    fxRgx+dhyZpqYDW8fCUF/5vXtABm9we+/edL7vFKfBuSipeZeAie+a+SO6aMLCIZ
    rMnH1PdJjzKTbCuWGA6K75y6X3vh2qenmh0/qVJljKktWfWq4Gp5u6EjcZ2mljzh
    kDMMrRikWq/ThLOFLHC1wjy3IQ7SLGfnM9qharpHTVinwU1g97RGdzFXiB+y69O1
    MlIlIXWeNJsU94pf1VidKz1hLJynw5cx7ifU+eSvbRU0ZLIHjuX3G4gVlwkUKPNN
    rdgXAgMBAAGjgYcwgYQwHQYDVR0OBBYEFBKXOJluZKJ+yy9XfVvmEBf3KspVMB8G
    A1UdIwQYMBaAFBKXOJluZKJ+yy9XfVvmEBf3KspVMA8GA1UdEwEB/wQFMAMBAf8w
    MQYDVR0RBCowKIILZXhhbXBsZS5jb22CDSouZXhhbXBsZS5jb22HBAoAAAGHBH8A
    AAEwDQYJKoZIhvcNAQELBQADggIBAEFfjgkurGI/ouVaJgCJYXmf+mJtexN6JYAB
    XuCcobhUM+t5bfYt9DbZhaNC3pvtT9OijujajnXmd20QJLgXbWJe3qMbiDYroXZ/
    ry8JHX9Nlp3wF3V/iGUQw4zLnFRalAeSyEAwg+nsoTwA6vkCompNFqzozh2ViSo+
    ucWkb71Ky7Fl3HmKxp9ohG/0REwLTMYYmlPaGvk4o3oWAH3jm6g0fuFiw3mC1Gvr
    nFiMva0JJ3LvaOZfoe2U2dSCFEleNqVMEvLNRF4Sd6SFR7/IhwjzvyxSw0DRSmmt
    FZPakLw1Dm2eibT0rhEhn+7fLVHbaTSgOg8diXKqI+kqxbReYbrGkfa4lIfrvAMi
    +ax9OQ54KIdQU3uqugyyTqqN1WERjvuv0uhWSF1sEk6AdpMWR8ymircQ72yEJmjI
    ycRC2okZKkhYwFHb4nYwhQJOtzkmIUc90xmXZK4EvC9SYbHDu/8RBFRqkF4BnE8y
    4eoywGFMRIJ4GdAzaI7rpQmtSvbZsiismBkTGU/IibM6udnnp5xD4R3HDDSdrkMH
    MpupMy3sYHTJ+pi4OdpvYBfQHy1Cq0RxRc5SCSb0Dn6IggUVgyG0QpT4SuIM5qt8
    3cp1JPJuZpbks3EkjqeyzKgfWsJWWi/9q68zbubW5oOgHk2548BO9x2njvEVfrTg
    rDo5uzq8
    -----END CERTIFICATE-----`,
				OutputFieldSourceFile:       "From file: test",
				OutputFieldCertificateIndex: "3",
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
