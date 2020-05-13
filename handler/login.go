package handler

import (
	"log"
	"net/http"

	"github.com/djomlaa/helpee/service"
	"github.com/gin-gonic/gin"
)

//LoginInput entry type
type LoginInput struct {
	Email string `json:"email" binding:"email"`
	Password string `json:"password"`
}

func (h *handler) login(ctx *gin.Context) {
	log.Println("Login endpoint")

	var lI LoginInput

	if err:=ctx.ShouldBind(&lI); err != nil {
		respondError(ctx, err, http.StatusBadRequest)
		return
	}

	l, err := h.Login(ctx, lI.Email, lI.Password)

	if err == service.ErrUserNotFound {
		respondError(ctx, err, http.StatusNotFound)
		return
	}

	if err != nil {
		respondError(ctx, err, http.StatusInternalServerError)
		return
	}
	respond(ctx, l, 200)
}