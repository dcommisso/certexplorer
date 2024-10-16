package cmd

import (
	"errors"
	"fmt"

	"github.com/dcommisso/cabundleinspect/certformatter"
	"github.com/spf13/cobra"
)

func (c *Configuration) GetListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := LoadFilesOrStdin(cmd, c)
			if err != nil {
				return err
			}

			formatter := c.certstore.NewFormatter()

			validSelectors := map[string]certformatter.Outputfield{
				"serial":    certformatter.OutputFieldSerialNumber,
				"issuer":    certformatter.OutputFieldIssuer,
				"subject":   certformatter.OutputFieldSubject,
				"validity":  certformatter.OutputFieldValidity,
				"notbefore": certformatter.OutputFieldNotBefore,
				"notafter":  certformatter.OutputFieldNotAfter,
				"skid":      certformatter.OutputFieldSKID,
				"akid":      certformatter.OutputFieldAKID,
				"san":       certformatter.OutputFieldSANs,
				"raw":       certformatter.OutputFieldRawCert,
				"source":    certformatter.OutputFieldSourceFile,
			}

			// fields order when no fields are selected
			orderedDefaultFields := []string{
				"serial",
				"issuer",
				"subject",
				"san",
				"validity",
				"akid",
				"skid",
				"source",
				"raw",
			}

			selectedFields, _ := cmd.Flags().GetStringSlice("fields")
			selectedCertIndexes, _ := cmd.Flags().GetIntSlice("certificates")

			// return error if selectedFields contains invalid field
			for _, selectedField := range selectedFields {
				if _, ok := validSelectors[selectedField]; !ok {
					return errors.New("invalid field")
				}
			}

			// if no field was selected use default
			if len(selectedFields) == 0 {
				selectedFields = orderedDefaultFields
			}

			// convert string selected fields to Outputfields
			selectedOutputField := []certformatter.Outputfield{}
			for _, field := range selectedFields {
				selectedOutputField = append(selectedOutputField, validSelectors[field])
			}

			certsToRender := []certformatter.FormattedCertificate{}
			// if no cert index was selected, get them all
			if len(selectedCertIndexes) == 0 {
				for i := 0; i < len(c.certstore.Certs); i++ {
					certsToRender = append(certsToRender, formatter.GetFormattedCertificate(i))
				}
			} else {
				for _, i := range selectedCertIndexes {
					if _, ok := c.certstore.Certs[i]; !ok {
						return errors.New(fmt.Sprintf("certificate index %v out of range", i))
					}
					certsToRender = append(certsToRender, formatter.GetFormattedCertificate(i))
				}
			}

			renderedOutput, err := formatter.ComposeFormattedCertificates(certsToRender, selectedOutputField)
			if err != nil {
				return err
			}
			cmd.Println(renderedOutput)

			return nil
		},
	}
	getListCmdSetFlags(cmd)
	return cmd
}

func getListCmdSetFlags(c *cobra.Command) {
	c.Flags().StringSliceP("fields", "f", []string{}, "Fields to show.")
	c.Flags().IntSliceP("certificates", "c", []int{}, "Certificate index numbers to show.")
}
