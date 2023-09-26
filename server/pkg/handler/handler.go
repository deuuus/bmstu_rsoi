package handler

import (
	"github.com/deuuus/bmstu-rsoi/pkg/service"
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

	api := router.Group("/api/v1")
	{
		persons := api.Group("/persons")
		{
			persons.GET("/:id", h.getPerson)
			persons.GET("", h.getAllPersons)
			persons.POST("", h.createPerson)
			persons.PATCH("/:id", h.updatePerson)
			persons.DELETE("/:id", h.deletePerson)

		}
	}

	return router
}
