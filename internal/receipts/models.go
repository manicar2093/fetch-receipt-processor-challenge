package receipts

import (
	"github.com/google/uuid"
)

type (
	Receipt struct {
		Id           uuid.UUID
		Retailer     string        `json:"retailer"`
		PurchaseDate Date          `json:"purchaseDate"`
		PurchaseTime string        `json:"purchaseTime"`
		Total        float64       `json:"total,string"`
		Items        []ReceiptItem `json:"items"`
	}

	ReceiptItem struct {
		ShortDescription string  `json:"shortDescription"`
		Price            float64 `json:"price,string"`
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
