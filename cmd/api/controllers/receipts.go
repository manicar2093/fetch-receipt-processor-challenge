package controllers

import (
	"net/http"

	"github.com/coditory/go-errors"
	"github.com/labstack/echo/v4"
	"github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
)

type Receipt struct {
	receiptService *receipts.DefaultService
}

func NewReceipt(receiptService *receipts.DefaultService) *Receipt {
	return &Receipt{
		receiptService: receiptService,
	}
}

func (c *Receipt) SetUpRoutes(group *echo.Group) {
	group.POST("/receipts/process", c.ProcessHandler)
	group.GET("/receipts/:id/points", c.FindByIdHandler)
}

// @Summary		Process a receipt
// @Description	Process a receipt calculating generated points
// @Tags			Receipt
// @Accept			json
// @Produce		json
// @Param			receipt_to_process	body		receipts.Receipt		true	"Receipt to be process"
// @Success		201					{object}	receipts.ProcessOutput	"Receipt created data"
// @Failure		500					"Something unidentified has occurred"
// @Router			/receipts/process [post]
func (c *Receipt) ProcessHandler(ctx echo.Context) error {
	var req receipts.Receipt
	if err := ctx.Bind(&req); err != nil {
		return errors.Wrap(err)
	}
	if err := ctx.Validate(&req); err != nil {
		return errors.Wrap(err)
	}

	res, err := c.receiptService.Process(req)
	if err != nil {
		return errors.Wrap(err)
	}
	return ctx.JSON(http.StatusCreated, res)
}

// @Summary		Returns a receipt registered points by receipt id
// @Description	Returns a receipt registered points by receipt id
// @Tags			Receipt
// @Accept			json
// @Produce		json
// @Param			id	path		string									true	"receipt UUID to get"
// @Success		200	{object}	receipts.FindPointsByReceiptIdOutput	"Receipt points data"
// @Failure		404	{object}	apperrors.MessagedError
// @Failure		500	"Something unidentified has occurred"
// @Router			/receipts/{id}/points [get]
func (c *Receipt) FindByIdHandler(ctx echo.Context) error {
	var req receipts.FindPointsByReceiptIdInput
	if err := ctx.Bind(&req); err != nil {
		return errors.Wrap(err)
	}
	if err := ctx.Validate(&req); err != nil {
		return errors.Wrap(err)
	}

	res, err := c.receiptService.FindPointsByReceiptId(req)
	if err != nil {
		return errors.Wrap(err)
	}
	return ctx.JSON(http.StatusOK, res)
}
