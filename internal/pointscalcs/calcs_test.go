package pointscalcs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/pointscalcs"
	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
)

var _ = Describe("Calcs", func() {

	Describe("ByRetailerName", func() {
		It("calculates point as one point for every alphanumeric character in the retailer name", func() {
			var (
				expectedReceipt = receipts.Receipt{
					Retailer: "Target",
				}
				expectedPoints = 6
			)

			got := pointscalcs.ByRetailerName(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})
	})

	Describe("ByRoundedTotal", func() {
		It("gives 50 points if the total is a round dollar amount with no cents", func() {
			var (
				expectedReceipt = receipts.Receipt{
					Total: 9.00,
				}
				expectedPoints = 50
			)

			got := pointscalcs.ByRoundedTotal(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})
	})

	Describe("ByTotalMultipleOf25", func() {
		It("gives 25 points if the total is a multiple of 0.25", func() {
			var (
				expectedReceipt = receipts.Receipt{
					Total: 9.00,
				}
				expectedPoints = 25
			)

			got := pointscalcs.ByTotalMultipleOf25(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})
	})

	Describe("ByEach2Items", func() {
		It("5 points for every two items on the receipt", func() {

		})
	})

	Describe("ByItemtrimmedDescription", func() {
		It("trimmes item description, vaidates is multiple of 3, multiply it by 3 and rounde it up to the nearest integer", func() {

		})
	})

	Describe("ByPurchaseDayIsOdd", func() {
		It("gives 6 points if the day in the purchase date is odd", func() {

		})
	})

	Describe("ByPurchaseTwoFourInterval", func() {
		It("gives 10 points if the time of purchase is after 2:00pm and before 4:00pm", func() {

		})
	})
})
