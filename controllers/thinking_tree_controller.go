package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"zero-thinking-backend/database"
	"zero-thinking-backend/models"

	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ThinkingTreeController struct{}

func NewThinkingTreeController() *ThinkingTreeController {
	return new(ThinkingTreeController)
}

func (ttc *ThinkingTreeController) List(c echo.Context) error {

	token := c.Get("token").(string)

	// TODO serviceへ移動
	db := database.GetDB()
	fields := []string{"id", "title", "insert_date"}
	query := []qm.QueryMod{
		qm.Where("user_id=?", token),
		qm.OrderBy(`id DESC`),
		qm.Select(fields...),
	}
	counts, errCount := models.ThinkingTrees(query...).Count(context.Background(), db)

	var thinkingTreeList []*models.ThinkingTree
	if errCount != nil {
		return errCount
	}

	if counts == 0 {
		thinkingTreeList = make([]*models.ThinkingTree, 0)
	} else {
		var errAll error
		thinkingTreeList, errAll = models.ThinkingTrees(query...).All(context.Background(), db)
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

func (ttc *ThinkingTreeController) Detail(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	// TODO serviceへ移動
	db := database.GetDB()
	fields := []string{"id", "title", "thinking_tree", "insert_date"}
	query := []qm.QueryMod{
		qm.Where("id=?", id),
		qm.Select(fields...),
	}
	counts, errCount := models.ThinkingTrees(query...).Count(context.Background(), db)

	var thinkingTree *models.ThinkingTree
	if errCount != nil {
		return errCount
	}

	if counts == 0 {
		thinkingTree = nil
	} else {
		var errOne error
		thinkingTree, errOne = models.ThinkingTrees(query...).One(context.Background(), db)
		if errOne != nil {
			return errOne
		}
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		thinkingTree,
	))
}

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
