package handler

import (
	"encoding/json"
	"fmt"
	"todos-go-backend/env"
	"todos-go-backend/middleware"
	"todos-go-backend/models"
	sv "todos-go-backend/services"
	"todos-go-backend/utils"

	"net/http"
	"net/url"

	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	
	var request models.RegisterRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, "field can't empty");
		return
	}
	
	resp, svErr := sv.RegisterSv(request)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func LoginHandler(c *gin.Context) {
	
	authParams := url.Values{}
	authParams.Add("response_type", "code")
	authParams.Add("client_id", env.CLIENT_ID)
	authParams.Add("redirect_uri", env.REDIRECT_URL)
	authParams.Add("state", "random_state_value") // Add proper state handling for security
	authParams.Add("disable_auto_login", "true")
	authParams.Add("scope", "profile openid email")

	authURLWithParams := env.AUTH_URL + "?" + authParams.Encode()
	c.Redirect(http.StatusMovedPermanently, authURLWithParams)
}

func LogoutHandler(c *gin.Context) {
	
	c.SetCookie("jwtToken", "", -1, "/", env.FRONTEND_URL, true, true)
	c.Redirect(http.StatusMovedPermanently, env.FRONTEND_URL)
}

func CallbackHandler(c *gin.Context) {
	
	code := c.Request.URL.Query().Get("code")
	tokenParams := url.Values{}
	tokenParams.Add("grant_type", "authorization_code")
	tokenParams.Add("code", code)
	tokenParams.Add("redirect_uri", env.REDIRECT_URL)
	tokenParams.Add("client_id", env.CLIENT_ID)
	tokenParams.Add("client_secret", env.CLIENT_SECRET)

	// Exchange authorization code for access token
	tokenResp, err := http.PostForm(env.TOKEN_URL, tokenParams)
	if err != nil {
		boom.BadRequest(c.Writer, "Failed to exchange code for access token")
		return
	}
	defer tokenResp.Body.Close()

	// Parse the access token from the response
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(tokenResp.Body).Decode(&tokenResponse); err != nil {
		boom.BadRequest(c.Writer, "Failed to exchange code for access token")
		return
	}

	// Use the access token to fetch user's profile
	profileReq, err := http.NewRequest("GET", env.PROFILE_URL, nil)
	if err != nil {
		boom.BadRequest(c.Writer, "Failed to create profile request")
		return
	}
	
	profileReq.Header.Set("Authorization", "Bearer "+ tokenResponse.AccessToken)
	client := &http.Client{}
	profileResp, err := client.Do(profileReq)
	if err != nil {
		boom.BadRequest(c.Writer, "Failed to fetch profile")
		return
	}
	defer profileResp.Body.Close()

	// Parse the user's profile from the response
	var profile map[string]string
	if err := json.NewDecoder(profileResp.Body).Decode(&profile); err != nil {
		boom.BadRequest(c.Writer, "Failed to parse profile response")
		return
	}

	userData := models.User{
		Id: profile["userId"],
		Username: profile["userId"],
		Email: profile["email"],
		LineToken: tokenResponse.AccessToken,
		DisplayName: profile["displayName"],
		PictureUrl: profile["pictureUrl"],
		UserId: profile["userId"],
	}

	jwtToken := middleware.JwtSign(&userData)
	redirectURL := fmt.Sprintf("%s/todos?jwtToken=%s", env.FRONTEND_URL, jwtToken)

	c.SetCookie("jwtToken", jwtToken, 60*60*3, "/", env.FRONTEND_URL, true, true)
	c.Redirect(http.StatusMovedPermanently, redirectURL)
}

func GetUserHandler(c *gin.Context) {
	
	userModel, err := utils.GetUserInfo(c)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
	}
	
	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&userModel);
}

func ChangePasswordHandler(c *gin.Context) {
	
	var request models.ChangePasswordRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, payloadErr.Error());
		return
	}

	userModel, err := utils.GetUserInfo(c)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
	}
	
	resp, svErr := sv.ChangePasswordSv(*userModel, request)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}