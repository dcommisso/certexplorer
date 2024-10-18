package cmd

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/dcommisso/certexplorer/certformatter"
	"github.com/fatih/color"
)

type validOutputs map[string]func(*certformatter.Certstore) *certformatter.Formatter

func (v validOutputs) getFullUsage(header string) string {
	sortedKeys := []string{}
	for name := range v {
		sortedKeys = append(sortedKeys, name)
	}
	slices.Sort(sortedKeys)

	return fmt.Sprintf("%s. One of: [%s]", header, strings.Join(sortedKeys, " |"))
}

func (v validOutputs) getFormatter(certstore *certformatter.Certstore, outputName string) (*certformatter.Formatter, error) {
	if formatter, ok := v[outputName]; !ok {
		return nil, errors.New(fmt.Sprintf("invalid output: %s", outputName))
	} else {
		return formatter(certstore), nil
	}
}

func getValidOuput() validOutputs {
	validOutputs := validOutputs{
		"nice": getNiceFormatter,
	}
	return validOutputs
}

// definition of nice formatter
func getNiceFormatter(certstore *certformatter.Certstore) *certformatter.Formatter {
	labelColor := color.New(color.FgHiBlue)
	formatter := certstore.NewFormatter()

	formatter.SetFieldFormatFunction(certformatter.OutputFieldSubject, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Subject: ")
		return label + c.GetSubject()
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldIssuer, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Issuer: ")
		return label + c.GetIssuer()
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldSerialNumber, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Serial Number:")
		return fmt.Sprintf("%s\n    %s", label, c.GetSerialNumber())
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldValidity, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Validity")
		return fmt.Sprintf("%s\n    Not Before: %s\n    Not After : %s", label, c.GetNotBefore(), c.GetNotAfter())
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldNotBefore, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Not Before:")
		return fmt.Sprintf("%s %s", label, c.GetNotBefore())
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldNotAfter, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Not After:")
		return fmt.Sprintf("%s %s", label, c.GetNotAfter())
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldSKID, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Subject Key Identifier")
		skid := c.GetSKID()
		if skid == "" {
			skid = "-"
		}
		return fmt.Sprintf("%s:\n    %s", label, skid)
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldAKID, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Authority Key Identifier")
		akid := c.GetAKID()
		if akid == "" {
			akid = "-"
		}
		return fmt.Sprintf("%s:\n    %s", label, akid)
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldSourceFile, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("From file:")
		return fmt.Sprintf("%s %s", label, c.Source)
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldSANs, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Subject Alternative Name")
		sans := c.GetSANs()
		if sans == "" {
			sans = "-"
		}
		return fmt.Sprintf("%s:\n    %s", label, sans)
	})

	formatter.SetFieldFormatFunction(certformatter.OutputFieldRawCert, func(c certformatter.Certificate) string {
		label := labelColor.Sprint("Raw Certificate:")
		rawCert := c.GetRawCert()
		rawCertLines := strings.Split(rawCert, "\n")
		rawCertFormattedLines := []string{}
		for _, line := range rawCertLines {
			rawCertFormattedLines = append(rawCertFormattedLines, "    "+line)
		}
		rawCertFormatted := strings.Join(rawCertFormattedLines, "\n")
		return fmt.Sprintf("%s\n%s", label, rawCertFormatted)
	})
	return formatter
}
