package handler

import (
	"github.com/djomlaa/helpee/middleware"
	"github.com/djomlaa/helpee/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	*service.Service
}

// SetRouter creates routing
func SetRouter(s *service.Service, e *gin.Engine) *gin.Engine {
	h := &handler{s}

	e.Use(middleware.Auth())

	e.POST("/login", h.login)
	e.GET("/users", h.users)
	e.POST("/users", h.createUser)

	return e
}

func respond(c *gin.Context, v interface{}, statusCode int) {
	c.JSON(statusCode, v)	
}

func respondError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}