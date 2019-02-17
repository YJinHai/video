package main

import (
	"github.com/gin-gonic/gin"
)

const DOWNLOADS_PATH = "videos/"

func main() {
	router := gin.Default()
	//router.Static("/", "./public")
	router.Use(MaxAllowed(20))
	router.POST("/upload", Upload)
	router.GET("/videos/:vid", Download)
	router.GET("/download",func(c *gin.Context){
		c.Header("Content-Type", "image/png")
		c.File("./videos/4.png")



	})

	router.Run(":8011")
}



