package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando o texto")
}
