package cmd

import (
	"errors"
	"fmt"

	"github.com/dcommisso/certexplorer/certformatter"
	"github.com/spf13/cobra"
)

func (c *Configuration) GetListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list certificates from files or standard input",
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

			validSelectors := getValidSelectors()

			// fields order when no fields are selected
			orderedDefaultFields, err := validSelectors.getDefaultOrder()
			if err != nil {
				return err
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
				selectedOutputField = append(selectedOutputField, validSelectors.getOutputField(field))
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
	validSelectors := getValidSelectors()
	c.Flags().StringSliceP("fields", "f", []string{}, validSelectors.getFullUsage("List of fields to show: `field1,field2,...`"))
	c.Flags().IntSliceP("certificates", "c", []int{}, "Certificate index numbers to show.")
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
