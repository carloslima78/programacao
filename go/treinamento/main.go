package main

import (
	"fmt"
	"time"
)

func worker() {
	i := 1

	// for infinito
	for {
		// Imprime a contagem atual
		fmt.Println("Contagem infinita:", i)
		i++

		// Simula o processamento contínuo
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Iniciando worker...")

	// Inicia o worker em uma goroutine
	go worker()
	fmt.Println("Worker iniciado.")

	// Espera um pouco antes de encerrar o programa (para observar a contagem)
	time.Sleep(time.Second * 10)
	fmt.Println("Programa encerrado.")
}
