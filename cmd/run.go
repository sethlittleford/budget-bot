package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	ValidArgs: months(),
	Args:      cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("run called with arg:", args[0])
		return nil
	},
}

func months() []string {
	m := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	months := make([]string, 0)
	for _, month := range m {
		months = append(months, month.String())                      // e.g. January
		months = append(months, strings.ToLower(month.String()))     // e.g. january
		months = append(months, month.String()[:3])                  // e.g. Jan
		months = append(months, strings.ToLower(month.String()[:3])) // e.g. jan
		months = append(months, strconv.Itoa(int(month)))            // e.g. 1
	}
	return months
}
