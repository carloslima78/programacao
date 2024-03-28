# Golang

## Forças

- Sintexe simples de aprender, curva de aprendizado menor que outras linguagens.
- Oferece performance, concorrência e paralelismo.
- Reduz o consumo de poder computacional, disponibliza mecanismos que colobora exigir menos recursos.
- Plataforma com recursos nativos de segurança.
- Bastante utilizado para comunicação entre microservices via GRPC.
- Acesso simples aos recursos e protocolos modernos de rede.
- Nasceu otimizado para multi-core.
- Baixa utilização de memória, inicia uma thread utilizando apenas 2k.
- Garbage Collector otimizado.
- O Go é uma linguagem compilada, porém não utiliza máquina virtual.


## Timeline

- **2007:** Início do projeto.
- **2009:** Open Sourced.
- **2012:** Versão estável v1.0.
- **2015:** Melhorias importantes v1.5.
- **2018:** Consolidação do ecossistema v1.11.
- **2022:** Introdução do Generics v1.18.
- **2024:** Novidades no HTTP Server para gerenciamento de rotas v1.22.

## Retrocompatilidade

O Golang oferece retrocompatiiblidade com todas as versões anteriores, ou seja, uma versão mais atualualizada do Go, vai conseguir executar aplicações compiladas em qualquer versão anterior independente da evolução e do tempo.

## Go é orientado a objetos?

O Go é uma linguagem fortemente tipada, porém, *não é orientada a objetos* e sim, orientada a dados.

## Go é Plataforma

- Programação.
- Testes
- Geração de documentos.
- Gerenciamento de dependências.
- Multiplataforma.
- Profilling.
- Security.
- Performance "out-of-box"
- Build & Deploy.

## Fundamentos

- Go Runtime + Código escrito = Binário

## Instalação

- Acesse  *go.dev*, selecione o sistema operacional e sucesso.
- Instale o plugin oficial para VSCode *Rich Go language support for Visual Studio Code*.


## Packages

Packages são como o Go organiza o código em um programa. Cada programa Go precisa de um pacote de entrada único. Um pacote Go não pode conter dois pacotes com nomes diferentes. 

Geralmente, um programa Go começa pela função main, que é a entrada do programa. A partir da função main, é possível chamar outras funções ou pacotes para executar diversas tarefas no programa.

[Link para pkg.go.dev que contém os pacotes padrão do Go](https://pkg.go.dev/std)

1. Crie um arquivo *main.go* e escreva o código abaixo:

```golang
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Olá Mundo")
}
```

2. Comando para executar um progrma Go

```golang
go run main.go
```

## Visibilidade

No Go, não existem as visibilidades presentes no encapsulamento da orientação a objetos, a visibilidade é controlada pelo uso de letras maiúsculas ou minúsculas no início dos nomes. 

Variáveis ou funções que começam com letra maiúscula são visíveis (exportadas), podendo ser acessadas por pacotes externos. Por exemplo, a variável Nome é exportada, enquanto a função calcularMedia não é. Veja:

```golang
package main

import (
	"fmt"
)

// Variável exportada
var Nome string = "João"

// Função não exportada
func calcularMedia(n1, n2 float64) float64 {
	return (n1 + n2) / 2
}

func main() {
	fmt.Println(Nome) // Saída: João
	// fmt.Println(calcularMedia(7.5, 8.2)) // Erro, função não exportada
}
```
Na função main, podemos imprimir a variável Nome, pois ela é exportada e visível em outros pacotes. No entanto, ao tentar chamar calcularMedia fora deste pacote, ocorrerá um erro, pois ela não é exportada e, portanto, não é visível em outros pacotes.


## Variáveis e Tipos de Dados

Variáveis são usadas para armazenar dados em um programa Go, já os tipos de dados são usados para especificar o tipo de dado que uma variável pode armazenar.

No Go, todo tipo de dado é fortemente tipado, e isso significa que cada variável deve ser declarada com um tipo de dado específico. 

Toda variável no Go é iniciada com um valor padrão, exemplo:

- int = 0
- string = ""
- bool = false
- float64 = 0.0


### Variável a nível de Package

Uma variável a nível de package é visível para todas as funções do pacote. Para declarar uma variável a nível de package, use a palavra-chave `var` seguida do tipo de dado e do nome da variável.

Por exemplo, a seguinte declaração cria uma variável a nível de package chamada `name` do tipo `string`:

```golang
package main

import (
	"fmt"
)

// declaração da variável name, que armazenará um texto
var name string 

func main() {

	// atribuição do valor "João" à variável name
	name = "João" 

	// impressão do texto "Olá mundo! Meu nome é João." no console
	fmt.Println("Olá mundo! Meu nome é", name, "!")
}
```

### Múltiplas Declarações a nível de Package

É possível declarar várias variáveis a nível de package de uma só vez, usando a seguinte sintaxe:


```golang
var (
	
	name string
	n1 int
	n2 int
)
```

Vale observar que não é possível declarar variáveis com essa sintaxe no escopo de uma função.

### Atribuições a nível de função

Para atribuir um valor a uma variável a nível de função, use a seguinte sintaxe:

```golang
message := "Olá Mundo"
fmt.Println(message)
```

### Múltiplas Atribuições

É possível atribuir valores a várias variáveis de uma só vez, usando a seguinte sintaxe:


```golang
var a, b, c = true, 2.3, "Olá"
fmt.Println(a, b, c)
```

### Troca de valores entre variáveis

Para trocar os valores de duas variáveis, use a seguinte sintaxe:


```golang
// declaração das variáveis x e y
var x, y = 10, 20
fmt.Println("Antes da troca:", x, y)
// troca os valores das variáveis x e y
x, y = y, x
fmt.Println("Depois da troca:", x, y)
```

## Condicionais

As estruturas condicionais são usadas para controlar o fluxo de execução de um programa Go. No Go, existem dois tipos de estruturas condicionais: `if` e `switch`.

### O if

O condicional if é usado para verificar se uma condição é verdadeira ou falsa. Em go, o if não utiliza parentteses para delimitar a condição.

 A sintaxe do if é a seguinte:

```golang
package main

import "fmt"

func main() {

	a, b := 10, 15

	// Comparando dois números inteiros
	if a > b {
		println("a é maior que b")
	} else if a < b {
		println("a é menor que b")
	} else {
		println("a é igual a b")
	}
}
```

### O switch

O condicional switch em Go permite avaliar várias condições e executar o bloco de código correspondente. Diferentemente de algumas linguagens, Go não requer break após cada caso. O exemplo abaixo ilustra isso:

```golang
package main

import "fmt"

func main() {

	numero := 2

	switch numero {

	case 1:
		fmt.Println("O número é 1")

	case 2:
		fmt.Println("O número é 2")

	default:
		fmt.Println("Número não reconhecido")
	}
}
```

Neste exemplo simples de switch em Go:

numero é avaliado em cada caso:

- Se numero for 1, imprime "O número é 1".
- Se numero for 2, imprime "O número é 2".
- Se numero não corresponder a nenhum caso, o bloco default é executado, imprimindo "Número não reconhecido".

Não é necessário utilizar break após cada caso, pois o comportamento padrão do switch em Go é sair do switch após executar o caso correspondente.

### Switch com condicionais

Neste exemplo o código abaixo ilustra como usar um switch em vez de múltiplos if e else if:

```golang
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
```

## Funções

Funções são blocos de código que podem ser reutilizados em um programa Go. Para declarar uma função, use a seguinte sintaxe:


```golang
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
```

### Funções com múltiplos retornos

Em Go é comum usar funções com mais de um retorno, especialmente para lidar com possíveis erros. Por exemplo:

```golang
package main

import (
	"errors"
	"fmt"
)

// Função divide recebe dois inteiros e retorna um inteiro e um erro
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// A função main é onde a execução do programa começa.
func main() {
	// Chamando a função divide com 10 como dividendo e 2 como divisor
	resultado, err := divide(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", resultado)
	}

	// Chamando a função divide com 10 como dividendo e 0 como divisor
	resultado, err = divide(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", resultado)
	}
}
```

Neste exemplo, a função **divide** tem dois retornos: resultado do tipo int e erro do tipo error. Se o divisor for zero, ela retorna um erro. No main, verificamos se há um erro após chamar a função divide e imprimimos a mensagem de erro.


### Array 

### Slice

### Struct

## Gerando Binário no Ubuntu

#### Gerando o arquivo binário

```bash
go build hello.go
```

#### Executando o arquivo binário

```bash
./hello
```

## Executando processos paralelos (multi-threads)

```golang
func contador(count int) {

	for i := 0; i < count; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println(i)
	}
}
```

```golang
func main() {

	// Thread 01 inicial com 2kb
	go contador(2)
	// Thread 02 inicial com 2kb
	go contador(2)
	// Thread 0 inicial com 2kb
	contador(2)
}
```

```bash
carlos@ubuntu-Lenovo-IdeaPad-Z400:~/Documents/Fontes - Projetos/programacao/go$ go run main.go
0
0
0
1
1
1
```


