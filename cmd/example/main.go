package main

import (
	"github.com/sp00kygh0st/example/pkg/foobar"
	"github.com/sp00kygh0st/example/pkg/xyzzy"
)

func main() {
	foobar.Foobar()

	x := xyzzy.NewXyzzy()
	x.Run(":8080")
}
