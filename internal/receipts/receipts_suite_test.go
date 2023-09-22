package receipts_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	T = GinkgoT()
)

func TestReceipts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Receipts Suite")
}
