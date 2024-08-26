package router

import (
	"goapi/api/router/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	// Public routes (do not require authentication)
	publicRoutes := r.Group("/v1")
	{
		publicRoutes.POST("/login")
	}

	// Protected routes (require authentication)
	protectedRoutes := r.Group("/v1/internal")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
	}

	return r
}
