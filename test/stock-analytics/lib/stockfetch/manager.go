package stockfetch

import (
	stock "github.com/stocks-playground/stock-analytics/protobuf/stock"
	"github.com/parnurzeal/gorequest"
)

type manager struct {
	url       string
	apiKey    string
	requester *gorequest.SuperAgent
}

type Manager interface {
	GetData(request *stock.StockPriceRequest) (*stock.StockPriceResponse, error)
}

func New(url, apiKey string) Manager {
	return &manager{
		url:       url,
		apiKey:    apiKey,
		requester: gorequest.New(),
	}
}
