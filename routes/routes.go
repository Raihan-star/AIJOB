package routes

import (
	"aijob/controllers"
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/aijobmeraih", controllers.FetchAllDataJobMeraih /*, middleware.IsAuthenticated*/) // GET JOB DATA MERAIH

	e.POST("/aijobmeraih", controllers.StoreDataJobMeraih /*, middleware.IsAuthenticated*/) // POST DATA JOB MERAIH
	//e.POST("/aijobmeraih", controllers.StoreDataJobMeraihQ /*, middleware.IsAuthenticated*/)
	e.PUT("/aijobmeraih", controllers.UpdateDataJobMeraih /*, middleware.IsAuthenticated*/)    // UPDATE DATA JOB MERAIH
	e.DELETE("/aijobmeraih", controllers.DeleteDataJobMeraih /*, middleware.IsAuthenticated*/) // DELETE DATA JIOB MERAIH
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)                        // KONVERSI PASSWORD -> HASH
	e.POST("/login", controllers.CheckLogin)                                                   // LOGIN

	e.POST("/visitorlogin", controllers.LoginDataVisitor)
	e.GET("/visitordata", controllers.FetchAllDataVisitor)

	e.GET("/auth/google", echo.WrapHandler(http.HandlerFunc(controllers.HandleGoogleLogin)))
	e.GET("/auth/google/callback", echo.WrapHandler(http.HandlerFunc(controllers.HandleGoogleCallback)))
	e.GET("/", echo.WrapHandler(http.HandlerFunc(controllers.HandleMain)))
	return e
}
