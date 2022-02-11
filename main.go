package main

import "github.com/kevinclb/block-explore/explorer"

func main() {
	//explorer.InitializeHandlers()
	//c := explorer.InstantiateClient()
	//balance, err := c.GetBalance()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Balance is %f BTC", balance)

	boolResult, err := explorer.TryCatcher(true)

	print("the boolean result was ", boolResult, "and the error was: ", err.Error())
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
