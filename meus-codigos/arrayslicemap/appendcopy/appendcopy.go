package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) // não funciona pois o primeiro argumento precisa ser um slice
	slice1 = append(slice1, 4, 5, 6) // O aapend se necessário "aumenta" o array interno(substitui ele)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // Copia quantos elementos couberem, nesse caso apenas 2
	fmt.Println(slice2)
}
