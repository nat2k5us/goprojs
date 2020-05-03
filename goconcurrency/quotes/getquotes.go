package quotes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetQuotesParallel() {
	numComplete := 0
	stockSymbols := []string{"AMZN", "AAPL", "DIS", "NVDA", "TSLA"}

	start := time.Now()
	for _, symbol := range stockSymbols {
		go func(symbol string) {
			url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + "NVBWKBA9R7V7WQR9"

			//fmt.Println(url)
			//data := getDataFromUrl(url)
			getDataFromUrl(url)
			//fmt.Println(data)
			numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		time.Sleep(1 * time.Millisecond)
	}
	elasped := time.Since(start)
	fmt.Println("execution time:", elasped)

}

func UnmarshalQuote(data []byte) (QuoteResponse, error) {
	var r QuoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *QuoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type QuoteResponse struct {
	The01Symbol           string `json:"01. symbol"`
	The02Open             string `json:"02. open"`
	The03High             string `json:"03. high"`
	The04Low              string `json:"04. low"`
	The05Price            string `json:"05. price"`
	The06Volume           string `json:"06. volume"`
	The07LatestTradingDay string `json:"07. latest trading day"`
	The08PreviousClose    string `json:"08. previous close"`
	The09Change           string `json:"09. change"`
	The10ChangePercent    string `json:"10. change percent"`
}

func getDataFromUrl(url string) string {

	var data string
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}

	return string(data)
}

func lineListSource(ctx context.Context, lines ...string) (
	<-chan string, <-chan error, error) {
	if len(lines) == 0 {
		// Handle an error that occurs before the goroutine begins.
		return nil, nil, fmt.Errorf("no lines provided")
	}
	out := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		for lineIndex, line := range lines {
			if line == "" {
				// Handle an error that occurs during the goroutine.
				errc <- fmt.Errorf("line %v is empty", lineIndex+1)
				return
			}
			// Send the data to the output channel but return early
			// if the context has been cancelled.
			select {
			case out <- line:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc, nil
}
