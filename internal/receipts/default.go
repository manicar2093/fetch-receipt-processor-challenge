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
	var pointsGenerated int
	for _, calc := range c.pointCalculators {
		pointsGenerated += calc(input)
	}
	toSave := ReceiptWithPoints{
		Receipt:         input,
		GeneratedPoints: pointsGenerated,
	}
	if err := c.receiptRepo.Save(&toSave); err != nil {
		return ProcessOutput{}, err
	}

	return ProcessOutput{
		Id: toSave.Id,
	}, nil
}

func (c *DefaultService) FindPointsByReceiptId(input FindPointsByReceiptIdInput) (FindPointsByReceiptIdOutput, error) {
	receiptWPointsFound, err := c.receiptRepo.FindById(input.ReceiptId)
	if err != nil {
		return FindPointsByReceiptIdOutput{}, err
	}

	return FindPointsByReceiptIdOutput{
		Points: receiptWPointsFound.GeneratedPoints,
	}, nil
}
