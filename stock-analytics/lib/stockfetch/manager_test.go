package stockfetch

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStockFetcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stock Fetcher Suite")
}

// below are stubs
var _ = BeforeSuite(func() {
})

var _ = AfterSuite(func() {
})
