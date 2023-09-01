package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//nolint:golint,gochecknoglobals
var (
	RootCmd = &cobra.Command{
		Use:               "meshchat",
		RunE:              runRoot,
		SilenceErrors:     true,
		DisableAutoGenTag: true,
	}
)

//nolint:golint,gochecknoinits
func init() {
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug logging")
}

func runRoot(cmd *cobra.Command, _ []string) error {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return err
	}

	if debug {
		fmt.Println("debug logging enabled")
	}

	return nil
}
