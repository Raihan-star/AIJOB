package controllers

import (
	"aijob/models"
	"net/http"

	"github.com/labstack/echo"
)

func FetchAllDataVisitor(c echo.Context) error {
	result, err := models.FetchAllDataVisitor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func LoginDataVisitor(c echo.Context) error {
	username_visitor := c.FormValue("username_visitor")
	institut_visitor := c.FormValue("institut_visitor")
	email_visitor := c.FormValue("email_visitor")

	result, err := models.LoginDataVisitor(username_visitor, institut_visitor, email_visitor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
