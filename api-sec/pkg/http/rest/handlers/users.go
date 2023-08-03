package handlers

import (
	"context"
	"net/http"

	"gopher-playground/api-sec/pkg/auth"
	"gopher-playground/api-sec/pkg/auth/accesscontrol"
	"gopher-playground/api-sec/pkg/auth/user"

	"github.com/gin-gonic/gin"
)

type criarUsuarioArgs struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(repo user.Repository, ac accesscontrol.AccessControl) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get args
		var args criarUsuarioArgs
		if err := c.Bind(&args); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Create user
		hashedPassword, err := auth.Encrypt(args.Password)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to encrypt password:"+err.Error())
		}

		_, err = repo.CreateUser(context.TODO(), args.Username, hashedPassword, args.Name)
		if err != nil {
			c.String(httpStatusFor(err), err.Error())
			return
		}

		// Add default user role
		ac.AssignRoleToUser(args.Username, accesscontrol.DefaultRole)

		c.Status(http.StatusCreated)
	}
}

func DeleteUser(repo user.Repository, ac accesscontrol.AccessControl) func(*gin.Context) {
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

		// Delete user role
		ac.RemoveUser(u.Username)

		c.Status(http.StatusNoContent)
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
