package pointscalcs

import (
	"math"
	"strings"
	"time"
	"unicode"

	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
	"github.com/manicar2093/fetch-receipt-processor-challenge/pkg/logger"
)

var log = logger.GetLogger()

func ByRetailerName(r receipts.Receipt) int {
	var points int
	for _, letter := range r.Retailer {
		if unicode.IsLetter(letter) {
			points += 1
		}
	}
	log.WithField("generated_points", points).Debug("ByRetailerName")
	return points
}

func ByRoundedTotal(r receipts.Receipt) int {
	points := 0
	if isInteger(r.Total) {
		points = 50
	}
	log.WithField("generated_points", points).Debug("ByRoundedTotal")
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
	log.WithField("generated_points", points).Debug("ByTotalMultipleOf25")
	return points
}

func ByEach2Items(r receipts.Receipt) int {
	var (
		itemsLen        = float64(len(r.Items)) / 2
		itemsLenFloored = math.Floor(itemsLen)
		points          = int(itemsLenFloored) * 5
	)
	log.WithField("generated_points", points).Debug("ByEach2Items")

	return points
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
	log.WithField("generated_points", points).Debug("ByItemTrimmedDescription")

	return points
}

func ByPurchaseDayIsOdd(r receipts.Receipt) int {
	points := 0
	if r.PurchaseDate.Day()%2 != 0 {
		points = 6
	}
	log.WithField("generated_points", points).Debug("ByPurchaseDayIsOdd")

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
	log.WithField("generated_points", points).Debug("ByPurchaseTwoFourInterval")

	return points
}

func isInteger(x float64) bool {
	return math.Ceil(x) == x
}
