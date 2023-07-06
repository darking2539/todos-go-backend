package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"todos-go-backend/models"

	"github.com/darahayes/go-boom"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("JWT_SECRET")

func JWTAuthUser(c *gin.Context) {
	
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")

	if token == "" {
		c.Abort()
		err := errors.New("token is empty")
		boom.BadRequest(c.Writer, err.Error())
		return
	}

	jwtToken, err := validateToken(token);
	if  err != nil {
		c.Abort()
		boom.BadRequest(c.Writer, err.Error())
		return
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
	
		userData := models.User{
			Id: fmt.Sprintf("%v", claims["id"]),
			Username: fmt.Sprintf("%v", claims["username"]),
			Email: fmt.Sprintf("%v", claims["email"]),
		}

		userDataBytes, err := json.Marshal(&userData)
		if err != nil {
			c.Abort()
			boom.BadRequest(c.Writer, err.Error());
			return
		}

		c.Request.Header.Set("X-User", string(userDataBytes))
		c.Next()
		
	} else {
		boom.BadRequest(c.Writer, jwtToken.Valid);
		c.Abort()
	}
}

func JWTGetUserMiddleware() gin.HandlerFunc {
	return JWTAuthUser
}

func JwtSign(payload models.User) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.Id
	atClaims["username"] = payload.Username
	atClaims["email"] = payload.Email
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token
}

func validateToken(token string) (*jwt.Token, error) {
	
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	return jwtToken, err
}
