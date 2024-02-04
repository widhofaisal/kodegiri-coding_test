package controller

import (
	"net/http"

	"kodegiri/kodingtest/config"
	"kodegiri/kodingtest/model"

	"github.com/labstack/echo/v4"
)

func Get_all_user(c echo.Context) error {
	var users []model.User
	if errFind := config.DB.Find(&users).Error; errFind != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseError{
			StatusCode: 500,
			Message:    "Failed to get all user data",
			Error:      errFind.Error(),
		})
	}

	// RESULT :
	type result struct {
		model.ResponseSuccess
		Data []model.User
	}
	res := result{
		model.ResponseSuccess{
			StatusCode: 200,
			Message:    "Success to get all user data",
		},
		users,
	}

	return c.JSON(http.StatusOK, res)
}
