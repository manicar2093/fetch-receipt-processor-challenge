package receipts

type (
	PointCalculator func(Receipt) int
	ReceiptRepo     interface {
		Save(*ReceiptWithPoints) error
	}
)
