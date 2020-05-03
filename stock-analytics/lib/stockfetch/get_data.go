package stockfetch

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	stock "github.com/marvin5064/stock-analytics/protobuf/stock"
)

type AlphavantageQueryReturn struct {
	Metadata         AlphavantageMetadata              `json:"Meta Data"`
	TimeSeriesPrices map[string]AlphavantageDailyPrice `json:"Time Series (Daily)"`
}
type AlphavantageMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}
type AlphavantageDailyPrice struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type SortableStockPriceList []*stock.StockPrice

func (l SortableStockPriceList) Len() int {
	return len(l)
}
func (l SortableStockPriceList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// latest to old
func (l SortableStockPriceList) Less(i, j int) bool {
	return l[i].GetTime() > l[j].GetTime()
}

func (m *manager) GetData(request *stock.StockPriceRequest) (*stock.StockPriceResponse, error) {
	_, body, errs := m.requester.
		Get("https://www.alphavantage.co/query").
		Param("function", "TIME_SERIES_DAILY").
		Param("symbol", request.GetSymbol()).
		Param("apikey", m.apiKey).
		End()

	if len(errs) != 0 {
		return nil, fmt.Errorf("%v", errs)
	}

	return parseDataReturn(body)
}

func parseDataReturn(body string) (*stock.StockPriceResponse, error) {
	dataReturned := &AlphavantageQueryReturn{}
	err := json.Unmarshal([]byte(body), dataReturned)
	if err != nil {
		return nil, err
	}
	grpcReturn := &stock.StockPriceResponse{
		Symbol: dataReturned.Metadata.Symbol,
	}

	var sortableConvertedPrices SortableStockPriceList
	for k, v := range dataReturned.TimeSeriesPrices {
		convertedPrice, err := convertPriceData(k, v)
		if err != nil {
			return nil, err
		}
		sortableConvertedPrices = append(sortableConvertedPrices, convertedPrice)
	}
	grpcReturn.Prices = sortableConvertedPrices
	return grpcReturn, nil
}

func convertPriceData(timeDate string, price AlphavantageDailyPrice) (*stock.StockPrice, error) {
	var err error
	stockPrice := &stock.StockPrice{}
	stockPrice.Close, err = strconv.ParseFloat(price.Close, 64)
	if err != nil {
		return nil, err
	}
	stockPrice.High, err = strconv.ParseFloat(price.High, 64)
	if err != nil {
		return nil, err
	}
	stockPrice.Low, err = strconv.ParseFloat(price.Low, 64)
	if err != nil {
		return nil, err
	}
	stockPrice.Open, err = strconv.ParseFloat(price.Open, 64)
	if err != nil {
		return nil, err
	}
	stockPrice.Volume, err = strconv.ParseUint(price.Volume, 10, 64)
	if err != nil {
		return nil, err
	}

	parsedTime, err := time.Parse("2006-01-02", timeDate)
	if err != nil {
		return nil, err
	}
	stockPrice.Time = uint64(parsedTime.Unix())

	return stockPrice, nil
}
