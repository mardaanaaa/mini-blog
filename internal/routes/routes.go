package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/auth"
)

func SetupRoutes(r *gin.Engine) {

	// Роуты
	students := r.Group("api/v1/auth")
	{
		students.POST("/login", auth.Login)
		students.POST("/register", auth.Register)
	}
}
