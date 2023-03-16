package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func main() {
	hello := "HELLO WORLD!"

	app := gin.Default()

	app.GET("/v1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": hello})
	})

	app.Run("localhost:3434")
}
