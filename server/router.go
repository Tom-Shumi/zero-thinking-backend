package server

import (
	"net/http"

	"zero-thinking-backend/config"
	"zero-thinking-backend/controllers"
	middlewareAuth "zero-thinking-backend/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: c.GetStringSlice("server.cors"),
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	version := router.Group("/" + c.GetString("server.version"))

	healthController := controllers.NewHealthController()
	version.GET("/health", healthController.Index)

	thinkingTreeController := controllers.NewThinkingTreeController()
	version.GET("/thinkingTree", thinkingTreeController.List, middlewareAuth.Auth())
	version.GET("/thinkingTree/:id", thinkingTreeController.Detail, middlewareAuth.Auth())
	version.POST("/thinkingTree", thinkingTreeController.Save, middlewareAuth.Auth())

	return router, nil
}
