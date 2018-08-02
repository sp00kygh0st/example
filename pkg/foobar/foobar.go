package foobar

import (
	"fmt"
)

type Thing struct{}

func NewThing() *Thing {
	thing := new(Thing)
	return thing
}

func Foobar() {
	fmt.Println("Foobar!")
}
