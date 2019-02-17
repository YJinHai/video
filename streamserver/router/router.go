package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"video/api/handler/comment"
	"video/api/handler/user"
	"video/api/handler/video"
	"video/api/router/middleware"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(mw...)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.POST("/user", user.Create)
	//g.POST("/user/login", user.Login)
	//r.GET("/user/:uname", api.GetUser)
	g.GET("/comment/:id", comment.Get)
	v1 := g.Group("/user/:uid")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.POST("/video", video.Create)
		v1.POST("/videos", user.Videos)
		v1.POST("/video/:vid/comment", comment.Create)
		v1.POST("/video/:vid/comments", video.Comments)
		//v1.POST("/video/:")

		//v1.POST("/comment", api.CreateComment)
		v1.POST("/comments", user.Comments)



	}
	g.Run(":8010")

	return g
}
