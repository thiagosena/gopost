package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagosena/gopost/internal/database"
	"github.com/thiagosena/gopost/internal/http"
)

func main() {
	connectionString := "postgresql://posts:p0stgr3s@db:5432/posts"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	g := gin.Default()
	http.Configure()
	http.SetRoutes(g)
	err = g.Run(":3000")
	if err != nil {
		panic(err)
	}
}
