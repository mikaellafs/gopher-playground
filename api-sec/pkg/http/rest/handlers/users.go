package handlers

import (
	"context"
	"net/http"

	"gopher-playground/api-sec/pkg/auth"
	"gopher-playground/api-sec/pkg/auth/user"

	"github.com/gin-gonic/gin"
)

type criarUsuarioBody struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(repo user.Repository) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get body
		var body criarUsuarioBody
		if err := c.Bind(&body); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Create user
		hashedPassword, err := auth.Encrypt(body.Password)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to encrypt password:"+err.Error())
		}

		_, err = repo.CreateUser(context.TODO(), body.Username, hashedPassword, body.Name)
		if err != nil {
			c.String(httpStatusFor(err), err.Error())
			return
		}

		c.Status(http.StatusCreated)
	}
}

func DeleteUser(repo user.Repository) func(*gin.Context) {
	return func(c *gin.Context) {
		ctxUser, exists := c.Get("user")
		if !exists {
			c.String(http.StatusInternalServerError, "Failed to get user from context")
			return
		}

		u, _ := ctxUser.(*user.User)
		err := repo.DeleteUser(context.TODO(), u.Username)
		if err != nil {
			c.String(httpStatusFor(err), err.Error())
			return
		}

		c.Status(http.StatusOK)
	}
}

// TODO: move to package
func httpStatusFor(err error) int {
	switch err {
	case user.ErrUserAlreadyExists:
		return http.StatusConflict
	case user.ErrUserNotFound:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
