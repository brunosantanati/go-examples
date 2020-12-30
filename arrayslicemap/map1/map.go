package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados, por isso a linha acima não funciona
	aprovados := make(map[int]string)
	fmt.Println(aprovados)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	// Iterar em um map
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682]) // Recuperar valor pela chave
	delete(aprovados, 95135745682)      // Excluir item do map
	fmt.Println(aprovados[95135745682]) // Não consegue recuperar porque o item com essa chave foi deletado
	fmt.Println(aprovados)
}
