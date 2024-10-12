package cmd

import "testing"

func TestList(t *testing.T) {
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
