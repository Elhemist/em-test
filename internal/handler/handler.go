package handler

import (
	"em-test/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/person")
	{
		user.POST("/", h.AddPerson)
		user.DELETE("/", h.DeletePerson)
		user.GET("/", h.CheckPerson)
		user.PATCH("/part", h.EditPerson)
	}

	return router
}
