package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array.
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println(a2, s3)

	// vc pode imaginar um slice como uma estrutura que tem um tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] // Slice de um slice
	fmt.Println(s2, s4)

	// Mudar o array e ver o reflexo disso nos slices
	a2[0] = 5
	a2[1] = 4
	a2[2] = 3
	a2[3] = 2
	a2[4] = 1

	fmt.Println(a2, s2)
	fmt.Println(a2, s3)
	fmt.Println(s2, s4)
}
