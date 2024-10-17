package cmd

import (
	"testing"

	"github.com/dcommisso/cabundleinspect/certformatter"
	"github.com/stretchr/testify/assert"
)

func getValidSelectorSet() validSelectorSet {
	return validSelectorSet{
		"validity": {
			outputField: certformatter.OutputFieldValidity,
			description: "validity of the certificate",
			priority:    4,
		},
		"subject": {
			outputField: certformatter.OutputFieldSubject,
			description: "subject",
			priority:    3,
		},
		"serial": {
			outputField: certformatter.OutputFieldSerialNumber,
			description: "serial number",
			priority:    1,
		},
		"issuer": {
			outputField: certformatter.OutputFieldIssuer,
			description: "issuer",
			priority:    2,
		},
		"notbefore": {
			outputField: certformatter.OutputFieldNotBefore,
			description: "Not Before date",
		},
	}
}

func getInvalidSelectorSet() validSelectorSet {
	return validSelectorSet{
		"validity": {
			outputField: certformatter.OutputFieldValidity,
			description: "validity of the certificate",
			priority:    5,
		},
		"subject": {
			outputField: certformatter.OutputFieldSubject,
			description: "subject",
			priority:    3,
		},
		"serial": {
			outputField: certformatter.OutputFieldSerialNumber,
			description: "serial number",
			priority:    5,
		},
		"issuer": {
			outputField: certformatter.OutputFieldIssuer,
			description: "issuer",
			priority:    2,
		},
		"notbefore": {
			outputField: certformatter.OutputFieldNotBefore,
			description: "Not Before date",
		},
	}
}

func TestGetFullUsage(t *testing.T) {
	vss := getValidSelectorSet()
	gotFullUsage := vss.getFullUsage("List of fields to show: field1,field2,...")
	expectedFullUsage := `List of fields to show: field1,field2,...

issuer      issuer
notbefore   Not Before date
serial      serial number
subject     subject
validity    validity of the certificate
`
	assert.Equal(t, gotFullUsage, expectedFullUsage)
}

func TestGetDefaultOrder(t *testing.T) {
	cases := map[string]struct {
		inputVss             validSelectorSet
		expectedDefaultOrder []string
		expectedError        string
	}{
		"valid vss": {
			inputVss:             getValidSelectorSet(),
			expectedDefaultOrder: []string{"serial", "issuer", "subject", "validity"},
		},
		"invalid vss": {
			inputVss:      getInvalidSelectorSet(),
			expectedError: "multiple validSelectors have the same priority",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			gotDefaultOrder, err := tc.inputVss.getDefaultOrder()

			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
				return
			}

			assert.Equal(t, tc.expectedDefaultOrder, gotDefaultOrder)
		})
	}
}
