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
	}
	rootCmdSubcommands(cmd, c)
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

// Add subcommands here
func rootCmdSubcommands(cmd *cobra.Command, c *Configuration) {
	cmd.AddCommand(c.GetListCmd())
}
