package receipts

import (
	"time"

	"github.com/google/uuid"
)

type (
	Receipt struct {
		Id           uuid.UUID
		Retailer     string        `json:"retailer"`
		PurchaseDate time.Time     `json:"purchaseDate"`
		PurchaseTime string        `json:"purchaseTime"`
		Total        float64       `json:"total"`
		Items        []ReceiptItem `json:"items"`
	}

	ReceiptItem struct {
		ShortDescription string  `json:"shortDescription"`
		Price            float64 `json:"price"`
	}

	ReceiptWithPoints struct {
		Receipt
		GeneratedPoints int
	}

	ProcessOutput struct {
		Id uuid.UUID
	}
)
