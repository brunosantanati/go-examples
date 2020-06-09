package main

import (
	"fmt"
	mathematics "math"
)

func main() {

	//declarar constante e variavel
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	fmt.Println("PI", PI)
	fmt.Println("raio", raio)

	// forma reduzida de criar uma variavel
	area := PI * mathematics.Pow(raio, 2)

	fmt.Println("A area da circunferencia eh", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
