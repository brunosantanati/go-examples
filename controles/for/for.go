package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("1) ####################################")
	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	fmt.Println("\n2) ####################################")
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n\n3) ####################################")
	fmt.Println("Misturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	fmt.Println("\n\n4) ####################################")
	for {
		// laço infinito
		fmt.Println("Para sempre...")
		//time.Sleep(time.Second) // Imprimir de 1 em 1 segundo
		time.Sleep(time.Second * 5) // Imprimir a cada 5 segundos
	}

	// Veremos o foreach no capítulo de array
}
