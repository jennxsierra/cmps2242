// Filename: go-basics.go

package main

import (
	"fmt"
)

// Hello returns the greeting message
func Hello() string {
	return "Hello, world!"
}

func main() {
	fmt.Println(Hello())
}
