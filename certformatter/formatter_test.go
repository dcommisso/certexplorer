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
