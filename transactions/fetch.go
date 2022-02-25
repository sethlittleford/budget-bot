package transactions

import (
	"fmt"
)

// Fetch gets all transactions for the given month from all financial
// institutions with which budget-bot is authorized via the Plaid API
func Fetch(month string) {
	// var err error
	// plaidClient, err = NewClient()
	// if err != nil {
	// 	log.Fatal("Error creating Plaid API Client: ", err)
	// }

	fmt.Println("Fetch() called with:", month)
}
