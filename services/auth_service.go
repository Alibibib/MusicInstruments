package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtKey = []byte("super_secret_key") // ❗замени в .env в будущем

type Claims struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}

func Register(user models.User) error {
	db := database.GetDB()

	// Проверка на существующего
	var existing models.User
	if err := db.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return fmt.Errorf("пользователь с таким email уже существует")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return db.Create(&user).Error
}

func Login(email, password string) (string, error) {
	db := database.GetDB()

	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", fmt.Errorf("неверный email или пароль")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("неверный email или пароль")
	}

	// Генерируем токен
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("недействительный токен")
	}
	return claims, nil
}
