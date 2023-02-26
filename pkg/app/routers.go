package app

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/myTest"
	"github.com/gin-gonic/gin"
)

func (a *App) initRouters() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/", a.handler.UserAdd)
		user.GET("/me", a.handler.ParseTokenAndRequiredSetUserID,
			a.handler.UserGet)
		user.GET("/all", a.handler.UserAll)
		user.POST("/token", a.handler.UserToken)
		user.PUT("/", a.handler.ParseTokenAndRequiredSetUserID,
			a.handler.UserUpdate)
		user.DELETE("/", a.handler.ParseTokenAndRequiredSetUserID,
			a.handler.UserDelete)
	}
	post := router.Group("/post")
	{
		post.POST("/", a.handler.ParseTokenAndRequiredSetUserID,
			a.handler.PostAdd)
		post.GET("/", a.handler.PostAll)
		post.GET("/:id", a.handler.PostGet)
		post.PUT("/:id", a.handler.PostUpdate)
		post.DELETE("/:id", a.handler.PostDelete)
	}

	comment := router.Group("/comment")
	{
		comment.POST("/", a.handler.ParseTokenAndRequiredSetUserID,
			a.handler.CommentAdd)
		comment.GET("/", a.handler.CommentAll)
		comment.GET("/:id", a.handler.CommentGet)
		comment.PUT("/:id", a.handler.CommentUpdate)
		comment.DELETE("/:id", a.handler.CommentDelete)
	}

	attachment := router.Group("/attachment")
	{
		attachment.POST("/", a.handler.AttachmentAdd)
		attachment.GET("/:id", a.MW.SetInt("id"),
			a.handler.AttachmentGet)
	}

	testGroup := router.Group("/test")
	{
		testGroup.POST("/", myTest.Test)
	}

	return router
}
