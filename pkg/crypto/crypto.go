package crypto

import (
	"encoding/json"
	"log"
	"net/http"
)

type BitcoinPrice struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type GetBtcResponse struct {
	Data BitcoinPrice `json:"data"`
}

func GetBtcPrice() BitcoinPrice {

	req, err := http.NewRequest("GET", "https://api.coinbase.com/v2/prices/spot?currency=USD", nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	var getBtcResponse GetBtcResponse

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&getBtcResponse); err != nil {
		log.Println(err)
	}

	return getBtcResponse.Data

}
