package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	pricefmt "github.com/shopspring/decimal"
)

func init() {
	mod := pricefmt.NewFromFloat(1.00)
	fake := pricefmt.NewFromFloat(420.69)
	fake = fake.DivRound(mod, 2)
	// FakePriceValue = &fake
	fakePercent := pricefmt.NewFromFloat(0.69).Mod(mod)
	fakePercent = fakePercent.DivRound(mod, 2)
	// FakePercentValue = &fakePercent
}

type AssetType string

const (
	assetTypeEquities AssetType = "EQ"
	assetTypeCrypto   AssetType = "CRYPTO"
	assetTypeFutures  AssetType = "FU"
	assetTypeOptions  AssetType = "OP"
	assetTypeCash     AssetType = "CASH"
)

var AssetTypes = struct {
	Equities AssetType
	Crypto   AssetType
	Futures  AssetType
	Options  AssetType
	Cash     AssetType
}{
	Equities: assetTypeEquities,
	Crypto:   assetTypeCrypto,
	Futures:  assetTypeFutures,
	Options:  assetTypeOptions,
	Cash:     assetTypeCash,
}

func (a *AssetType) MarshalJSON() ([]byte, error) {
	if a == nil {
		return nil, errors.New("Invalid")
	}
	switch *a {
	case assetTypeEquities:
		return json.Marshal("EQ")
	case assetTypeCrypto:
		return json.Marshal("CRYPTO")
	case assetTypeOptions:
		return json.Marshal("OP")
	case assetTypeFutures:
		return json.Marshal("FU")
	case assetTypeCash:
		return json.Marshal("CASH")

	default:
		return json.Marshal(string(*a))
		// e := errors.New("UNKNOWN Asset Type")
		// log.WithField("asset-type", *a).WithError(e).Error("Failed to marshal asset type")
		// return nil, e
	}
}

type Item struct {
	AssetType    AssetType         `json:"AssetType,omitempty"`
	TotalValue   *pricefmt.Decimal `json:"TotalValue,omitempty"`
	ValuePercent *pricefmt.Decimal `json:"ValuePercent,omitempty"`
}

type Holding struct {
	Type  string `json:"Type"` // enum of "Assets" and "Sectors"
	Items []Item `json:"Items"`
}
type Account struct {
	AccountID                string            `json:"AccountID"`
	DisplayName              string            `json:"DisplayName,omitempty"`
	AccountProfitLossPercent *pricefmt.Decimal `json:"AccountProfitLossPercent,omitempty"`
	AccountProfitLoss        *pricefmt.Decimal `json:"AccountProfitLoss,omitempty"`
	TotalValue               *pricefmt.Decimal `json:"TotalValue,omitempty"`
	EnabledAssets            []string          `json:"EnabledAssets,omitempty"`
}
type AccountSummary struct {
	EnabledAssets     []string          `json:"EnabledAssets,omitempty"`
	NumberOfAccounts  int               `json:"NumberOfAccounts,omitempty"`
	TotalValue        *pricefmt.Decimal `json:"TotalValue,omitempty"`
	ProfitLoss        *pricefmt.Decimal `json:"ProfitLoss,omitempty"`
	ProfitLossPercent *pricefmt.Decimal `json:"ProfitLossPercent,omitempty"`
	Accounts          []Account         `json:"Accounts,omitempty"`
}

type AccountOverview struct {
	TotalPortfolioValue             *pricefmt.Decimal `json:"TotalPortfolioValue,omitempty"`
	TotalPortfolioProfitLoss        *pricefmt.Decimal `json:"TotalPortfolioProfitLoss,omitempty"`
	TotalPortfolioProfitLossPercent *pricefmt.Decimal `json:"TotalPortfolioProfitLossPercent,omitempty"`
	Holdings                        []Holding         `json:"Holdings,omitempty"`
	AccountSummaries                []AccountSummary  `json:"AccountSummaries,omitempty"`
}

// NewAccountOverview returns all of the fields.
func NewAccountOverview() (AccountOverview, error) {

	out := AccountOverview{}
	out.FakeIt()
	return out, nil
}

func (a *AccountOverview) FakeIt() {

	a.TotalPortfolioValue = FakeDecimal(a.TotalPortfolioValue, fake)
	a.TotalPortfolioProfitLoss = FakeDecimal(a.TotalPortfolioProfitLoss, FakePriceValue)
	a.TotalPortfolioProfitLossPercent = util.FakeDecimal(a.TotalPortfolioProfitLossPercent, FakePriceValue)

	item1 := Item{AssetType: "EQ", TotalValue: util.RandomDecimal2(), ValuePercent: util.RandomDecimal2()}
	item2 := Item{AssetType: "OP", TotalValue: util.RandomDecimal2(), ValuePercent: util.RandomDecimal2()}
	item3 := Item{AssetType: "FU", TotalValue: util.RandomDecimal2(), ValuePercent: util.RandomDecimal2()}
	item4 := Item{AssetType: "CRTPTO", TotalValue: util.RandomDecimal2(), ValuePercent: util.RandomDecimal2()}
	item5 := Item{AssetType: "CASH", TotalValue: util.RandomDecimal2(), ValuePercent: util.RandomDecimal2()}
	FakeHoldings := []Holding{{"Assets", []Item{item1, item2, item3, item4, item5}},
		{"Sectors", nil},
	}
	a.Holdings = FakeHoldings
	acctInvestments := Account{AccountID: "SIM2323223", DisplayName: "Investments", AccountProfitLossPercent: util.RandomDecimal2(), AccountProfitLoss: util.RandomDecimal2(), TotalValue: util.RandomDecimal2(), EnabledAssets: []string{"EQ", "OP", "CRYPTO"}}
	acctCrypto := Account{AccountID: "SIM2323224", DisplayName: "Crypto", AccountProfitLossPercent: util.RandomDecimal2(), AccountProfitLoss: util.RandomDecimal2(), TotalValue: util.RandomDecimal2(), EnabledAssets: []string{"CRYPTO"}}
	acctDayTrades := Account{AccountID: "SIM2323225", DisplayName: "Day Trading", AccountProfitLossPercent: util.RandomDecimal2(), AccountProfitLoss: util.RandomDecimal2(), TotalValue: util.RandomDecimal2(), EnabledAssets: []string{"EQ", "OP", "CRYPTO"}}
	acctFutures := Account{AccountID: "SIM2323226", DisplayName: "Futures", AccountProfitLossPercent: util.RandomDecimal2(), AccountProfitLoss: util.RandomDecimal2(), TotalValue: util.RandomDecimal2(), EnabledAssets: []string{"FU"}}
	FakeAccountSummaries := []AccountSummary{
		{EnabledAssets: []string{"EQ"}, NumberOfAccounts: 3, TotalValue: util.RandomDecimal2(), ProfitLoss: util.RandomDecimal2(), ProfitLossPercent: util.RandomDecimal2(), Accounts: []Account{acctInvestments, acctCrypto, acctDayTrades}},
		{EnabledAssets: []string{"FU"}, NumberOfAccounts: 1, TotalValue: util.RandomDecimal2(), ProfitLoss: util.RandomDecimal2(), ProfitLossPercent: util.RandomDecimal2(), Accounts: []Account{acctFutures}},
	}
	a.AccountSummaries = FakeAccountSummaries

}

// func (a AccountOverview) decimalUpdateIfNotNil(existing, update *pricefmt.Decimal) *pricefmt.Decimal {
// 	return existing
// }

// TODO: GetChanges return updated fields only
func (a AccountOverview) GetChanges(old AccountOverview) AccountOverview {
	return a
}

// type outputWriter interface {
// 	WriteError(code, message string, fatal bool) error
// 	WriteObject(...interface{}) error
// }

func NewAccountOverviewStream(ctx context.Context, accountOverview <-chan interface{}) <-chan interface{} {
	outputCh := make(chan interface{})
	currentAccountOverview := AccountOverview{}
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Debug("AccountOverviewService.Run ctx.Done")
				return
			case unk := <-accountOverview:
				xType := fmt.Sprintf("%T", unk)
				fmt.Println(xType)
				switch asTyp := unk.(type) {
				case nil:
					log.Warn("AccountDetailService chan nil")
					outputCh <- errors.New("stream closed")
					return
				case error, v3.WebAPIError:
					fmt.Println("AccountOverviewService", asTyp)
					outputCh <- errors.New("error from BODBalances?")
					return
				case v3.Balance:
					newAccountOverview, err := NewAccountOverview()
					if err != nil {
						log.WithError(err).Error("error creating account overview")
						outputCh <- errors.New("error calculating account overview")
						return
					}
					outputCh <- newAccountOverview.GetChanges(currentAccountOverview)
					currentAccountOverview = newAccountOverview
				}
			}
		}
	}()
	return outputCh
}
func NewFakeAccountOverviewStream(ctx context.Context) <-chan interface{} {
	outputCh := make(chan interface{})
	overview := AccountOverview{}
	overview.FakeIt()
	go func() {
		for {
			// output
			select {
			case <-ctx.Done():
				return
			case outputCh <- overview:
				// sent
			}
			// wait
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second * 10):
				overview.FakeIt()
			}
		}
	}()
	return outputCh
}

func main() {
	SayHello("Sam")
}
