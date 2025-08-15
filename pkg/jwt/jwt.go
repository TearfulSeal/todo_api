package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte
var jwtExpiration time.Duration

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func Init(secret string, expiration time.Duration){
	jwtSecret = []byte(secret)
	jwtExpiration = expiration 
}

func GenerateToken(userID uint) (string, error){
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (uint, error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func (token *jwt.Token) (any,error){
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil 
	}
	return 0, errors.New("invalid token")
}
