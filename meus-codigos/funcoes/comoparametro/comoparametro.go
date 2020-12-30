package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func divisao(a, b int) int {
	return a / b
}

func exec(funcao func(int, int) int, p1, p2 int) int { // Assinatura que recebe uma função e outros parâmetros
	fmt.Println("Fazendo conta...\nResultado:")
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4) // Passando uma função para outra
	fmt.Println(resultado)

	resultado2 := exec(divisao, 4, 4) // Passando uma função para outra
	fmt.Println(resultado2)
}
