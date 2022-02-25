package transactions

import (
	"errors"
	"os"

	"github.com/plaid/plaid-go/v2/plaid"
)

type Client struct {
	*plaid.APIClient
}

func NewClient() (*Client, error) {
	config := plaid.NewConfiguration()

	clientID, ok := os.LookupEnv("PLAID_CLIENT_ID")
	if !ok {
		return nil, errors.New("PLAID_CLIENT_ID is not present in the environment")
	}
	devSecret, ok := os.LookupEnv("PLAID_DEVELOPMENT_SECRET")
	if !ok {
		return nil, errors.New("PLAID_DEVELOPMENT_SECRET is not present in the environment")
	}

	config.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	config.AddDefaultHeader("PLAID-SECRET", devSecret)

	config.UseEnvironment(plaid.Development)

	return &Client{
		plaid.NewAPIClient(config),
	}, nil
}
