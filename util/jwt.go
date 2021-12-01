package util

import (
	"fmt"
	"gin/log"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(key string, id string) string {
	claim := jwt.StandardClaims{
		Audience:  "localhost:8050",
		ExpiresAt: time.Now().Unix() + 10*60*1000*1000*1000,
		Id:        id,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "test admin",
		Subject:   "localhost",
		//NotBefore: time.Now().UnixNano(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	log.Info(claim)
	token.Claims = claim
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}
func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	claims := token.Claims
	if token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}
