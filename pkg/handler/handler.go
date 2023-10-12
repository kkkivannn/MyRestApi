package handler

import (
	"RestaurantRestApi/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signIn", h.signIn)
			auth.POST("/signUp", h.singUp)
			auth.POST("/refresh", h.refreshToken)
		}

		client := api.Group("/client", h.Identity)
		{
			client.GET("/profile", h.profile)
		}

		admin := api.Group("/admin")
		{
			authEmployee := admin.Group("/employee")
			{
				authEmployee.POST("/signIn", h.signInEmployee)
				authEmployee.POST("/signUp", h.signUpEmployee)
				authEmployee.GET("/profile", h.employeeProfile)
			}
		}

	}
	return router
}
