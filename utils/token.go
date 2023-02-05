package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Similadayo/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/jinzhu/gorm"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type PasswordResetClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(user_id int, username string, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UserID:   user_id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AddToBlacklist function for adding token to the blacklist
func AddToBlacklist(db *gorm.DB, token string, expiresAt time.Time) error {
	blacklistedToken := &models.BlacklistToken{
		Token:     token,
		ExpiresAt: expiresAt,
	}
	if err := db.Create(blacklistedToken).Error; err != nil {
		return err
	}
	return nil
}

// IsTokenBlacklisted function for checking if the token is blacklisted
func IsTokenBlacklisted(db *gorm.DB, token string) (bool, error) {
	var blacklistedToken models.BlacklistToken
	if err := db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&blacklistedToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	t, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// GeneratePasswordResetToken generates a new JWT token that can be used to reset a user's password
func GeneratePasswordResetToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24) // token expires in 24 hours
	claims := &PasswordResetClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func VerifyPasswordResetToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	t, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
