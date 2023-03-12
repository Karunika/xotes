package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/rs/cors"
)

func main() {
	router := gin.New()

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello!!",
		})
	})

	router.Run(":3000")
}
