package main

import "fmt"

func main() {
	// Criar e já inicializar o map
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0 // Add mais um elemento no map
	delete(funcsESalarios, "Inexistente")   // Tentar excluir elemento que não existe, não faz nada

	// Iterar pelo map
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	fmt.Println("###########################")
	// Iterar somente pela chave
	for nome := range funcsESalarios {
		fmt.Println(nome)
	}

	fmt.Println("###########################")
	// Iterar somente pelo valor
	for _, salario := range funcsESalarios {
		fmt.Println(salario)
	}
}
