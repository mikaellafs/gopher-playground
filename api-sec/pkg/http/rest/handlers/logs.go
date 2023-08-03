package handlers

import (
	"net/http"

	"gopher-playground/api-sec/pkg/log"

	"github.com/gin-gonic/gin"
)

func ListAllLogs(repo log.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		result := repo.ReadAll()

		c.JSON(http.StatusOK, result)
	}
}
