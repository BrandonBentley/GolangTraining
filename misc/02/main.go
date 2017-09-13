package main

import (
	"fmt"
	"sort"
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
	slice1 := []string {"this", "is", "a", "slice"}
	r := Rectangle {2 , 6}
	sort.Strings(slice1)
	fmt.Println(anotherArea(r))
	fmt.Println(slice1)
}
