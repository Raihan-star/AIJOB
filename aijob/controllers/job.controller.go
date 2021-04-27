package controllers

import (
	"aijob/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GET DATA JOB ALL
func FetchAllDataJobAll(c echo.Context) error {
	result, err := models.FetchAllDataJobAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

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
	code_Job := c.FormValue("code_job")
	nama_Job := c.FormValue("nama_job")
	keterangan := c.FormValue("keterangan")
	bukti := c.FormValue("bukti")
	status := c.FormValue("status")
	tanggal := c.FormValue("tanggal")

	result, err := models.StoreDataJobMeraih(code_Job, nama_Job, keterangan, bukti, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// UPDATE DATA JOB MERAIH
func UpdateDataJobMeraih(c echo.Context) error {
	id := c.FormValue("id")
	code_Job := c.FormValue("code_job")
	nama_Job := c.FormValue("nama_job")
	keterangan := c.FormValue("keterangan")
	bukti := c.FormValue("bukti")
	status := c.FormValue("status")
	tanggal := c.FormValue("tanggal")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateDataJobMeraih(conv_id, code_Job, nama_Job, keterangan, bukti, status, tanggal)
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
