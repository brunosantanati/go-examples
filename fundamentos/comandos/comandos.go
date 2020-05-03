package main

import "fmt"

func main() {
	//fmt.Printf("Outro programa em %s!") //"go vet comandos.go" shows this message: ./comandos.go:6:2: Printf format %s reads arg #1, but call has 0 args
	fmt.Printf("Outro programa em %s!!!\n", "Go")
}

/*
COMMANDS EXECUTED IN THIS CLASS
go version
go env
go doc cmd/vet
go vet comandos.go
go help get
godoc -http=:6060
go build comandos.go
go run comandos.go
go get -u github.com/go-sql-driver/mysql
*/
