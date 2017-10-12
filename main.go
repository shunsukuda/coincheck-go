package main

import cc "github.com/shunsukuda/coincheck-go"

func main() {
	_ := new(cc.CoinCheck).NewClient("", "")
}
