package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // [...] compilador conta quantos elementos tem no array e define o tamanho fixo dele

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero) // Imprime o índice e o elemento naquele índice
	}

	fmt.Println("#####################################")
	for _, num := range numeros { // Ignorar a variável referente ao índice
		fmt.Println(num) // Imprime elementos
	}

	fmt.Println("#####################################")
	for i := range numeros { // Se tem uma variável só, pega o índice
		fmt.Println(i) // Imprime os índices
	}
}
