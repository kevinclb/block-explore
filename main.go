package main

import (
	"github.com/kevinclb/block-explore/explorer"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get-btc", explorer.GetBTCBuyPriceHandler)
	http.HandleFunc("/get-eth", explorer.GetBTCSellPriceHandler)
	http.HandleFunc("get-btc-sell", explorer.GetETHBuyPriceHandler)
	http.HandleFunc("/get-eth-sell", explorer.GetETHSellPriceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
