package stockfetch

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#parseDataReturn", func() {
	var (
		input string
	)
	Describe("#DrawShape", func() {
		BeforeEach(func() {
			input = `
			{
				"Meta Data": {
					"1. Information": "Daily Prices (open, high, low, close) and Volumes",
					"2. Symbol": "0941.HK",
					"3. Last Refreshed": "2018-08-10",
					"4. Output Size": "Compact",
					"5. Time Zone": "US/Eastern"
				},
				"Time Series (Daily)": {
					"2018-08-10": {
						"1. open": "71.4500",
						"2. high": "71.9000",
						"3. low": "71.3500",
						"4. close": "71.8500",
						"5. volume": "14193465"
					},
					"2018-08-09": {
						"1. open": "70.5000",
						"2. high": "71.8000",
						"3. low": "70.1500",
						"4. close": "71.3000",
						"5. volume": "16674602"
					}
				}
			}`
		})

		Context("success case", func() {
			It("should return", func() {
				rec, err := parseDataReturn(input)
				Expect(err).To(BeNil())
				Expect(len(rec.GetPrices())).To(Equal(2))
				Expect(rec.GetPrices()[0].GetTime() > rec.GetPrices()[1].GetTime()).To(BeTrue())
			})
		})
	})
})
