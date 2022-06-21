package order

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"learn/kafka/api/v1/order/request"
	"learn/kafka/business/order"
)

type Controller struct {
	service order.Service
}

func NewController(service order.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) CreateNewOrder(ctx echo.Context) error {
	newOrder := new(request.NewOrder)

	if err := ctx.Bind(newOrder); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"code": "400", "message": "invalid body format"})
	}

	if err := c.service.CreateNewOrder(newOrder.ToBusinessNewOrder()); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"code": "400", "message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"code": "201", "message": "success create new order"})
}
