package foobar

import (
	"fmt"

	"example/internal/pkg/other"
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
