package main

import "fmt"

func main() {
	nome := "Bruno"
	fmt.Println("Nome:", nome)
	mudarNome(&nome)
	fmt.Println("Nome alterado:", nome)
}

func mudarNome(n *string) {
	*n = "Jesus"
}
