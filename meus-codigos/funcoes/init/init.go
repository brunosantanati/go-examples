package main

import "fmt"

//Para testar no terminal, dentro da pasta init rodar: go run *.go
//Assim serão levados em conta tanto o arquivo init.go quanto o init2.go na execução
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
