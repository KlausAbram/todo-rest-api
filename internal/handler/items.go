package handler

import (
	"github.com/gin-gonic/gin"
	todoapi "github.com/klaus-abram/todo-rest-api"
	"net/http"
	"strconv"
)

func (hnd *Handler) createItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input todoapi.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	itemId, err := hnd.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"ItemId": itemId,
	})
}

func (hnd *Handler) getAllItems(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	items, err := hnd.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (hnd *Handler) getItemById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid item id")
		return
	}

	item, err := hnd.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (hnd *Handler) updateItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "error with item id")
		return
	}

	var input todoapi.UpdateItemInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = hnd.services.TodoItem.Update(userId, itemId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (hnd *Handler) deleteItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid item id")
		return
	}

	if err = hnd.services.TodoItem.Delete(userId, itemId); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
