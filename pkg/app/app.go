package app

import (
	"context"
	"database/sql"
	"github.com/Abdullayev65/gin_bun_project/pkg/hendler"
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"github.com/Abdullayev65/gin_bun_project/pkg/service"
	"github.com/Abdullayev65/gin_bun_project/pkg/utill"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"time"

	"github.com/uptrace/bun/extra/bundebug"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type App struct {
	handler *hendler.Handler
	router  *gin.Engine
}

func New() *App {
	db, ctx := database(false)
	s := service.New(db, ctx)
	t := utill.NewToken("salat", time.Hour*6)
	return &App{handler: hendler.New(s, t)}
}

func (a *App) SetUp() {
	a.router = a.initRouters()
}

func (a *App) Run() {
	err := a.router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

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

	return router
}

func database(resetModel bool) (*bun.DB, context.Context) {
	dsn := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(pgdb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	ctx := context.Background()
	if resetModel {
		db.ResetModel(ctx, (*models.User)(nil))
		db.ResetModel(ctx, (*models.Post)(nil))
	}
	return db, ctx
}
