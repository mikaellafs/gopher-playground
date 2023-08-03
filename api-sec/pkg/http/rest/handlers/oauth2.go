package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"gopher-playground/api-sec/pkg/auth/accesscontrol"
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	api_oauth2 "gopher-playground/api-sec/pkg/auth/oauth2"
	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"
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

func HandleGoogleCallback(repo user.Repository, authMode authmode.AuthMode, ac accesscontrol.AccessControl) func(*gin.Context) {
	return func(c *gin.Context) {
		googleOauthConfig := api_oauth2.GetConfig()

		code := c.Query("code")
		if code == "" {
			c.String(http.StatusBadRequest, "Code not found")
			return
		}

		gtoken, err := googleOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to exchange token: %s", err.Error()))
			return
		}

		// Get user info
		userInfo, err := api_oauth2.GetUserInfo(gtoken.AccessToken)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Create user if it does not exist
		randPw, err := token.GenerateRandTokenString()
		if err != nil {
			c.String(httpStatusFor(err), "failed to create user:"+err.Error())
			return
		}

		_, err = repo.CreateUser(context.TODO(), userInfo.Email, randPw, userInfo.Name)
		if err != nil && err != user.ErrUserAlreadyExists {
			c.String(httpStatusFor(err), err.Error())
			return
		}

		// Add default user role
		ac.AssignRoleToUser(userInfo.Email, accesscontrol.DefaultRole)

		// Generate token for user
		t, err := authMode.GenerateToken(c, userInfo.Email)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if t != nil {
			c.JSON(http.StatusOK, gin.H{
				"token": *t,
			})
		}
	}
}
