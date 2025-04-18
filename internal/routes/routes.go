package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/db"
	"rest-project/internal/handler"
	"rest-project/internal/middleware"
	"rest-project/internal/repository"
	"rest-project/internal/services"
)

func InitRoutes(router *gin.Engine) {
	// Подключение к базе данных
	database := db.GetDB()

	// Репозитории
	postRepo := repository.NewPostRepository(database)
	commentRepo := repository.NewCommentRepository(database)

	// Сервисы
	postService := services.NewPostService(postRepo)
	commentService := services.NewCommentService(commentRepo)

	// Обработчики
	postHandler := handler.NewPostHandler(postService)
	commentHandler := handler.NewCommentHandler(commentService)

	// Public Routes (например, для получения постов)
	router.GET("/posts", postHandler.GetAllPosts)
	router.GET("/posts/:id", postHandler.GetPostByID)
	router.GET("/comments/post/:post_id", commentHandler.GetCommentsForPost)

	// Routes for authenticated users (роль: user или admin)
	authGroup := router.Group("/api/user")
	authGroup.Use(middleware.RoleMiddleware("user", "admin"))
	{
		authGroup.POST("/posts", postHandler.CreatePost)
		authGroup.POST("/comments/post/:post_id", commentHandler.CreateComment)
		// другие пользовательские маршруты
	}

	// Routes for admins only
	adminGroup := router.Group("/api/admin")
	adminGroup.Use(middleware.RoleMiddleware("admin"))
	{
		adminGroup.DELETE("/posts/:id", postHandler.DeletePost) // допустим ты хочешь добавить удаление постов
		// другие маршруты для админов
	}
}
