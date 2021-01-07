package main

import (
	"fmt"

	"github.com/brunosantanati/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	// fmt.Println(area._TrianguloEq(5.0, 2.0)) // Função privada, não consigo usá-la a partir de outro pacote
}
