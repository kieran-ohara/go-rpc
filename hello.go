package main

import "fmt"

import "github.com/kieran-ohara/gogreet"
import "github.com/kieran-ohara/gogreet/greetrpcdist"

func main() {
	message, err := gogreet.Hello("")
	if err != nil {
		fmt.Println("there was an error")
	}
	fmt.Println(message)

  greetDevice := greetrpcdist.Greet {
    Name: "this is a name",
  }
	fmt.Println(greetDevice.Name)
}
