package certformatter

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
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
	fieldsFormatFunctions map[Outputfield]func(c Certificate) string
	// CertstoreFormatFunction formats the certstore or part of it
	composeFunction func(certs []FormattedCertificate, orderedFieldsToRender []Outputfield) (string, error)
}

func (c *Certstore) NewFormatter() *Formatter {
	return &Formatter{
		certstore: c,
		fieldsFormatFunctions: map[Outputfield]func(c Certificate) string{
			OutputFieldSubject:      formatSubject,
			OutputFieldIssuer:       formatIssuer,
			OutputFieldSerialNumber: formatSerialNumber,
			OutputFieldValidity:     formatValidity,
			OutputFieldNotBefore:    formatNotBefore,
			OutputFieldNotAfter:     formatNotAfter,
			OutputFieldSKID:         formatSKID,
			OutputFieldAKID:         formatAKID,
			OutputFieldSANs:         formatSANs,
			OutputFieldRawCert:      formatRawCert,
			OutputFieldSourceFile:   formatSourceFile,
		},
		composeFunction: composeCertificates,
	}
}

// SetFieldFormatFunction substitutes the default format function for a field
// with the provided one.
func (f *Formatter) SetFieldFormatFunction(field Outputfield, formatFunc func(c Certificate) string) {
	f.fieldsFormatFunctions[field] = formatFunc
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
	for field, formatFunction := range f.fieldsFormatFunctions {
		if len(selectedFields) == 0 || slices.Contains(selectedFields, field) {
			fcToReturn[field] = formatFunction(f.certstore.Certs[certIndex])
		}
	}
	fcToReturn[OutputFieldCertificateIndex] = strconv.Itoa(certIndex)
	return fcToReturn, nil
}

func (f *Formatter) ComposeFormattedCertificates(formattedCertificates []FormattedCertificate, orderedFieldsToRender []Outputfield) (string, error) {
	return f.composeFunction(formattedCertificates, orderedFieldsToRender)
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

// default OutputFieldSKID format function
func formatSKID(c Certificate) string {
	label := "Subject Key Identifier"
	skid := c.GetSKID()
	if skid == "" {
		skid = "-"
	}
	return fmt.Sprintf("%s:\n    %s", label, skid)
}

// default OutputFieldAKID format function
func formatAKID(c Certificate) string {
	label := "Authority Key Identifier"
	akid := c.GetAKID()
	if akid == "" {
		akid = "-"
	}
	return fmt.Sprintf("%s:\n    %s", label, akid)
}

// default OutputFieldSourceFile format function
func formatSourceFile(c Certificate) string {
	return fmt.Sprintf("From file: %s", c.Source)
}

// default OutputFieldSANs format function
func formatSANs(c Certificate) string {
	label := "Subject Alternative Name"
	sans := c.GetSANs()
	if sans == "" {
		sans = "-"
	}
	return fmt.Sprintf("%s:\n    %s", label, sans)
}

// default OutputFieldRawCert format function
func formatRawCert(c Certificate) string {
	rawCert := c.GetRawCert()
	rawCertLines := strings.Split(rawCert, "\n")
	rawCertFormattedLines := []string{}
	for _, line := range rawCertLines {
		rawCertFormattedLines = append(rawCertFormattedLines, "    "+line)
	}
	rawCertFormatted := strings.Join(rawCertFormattedLines, "\n")
	return fmt.Sprintf("Raw Certificate:\n%s", rawCertFormatted)
}

// default composeFunction
func composeCertificates(certs []FormattedCertificate, orderedFieldsToRender []Outputfield) (string, error) {

	b := bytes.Buffer{}
	w := tabwriter.NewWriter(&b, 0, 0, 0, ' ', tabwriter.AlignRight)

	if len(certs) == 0 {
		return "", errors.New("certs cannot be empty")
	}

	if len(orderedFieldsToRender) == 0 {
		return "", errors.New("orderedFieldsToRender cannot be empty")
	}

	for _, cert := range certs {
		certLines := []string{}
		for _, field := range orderedFieldsToRender {
			fieldLines := strings.Split(cert[field], "\n")
			certLines = append(certLines, fieldLines...)
		}
		fmt.Fprintf(w, "[%s] \t%s\n", cert[OutputFieldCertificateIndex], certLines[0])
		for i := 1; i < len(certLines); i++ {
			fmt.Fprintf(w, "\t%s\n", certLines[i])
		}
	}
	w.Flush()
	return strings.TrimSpace(b.String()), nil
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
