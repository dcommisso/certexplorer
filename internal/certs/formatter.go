package certs

import (
	"errors"
	"slices"
	"strconv"
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
			OutputFieldSubject: formatSubject,
			OutputFieldIssuer:  formatIssuer,
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
