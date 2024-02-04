package controller

import (
	"net/http"
	
	"kodegiri/kodingtest/model"
	
	"github.com/labstack/echo/v4"
)

func Hello_world(c echo.Context) error {
	return c.JSON(http.StatusOK, model.ResponseSuccess{
		StatusCode: 200,
		Message:    "Hello Kodegiri, this is a coding test project",
	})
}
