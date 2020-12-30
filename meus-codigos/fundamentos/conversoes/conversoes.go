package main

import (
	"fmt"
	c "strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(int(x) / y)     //converter para int
	fmt.Println(x / float64(y)) //converter para float64

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado
	//fmt.Println("Teste" + 123) //invalid operation: "Teste" + 123 (mismatched types untyped string and untyped int)
	fmt.Println("Teste " + string(123)) //mostra o caracter correspondente ao código 123 na tabela ASCII

	//int para string
	fmt.Println("Teste " + c.Itoa(123))
	fmt.Println("Teste", 123)

	//string para int
	num, _ := c.Atoi("123") //converte a string "123" para int e atribui a variável num. O _ seria para receber o erro, mas estamos ignorando por enquanto
	fmt.Println(num - 122)

	//string to bool
	b, _ := c.ParseBool("true")
	//b, _ := c.ParseBool("1") //1 mesmo que true, qualquer outro número é false
	if b {
		fmt.Println("Verdadeiro")
	}
}
