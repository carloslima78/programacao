package main

import "fmt"

func main() {
	// Cria um array de tamanho fixo de strings
	var nomes [3]string

	// Atribui valores aos elementos do array
	nomes[0] = "João"
	nomes[1] = "Maria"
	nomes[2] = "Pedro"

	// Imprime o array
	fmt.Println(nomes)

	// Acessando elementos do array
	fmt.Println(nomes[0]) // Imprime "João"
}
