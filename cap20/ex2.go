package main

import (
	"fmt"
)

type Human interface {
	speak(s []byte)
}

type Person struct {
	Name string
}

func (p *Person) speak(s []byte) {
	fmt.Println(p.Name, "says", string(s))
}

func saySomething(h Human) {
	h.speak([]byte("Somenthing"))
}

func main() {
	p := Person{"Renato"}

	p.speak([]byte("Shortcut to &p"))

	(&p).speak([]byte("Right way"))

	saySomething(&p)
}
