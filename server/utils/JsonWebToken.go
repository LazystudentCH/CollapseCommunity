package utils

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"server/models"
	"time"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	ID uint `json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Phone string `json:"phone"`
	Gender string `json:"gender"`
	IsAdmin bool `json:"is_admin"`
}


func CreateToken(user *models.User) (tokenString string, err error) {
	claims := &JwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * time.Duration(Config.JwtExpireTime)).Unix()),
			Issuer:    Config.Issuer,
		},
		user.ID,
		user.Name,
		user.Avatar,
		user.Phone,
		user.Gender,
		user.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(Config.JwtSecret))
	return
}

func ParseToken(r *http.Request) (*models.User,bool) {
	var token *jwt.Token
	tokenStr := r.Header.Get("token")
	token, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return []byte(Config.JwtSecret), nil
	})
	if err != nil {
		return nil,false
	}

	user := &models.User{}
	data, _ :=json.Marshal(token.Claims)
	json.Unmarshal(data,user)
	return user,token.Claims.Valid() == nil
}
