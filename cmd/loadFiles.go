package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

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
		b, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}
		config.certstore.Load(b, stdinLabel)
	}
	return nil
}
