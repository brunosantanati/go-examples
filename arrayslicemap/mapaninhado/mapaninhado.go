package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João":    11325.45,
			"João Roberto": 10000.00,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	fmt.Println(funcsPorLetra)
	delete(funcsPorLetra, "P") // Excluir todos funcionários que tem nome com P
	fmt.Println(funcsPorLetra)

	fmt.Println("##################################")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	// For dentro de for para percorrer map aninhado
	fmt.Println("##################################")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
