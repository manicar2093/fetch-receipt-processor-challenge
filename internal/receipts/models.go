package receipts

import (
	"time"

	"github.com/google/uuid"
)

type (
	Receipt struct {
		Id           uuid.UUID
		Retailer     string
		PurchaseDate time.Time
		PurchaseTime string
		Total        float64
		Items        []ReceiptItem
	}

	ReceiptItem struct {
		ShortDescription string
		Price            float64
	}

	ReceiptWithPoints struct {
		Receipt
		GeneratedPoints int
	}

	ProcessOutput struct {
		Id uuid.UUID
	}
)
