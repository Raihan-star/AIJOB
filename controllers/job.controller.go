package controllers

import (
	"aijob/models"
	"net/http"

	"github.com/labstack/echo"
)

// GET DATA JOB MERAIH
func FetchAllDataJobMeraih(c echo.Context) error {
	result, err := models.FetchAllDataJobMeraih()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// PUT DATA JOB MERAIH
func StoreDataJobMeraih(c echo.Context) error {
	var request models.DataJob
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	result, err := models.StoreDataJobMeraih(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// UPDATE DATA JOB MERAIH
func UpdateDataJobMeraih(c echo.Context) error {
	var request models.DataJob
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	result, err := models.UpdateDataJobMeraih(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// DELETE DATA JOB MERAIH
func DeleteDataJobMeraih(c echo.Context) error {
	var request models.DataJob
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	result, err := models.DeleteDataJobMeraih(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
