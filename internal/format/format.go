package format

import (
	"encoding/hex"
	"strings"
)

func ToColonNotation(hexNumber []byte) string {
	hexString := hex.EncodeToString(hexNumber)

	var splitted []string

	for i := 0; i < len(hexString); i += 2 {
		splitted = append(splitted, hexString[i:i+2])
	}

	return strings.Join(splitted, ":")
}
