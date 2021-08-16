package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func (hnd *Handler) createList(ctx *gin.Context) {
	///test functional
	id, _ := ctx.Get(userCtx)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":       id,
		"sign_key": os.Getenv("sign_key"),
	})
	///delete
}

func (hnd *Handler) getAllLists(ctx *gin.Context) {

}

func (hnd *Handler) getListById(ctx *gin.Context) {

}

func (hnd *Handler) updateList(ctx *gin.Context) {

}

func (hnd *Handler) deleteList(ctx *gin.Context) {

}
