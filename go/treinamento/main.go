package main

import "fmt"

func enviarDados(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Envia valores para o channel
		fmt.Println("Valor enviado:", i)
	}
	close(ch) // Fecha o channel após enviar todos os valores
}

func receberDados(ch <-chan int) {
	for {
		valor, ok := <-ch // Recebe valores do channel
		if !ok {
			fmt.Println("Channel fechado.")
			return
		}
		fmt.Println("Valor recebido:", valor)
	}
}

func main() {
	// Criar um channel
	dados := make(chan int)

	// Executar a função para enviar dados em uma goroutine
	go enviarDados(dados)

	// Executar a função para receber dados em outra goroutine
	go receberDados(dados)

	// Aguardar um pouco para não terminar imediatamente
	fmt.Println("Aguardando...")
	fmt.Scanln()
}
