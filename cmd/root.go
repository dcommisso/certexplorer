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

	"github.com/dcommisso/certexplorer/certformatter"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (c *Configuration) GetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certexplorer",
		Short: "certexplorer allow to analyze big cabundle files",
		Long: `certexplorer is able to read certificates from multiple
files. The output is flexible and it's possible to choose
the certificate fields to show.`,
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
	c.Flags().IntSliceP("certificates", "c", []int{}, "Certificate index numbers to show.")
	c.Flags().StringP("output", "o", "nice", validOutputs.getFullUsage("Output `format`"))
	c.Flags().Bool("no-color", false, "Disable colors.")
}
