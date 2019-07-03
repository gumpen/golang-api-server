package main

import (
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	Version  string
	Revision string
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())

	router.StaticFile("/", "./index.html")

	router.GET("/api/greeting", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello, world",
		})
	})

	return router
}

func main() {
	fmt.Println("Greetings Server : Version:" + Version + " Revision:" + Revision)

	endless.ListenAndServe(":8080", setupRouter())
}
