package explorer

import (
	"github.com/fabioberger/coinbase-go"
	"github.com/fabioberger/coinbase-go/config"
	"os"
)

//InstantiateClient creates an instance of the Coinbase client.
//The coinbase client can be further used to run commands with your Coinbase account.

func InstantiateClient() coinbase.Client {
	config.BaseUrl = "https://api.coinbase.com/v1/"
	return coinbase.ApiKeyClient(os.Getenv("CB_API_KEY"), os.Getenv("CB_SECRET"))
}

//TODO: Write more client code here.
