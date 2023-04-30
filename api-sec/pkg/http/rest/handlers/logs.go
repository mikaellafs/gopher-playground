package handlers

import (
	"net/http"

	"gopher-playground/api-sec/pkg/log"

	"github.com/gin-gonic/gin"
)

type LogsHandler struct {
	repo log.Repository
}

func NewLogsHandler(r log.Repository) *LogsHandler {
	return &LogsHandler{
		repo: r,
	}
}

func (h *LogsHandler) ListAll(c *gin.Context) {
	result := h.repo.ReadAll()

	c.JSON(http.StatusOK, result)
}
