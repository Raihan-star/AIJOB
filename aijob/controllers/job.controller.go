package controllers

import (
	"aijob/models"
	"net/http"
	"strconv"

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
	namajob := c.FormValue("namajob")
	keterangan := c.FormValue("keterangan")
	bukti := c.FormValue("bukti")
	status := c.FormValue("status")
	tanggal := c.FormValue("tanggal")

	result, err := models.StoreDataJobMeraih(namajob, keterangan, bukti, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
func StoreDataJobMeraihQ(c echo.Context) error {
	namajob := c.QueryParam("namajob")
	keterangan := c.QueryParam("keterangan")
	bukti := c.QueryParam("bukti")
	status := c.QueryParam("status")
	tanggal := c.QueryParam("tanggal")

	result, err := models.StoreDataJobMeraihQ(namajob, keterangan, bukti, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// UPDATE DATA JOB MERAIH
func UpdateDataJobMeraih(c echo.Context) error {
	id := c.FormValue("id")
	namajob := c.FormValue("namajob")
	keterangan := c.FormValue("keterangan")
	bukti := c.FormValue("bukti")
	status := c.FormValue("status")
	tanggal := c.FormValue("tanggal")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateDataJobMeraih(conv_id, namajob, keterangan, bukti, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// DELETE DATA JOB MERAIH
func DeleteDataJobMeraih(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteDataJobMeraih(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
