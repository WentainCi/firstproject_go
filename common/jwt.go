package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/xiaotian/synk/model"
)

// jwt加密秘钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 发放token的方法
func ReleaseToken(user model.User) (string, error) {
	//过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	Claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "oceanlearn.tech",     //谁发放的这个token
			Subject:   "user token",          //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
