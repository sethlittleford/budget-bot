package cmd

import (
	"github.com/sethlittleford/budget-bot/transactions"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <month> <year>",
	Short: "Fetch, categorize, and write transactions to Google Sheets for the given month",
	Long: `Fetch all transactions for the given month and year from all financial
institutions with which budget-bot is authorized.

Categorize these transactions into custom budgets.

Write the transaction data & category to Google Sheets.

The <month> argument must be a valid month taking any of the following formats, e.g.:
	January
	january
	Jan
	jan
	1
	
The <year> argument must be any year >= 2022 in the format YYYY`,
	Example: `budget-bot run January 2022
budget-bot run 1 2022
budget-bot run December 2023
budget-bot run 12 2023`,
	ValidArgs: append(transactions.Months(), transactions.Years()...),
	Args:      cobra.ExactValidArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		transactions.Fetch(args[0], args[1])
		return nil
	},
}
