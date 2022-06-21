package api

import (
	echo "github.com/labstack/echo/v4"

	"learn/kafka/api/v1/order"
	"learn/kafka/api/v1/welcome"
)

type Routes struct {
	Welcome *welcome.Controller
	Order   *order.Controller
}

func (r *Routes) New(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/welcome", r.Welcome.Index)
	v1.POST("/order", r.Order.CreateNewOrder)
}
