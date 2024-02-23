package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.POST("/posts", CreatePosts)
	g.DELETE("/posts/:id", DeletePosts)
	g.GET("/posts", GetAllPosts)
	g.GET("/posts/:id", GetPosts)

	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	g.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "olá mundo")
	})
}
