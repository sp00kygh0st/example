package xyzzy

import (
	"github.com/gin-gonic/gin"
)

type Xyzzy struct {
	engine *gin.Engine
}

func (x *Xyzzy) Index(c *gin.Context) {
	c.String(200, "OK")
}

func (x *Xyzzy) Run(addr string) {
	x.engine.Run(addr)
}

func NewXyzzy() *Xyzzy {
	x := new(Xyzzy)
	x.engine = gin.Default()
	return x
}
