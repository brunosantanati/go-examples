package main

import "fmt"

// Go não tem aritmética de ponteiros
func main() {

	i := 1

	var p *int = nil // criando um ponteiro
	p = &i           // pegando o endereço de memória de i
	*p++             // desreferenciando (pegando o valor), para poder incrementar
	i++

	// p++ //isso não funciona, pois tenta incrementar o endereço de memória

	fmt.Println(p, *p, i, &i)
}
