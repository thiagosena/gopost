package http

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.POST("/posts", CreatePosts)
	g.DELETE("/posts/:id", DeletePosts)
	g.GET("/posts", GetAllPosts)
	g.GET("/posts/:id", GetPosts)
}
