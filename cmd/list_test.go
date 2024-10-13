package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func executeTest(cmd *cobra.Command, args ...string) (std, stderr string, reterr error) {
	cmd.SetArgs(args)
	stdBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	cmd.SetOut(stdBuf)
	cmd.SetErr(stderrBuf)
	err := cmd.Execute()

	return strings.TrimSpace(stdBuf.String()), strings.TrimSpace(stderrBuf.String()), err
}

func getTestCerts() string {
	return `-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----
# emSign Root CA - G1
-----BEGIN CERTIFICATE-----
MIIDlDCCAnygAwIBAgIKMfXkYgxsWO3W2DANBgkqhkiG9w0BAQsFADBnMQswCQYD
VQQGEwJJTjETMBEGA1UECxMKZW1TaWduIFBLSTElMCMGA1UEChMcZU11ZGhyYSBU
ZWNobm9sb2dpZXMgTGltaXRlZDEcMBoGA1UEAxMTZW1TaWduIFJvb3QgQ0EgLSBH
MTAeFw0xODAyMTgxODMwMDBaFw00MzAyMTgxODMwMDBaMGcxCzAJBgNVBAYTAklO
MRMwEQYDVQQLEwplbVNpZ24gUEtJMSUwIwYDVQQKExxlTXVkaHJhIFRlY2hub2xv
Z2llcyBMaW1pdGVkMRwwGgYDVQQDExNlbVNpZ24gUm9vdCBDQSAtIEcxMIIBIjAN
BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAk0u76WaK7p1b1TST0Bsew+eeuGQz
f2N4aLTNLnF115sgxk0pvLZoYIr3IZpWNVrzdr3YzZr/k1ZLpVkGoZM0Kd0WNHVO
8oG0x5ZOrRkVUkr+PHB1cM2vK6sVmjM8qrOLqs1D/fXqcP/tzxE7lM5OMhbTI0Aq
d7OvPAEsbO2ZLIvZTmmYsvePQbAyeGHWDV/D+qJAkh1cF+ZwPjXnorfCYuKrpDhM
tTk1b+oDafo6VGiFbdbyL0NVHpENDtjVaqSW0RM8LHhQ6DqS0hdW5TUaQBw+jSzt
Od9C4INBdN+jzcKGYEho42kLVACL5HZpIQ15TjQIXhTCzLG3rdd8cIrHhQIDAQAB
o0IwQDAdBgNVHQ4EFgQU++8Nhp6w492pufEhF38+/PB3KxowDgYDVR0PAQH/BAQD
AgEGMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAFn/8oz1h31x
PaOfG1vR2vjTnGs2vZupYeveFix0PZ7mddrXuqe8QhfnPZHr5X3dPpzxz5KsbEjM
wiI/aTvFthUvozXGaCocV685743QNcMYDHsAVhzNixl03r4PEuDQqqE/AjSxcM6d
GNYIAwlG7mDgfrbESQRRfXBgvKqy/3lyeqYdPV8q+Mri/Tm3R7nrft8EI6/6nAYH
6ftjk4BAtcZsCjEozgyfz7MjNYBBjWzEN3uBL4ChQEKF6dk4jeihU80Bv2noWgby
RQuQ+q7hv53yrlc8pa6yVvSLZUDp/TGBLPQ5Cdjua6e0ph0VpZj3AYHYhX3zUVxx
iN66zB+Afko=
-----END CERTIFICATE-----

# vTrus ECC Root CA
-----BEGIN CERTIFICATE-----
MIICDzCCAZWgAwIBAgIUbmq8WapTvpg5Z6LSa6Q75m0c1towCgYIKoZIzj0EAwMw
RzELMAkGA1UEBhMCQ04xHDAaBgNVBAoTE2lUcnVzQ2hpbmEgQ28uLEx0ZC4xGjAY
BgNVBAMTEXZUcnVzIEVDQyBSb290IENBMB4XDTE4MDczMTA3MjY0NFoXDTQzMDcz
MTA3MjY0NFowRzELMAkGA1UEBhMCQ04xHDAaBgNVBAoTE2lUcnVzQ2hpbmEgQ28u
LEx0ZC4xGjAYBgNVBAMTEXZUcnVzIEVDQyBSb290IENBMHYwEAYHKoZIzj0CAQYF
K4EEACIDYgAEZVBKrox5lkqqHAjDo6LN/llWQXf9JpRCux3NCNtzslt188+cToL0
v/hhJoVs1oVbcnDS/dtitN9Ti72xRFhiQgnH+n9bEOf+QP3A2MMrMudwpremIFUd
e4BdS49nTPEQo0IwQDAdBgNVHQ4EFgQUmDnNvtiyjPeyq+GtJK97fKHbH88wDwYD
VR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwCgYIKoZIzj0EAwMDaAAwZQIw
V53dVvHH4+m4SVBrm2nDb+zDfSXkV5UTQJtS0zvzQBm8JsctBp61ezaf9SXUY2sA
AjEA6dPGnlaaKsyh2j/IZivTWJwghfqrkYpwcBE4YGQLYgmRWAD5Tfs0aNoJrSEG
GJTO
-----END CERTIFICATE-----`
}

func getTestdataDir() string {
	return "../internal/certformatter/testdata/"
}

func TestFileLoad(t *testing.T) {
	cases := map[string]struct {
		inputParams               []string
		inputStdin                string
		expectedError             string
		expectedCertstoreElements int

		expectedFirstCertificateSerial string
		expectedFirstCertificateSource string

		expectedLastCertificateSerial string
		expectedLastCertificateSource string
	}{
		"one big file": {
			inputParams:                    []string{"list", getTestdataDir() + "tls-ca-bundle.pem"},
			expectedCertstoreElements:      142,
			expectedFirstCertificateSerial: "5e:c3:b7:a6:43:7f:a4:e0",
			expectedFirstCertificateSource: getTestdataDir() + "tls-ca-bundle.pem",

			expectedLastCertificateSerial: "43:e3:71:13:d8:b3:59:14:5d:b7:ce:8c:fd:35:fd:6f:bc:05:8d:45",
			expectedLastCertificateSource: getTestdataDir() + "tls-ca-bundle.pem",
		},
		"multiple files": {
			inputParams:                    []string{"list", getTestdataDir() + "tls-ca-bundle.pem", getTestdataDir() + "example.com.crt"},
			expectedCertstoreElements:      143,
			expectedFirstCertificateSerial: "5e:c3:b7:a6:43:7f:a4:e0",
			expectedFirstCertificateSource: getTestdataDir() + "tls-ca-bundle.pem",

			expectedLastCertificateSerial: "39:28:ed:76:45:6f:84:d0:77:9a:cb:0c:0f:e2:f4:d3:87:e5:b3:64",
			expectedLastCertificateSource: getTestdataDir() + "example.com.crt",
		},
		"multiple certs from stdin": {
			inputParams:                    []string{"list"},
			inputStdin:                     getTestCerts(),
			expectedCertstoreElements:      3,
			expectedFirstCertificateSerial: "39:28:ed:76:45:6f:84:d0:77:9a:cb:0c:0f:e2:f4:d3:87:e5:b3:64",
			expectedFirstCertificateSource: "[stdin]",

			expectedLastCertificateSerial: "6e:6a:bc:59:aa:53:be:98:39:67:a2:d2:6b:a4:3b:e6:6d:1c:d6:da",
			expectedLastCertificateSource: "[stdin]",
		},
		"invalid file": {
			inputParams:   []string{"list", getTestdataDir() + "tls-ca-bundle.pem", "/not/existent/file"},
			expectedError: "Error: open /not/existent/file: no such file or directory",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			config := NewConfiguration()
			rootCmd := config.GetRootCmd()

			if tc.inputStdin != "" {
				rootCmd.SetIn(strings.NewReader(tc.inputStdin))
			}

			_, stderr, _ := executeTest(rootCmd, tc.inputParams...)

			if stderr != tc.expectedError {
				t.Errorf("expected error: %v - got: %v", tc.expectedError, stderr)
			}

			if tc.expectedError != "" {
				return
			}

			gotCertstoreElements := len(config.certstore.Certs)
			if gotCertstoreElements != tc.expectedCertstoreElements {
				t.Errorf("expected number of certificates: %v - got: %v", tc.expectedCertstoreElements, gotCertstoreElements)
			}

			gotFirstCertificateSerial := config.certstore.Certs[0].GetSerialNumber()
			if gotFirstCertificateSerial != tc.expectedFirstCertificateSerial {
				t.Errorf("expected first certificate serial: %v - got: %v", tc.expectedFirstCertificateSerial, gotFirstCertificateSerial)
			}

			lastCertificateIndex := len(config.certstore.Certs) - 1
			gotLastCertificateSerial := config.certstore.Certs[lastCertificateIndex].GetSerialNumber()
			if gotLastCertificateSerial != tc.expectedLastCertificateSerial {
				t.Errorf("expected last certificate serial: %v - got: %v", tc.expectedLastCertificateSerial, gotLastCertificateSerial)
			}

			gotFirstCertificateSource := config.certstore.Certs[0].Source
			if gotFirstCertificateSource != tc.expectedFirstCertificateSource {
				t.Errorf("expected first certificate source: %v - got: %v", tc.expectedFirstCertificateSource, gotFirstCertificateSource)
			}

			gotLastCertificateSource := config.certstore.Certs[lastCertificateIndex].Source
			if gotLastCertificateSource != tc.expectedLastCertificateSource {
				t.Errorf("expected last certificate source: %v - got: %v", tc.expectedLastCertificateSource, gotLastCertificateSource)
			}
		})
	}
}

func TestList(t *testing.T) {
	t.Skip("Test skipped")
	cases := map[string]struct {
		inputParams   []string
		expectedOut   string
		expectedError string
	}{
		"no flags": {
			inputParams: []string{"list", getTestdataDir() + "fewCertificates.pem"},
			expectedOut: `[0] Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES
    Subject Alternative Name:
        -
    Validity:
        Not Before: 2008-10-29 15:59:56 +0000 UTC
        Not After : 2030-01-01 00:00:00 +0000 UTC

[1] Subject: CN=AC RAIZ FNMT-RCM SERVIDORES SEGUROS,OU=Ceres,O=FNMT-RCM,C=ES,2.5.4.97=#130f56415445532d51323832363030344a
    Subject Alternative Name:
        -
    Validity:
        Not Before: 2018-12-20 09:37:33 +0000 UTC
        Not After : 2043-12-20 09:37:33 +0000 UTC

[2] Subject: SERIALNUMBER=G63287510,CN=ANF Secure Server Root CA,OU=ANF CA Raiz,O=ANF Autoridad de Certificacion,C=ES
    Subject Alternative Name:
        -
    Validity:
        Not Before: 2019-09-04 10:00:38 +0000 UTC
        Not After : 2039-08-30 10:00:38 +0000 UTC

[3] Subject: CN=example.com
    Subject Alternative Name:
        DNS:example.com, DNS:*.example.com, IP Address:10.0.0.1, IP Address:127.0.0.1
    Validity:
        Not Before: 2024-10-07 15:44:12 +0000 UTC
        Not After : 2034-10-05 15:44:12 +0000 UTC
`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			config := NewConfiguration()
			cmd := config.GetRootCmd()

			std, stderr, _ := executeTest(cmd, tc.inputParams...)

			if stderr != tc.expectedError {
				t.Errorf("expected error: %v - got: %v", tc.expectedError, stderr)
			}

			if tc.expectedError != "" {
				return
			}

			if std != tc.expectedOut {
				t.Errorf("expected output: %v - got: %v", tc.expectedOut, std)
			}
		})
	}
}
