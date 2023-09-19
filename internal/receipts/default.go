package receipts

type (
	DefaultService struct {
		receiptRepo      ReceiptRepo
		pointCalculators []PointCalculator
	}
)

func NewDefaultService(receiptRepo ReceiptRepo, pointCalcs ...PointCalculator) *DefaultService {
	return &DefaultService{
		receiptRepo:      receiptRepo,
		pointCalculators: pointCalcs,
	}
}

func (c *DefaultService) Process(input Receipt) (ProcessOutput, error) {
	panic("not implemented")
}
