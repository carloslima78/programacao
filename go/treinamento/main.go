package main

import (
	"fmt"
)

func hello(nome string) {
	fmt.Println("Olá", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func main() {

	hello("João")

	fmt.Println("Sua idade é:", sum(10, 20), "anos")
}
