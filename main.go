package main

import (
	"fmt"

	greeting "github.com/runcess/module1-greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}
