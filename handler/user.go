package handler

import (
	"log"
	"net/http"

	"github.com/djomlaa/helpee/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	v "github.com/djomlaa/helpee/validator"
)

var validate *validator.Validate

func (h *handler) users(ctx *gin.Context) {
	log.Println("Users endpoint")
	uu, err := h.Users()
	if err != nil {
		respondError(ctx, err, http.StatusInternalServerError)
		return
	}
	respond(ctx, uu, 200)
}


func (h *handler) createUser(ctx *gin.Context) {
	log.Println("Create User endpoint")

	validate = validator.New()
	validate.RegisterValidation("dateOfBirth", v.ValidateDateOfBirth)

	var user service.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		respondError(ctx, err, http.StatusInternalServerError)		
		return
	}

	err := validate.Struct(user)
	if err != nil {
		respondError(ctx, err, http.StatusInternalServerError)		
		return
	}

	err = h.CreateUser(ctx, user)

	if err == service.ErrEmailTaken {
		respondError(ctx, err,http.StatusConflict)
		return
	}

	if err == service.ErrUsernameTaken {
		respondError(ctx, err, http.StatusConflict)
		return
	}

	if err != nil {
		respondError(ctx, err, http.StatusInternalServerError)
		return
	}
	respond(ctx, nil, 200)
}