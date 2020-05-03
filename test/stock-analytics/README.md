[![Build Status](https://travis-ci.org/stocks-playground/stock-analytics.svg?branch=master)](https://travis-ci.org/stocks-playground/stock-analytics)
# stock-analytics

## prerequisite
```
cp config/config.sample.json config/config.json
```

## run
```
make run
```

## check output
make sure you have:
```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

```
grpcurl -plaintext -d '{"symbol": "0941.HK"}' localhost:8080 StockManager.GetStockPrices > tmp.out
```
