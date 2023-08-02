package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	api_oauth2 "gopher-playground/api-sec/pkg/auth/oauth2"
	"gopher-playground/api-sec/pkg/env"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func HandleHomePage(c *gin.Context) {
	data := struct {
		GoogleClientID string
	}{
		GoogleClientID: os.Getenv(env.GOOGLE_CLIENT_ID),
	}

	c.HTML(http.StatusOK, "home.html", data)
}

func HandleGoogleLogin(c *gin.Context) {
	googleOauthConfig := api_oauth2.GetConfig()

	url := googleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c *gin.Context) {
	googleOauthConfig := api_oauth2.GetConfig()

	code := c.Query("code")
	if code == "" {
		c.String(http.StatusBadRequest, "Code not found")
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to exchange token: %s", err.Error()))
		return
	}

	// Get user info
	user, err := api_oauth2.GetUserInfo(token.AccessToken)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Create user if it does not exist

	// Store the token and proceed with further actions (e.g., accessing resources).
	c.String(http.StatusOK, "Successfully logged in with Google! User: %v", user)
}
