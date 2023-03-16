package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("eeEKICQI+I++I+I++++__232lskdjbruurARRARAR")

func CreateToken(email string) string {

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ParseToken(tokenStr string) (*Claims, bool) {
	var claims *Claims
	var ok bool
	token, _ := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok = token.Claims.(*Claims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
