package main

import "fmt"

// Definindo a interface Animal
type Animal interface {
	Som() string
	Andar() string
}

// Definindo uma struct Dog que implementa a interface Animal
type Dog struct {
	Nome string
}

// Implementação do método Som para Dog
func (d Dog) Som() string {
	return "Au Au"
}

// Implementação do método Andar para Dog
func (d Dog) Andar() string {
	return fmt.Sprintf("%s está caminhando...", d.Nome)
}

func main() {
	// Criando uma instância de Dog
	cachorro := Dog{Nome: "Rex"}

	// Usando métodos da interface Animal
	fmt.Println("Som do", cachorro.Nome+":", cachorro.Som())
	fmt.Println("Movimento do", cachorro.Nome+":", cachorro.Andar())
}
