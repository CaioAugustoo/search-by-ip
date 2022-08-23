# Search By IP

Aplicação de linha de comando que busca servidores por IP ou Name Servers.

## Sobre o projeto

Aplicação de linha de comando que busca servidores por IP ou Name Servers desenvolvida em GO.

## Objetivo do projeto

Este projeto foi desenvolvido com o intuito de aplicar e reforçar os conhecimentos adquiridos na linguagem GO.

## Tecnologias

Para a realização desse projeto foram utilizadas as seguintes tecnologias/linguagens:

- [GO](https://go.dev/)

## Instalação

Para que este rode em sua máquina, siga os passos abaixo:

```bash
# Clone o repositório em alguma pasta em sua máquina
$ git clone https://github.com/CaioAugustoo/search-by-ip

# Entre no repositório
$ cd search-by-ip

# Rode a aplicação
$ go run main.go

# Você pode gerar o build da aplicação
$ go build

# Retornando IP de amazon.com.br
$ go run main.go ip --host amazon.com.br

# Retornando servidores de amazon.com.br
$ go run main.go server --host amazon.com.br

Caso nenhum argumento seja passado, retornará dados do endereço: caioaugusto.dev
$ go run main.go server

# ou

$ go run main.go ip
```
