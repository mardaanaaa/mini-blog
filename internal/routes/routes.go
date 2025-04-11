package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/db"
	"rest-project/internal/handler"
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

	// Роуты для постов
	router.GET("/posts", postHandler.GetAllPosts)
	router.GET("/posts/:id", postHandler.GetPostByID)
	router.POST("/posts", postHandler.CreatePost)

	// Роуты для комментариев
	router.GET("/posts/:post_id/comments", commentHandler.GetCommentsForPost)
	router.POST("/posts/:post_id/comments", commentHandler.CreateComment)
}
