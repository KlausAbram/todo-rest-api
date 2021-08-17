package handler

import (
	"github.com/gin-gonic/gin"
	todoapi "github.com/klaus-abram/todo-rest-api"
	"net/http"
	"strconv"
)

func (hnd *Handler) createList(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	var input todoapi.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	listId, err := hnd.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"ListId": listId,
	})
}

func (hnd *Handler) getAllLists(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	lists, err := hnd.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (hnd *Handler) getListById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	list, err := hnd.services.TodoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (hnd *Handler) updateList(ctx *gin.Context) {

}

func (hnd *Handler) deleteList(ctx *gin.Context) {

}
