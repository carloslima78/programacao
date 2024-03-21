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

O Go é uma linguagem fortemente tipada, porém, não é orientada a objetos e sim, orientado a dados.

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

## Comandos Principais

- Comando para executar um arquivo Go

```golang
go run hello.go
```

## Principais Tipos de Dados

### String 

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