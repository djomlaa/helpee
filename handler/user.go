package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/djomlaa/helpee/service"
	v "github.com/djomlaa/helpee/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func (h *handler) users(ctx *gin.Context) {
	log.Println("Users endpoint")
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	if err != nil {
		respondError(ctx, err, http.StatusBadRequest)
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "5"))
	if err != nil {
		respondError(ctx, err, http.StatusBadRequest)
	}
	uu, err := h.Users(ctx, page, size)
	if err != nil {
		respondError(ctx, err, http.StatusInternalServerError)
		return
	}
	respond(ctx, uu, 200)
}


func (h *handler) user(ctx *gin.Context) {
	log.Println("User endpoint")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		respondError(ctx, err, http.StatusBadRequest)
	}

	u, err := h.User(ctx, id)
	if err != nil {
		respondError(ctx, err, http.StatusInternalServerError)
		return
	}
	respond(ctx, u, 200)
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