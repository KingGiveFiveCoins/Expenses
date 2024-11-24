package utils

import "github.com/golang-jwt/jwt"

var jwtKey []byte

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func InitUtils(jwt []byte) {
	jwtKey = jwt
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
