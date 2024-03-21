package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"server/entity"
	"server/repository"
	"time"
)

const (
	salt = "4444lolkekcheburek4444"
)

type AuthService struct {
	rep repository.AuthRepositoryImpl
}

func NewAuthService(rep repository.AuthRepositoryImpl) *AuthService {
	return &AuthService{rep: rep}
}

func (as *AuthService) SignUp(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.rep.SignUp(user)
}

func (as *AuthService) GenerateToken(username string, password string) (string, int, error) {
	user, err := as.rep.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", 0, errors.New("cant generate token")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("jwtKey")))
	if err != nil {
		return "", 0, err
	}
	return tokenStr, user.Id, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (as *AuthService) ParseToken(tokenString string) (int, string, error) {
	// Парсим по ключу
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("jwtKey")), nil
	})

	if err != nil {
		log.Println("Invalid token: ", err)
		return 0, "", err
	}

	// Проверяем, если токен валиден
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFloat := claims["user_id"].(float64)
		userId := int(userIdFloat)
		userRole := claims["role"].(string)
		return userId, userRole, nil
	} else {
		return -1, "", err
	}

}

func (as *AuthService) CreateAdmin() (int, error) {
	password := os.Getenv("admin_password")
	passwordHash := generatePasswordHash(password)
	adminId, err := as.rep.CreateAdmin("admin", passwordHash)
	if err != nil {
		return -1, err
	}
	return adminId, nil
}
