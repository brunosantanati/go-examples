package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Pq vc não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc!", 1)

	//go fale("Maria", "Ei...", 500) // Executar usando go cria uma goroutine(parecido com uma thread "leve")
	//go fale("João", "Opa...", 500)
	//time.Sleep(time.Second * 5)

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}
