package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
		time.Sleep(time.Second) // Imprimir de 1 em 1 segundo
		//time.Sleep(time.Second * 5) // Imprimir a cada 5 segundos

		// Break Statement in Golang
		//https://www.golearningsource.com/fundamentals/break-statement-in-golang/
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		numeroAleatorio := r.Intn(4)
		if numeroAleatorio == 3 {
			fmt.Println("Não era para sempre...")
			break // Sai do for
		}
	}

	// Usar uma label para o loop
	fmt.Println("\n5) ####################################")
	// label for outer loop
outer:
	for a := 0; a < 3; a++ {
		fmt.Print(strconv.Itoa(a) + " ")
		if a == 1 {
			fmt.Println("break outer called")
			// labelled break statement
			break outer
		}
		// label for inner loop
	inner:
		for b := 0; b < 3; b++ {
			fmt.Print(strconv.Itoa(b) + " ")
			if b == 1 {
				fmt.Println("break inner called")
				// labelled break statement
				break inner
			}
		}
	}
	fmt.Println("outside loop")

	// Veremos o foreach no capítulo de array
}
