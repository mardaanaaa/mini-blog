package main

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/db"
	"rest-project/internal/routes"
)

func main() {
	// Инициализация базы данных
	db.ConnectDatabase()

	// Инициализация маршрутов
	r := gin.Default()
	routes.InitRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
