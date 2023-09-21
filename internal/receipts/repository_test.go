package receipts_test

import (
	"net/http"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
	"github.com/manicar2093/fetch-receipt-processor-challenge/mocks"
	"github.com/manicar2093/fetch-receipt-processor-challenge/pkg/apperrors"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var _ = Describe("Repository", Ordered, func() {

	var (
		uuidGenMock           *mocks.UUIDGen
		expectedGeneratedUUID uuid.UUID
		concurrentMap         cmap.ConcurrentMap[string, *receipts.ReceiptWithPoints]
		expectedReceiptSaved  receipts.ReceiptWithPoints
		repo                  *receipts.CMapRepo
	)

	BeforeAll(func() {
		uuidGenMock = mocks.NewUUIDGen(T)
		expectedGeneratedUUID = uuid.New()
		concurrentMap = cmap.New[*receipts.ReceiptWithPoints]()
		expectedReceiptSaved = receipts.ReceiptWithPoints{
			Receipt: receipts.Receipt{
				Retailer:     faker.Name(),
				PurchaseDate: time.Date(2023, time.September, 21, 0, 0, 0, 0, time.UTC),
				PurchaseTime: "14:22",
				Total:        faker.Latitude(),
				Items: []receipts.ReceiptItem{
					{ShortDescription: faker.Paragraph(), Price: faker.Latitude()},
					{ShortDescription: faker.Paragraph(), Price: faker.Latitude()},
				},
			},
			GeneratedPoints: int(faker.Latitude()),
		}
		repo = receipts.NewCMapRepo(uuidGenMock.Execute, concurrentMap)
	})

	Describe("Save", func() {
		It("stores a new receipt with points and generate a uuid to identify it", func() {
			uuidGenMock.EXPECT().Execute().Return(expectedGeneratedUUID)

			Expect(repo.Save(&expectedReceiptSaved)).To(Succeed())
			Expect(expectedReceiptSaved.Id).To(Equal(expectedGeneratedUUID))
		})
	})

	Describe("FindById", func() {
		It("return a receipt by its id", func() {
			got, err := repo.FindById(expectedGeneratedUUID)

			Expect(err).ToNot(HaveOccurred())
			Expect(got).To(Equal(&expectedReceiptSaved))
		})

		When("is not found", func() {
			It("return a ErrReceiptNotFound", func() {
				var (
					expectedUuid = uuid.New()
				)

				got, err := repo.FindById(expectedUuid)

				Expect(got).To(BeNil())
				asHandleableErr := err.(apperrors.HandleableError)
				Expect(asHandleableErr.Error()).To(Equal("receipt was not found with given id"))
				Expect(asHandleableErr.StatusCode()).To(Equal(http.StatusNotFound))
			})
		})
	})
})
