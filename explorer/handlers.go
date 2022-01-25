package explorer

import (
	"fmt"
	"net/http"
)

func GetBTCBuyPriceHandler (w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting btc buy price")
	body, _ := GetMarketData(BTCBuy)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}

func GetBTCSellPriceHandler (w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting btc sell price")
	body, _ := GetMarketData(BTCSell)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}
func GetETHBuyPriceHandler (w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting eth buy price")
	body, _ := GetMarketData(ETHBuy)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}

func GetETHSellPriceHandler (w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting eth sell price")
	body, _ := GetMarketData(ETHSell)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}