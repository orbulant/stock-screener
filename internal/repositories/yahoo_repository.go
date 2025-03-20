package repositories

import (
	"fmt"

	"resty.dev/v3"
)

const (
	YahooFinanceBaseUrl = "https://query1.finance.yahoo.com/v6/finance/options/MSFT"
	CatUrl              = "https://catfact.ninja/fact"
)

func FetchStockDataFromAPI() (any, error) {
	baseUrl := CatUrl
	url := baseUrl

	// Using resty, a simple HTTP and REST client library for Go
	c := resty.New()
	defer c.Close()

	resp, err := c.R().Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Process the response to JSON

	response := resp.String()

	fmt.Println(response)

	return nil, nil
}
