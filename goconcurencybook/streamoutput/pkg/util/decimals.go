package util

import (
	"fmt"
	"math/rand"

	pricefmt "github.com/shopspring/decimal"
)

func SayHello() {
	fmt.Println("Hello there!")
}

func Decimal(val string) *pricefmt.Decimal {
	a, _ := pricefmt.NewFromString(val)
	return &a
}

func RandomDecimal(value_1, value_2 float32) string {
	randomFloat := value_1 + value_2*rand.Float32()
	return fmt.Sprintf("%f", randomFloat)
}

func RandomDecimal2() *pricefmt.Decimal {
	randomFloat := 0.0 + 99999.01*rand.Float32()
	rndStr := fmt.Sprintf("%f", randomFloat)
	a, _ := pricefmt.NewFromString(rndStr)
	return &a
}

func FakeDecimal(val, fake *pricefmt.Decimal) *pricefmt.Decimal {
	if val == nil {
		return fake
	}
	structVal := *val
	if structVal.Equal(pricefmt.Zero) {
		return fake
	}
	return val
}

func FakeInt(val string) string {
	if val == "" || val == "0" {
		return fmt.Sprintf("%d", rand.Int31n(10))
	}
	return val
}
