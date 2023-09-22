package receipts

import (
	"github.com/google/uuid"
)

type (
	Receipt struct {
		Id           uuid.UUID
		Retailer     string        `json:"retailer" validate:"required"`
		PurchaseDate Date          `json:"purchaseDate" validate:"required"`
		PurchaseTime string        `json:"purchaseTime" validate:"required"`
		Total        float64       `json:"total,string" validate:"required"`
		Items        []ReceiptItem `json:"items" validate:"min_len:1"`
	}

	ReceiptItem struct {
		ShortDescription string  `json:"shortDescription" validate:"required"`
		Price            float64 `json:"price,string" validate:"required"`
	}

	ReceiptWithPoints struct {
		Receipt
		GeneratedPoints int
	}

	ProcessOutput struct {
		Id uuid.UUID `json:"id"`
	}

	FindPointsByReceiptIdInput struct {
		ReceiptId uuid.UUID `param:"id"`
	}

	FindPointsByReceiptIdOutput struct {
		Points int `json:"points"`
	}
)
