package main

import (
	"log"
	"net/http"
)

// Rodar no terminal: go run *.go
// URL todos usuários: http://localhost:3000/usuarios/
// URL usuário específico(passar id): http://localhost:3000/usuarios/1

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
