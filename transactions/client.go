package transactions

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/plaid/plaid-go/v2/plaid"
)

func init() {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/budget-bot/.env")) // must be the absolute path for the .env file
	if err != nil {
		log.Fatal("Error loading .env file from root: ", err)
	}
}

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
