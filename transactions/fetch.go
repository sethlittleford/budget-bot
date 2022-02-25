package transactions

import "log"

var plaidClient *Client

func init() {
	var err error
	plaidClient, err = NewClient()
	if err != nil {
		log.Fatal("Error creating Plaid API Client: ", err)
	}
}
