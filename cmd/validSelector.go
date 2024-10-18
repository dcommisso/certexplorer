package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"text/tabwriter"

	"github.com/dcommisso/certexplorer/certformatter"
)

// key of the map is the name of the selector passed on command line
type validSelectorSet map[string]struct {
	// the correspondent Outputfield passed to certformatter library.
	outputField certformatter.Outputfield
	// it will appear to help description
	description string
	// priority defines the order (from lower to higher) in which the fields will be
	// outputted by default. If a field has not a priority value, or has a priority
	// 0, it will not present in the default output. Two validSelectors can't have
	// the same priority number.
	priority int
}

func (v validSelectorSet) getOutputField(selectorName string) certformatter.Outputfield {
	return v[selectorName].outputField
}

func (v validSelectorSet) getFieldDescription(selectorName string) string {
	return v[selectorName].description
}

func (v validSelectorSet) getFullUsage(header string) string {
	b := new(bytes.Buffer)
	w := tabwriter.NewWriter(b, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%v\n\n", header)
	selectorsAlphabeticalOrder := []string{}
	for selector := range v {
		selectorsAlphabeticalOrder = append(selectorsAlphabeticalOrder, selector)
	}
	slices.Sort(selectorsAlphabeticalOrder)
	for _, field := range selectorsAlphabeticalOrder {
		fmt.Fprintf(w, "%v\t%v\n", field, v.getFieldDescription(field))
	}
	w.Flush()
	return b.String()
}

func (v validSelectorSet) getDefaultOrder() ([]string, error) {
	orderedSelectors := []string{}
	priorities := map[int]string{}
	for name := range v {
		currentPrio := v[name].priority

		// priority 0 will be discarded
		if currentPrio == 0 {
			continue
		}

		//check if we already stored a selector with the same priority
		if _, ok := priorities[currentPrio]; ok {
			return []string{}, errors.New("multiple validSelectors have the same priority")
		}

		priorities[currentPrio] = name
	}

	// we need to start from 1 since the lowest priority is 1 and the highest
	// len(v) + 1
	for i := 1; i < len(v)+1; i++ {
		if selectorName, ok := priorities[i]; ok {
			orderedSelectors = append(orderedSelectors, selectorName)
		}
	}
	return orderedSelectors, nil
}

func getValidSelectors() validSelectorSet {
	return validSelectorSet{
		"serial": {
			outputField: certformatter.OutputFieldSerialNumber,
			description: "Serial Number",
			priority:    1,
		},
		"issuer": {
			outputField: certformatter.OutputFieldIssuer,
			description: "Issuer",
			priority:    2,
		},
		"subject": {
			outputField: certformatter.OutputFieldSubject,
			description: "Subject",
			priority:    3,
		},
		"validity": {
			outputField: certformatter.OutputFieldValidity,
			description: "Validity of the certificate",
			priority:    5,
		},
		"notbefore": {
			outputField: certformatter.OutputFieldNotBefore,
			description: "Not Before date of certificate",
			priority:    0,
		},
		"notafter": {
			outputField: certformatter.OutputFieldNotAfter,
			description: "Not After date of certificate",
			priority:    0,
		},
		"skid": {
			outputField: certformatter.OutputFieldSKID,
			description: "Subject Key Identifier",
			priority:    7,
		},
		"akid": {
			outputField: certformatter.OutputFieldAKID,
			description: "Authority Key Identifier",
			priority:    6,
		},
		"san": {
			outputField: certformatter.OutputFieldSANs,
			description: "Subject Alternative Names",
			priority:    4,
		},
		"raw": {
			outputField: certformatter.OutputFieldRawCert,
			description: "Raw certificate",
			priority:    9,
		},
		"source": {
			outputField: certformatter.OutputFieldSourceFile,
			description: "The file containing the certificate",
			priority:    8,
		},
	}
}
