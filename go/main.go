package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	router.POST("/user/create")
	router.POST("/user/action")

	router.Run("localhost:8080")
}
