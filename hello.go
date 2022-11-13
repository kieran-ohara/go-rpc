package main

import "fmt"

import "github.com/kieran-ohara/gogreet"

func main() {
	message, err := gogreet.Hello("")
	if err != nil {
		fmt.Println("there was an error")
	}
	fmt.Println(message)
}
