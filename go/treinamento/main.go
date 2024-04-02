package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Email  string
	Phone  string
	Status bool
	Father *Client // Autorelacionamento: Ponteiro para outro Client
}

func main() {
	// Criando instâncias de Client
	child := Client{
		Name:   "João",
		Age:    30,
		Email:  "joao@example.com",
		Phone:  "123-456-7890",
		Status: true,
	}

	father := Client{
		Name:   "Pedro",
		Age:    25,
		Email:  "pedro@example.com",
		Phone:  "987-654-3210",
		Status: true,
	}

	// Estabelecendo a relação "pai-filho"
	child.Father = &father

	// Imprimindo informações
	fmt.Println("Cliente 1:")
	fmt.Println("Nome:", child.Name)
	fmt.Println("Idade:", child.Age)
	fmt.Println("E-mail:", child.Email)
	fmt.Println("Telefone:", child.Phone)
	fmt.Println("Status:", child.Status)
	if child.Father != nil {
		fmt.Println("Pai:", child.Father.Name)
	} else {
		fmt.Println("Pai: Nenhum")
	}

	fmt.Println("\nCliente 2:")
	fmt.Println("Nome:", father.Name)
	fmt.Println("Idade:", father.Age)
	fmt.Println("E-mail:", father.Email)
	fmt.Println("Telefone:", father.Phone)
	fmt.Println("Status:", father.Status)
}
