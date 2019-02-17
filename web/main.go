package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("./templates/*")
	router.LoadHTMLFiles("templates/home.html", "templates/userhome.html")

	router.GET("/",func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"Name": "Main website",
		})
	})
	router.GET("/userhome",func(c *gin.Context) {
		c.HTML(http.StatusOK, "userhome.html", gin.H{
			"Name": "Main website",
		})
	})

	router.Run(":8080")
}