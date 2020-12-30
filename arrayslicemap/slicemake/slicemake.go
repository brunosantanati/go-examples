package main

import "fmt"

func main() {
	s := make([]int, 10) // Criando um slice com make
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // Criado slice com tamanho 10 e o array interno com tamanho 20
	fmt.Println(s, len(s), cap(s)) // len pega o tamanho do slice(10) e cap pega a capacidade máxima que é o tamanho do array(20)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // Adiciona mais 10 alementos para atingir a capacidade máxima
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)               // Adicionando um elemento mesmo que o slice já atingiu a capacidade máxima
	fmt.Println(s, len(s), cap(s)) // Podemos ver que o tamanho do slice cresceu para 21 e a capacidade do array interno dobrou de tamanho(usou outro array internamente)
}
