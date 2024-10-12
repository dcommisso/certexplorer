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
	"io"
	"os"

	"github.com/spf13/cobra"
)

func (c *Configuration) GetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cabundleinspect",
		Short: "cabundleinspect allow to analyze big cabundle files",
		Long: `cabundleinspect is able to read certificates from multiple
files. The output is flexible and it's possible to choose
the certificate fields to show.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := LoadFilesOrStdin(cmd, c)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func Execute() {
	config := NewConfiguration()
	rootCmd := config.GetRootCmd()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func LoadFilesOrStdin(cmd *cobra.Command, config *Configuration) error {
	const stdinLabel = "[stdin]"

	inputFiles := cmd.Flags().Args()
	if len(inputFiles) > 0 {
		for _, fname := range inputFiles {
			b, err := os.ReadFile(fname)
			if err != nil {
				return err
			}
			config.certstore.Load(b, fname)
		}
	} else {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		config.certstore.Load(b, stdinLabel)
	}
	return nil
}
