package handlers

import (
	"context"
	"net/http"

	"gopher-playground/api-sec/pkg/auth"
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"

	"github.com/gin-gonic/gin"
)

func Login(userRepo user.Repository, authMode authmode.AuthMode, tstore token.TokenStore) func(*gin.Context) {
	return func(c *gin.Context) {
		// Parse 'application/x-www-form-urlencoded'
		err := c.Request.ParseForm()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Get username and password
		username := c.Request.Form.Get("username")
		password := c.Request.Form.Get("password")

		// Check authenticity
		user, err := userRepo.ReadUser(context.TODO(), username)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if !auth.CompareEqual(user.Password, password) {
			c.String(http.StatusUnauthorized, "Username and password don't match")
			return
		}

		// Generate auth token
		t, err := authMode.GenerateToken(c, username)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Send response
		if t != nil {
			c.JSON(http.StatusOK, gin.H{
				"token": *t,
			})
		}
	}
}

func Logout(authMode authmode.AuthMode) func(*gin.Context) {
	return func(c *gin.Context) {
		authMode.Logout(c)
		c.Status(http.StatusNoContent)
	}
}
