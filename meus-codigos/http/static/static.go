package main

import (
	"log"
	"net/http"
)

// Rodar no terminal: go run static.go
// Acessar: http://localhost:3000

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
