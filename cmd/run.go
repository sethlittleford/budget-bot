package cmd

import (
	"github.com/sethlittleford/budget-bot/transactions"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <month>",
	Short: "Fetch, categorize, and write transactions to Google Sheets for the given month",
	Long: `Fetch all transactions for the given month from all financial
institutions with which budget-bot is authorized.

Categorize these transactions into custom budgets.

Write the transaction data & category to Google Sheets.

The <month> argument must be a valid month taking any of the following formats, e.g.:
	January
	january
	Jan
	jan
	1`,
	Example: `budget-bot run January
budget-bot run 1
budget-bot run December
budget-bot run 12`,
	ValidArgs: transactions.Months(),
	Args:      cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		transactions.Fetch(args[0])
		return nil
	},
}
