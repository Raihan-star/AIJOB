package controllers

import (
	"aijob/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func FetchAllDataJob(c echo.Context) error {
	result, err := models.FetchAllDataJob()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreDataJob(c echo.Context) error {
	code_Job := c.FormValue("code_job")
	nama_Job := c.FormValue("nama_job")
	keterangan := c.FormValue("keterangan")
	bukti := c.FormValue("bukti")
	status := c.FormValue("status")
	tanggal := c.FormValue("tanggal")

	result, err := models.StoreDataJob(code_Job, nama_Job, keterangan, bukti, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateDataJob(c echo.Context) error {
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

	result, err := models.UpdateDataJob(conv_id, code_Job, nama_Job, keterangan, bukti, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteDataJob(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteDataJob(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
