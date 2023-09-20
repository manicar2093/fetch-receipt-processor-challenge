package pointscalcs

import (
	"math"

	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
)

func ByRetailerName(r receipts.Receipt) int {
	return len(r.Retailer)
}

func ByRoundedTotal(r receipts.Receipt) int {
	points := 0
	if isInteger(r.Total) {
		points = 50
	}
	return points
}

func ByTotalMultipleOf25(r receipts.Receipt) int {
	var (
		points = 0
		divRes = r.Total / .25
	)
	if isInteger(divRes) {
		points = 25
	}
	return points
}

func isInteger(x float64) bool {
	return math.Ceil(x) == x
}
