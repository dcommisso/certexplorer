/*
Copyright © 2024 Domenico Commisso dcommiss@redhat.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dcommisso/certexplorer/certformatter"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (c *Configuration) GetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     getUse(),
		Short:   getShort(),
		Long:    getLongDescription(),
		Example: getExample(),
		Version: getVersion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			noColor, _ := cmd.Flags().GetBool("no-color")
			if noColor {
				color.NoColor = true
			}

			err := LoadFilesOrStdin(cmd, c)
			if err != nil {
				return err
			}

			formatter := c.certstore.NewFormatter()
			validSelectors := getValidSelectors()
			validOutputs := getValidOuput()
			selectedOutput, _ := cmd.Flags().GetString("output")
			formatter, err = validOutputs.getFormatter(c.certstore, selectedOutput)
			if err != nil {
				return err
			}

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
	getRootCmdSetFlags(cmd)
	return cmd
}

func Execute() {
	config := NewConfiguration()
	rootCmd := config.GetRootCmd()
	// needed because, for some mysterious reasons, cmd.Print default to stderr.
	// Put it here because this function it's not executed during test.
	rootCmd.SetOut(os.Stdout)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getRootCmdSetFlags(c *cobra.Command) {
	validSelectors := getValidSelectors()
	validOutputs := getValidOuput()
	c.Flags().StringSliceP("fields", "f", []string{}, validSelectors.getFullUsage("List of fields to show: `field1,field2,...`"))
	c.Flags().IntSliceP("certificates", "c", []int{}, "Certificate index numbers to show: `index1,index2...`")
	c.Flags().StringP("output", "o", "nice", validOutputs.getFullUsage("Output `format`"))
	c.Flags().Bool("no-color", false, "Disable colors.")
}

func getUse() string {
	return "certexplorer [file1 file2 ...]"
}

func getShort() string {
	return "certexplorer (nicely) shows certificates with flexible output options"
}

func getLongDescription() string {

	longDescription := `
certexplorer

certexplorer helps to go through cabundle files full of certificates as well as
quickly view a single certificate. By default the certificates are nicely
formatted, but it's possible to disable colors or get them in a plain output for
easily manage the raw certificates somewhere else. The output options permit to
select which certificates and which fields to show (in the order given). The
certificates can be read from multiple files or standard input.

The default output ("nice") also color the expiration date (red for expired
certificates, yellow for certificates expiring in less than a month).
`
	return strings.TrimSpace(longDescription)
}

func getExample() string {
	examples := `
# Get all the default fields for all the certificates in provided files
$ certexplorer certificate.pem cabundle-full-of-certificates.pem

# Get only the serial number and the validity of the certificates in provided files
$ certexplorer certificate.pem cabundle-full-of-certificates.pem -f serial,validity

# Get only the serial number and the validity of the certificates 119 and 3 in provided files
$ certexplorer certificate.pem cabundle-full-of-certificates.pem -f serial,validity -c 119,3

# Get the expiration date and the subject of a certificate from standard input
echo "-----BEGIN CERTIFICATE----- ..." | certexplorer -f notafter,subject

# Get some certificates from cabundle in plain output (useful for get the raw certificate without indentation)
$ certexplorer cabundle-full-of-certificates.pem -c 1,5 -o plain
`
	return strings.TrimSpace(examples)
}

func getVersion() string {
	return "v1.0.0"
}
