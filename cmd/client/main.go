package main

import (
	"PaxfulGoBot/pkg/crypto"
	"fmt"
)

func main() {
	btcPrice := crypto.GetBtcPrice()
	btcPricep := &btcPrice
	btcPricep.Currency = "asdfasf gadf "
	fmt.Println(btcPrice)
}
