package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ACCESS_ORIGIN string
	MYSQL_URL     string
	JWT_SECRET    string
	PORT          string
	CLIENT_ID     string
	CLIENT_SECRET string
	REDIRECT_URL  string
	AUTH_URL      string
	TOKEN_URL     string
	PROFILE_URL   string
	FRONTEND_URL  string
)

func init() {
	godotenv.Load()
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	REDIRECT_URL = os.Getenv("REDIRECT_URL")
	AUTH_URL = "https://access.line.me/oauth2/v2.1/authorize"
	TOKEN_URL = "https://api.line.me/oauth2/v2.1/token"
	PROFILE_URL = "https://api.line.me/v2/profile"
	FRONTEND_URL= os.Getenv("FRONTEND_URL")
}
