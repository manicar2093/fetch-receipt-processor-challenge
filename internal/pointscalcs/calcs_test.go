package pointscalcs_test

import (
	"time"

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

		Context("just count alphanumeric values", func() {
			It("returns total count", func() {
				var (
					expectedReceipt = receipts.Receipt{
						Retailer: "M&M Corner Market",
					}
					expectedPoints = 14
				)

				got := pointscalcs.ByRetailerName(expectedReceipt)

				Expect(got).To(Equal(expectedPoints))
			})
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
			var (
				expectedReceipt = receipts.Receipt{
					Items: []receipts.ReceiptItem{
						{},
						{},
						{},
						{},
						{},
					},
				}
				expectedPoints = 10
			)

			got := pointscalcs.ByEach2Items(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})
	})

	Describe("ByItemTrimmedDescription", func() {
		It("trimmes item description, vaidates is multiple of 3, multiply it by 3 and round it up to the nearest integer", func() {
			var (
				expectedReceipt = receipts.Receipt{
					Items: []receipts.ReceiptItem{
						{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
						{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
						{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
						{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
						{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
					},
				}
				expectedPoints = 6
			)

			got := pointscalcs.ByItemTrimmedDescription(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})
	})

	Describe("ByPurchaseDayIsOdd", func() {
		It("gives 6 points if the day in the purchase date is odd", func() {
			var (
				expectedReceipt = receipts.Receipt{
					PurchaseDate: receipts.Date(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)),
				}
				expectedPoints = 6
			)

			got := pointscalcs.ByPurchaseDayIsOdd(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})
	})

	Describe("ByPurchaseTwoFourInterval", func() {
		It("gives 10 points if the time of purchase is after 2:00pm and before 4:00pm", func() {
			var (
				expectedReceipt = receipts.Receipt{
					PurchaseTime: "14:33",
				}
				expectedPoints = 10
			)

			got := pointscalcs.ByPurchaseTwoFourInterval(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})
	})

	Describe("ByNotDuplicatedItems", func() {
		It("returns 5 per items if there no duplicated items", func() {
			var (
				expectedReceipt = receipts.Receipt{
					Items: []receipts.ReceiptItem{
						{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
						{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
						{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
						{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
						{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
					},
				}
				expectedPoints = 25
			)

			got := pointscalcs.ByNotDuplicatedItems(expectedReceipt)

			Expect(got).To(Equal(expectedPoints))
		})

		When("there is a duplicated item", func() {
			It("returns 0 point", func() {
				var (
					expectedReceipt = receipts.Receipt{
						Items: []receipts.ReceiptItem{
							{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
							{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
							{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
							{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
							{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
						},
					}
					expectedPoints = 0
				)

				got := pointscalcs.ByNotDuplicatedItems(expectedReceipt)

				Expect(got).To(Equal(expectedPoints))
			})
		})
	})
})
