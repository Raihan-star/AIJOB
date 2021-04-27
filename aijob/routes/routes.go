package routes

import (
	"aijob/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/cors"
)

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	},
)

func Init() *echo.Echo {
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/aijoball", controllers.FetchAllDataJobAll /*, middleware.IsAuthenticated*/)       // GET JOB DATA ALL
	e.GET("/aijobmeraih", controllers.FetchAllDataJobMeraih /*, middleware.IsAuthenticated*/) // GET JOB DATA MERAIH

	e.POST("/aijobmeraih", controllers.StoreDataJobMeraih /*, middleware.IsAuthenticated*/)    // POST DATA JOB MERAIH
	e.PUT("/aijobmeraih", controllers.UpdateDataJobMeraih /*, middleware.IsAuthenticated*/)    // UPDATE DATA JOB MERAIH
	e.DELETE("/aijobmeraih", controllers.DeleteDataJobMeraih /*, middleware.IsAuthenticated*/) // DELETE DATA JIOB MERAIH
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)                        // KONVERSI PASSWORD -> HASH
	e.POST("/login", controllers.CheckLogin)                                                   // LOGIN

	e.POST("/visitorlogin", controllers.LoginDataVisitor)
	e.GET("/visitordata", controllers.FetchAllDataVisitor)

	e.GET("/", echo.WrapHandler(http.HandlerFunc(controllers.HandleGoogleMain)))
	e.GET("/gologin", echo.WrapHandler(http.HandlerFunc(controllers.HandleGoogleLogin)))
	e.GET("/callback", echo.WrapHandler(http.HandlerFunc(controllers.HandleGoogleCallback)))
	return e
}
