package main

import (
	"fmt"
	"url-shortener-go/shorten"
)

func main() {
	str, _ := shorten.CreateKey()
	fmt.Println(str)
	fmt.Println(len(str))
}


