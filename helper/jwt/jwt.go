package jwt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = "secret"

type UserClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

var UserData UserClaims

func GenerateToken(id int, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	}

	UserData = UserClaims{
		ID:   id,
		Role: role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := parseToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errReponse := errors.New("sign in to proceed")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errReponse
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errReponse
		}
		return []byte(secret), nil
	})
	fmt.Println(token)
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errReponse
	}
	fmt.Println(token.Claims.(jwt.MapClaims))
	return token.Claims.(jwt.MapClaims), nil
}
