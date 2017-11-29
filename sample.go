package main

import (
	"fmt"

	cc "github.com/shunsukuda/coincheck-go/Coincheck"
)

func main() {
	access_key := "ZKMSZf_y_Py1MN7t"
	api_secret := "rWsk_Ytg-rQJFCAqZuZq2BVYQ3If1J0N"
	client := new(cc.CoinCheck).NewClient(access_key, api_secret)
	// um, raw := client.GetTicker(1)
	// um, raw := client.GetTrades(1)
	// um, raw := client.GetOrderBooks(1)
	// fmt.Printf("%#v %s\n", um, raw)
	um, _ := client.GetLeverageBalance(1)
	fmt.Println(um)
}
