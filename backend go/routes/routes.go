package routes

import (
	"aijob/controllers"
	"aijob/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "SELAMAT DATANG DI AI JOB")
	})
	e.GET("/aijob", controllers.FetchAllDataJob, middleware.IsAuthenticated)
	e.POST("/aijob", controllers.StoreDataJob, middleware.IsAuthenticated)
	e.PUT("/aijob", controllers.UpdateDataJob, middleware.IsAuthenticated)
	e.DELETE("/aijob", controllers.DeleteDataJob, middleware.IsAuthenticated)
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.GET("/gologin", controllers.GoogleCallback)
	e.POST("/ailogin", controllers.CheckLogin)

	return e
}
