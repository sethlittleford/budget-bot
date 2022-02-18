package cmd

import (
	"fmt"

	"github.com/sethlittleford/budget-bot/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display current budget-bot version",
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := version.Version()
		if err != nil {
			return err
		}
		fmt.Print(v)
		return nil
	},
}
