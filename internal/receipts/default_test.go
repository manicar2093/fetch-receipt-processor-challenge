package receipts_test

import (
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"template/internal/receipts"
)

var _ = Describe("Default", func() {

	var (
		receiptRepositoryMock *mocks.ReceiptsRepositoryMock
		pointCalculatorMock1  *mocks.PointCaluculator
		pointCalculatorMock2  *mocks.PointCaluculator
		service               *receipts.DefaultService
	)

	BeforeEach(func() {
		receiptRepositoryMock = mocks.NewReceiptsRepositoryMock(T)
		pointCalculatorMock1 = mocks.NewPointCalculator(T)
		pointCalculatorMock2 = mocks.NewPointCalculator(T)
		service = receipts.DefaultService(receiptRepositoryMock, pointCalculatorMock1.Execute, pointCalculatorMock2.Execute)
	})

	Describe("Process", func() {
		It("saves a receipt and calculates its generated points", func() {
			var (
				expectedInput = receipts.Receipt{
					Retailer:     "Walgreens",
					PurchaseDate: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.Local),
					PurchaseTime: "08:13",
					Total:        2.65,
					Items: []receipts.ReceiptItem{
						{
							ShortDescription: "Pepsi - 12-oz",
							Price:            1.25,
						},
						{
							ShortDescription: "Dasani",
							Price:            1.40,
						},
					},
				}
				expectedPointsCalc1Return = 5
				expectedPointsCalc2Return = 5
				expectedGeneratedPoints   = expectedPointsCalc1Return + expectedPointsCalc2Return
				expectedReceiptToSave     = receipts.ReceiptWithPoints{
					Receipt:         expectedInput,
					GeneratedPoints: expectedGeneratedPoints,
				}
				expectedReceiptId = uuid.New()
			)
			pointCalculatorMock1.EXPECT().Execute(expectedInput).Return(expectedPointsCalc1Return)
			pointCalculatorMock2.EXPECT().Execute(expectedInput).Return(expectedPointsCalc2Return)
			receiptRepositoryMock.EXPECT().Save(&expectedReceiptToSave).RunAndReturn(func(r *receipts.ReceiptWithPoints) error {
				r.Id = expectedReceiptId
				return nil
			})

			got, err := service.Process(expectedInput)

			Expect(err).ToNot(HaveOccurred())
			Expect(got.Id).To(Equal(expectedReceiptId))
		})
	})

})
