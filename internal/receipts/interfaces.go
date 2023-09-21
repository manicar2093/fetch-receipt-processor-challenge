package receipts

import "github.com/google/uuid"

type (
	PointCalculator func(Receipt) int
	ReceiptRepo     interface {
		Save(*ReceiptWithPoints) error
		FindById(uuid.UUID) (*ReceiptWithPoints, error)
	}
)
