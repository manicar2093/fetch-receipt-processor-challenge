package receipts_test

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
	"github.com/manicar2093/fetch-receipt-processor-challenge/mocks"
)

var _ = Describe("Default", func() {

	var (
		receiptRepositoryMock *mocks.ReceiptRepo
		pointCalculatorMock1  *mocks.PointCalculator
		pointCalculatorMock2  *mocks.PointCalculator
		service               *receipts.DefaultService
	)

	BeforeEach(func() {
		receiptRepositoryMock = mocks.NewReceiptRepo(T)
		pointCalculatorMock1 = mocks.NewPointCalculator(T)
		pointCalculatorMock2 = mocks.NewPointCalculator(T)
		service = receipts.NewDefaultService(receiptRepositoryMock, pointCalculatorMock1.Execute, pointCalculatorMock2.Execute)
	})

	Describe("Process", func() {
		It("saves a receipt and calculates its generated points", func() {
			var (
				expectedInput = receipts.Receipt{
					Retailer:     "Walgreens",
					PurchaseDate: receipts.Date(time.Date(2022, time.January, 2, 0, 0, 0, 0, time.Local)),
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

	Describe("FindPointsByReceiptId", func() {
		It("returns a receipt generated points by id", func() {
			var (
				expectedReceiptId = uuid.New()
				expectedInput     = receipts.FindPointsByReceiptIdInput{
					ReceiptId: expectedReceiptId,
				}
				expectedGeneratedPoints         = int(faker.Latitude())
				expectedReceiptWithPointsReturn = receipts.ReceiptWithPoints{
					Receipt: receipts.Receipt{
						Id: expectedReceiptId,
					},
					GeneratedPoints: expectedGeneratedPoints,
				}
			)
			receiptRepositoryMock.EXPECT().FindById(expectedReceiptId).Return(&expectedReceiptWithPointsReturn, nil)

			got, err := service.FindPointsByReceiptId(expectedInput)

			Expect(err).ToNot(HaveOccurred())
			Expect(got.Points).To(Equal(expectedGeneratedPoints))
		})
	})
})
