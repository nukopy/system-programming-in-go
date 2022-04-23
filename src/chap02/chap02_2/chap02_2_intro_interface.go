package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s!\n", g.name)
}

func main() {
	var talker Talker = &Greeter{name: "wozozo"}
	talker.Talk()
}