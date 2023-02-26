package mw

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type MW struct {
}

func (mw *MW) SetInt(name string) func(*gin.Context) {
	return func(c *gin.Context) {
		param := c.Param(name)
		i, err := strconv.Atoi(param)
		if err != nil {
			c.String(400, "param "+name+" not found")
			c.Abort()
			return
		}
		c.Set(name, i)
	}
}
