package main

import "fmt"

type Client struct {
	Name string
	Age  int
}

// Método de Pessoa para exibir os detalhes
func (c Client) ShowDetails() {
	fmt.Println("Nome:", c.Name)
	fmt.Println("Idade:", c.Age)
}

func main() {
	client := Client{Name: "Ana", Age: 30}
	client.ShowDetails()
}
