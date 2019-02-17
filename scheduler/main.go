package main

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

const DOWNLOADS_PATH = "videos/"

func main() {
	router := gin.Default()
	jobrunner.Start()
	router.DELETE("/videos/:vid", DeleteVideo)


	router.Run(":8012")
}



