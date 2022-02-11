package main

import (
	"fmt"
	"github.com/kevinclb/block-explore/explorer"
	"log"
)

func main() {
	explorer.InitializeHandlers()
	c := explorer.InstantiateClient()
	balance, err := c.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance is %f BTC", balance)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
