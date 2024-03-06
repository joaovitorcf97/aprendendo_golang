package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Printf("Hello %s", texto)
	}("world!")
}
