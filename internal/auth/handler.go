package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-project/internal/db"
	"rest-project/internal/models" // Убедитесь, что модель User импортирована
)

// Login обработчик для логина
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var u models.User // Используем структуру User
	// Ищем пользователя в базе данных
	db.DB.Where("username = ?", req.Username).First(&u)
	if u.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Сравниваем хеш пароля
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login or password incorrect"})
		return
	}

	// Генерация JWT токена
	token, err := GenerateJWT(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register обработчик для регистрации нового пользователя
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Проверка, существует ли уже пользователь с таким именем
	var existing models.User
	if err := db.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Создание нового пользователя
	u := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if err := db.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
