package route

import (
	"kodegiri/kodingtest/constant"
	"kodegiri/kodingtest/controller"
	"kodegiri/kodingtest/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(mid.CORS())
	e.Use(middleware.MiddlewareLogging)

	// PUBLIC ENDPOINT
	e.GET("/", controller.Hello_world)
	e.POST("/login", controller.Login)

	// PRIVATE ENDPOINT
	eAuth := e.Group("/auth")
	eAuth.Use(JWTMiddleware())
	eAuth.GET("/users", controller.Get_all_user)

	return e
}

func JWTMiddleware() echo.MiddlewareFunc {
	config := mid.JWTConfig{
		SigningKey: []byte(constant.SECRET_JWT),
	}
	return mid.JWTWithConfig(config)
}
