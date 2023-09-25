package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
	"github.com/manicar2093/fetch-receipt-processor-challenge/cmd/api/controllers"
	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/pointscalcs"
	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
	"github.com/manicar2093/fetch-receipt-processor-challenge/pkg/config"
	"github.com/manicar2093/fetch-receipt-processor-challenge/pkg/logger"
	"github.com/manicar2093/fetch-receipt-processor-challenge/pkg/validator"
	cmap "github.com/orcaman/concurrent-map/v2"

	"github.com/labstack/echo/v4"
)

var (
	port = flag.String("port", "5001", "app port to use")
)

// @title		fetch-receipt-processor-challenge
// @version	0.2.0
func main() {
	var (
		echoInstance    = echo.New()
		baseEndpoint    = ""
		gookitValidator = validator.NewGooKitValidator()
		concurrentMap   = cmap.New[*receipts.ReceiptWithPoints]()
		receiptRepo     = receipts.NewCMapRepo(uuid.New, concurrentMap)
		receiptService  = receipts.NewDefaultService(
			receiptRepo,
			pointscalcs.ByRetailerName,
			pointscalcs.ByRoundedTotal,
			pointscalcs.ByTotalMultipleOf25,
			pointscalcs.ByEach2Items,
			pointscalcs.ByItemTrimmedDescription,
			pointscalcs.ByPurchaseDayIsOdd,
			pointscalcs.ByPurchaseTwoFourInterval,
		)
		receiptController = controllers.NewReceipt(receiptService)
		server            = NewServer(echoInstance, gookitValidator, baseEndpoint, receiptController)
	)

	flag.Parse()

	fmt.Println("Environment:", config.Instance.Environment)
	logger.GetLogger().Fatal(server.Start(fmt.Sprintf(":%s", *port)))
}
