package app

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/hendler"
	"github.com/Abdullayev65/gin_bun_project/pkg/mw"
	"github.com/Abdullayev65/gin_bun_project/pkg/service"
	"github.com/Abdullayev65/gin_bun_project/pkg/utill"
	"github.com/gin-gonic/gin"
	"time"
)

type App struct {
	handler *hendler.Handler
	router  *gin.Engine
	MW      *mw.MW
}

func New() *App {
	db, ctx := database()
	s := service.New(db, ctx)
	t := utill.NewToken("salat", time.Hour*26)
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
