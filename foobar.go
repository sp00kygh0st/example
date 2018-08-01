package sandbox

import "fmt"

type Foobar struct {
	value string
}

func NewFoobar() *Foobar {
	return &Foobar{}
}

func (*Foobar) DoThing() {
	fmt.Println("Doing the thing!")
}
