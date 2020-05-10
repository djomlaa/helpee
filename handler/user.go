package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *handler) users(ctx *gin.Context) {
	log.Println("Users endpoint")
	uu, err := h.Users()
	if err != nil {
		respondError(ctx, err)
		return
	}
	respond(ctx, uu, 200)
}