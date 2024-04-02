package main

import "fmt"

// Definindo um tipo Veiculo básico
type Veiculo struct {
	Modelo string
	Ano    int
}

// Método para parar um veículo
func (v Veiculo) Parar() {
	fmt.Printf("O veículo %s está parado.\n", v.Modelo)
}

// Definindo um tipo Carro que incorpora Veiculo
type Carro struct {
	Veiculo // Incorporação do tipo Veiculo
	Motor   string
}

// Método para o carro andar
func (c Carro) Andar() {
	fmt.Printf("O carro %s está andando com seu motor %s...\n", c.Modelo, c.Motor)
}

// Definindo um tipo Barco que incorpora Veiculo
type Barco struct {
	Veiculo // Incorporação do tipo Veiculo
	Helice  string
}

// Método para o barco navegar
func (b Barco) Navegar() {
	fmt.Printf("O barco %s está navegando com sua hélice %s...\n", b.Modelo, b.Helice)
}

func main() {

	// Criando um Carro
	carro := Carro{
		Veiculo: Veiculo{Modelo: "Fusca", Ano: 1972},
		Motor:   "V8",
	}

	// Criando um Barco
	barco := Barco{
		Veiculo: Veiculo{Modelo: "Lancha", Ano: 2020},
		Helice:  "Marinha 3000",
	}

	// Chamando o método Andar para o Carro
	carro.Andar()

	// Chamando o método Navegar para o Barco
	barco.Navegar()

	// Chamando o método Parar para ambos os veículos
	carro.Parar()
	barco.Parar()
}
