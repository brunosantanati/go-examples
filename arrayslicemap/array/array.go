package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10 // índice inválido, gera erro de compilação: invalid array index 3
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %f\n", media)   // Média com muitas casas decimais
	fmt.Printf("Média %.2f\n", media) // Média com duas casas decimais
}
