package repositories

import (
	"github.com/orbulant/random-user-agent/useragent"
	"resty.dev/v3"
)

const (
	YahooFinanceBaseUrl = "https://query1.finance.yahoo.com/v6/finance/options/MSFT"
)

func FetchStockDataFromAPI() (any, error) {
	client := resty.New()

	userAgent := useragent.New().GetRandomAgent()

	res, err := client.R().SetHeader("User-Agent", userAgent).Get(YahooFinanceBaseUrl)

	if err != nil {
		return nil, err
	}

	return res.String(), nil
}
