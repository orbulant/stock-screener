package repositories

import (
	"fmt"

	"github.com/orbulant/random-user-agent/useragent"
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

	yres, yerr := c.R().Get(YahooFinanceBaseUrl)

	if err != nil || yerr != nil {
		return nil, err
	}

	// Process the response to JSON

	response := resp.String()
	yresponse := yres.String()

	fmt.Println(response)
	fmt.Println(yresponse)

	ua := useragent.New().GetRandomAgent()
	fmt.Println(ua)

	return nil, nil
}
