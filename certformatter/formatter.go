package certformatter

import (
	"encoding/hex"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Outputfield int

const (
	OutputFieldSerialNumber = iota
	OutputFieldIssuer
	OutputFieldSubject
	OutputFieldValidity
	OutputFieldNotBefore
	OutputFieldNotAfter
	OutputFieldSKID
	OutputFieldAKID
	OutputFieldSANs
	OutputFieldRawCert
	OutputFieldSourceFile
	OutputFieldCertificateIndex
)

type FormattedCertificate map[Outputfield]string

type Formatter struct {
	certstore *Certstore
	//FieldsFormatFunctions contains the format function of each field
	FieldsFormatFunctions map[Outputfield]func(c Certificate) string
	// CertstoreFormatFunction formats the certstore or part of it
	CertstoreFormatFunction func([]FormattedCertificate) string
}

func (c *Certstore) NewFormatter() *Formatter {
	return &Formatter{
		certstore: c,
		FieldsFormatFunctions: map[Outputfield]func(c Certificate) string{
			OutputFieldSubject:      formatSubject,
			OutputFieldIssuer:       formatIssuer,
			OutputFieldSerialNumber: formatSerialNumber,
			OutputFieldValidity:     formatValidity,
			OutputFieldNotBefore:    formatNotBefore,
			OutputFieldNotAfter:     formatNotAfter,
		},
	}
}

// GetFormattedCertificate returns a FormattedCertificate with the fields
// rendered using the functions defined in FieldsFormatFunctions. If
// selectedFields parameter is defined only the selected fields are returned,
// otherwise all the fields are returned.
func (f *Formatter) GetFormattedCertificate(certIndex int, selectedFields ...Outputfield) (FormattedCertificate, error) {

	// OutputFieldCertificateIndex is metadata. It's always present and should
	// not be selected.
	if slices.Contains(selectedFields, OutputFieldCertificateIndex) {
		return FormattedCertificate{}, errors.New("invalid OutputField")
	}

	fcToReturn := make(FormattedCertificate)
	for field, formatFunction := range f.FieldsFormatFunctions {
		if len(selectedFields) == 0 || slices.Contains(selectedFields, field) {
			fcToReturn[field] = formatFunction(f.certstore.Certs[certIndex])
		}
	}
	fcToReturn[OutputFieldSourceFile] = f.certstore.Certs[certIndex].Source
	fcToReturn[OutputFieldCertificateIndex] = strconv.Itoa(certIndex)
	return fcToReturn, nil
}

// default OutputFieldSubject format function
func formatSubject(c Certificate) string {
	return "Subject: " + c.GetSubject()
}

// default OutputFieldIssuer format function
func formatIssuer(c Certificate) string {
	return "Issuer: " + c.GetIssuer()
}

// default OutputFieldSerialNumber format function
func formatSerialNumber(c Certificate) string {
	return fmt.Sprintf("Serial Number:\n    %s", c.GetSerialNumber())
}

// default OutputFieldValidity format function
func formatValidity(c Certificate) string {
	return fmt.Sprintf("Validity\n    Not Before: %s\n    Not After : %s", c.GetNotBefore(), c.GetNotAfter())
}

// default OutputFieldNotBefore format function
func formatNotBefore(c Certificate) string {
	return fmt.Sprintf("Not Before: %s", c.GetNotBefore())
}

// default OutputFieldNotAfter format function
func formatNotAfter(c Certificate) string {
	return fmt.Sprintf("Not After : %s", c.GetNotAfter())
}

// ToColonNotation adds colon to hex number. Example:
// F77DC5FDC4E89A1B7764A7F51DA0CCBF87609A6D ->
// F7:7D:C5:FD:C4:E8:9A:1B:77:64:A7:F5:1D:A0:CC:BF:87:60:9A:6D
func ToColonNotation(hexNumber []byte) string {
	hexString := hex.EncodeToString(hexNumber)

	var splitted []string

	for i := 0; i < len(hexString); i += 2 {
		splitted = append(splitted, hexString[i:i+2])
	}

	return strings.Join(splitted, ":")
}
