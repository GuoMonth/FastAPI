package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	// Create the route
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	// bind the route
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "Hi,I am FastAPI,Just CURD!😀")
	})

	r.Run(":3721")
}
