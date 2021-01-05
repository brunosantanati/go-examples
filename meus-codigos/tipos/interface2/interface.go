package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() { // Receiver com ponteiro para poder alterar
	f.turboLigado = true // Altera valor
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0} // Usar &ferrari
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
