package controllers

import (
	"context"
	"net/http"
	"time"

	"zero-thinking-backend/database"
	"zero-thinking-backend/models"

	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// HealthController controller for health request
type ThinkingTreeController struct{}

// NewHealthController is constructer for HealthController
func NewThinkingTreeController() *ThinkingTreeController {
	return new(ThinkingTreeController)
}

// Index is index route for health
func (ttc *ThinkingTreeController) List(c echo.Context) error {

	// TODO serviceへ移動
	db := database.GetDB()
	counts, errCount := models.ThinkingTrees().Count(context.Background(), db)

	var thinkingTreeList []*models.ThinkingTree
	if errCount != nil {
		return errCount
	}

	if counts == 0 {
		thinkingTreeList = make([]*models.ThinkingTree, 0)
	} else {
		var errAll error
		thinkingTreeList, errAll = models.ThinkingTrees().All(context.Background(), db)
		if errAll != nil {
			return errAll
		}
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		thinkingTreeList,
	))
}

// Index is index route for health
func (ttc *ThinkingTreeController) Save(c echo.Context) error {

	// TODO serviceへ移動
	db := database.GetDB()
	thinkingTree := models.ThinkingTree{
		ID:           1,
		UserID:       "1",
		Title:        "test-title",
		ThinkingTree: "test-thinking-tree",
		InsertDate:   time.Now(),
	}
	thinkingTree.Insert(context.Background(), db, boil.Infer())

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}
