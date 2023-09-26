package handler

import (
	"github.com/deuuus/bmstu-rsoi/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"token", "Origin", "X-Requested-With", "Content-Type", "Accept"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}

	router.Use(cors.New(config))

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
