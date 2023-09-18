package handler

import (
	"github.com/gin-gonic/gin"
	"todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetAllList)
			lists.GET("/:id", h.GetListById)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.GetAllItem)
				items.GET("/:item_id", h.GetItemById)
				items.PUT("/:item_id", h.UpdateItemById)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}
	return router
}
