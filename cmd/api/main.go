package main

import (
	"flag"
	"fmt"
	"template/cmd/api/controllers"
	"template/pkg/config"
	"template/pkg/logger"
	"template/pkg/validator"

	"github.com/labstack/echo/v4"
)

var (
	port = flag.String("port", "5001", "app port to use")
)

//	@title		fetch-receipt-processor-challenge
//	@version	0.0.0
func main() {
	var (
		echoInstance      = echo.New()
		baseEndpoint      = "/api/v1"
		gookitValidator   = validator.NewGooKitValidator()
		initialController = controllers.NewInitial()
		server            = NewServer(echoInstance, gookitValidator, baseEndpoint, initialController)
	)

	flag.Parse()

	fmt.Println("Environment:", config.Instance.Environment)
	logger.GetLogger().Fatal(server.Start(fmt.Sprintf(":%s", *port)))
}
