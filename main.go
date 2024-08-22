package main

import "fmt"

func main() {
	msg := sayHello("Test")
	println(msg)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s Gopher", name)
}
