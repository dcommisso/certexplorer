package certs

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		"single field selected": {
			inputCertIndex:      0,
			inputSelectedFields: []Outputfield{OutputFieldSubject},
			expectedOutput: FormattedCertificate{
				OutputFieldSubject:          "Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldSourceFile:       "test",
				OutputFieldCertificateIndex: "0",
			},
		},
		"all fields selected": {
			inputCertIndex: 0,
			expectedOutput: FormattedCertificate{
				OutputFieldSubject:          "Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldIssuer:           "Issuer: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldSourceFile:       "test",
				OutputFieldCertificateIndex: "0",
			},
		},
		"OutputFieldSourceFile as invalid selected field": {
			inputCertIndex:      0,
			inputSelectedFields: []Outputfield{OutputFieldSourceFile},
			expectedError:       "invalid OutputField",
		},
		"OutputFieldCertificateIndex as invalid selected field": {
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

			if !reflect.DeepEqual(got, tc.expectedOutput) {
				t.Errorf("expected: %v - got: %v", tc.expectedOutput, got)
			}
		})
	}
}
