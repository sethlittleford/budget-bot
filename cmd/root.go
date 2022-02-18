package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "budget-bot",
	Short: "A CLI to automate keeping a financial budget",
	Long: `A CLI that automates fetching financial transactions
from users' bank accounts via the Plaid API, categorizes
those transaction, and stores them in human-readable format
in a spreadsheet via the Google Sheets API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Define your flags and configuration settings.
	
	// Add Commands to root
	rootCmd.AddCommand(versionCmd)
}
