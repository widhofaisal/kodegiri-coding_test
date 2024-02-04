package controller

import (
	"log"
	"net/http"

	"kodegiri/kodingtest/config"
	"kodegiri/kodingtest/model"
	"kodegiri/kodingtest/utils"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

// func Sign_in(c echo.Context) error {

// }

func Login(c echo.Context) error {
	var user model.User

	// define struct only for binding
	var bindingUser struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// binding struct
	if err_bind := c.Bind(&bindingUser); err_bind != nil {
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// check if request body empty
	if bindingUser.Username == "" || bindingUser.Password == "" {
		log.Print(color.RedString("request body not valid"))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// check username validity
	if err_username := config.DB.Where("username=?", bindingUser.Username).First(&user).Error; err_username != nil {
		log.Print(color.RedString(err_username.Error()))
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  401,
			"message": "unauthorized",
		})
	}

	// if user.Role != "admin" {
	// 	log.Print(color.RedString("only admins are allowed"))
	// 	return c.JSON((http.StatusUnauthorized), map[string]interface{}{
	// 		"status":  401,
	// 		"error":   "unauthorized",
	// 		"message": "only admins are allowed",
	// 	})
	// }

	// verify the password
	if !utils.CheckPasswordHash(bindingUser.Password, user.Password) {
		log.Print(color.RedString("password is incorrect"))
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  401,
			"message": "unauthorized",
		})
	}

	// create token
	token, err_token := utils.Create_token(user.ID, user.Username, user.Role)
	if err_token != nil {
		log.Print(color.RedString(err_token.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   200,
		"message":  "success to login",
		"token":    token,
		"fullname": user.Fullname,
		"username": user.Username,
	})
}
