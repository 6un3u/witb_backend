package router

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/6un3u/witb_backend/docs"
	"github.com/6un3u/witb_backend/handlers"
	"github.com/6un3u/witb_backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Kyobo BookStock Swagger API
// @version 1.0
// @host localhost:4000
// @BasePath /api
func Router() *echo.Echo {
	e := echo.New()

	// Set Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} ${status} ${method} ${uri}" + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))
	// Set access control for Dev
	devOnlyMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if e.Debug {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden, "Access forbidden")
		}
	}

	e.Validator = utils.NewValidator()

	e.GET("/healthy", func(c echo.Context) error {
		currentTime := time.Now()
		return c.String(http.StatusOK, fmt.Sprintf("%s", currentTime.Local().Format("2006-01-02 15:04:05")))
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler, devOnlyMiddleware)

	api := e.Group("/api")
	{
		api.POST("/search", handlers.SearchHandler)
		api.POST("/stock", handlers.StockHandler)
	}

	return e
}
