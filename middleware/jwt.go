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
	
	tokenCookie, _ := c.Cookie("jwtToken")

	s := c.Request.Header.Get("Authorization")
	tokenHeader := strings.TrimPrefix(s, "Bearer ")

	token := tokenCookie
	if tokenHeader != "" && tokenCookie == ""{
		token = tokenHeader
	}
	
	if token == "" {
		c.Abort()
		err := errors.New("token is empty")
		boom.Unathorized(c.Writer, err.Error())
		return
	}

	jwtToken, err := validateToken(token);
	if  err != nil {
		c.Abort()
		boom.Unathorized(c.Writer, err.Error())
		return
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
	
		userData := models.User{
			Id: fmt.Sprintf("%v", claims["id"]),
			Username: fmt.Sprintf("%v", claims["username"]),
			Email: fmt.Sprintf("%v", claims["email"]),
			LineToken: fmt.Sprintf("%v", claims["lineToken"]),
			DisplayName: fmt.Sprintf("%v", claims["displayName"]),
			PictureUrl: fmt.Sprintf("%v", claims["pictureUrl"]),
			UserId: fmt.Sprintf("%v", claims["userId"]),
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
		boom.Unathorized(c.Writer, jwtToken.Valid);
		c.Abort()
	}
}

func JWTGetUserMiddleware() gin.HandlerFunc {
	return JWTAuthUser
}


func JwtSign(payload *models.User) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.Id
	atClaims["username"] = payload.Username
	atClaims["email"] = payload.Email
	atClaims["lineToken"] = payload.LineToken
	atClaims["displayName"] = payload.DisplayName
	atClaims["pictureUrl"] = payload.PictureUrl
	atClaims["userId"] = payload.UserId
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
