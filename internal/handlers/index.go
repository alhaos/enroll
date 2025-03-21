package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) IndexGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", nil)
}
