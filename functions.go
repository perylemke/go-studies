package main

import (
	"fmt"
)

// Retorno simples
func sum(x int, y int) int {
	return x + y
}

// Retorno duplo
func change(x string, y string) (string, string) {
	return y, x
}

// Retorno assinado
func split(x, y int) (result, rest int) { // os dois retornos são inteiros nesse exemplo
	rest = x % y
	result = x / y
	return
}

func executeFunction(f func(string) string, value string) {
	aux := f(value)
	fmt.Printf(aux)
}

func printValueByRef(value *string) {
	fmt.Printf("Valor por referencia = %v\r\n", *value)
}

func main() {
	fmt.Printf("Funções!\r\n")

	fmt.Printf("Soma 1+1 = %v\r\n", sum(1, 1))

	b, a := change("A", "B")
	fmt.Printf("troca A, B = %v, %v\r\n", b, a)

	result, rest := split(5, 2)
	fmt.Printf("A divisão de 5 por 2 é = %v\r\n", result)
	fmt.Printf("O resto da divisão de 5 por 2 é %v\r\n", rest)

	// Função anônima que vamos passar para printFunc
	ola := func(v string) string {
		return "Olá " + v + "!\r\n"
	}

	executeFunction(ola, "Pery")

	value := "Esse valor não vai ser copiado, só estamos passando o ponteiro"
	printValueByRef(&value)
}
