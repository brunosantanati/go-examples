# Course [Go (Golang): Explorando a Linguagem do Google](https://www.udemy.com/course/curso-go/)  

## Links  
[Documentation](https://go.dev/doc/code)  
[Golang Pointers - Fully Understanding Pointers in Go](https://www.youtube.com/watch?v=XOz_Xu5WanQ)  

## Linux Installation
```
1-Download the Go binaries from the website
2-sudo tar -C /usr/local/ -xzf go1.13.4.linux-amd64.tar.gz
3-vi ~/.profile
export PATH=$PATH:/usr/local/go/bin
4-source ~/.profile
```

## Windows Installation
```
1-Download Go installer from the website (for example this link https://golang.org/doc/install?download=go1.15.6.windows-amd64.msi)
2-Install Go using the 'wizard' (it automatically put the C:\Go\bin in the PATH)
```

## Plugins VSCode
```
1-Go
2-Code Runner
```

## Commands
```
go version
It shows the installed version

go env
It shows the environment variables
GOROOT="/usr/local/go" -> location where Go was installed and where stands all tools to execute Go
GOPATH="/home/bruno/go" -> location where stands all projects

go env GOPATH
go env GOROOT
It shows a specific environment variable, in the first example the GOPATH and in the second example the GOROOT

Structure for GOPATH
/home/bruno/go/bin -> when you create a executable Go program it stands here.
/home/bruno/go/pkg -> It contains the compiled files.
/home/bruno/go/src -> It contains the source code files.

go doc cmd/vet
It shows Go documentation about Vet

go vet comandos.go
Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string.

go help get
It shows information about a command, in this case about get command.

godoc -http=:6060
It enables offline Go documentation in the address http://localhost:6060
When I ran this, Linux suggested that I should run "sudo apt install golang-golang-x-tools", and I did it. But it remains with error.

go build comandos.go
It compiles the code. The example above generates a comandos file that could be executed using ./comandos

go run comandos.go
OR
go run "/home/bruno/Documentos/go/fundamentos/comandos/comandos.go"
It compiles and executes a go program

go run *.go
There were two files in the directory(funcoes.go and main.go), but only one has the main method. I used this command to run the main file.
OR
On Windows, to execute multiple files:
go run arquivo1.go arquivo2.go

go get -u github.com/go-sql-driver/mysql
It installs the mysql package at /home/bruno/go/src/github.com/go-sql-driver/mysql

go get -u github.com/brunosantanati/goarea
It gets the code from https://github.com/brunosantanati/goarea and installs it in C:\Users\55119\go\pkg\windows_amd64\github.com\brunosantanati as a binary file and puts the source code in C:\Users\55119\go\src\github.com\brunosantanati\goarea

go test
It runs the unit tests(you should be located in the folder where the test file is)
go test -v
Ir runs the unit tests in verbose mode(It prints all t.Logf statements)
go test ./...
It runs the unit tests(from the current folder and It considers also all subfolders)
go test std
It runs the unit tests for golang packages

UNIT TEST COVER
Should run the command bellow from C:\Users\55119\go\src\github.com\brunosantanati\matematica
go test -cover
go test -coverprofile=resultado.out
go tool cover -func=resultado.out
go tool cover -html=resultado.out

SERVER
go run static.go
It runs the static.go which starts a server which serves static content. The page should be accessed by this url http://localhost:3000
go run dinamico.go
It runs the dynamic.go which starts a server which serves dynamic content. The page should be accessed by this url http://localhost:3000/horaCerta

FROM VS CODE
ctrl + alt + n
It executes Go program
ctrl + alt + m
It stops Go program execution
```

## Getting started examples
```
/home/bruno/Documentos/go -> folder(could be any folder) where I put my getting started examples

cd /home/bruno/Documentos/go && code .
Commands to Open the getting started examples in the VS CODE
```
