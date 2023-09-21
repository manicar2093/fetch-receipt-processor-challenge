package pointscalcs

import (
	"math"
	"strings"
	"time"

	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
	"github.com/manicar2093/fetch-receipt-processor-challenge/pkg/logger"
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

func ByEach2Items(r receipts.Receipt) int {
	var (
		itemsLen        = float64(len(r.Items)) / 2
		itemsLenFloored = math.Floor(itemsLen)
	)

	return int(itemsLenFloored) * 5
}

func ByItemTrimmedDescription(r receipts.Receipt) int {
	var points int

	for _, item := range r.Items {
		trimmedDescription := strings.Trim(item.ShortDescription, " ")
		if len(trimmedDescription)%3 != 0 {
			continue
		}
		points += int(math.Ceil(item.Price * .2))
	}

	return points
}

func ByPurchaseDayIsOdd(r receipts.Receipt) int {
	points := 0
	if r.PurchaseDate.Day()%2 != 0 {
		points = 6
	}
	return points
}

func ByPurchaseTwoFourInterval(receipt receipts.Receipt) int {
	var (
		timeFormat = "15:04"
		initTime   = time.Date(0, time.January, 1, 14, 0, 0, 0, time.UTC)
		finalTime  = time.Date(0, time.January, 1, 16, 0, 0, 0, time.UTC)
		points     = 0
	)
	t, err := time.Parse(timeFormat, receipt.PurchaseTime)
	if err != nil {
		logger.GetLogger().WithField("error", err).Error("returning 0 as default value")
		return points
	}
	if t.After(initTime) && t.Before(finalTime) {
		points = 10
	}
	return points
}

func isInteger(x float64) bool {
	return math.Ceil(x) == x
}
