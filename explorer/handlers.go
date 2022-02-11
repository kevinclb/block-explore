package explorer

import (
	"fmt"
	"net/http"
)

func GetBTCBuyPriceHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting btc buy price")
	body, _ := GetPriceData(BTCBuy)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}

func GetBTCSellPriceHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting btc sell price")
	body, _ := GetPriceData(BTCSell)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}
func GetETHBuyPriceHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting eth buy price")
	body, _ := GetPriceData(ETHBuy)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}

func GetETHSellPriceHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting eth sell price")
	body, _ := GetPriceData(ETHSell)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		return
	}
}

func InitializeHandlers() {
	http.HandleFunc("/get-btc", GetBTCBuyPriceHandler)
	http.HandleFunc("/get-eth", GetBTCSellPriceHandler)
	http.HandleFunc("get-btc-sell", GetETHBuyPriceHandler)
	http.HandleFunc("/get-eth-sell", GetETHSellPriceHandler)

	TryCatcher(true)
}

func TryCatcher(isThereAnError bool) (bool, error) {
	if isThereAnError == true {
		fmt.Println("There is an error!")
		return isThereAnError, ErrorThrower{customData: 3}
	} else {
		return isThereAnError, ErrorThrower{customData: 2}
	}
}

type ErrorThrower struct {
	customData int
}

func (e ErrorThrower) Error() string {
	if e.customData == 3 {
		return "i was told to create this error (ErrorThrower had 3)"
	} else {
		return "i was not told to create this error (ErrorThrower did not have 3)"
	}
}
