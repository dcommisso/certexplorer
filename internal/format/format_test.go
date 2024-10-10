package format

import (
	"math/big"
	"testing"
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
