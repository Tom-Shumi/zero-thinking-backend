package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HealthController controller for health request
type ThinkingTreeController struct{}

// NewHealthController is constructer for HealthController
func NewThinkingTreeController() *ThinkingTreeController {
	return new(ThinkingTreeController)
}

// Index is index route for health
func (ttc *ThinkingTreeController) List(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}

// Index is index route for health
func (ttc *ThinkingTreeController) Save(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}
