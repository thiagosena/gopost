package http

import (
	"errors"
	"net/http"

	"github.com/thiagosena/gopost/internal"
	"github.com/thiagosena/gopost/internal/database"
	"github.com/thiagosena/gopost/internal/post"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var service post.Service

func Configure() {
	service = post.Service{
		Repository: post.Repository{
			Conn: database.Conn,
		},
	}
}

func CreatePosts(ctx *gin.Context) {
	var dto internal.Post
	if err := ctx.BindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.Create(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func DeletePosts(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	if err := service.Delete(id); err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, post.ErrPostNotFound) {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func GetPosts(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	p, err := service.FindOneByID(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, post.ErrPostNotFound) {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func GetAllPosts(ctx *gin.Context) {
	p, err := service.FindAll()
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, post.ErrPostNotFound) {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, p)
}
