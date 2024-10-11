package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestStableFileLoad(t *testing.T) {
	cases := map[string]struct {
		inputParams                    []string
		expectedError                  string
		expectedCertstoreElements      int
		expectedFirstCertificateSerial string
		expectedLastCertificateSerial  string
	}{
		"one big file": {
			inputParams:                    []string{".internal/certs/testdata/tls-ca-bundle.pem"},
			expectedCertstoreElements:      142,
			expectedFirstCertificateSerial: "5e:c3:b7:a6:43:7f:a4:e0",
			expectedLastCertificateSerial:  "43:e3:71:13:d8:b3:59:14:5d:b7:ce:8c:fd:35:fd:6f:bc:05:8d:45",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			config := NewConfiguration()
			rootCmd := config.GetRootCmd()

			_, stderr, err := executeTest(rootCmd, tc.inputParams...)
			if err != nil {
				t.Errorf("execute error: %v", err)
			}

			if stderr != tc.expectedError {
				t.Errorf("expected error: %v - got: %v", tc.expectedError, stderr)
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
				t.Errorf("expected lastcertificate serial: %v - got: %v", tc.expectedLastCertificateSerial, gotLastCertificateSerial)
			}
		})
	}
}

func executeTest(cmd *cobra.Command, args ...string) (std, stderr string, reterr error) {
	cmd.SetArgs(args)
	stdBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	cmd.SetOut(stdBuf)
	cmd.SetErr(stderrBuf)
	err := cmd.Execute()

	return strings.TrimSpace(stdBuf.String()), strings.TrimSpace(stderrBuf.String()), err
}
