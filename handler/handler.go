package handler

import (
	"net/http"

	"github.com/djomlaa/helpee/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	*service.Service
}

// SetRouter creates routing
func SetRouter(s *service.Service, e *gin.Engine) *gin.Engine {
	h := &handler{s}

	e.GET("/users", h.users)

	return e
}

func respond(c *gin.Context, v interface{}, statusCode int) {
	c.JSON(statusCode, v)	
}

func respondError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}