package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
)
func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
