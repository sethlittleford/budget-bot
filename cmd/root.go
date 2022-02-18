package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "budget-bot",
	Short: "A CLI to automate keeping a financial budget",
	Long: `A CLI that automates fetching financial transactions
from users' bank accounts via the Plaid API, then categorizing
those transaction and storing them in a human-readable format
in a spreadsheet via the Google Sheets API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define your flags and configuration settings.
}


