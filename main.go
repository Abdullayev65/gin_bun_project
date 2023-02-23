package main

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/app"
)

func main() {
	a := app.New()
	a.SetUp()
	a.Run()
}
