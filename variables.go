package main

import (
	"fmt"
)

var ( // declarando variaveis
	nome                   = "Pery"
	idade                  = 27
	url, path, query, page = "sudocast.com.br", "/search", "&q=golang", 1
)

const pi = 3.141592

func main() {
	// Declarando multiplas variaveis na mesma linha e atribuindo valor
	var x, y int = 1, 2

	// Declarando sem atribuir valor (Go sempre assume um valor automaticamente)
	var a string
	var s string

	/*
		A variavel "a" vai receber um valor mas vamos deixar
		a variavel "s" sem valor propositalmente no fim do programa
		vamos ver que uma string vazia foi atribuida a "s" automaticamente
	*/

	a = "texto 1"

	// Criação de variável utilizando : que substitui a palavra chave var quando se cria
	// uma variavel atribuindo um valor a ela na mesma instrução
	b := "texto 2"

	ola := func() {
		fmt.Printf("Olá da função anônima!\r\n")
	}

	fmt.Printf("a tipo: %T\r\n", a)
	fmt.Printf("b tipo  %T\r\n", b)
	fmt.Printf("π tipo: %T\r\n", pi)
	fmt.Printf("ola tipo: %T\r\n", ola)

	fmt.Printf("valor de a = %v\r\n", a)
	fmt.Printf("valor de b = %v\r\n", b)
	fmt.Printf("valor de π = %v\r\n", pi)

	fmt.Printf("valor de x = %v\r\n", x)
	fmt.Printf("valor de y = %v\r\n", y)

	fmt.Printf("valor de s = %q\r\n", s)

	ola()
}
