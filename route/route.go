package route

import (
	"kodegiri/kodingtest/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", controller.Hello_world)
	e.GET("/users", controller.Get_all_user)

	return e
}
