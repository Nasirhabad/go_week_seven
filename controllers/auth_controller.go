package controllers

import (
	"net/http"

	"week7/config"
	"week7/models"
	"week7/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	config.DB.Create(&user)
	return c.JSON(http.StatusCreated, "User registered successfully")
}

func Login(c echo.Context) error {
	var user models.User
	var input models.User
	if err := c.Bind(&input); err != nil {
		return err
	}
	config.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}
	token, _ := utils.GenerateToken(user.Email)
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
