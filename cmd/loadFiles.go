package cmd

import (
	"errors"
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
		return nil
	}

	// check if there's stuff on stdin
	stdinstat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if stdinstat.Mode()&os.ModeNamedPipe != 0 || config.testMode {
		b, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}
		config.certstore.Load(b, stdinLabel)
		return nil
	}

	return errors.New("no files or stdin provided\n")
}
