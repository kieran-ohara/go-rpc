package main

import "github.com/kieran-ohara/gogreet"
import "github.com/kieran-ohara/gogreet/greetrpcdist"
import "github.com/kieran-ohara/gogreet/server"

import "fmt"

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

  server.RpcServe()
}
