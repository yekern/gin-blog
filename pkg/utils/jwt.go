package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"time"
)

type Jwt struct {
	secret  string
	expired time.Duration
}

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

type Token struct {
	Token  string `json:"token"`
	Header string `json:"header"`
}

func NewJWT() *Jwt {
	return &Jwt{
		secret:  viper.GetString("jwt.secret"),
		expired: viper.GetDuration("jwt.expired"),
	}
}

// CreateToken 创建Token
func (j *Jwt) CreateToken(userId int64) *Token {
	claims := &CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.expired * time.Second).Unix(),
			Issuer:    "System",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		log.Println(err.Error())
		zap.L().Error("token生成失败", zap.Error(err))
	}
	return &Token{
		Token:  tokenString,
		Header: "Bearer",
	}
}

// Decode 解析token
func (j *Jwt) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
