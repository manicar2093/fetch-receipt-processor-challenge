package receipts

import (
	"net/http"

	"github.com/manicar2093/fetch-receipt-processor-challenge/pkg/apperrors"
)

var (
	ErrReceiptNotFound = apperrors.NewMessagedError("receipt was not found with given id", http.StatusNotFound)
)
