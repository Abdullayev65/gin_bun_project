package app

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/controller"
	"github.com/gin-gonic/gin"
)

type App struct {
	controller *controller.Controller
}

func New() *App {
	return &App{controller: controller.New()}
}
func (a *App) SetUp() {
	a.initRouters()
}
func (a *App) Run() {

}

func (a *App) initRouters() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/", a.controller.UserController.Add)
		user.GET("/me", a.controller.UserController.Get)
		user.GET("/all", a.controller.UserController.All)
		user.PUT("/", a.controller.UserController.Update)
		user.DELETE("/", a.controller.UserController.Delete)
	}
	post := router.Group("/post")
	{
		post.POST("/", a.controller.PostController.Add)
		post.GET("/", a.controller.PostController.All)
		post.GET("/:id", a.controller.PostController.Get)
		post.PUT("/:id", a.controller.PostController.Update)
		post.DELETE("/:id", a.controller.PostController.Delete)
	}

	return router
}
