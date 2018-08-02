package foobar

import (
	"fmt"

	"github.com/sp00kygh0st/example/internal/pkg/other"
)

type Thing struct{}

func NewThing() *Thing {
	thing := new(Thing)
	return thing
}

func Foobar() {
	fmt.Println("Foobar!")

	other.Other()
}
