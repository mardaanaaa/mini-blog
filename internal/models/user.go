package models

// User модель для хранения данных пользователя
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}
