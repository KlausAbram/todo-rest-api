package handler

import (
	"github.com/gin-gonic/gin"
	todoapi "github.com/klaus-abram/todo-rest-api"
	"net/http"
)

func (hnd *Handler) signUp(ctx *gin.Context) {
	var input todoapi.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	id, err := hnd.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (hnd *Handler) signIn(ctx *gin.Context) {

}
