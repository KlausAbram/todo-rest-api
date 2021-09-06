package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/klaus-abram/todo-rest-api/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (hnd *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", hnd.signIn)
		auth.POST("/sign-up", hnd.signUp)
	}

	api := router.Group("/api", hnd.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", hnd.createList)
			lists.GET("/", hnd.getAllLists)
			lists.GET("/:id", hnd.getListById)
			lists.PUT("/:id", hnd.updateList)
			lists.DELETE("/:id", hnd.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", hnd.createItem)
				items.GET("/", hnd.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", hnd.getItemById)
			items.PUT("/:id", hnd.updateItem)
			items.DELETE("/:id", hnd.deleteItem)
		}
	}

	return router
}
