package main

import "fmt"

import "github.com/kieran-ohara/gogreet"

func main() {
  message := gogreet.Hello("Kieran")
  fmt.Println(message)
}
