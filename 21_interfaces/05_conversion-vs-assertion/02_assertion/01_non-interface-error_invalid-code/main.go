package main

import "fmt"

func main() {
//	name := "Sydney"
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
