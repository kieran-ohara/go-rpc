package main

import "fmt"

import "github.com/kieran-ohara/gogreet"
import "github.com/kieran-ohara/gogreet/greetrpcdist"

func main() {
	message, err := gogreet.Hello("this is a Hello name")
	if err != nil {
		fmt.Println("there was an error")
	}
	fmt.Println(message)

  greetDevice := greetrpcdist.Greet {
    Name: "this is Greet name",
  }
	fmt.Println(greetDevice.Name)
}
