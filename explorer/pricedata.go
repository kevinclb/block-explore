package explorer

import (
	_ "github.com/fabioberger/coinbase-go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//Creating a back end market data price fetcher
//by taking advantage of free, public, key-less
//Coinbase API.

//the base coinbase url along with constant paths
//which may be appended to the URL
//depending on the specific market data desired.
const (
	coinbaseURL = "https://api.coinbase.com/v2/prices/"
	BTCBuy      = "BTC-USD/buy"
	ETHBuy      = "ETH-USD/buy"
	BTCSell     = "BTC-USD/sell"
	ETHSell     = "ETH-USD/sell"
)

//GetPriceData takes a url
func GetPriceData(urlPath string) ([]byte, string) {
	requestURL := urlSpecifier(urlPath)
	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		log.Fatal("error creating new request: ", err)
	}
	req.Header.Add("Accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error: ", err)
		}
	}(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	return body, string(body)
}
