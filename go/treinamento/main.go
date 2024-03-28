package main

import "fmt"

func main() {
	a, b := 10, 15

	// Switch para comparar dois números inteiros
	switch {
	case a > b:
		fmt.Println("a é maior que b")
	case a < b:
		fmt.Println("a é menor que b")
	default:
		fmt.Println("a é igual a b")
	}
}
