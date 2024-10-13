package certs

import (
	"reflect"
	"testing"
)

func TestGetFormattedCertificate(t *testing.T) {
	cs := NewCertstore()
	cs.Load(getSampleCert("multiple certificates with comments"), "test")
	certformatter := cs.NewFormatter()

	cases := map[string]struct {
		inputCertIndex      int
		inputSelectedFields []Outputfield
		expected            FormattedCertificate
	}{
		"single field selected": {
			inputCertIndex:      0,
			inputSelectedFields: []Outputfield{OutputFieldSubject},
			expected: FormattedCertificate{
				OutputFieldSubject:          "Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldSourceFile:       "test",
				OutputFieldCertificateIndex: "0",
			},
		},
		"all fields selected": {
			inputCertIndex: 0,
			expected: FormattedCertificate{
				OutputFieldSubject:          "Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldIssuer:           "Issuer: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
				OutputFieldSourceFile:       "test",
				OutputFieldCertificateIndex: "0",
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, _ := certformatter.GetFormattedCertificate(tc.inputCertIndex, tc.inputSelectedFields...)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected: %v - got: %v", tc.expected, got)
			}
		})
	}
}
