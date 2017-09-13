package main

import (
	"fmt"
)

type shape interface {
	area() int
}

type Rectangle struct {
	height int
	width int
}

func (r Rectangle) area() int {
	return r.height *r.width
}

func anotherArea(s shape) int {
	return s.area()
}

func main() {
	r := Rectangle {2 , 6}
	fmt.Print(anotherArea(r))
}
