package api

import (
	"learn/kafka/api/v1/welcome"

	echo "github.com/labstack/echo/v4"
)

type Routes struct {
	Welcome *welcome.Controller
}

func (r *Routes) New(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/welcome", r.Welcome.Index)
}
