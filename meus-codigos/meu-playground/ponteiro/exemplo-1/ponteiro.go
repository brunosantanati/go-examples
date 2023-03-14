package main

import "fmt"

// This function takes a pointer to an integer as a parameter
func incrementByPointer(numPtr *int) {
	*numPtr = *numPtr + 1
}

func main() {
	num := 5
	fmt.Println("Before incrementing:", num)

	// Pass the address of the num variable to the incrementByPointer function
	incrementByPointer(&num)
	fmt.Println("After incrementing:", num)
}
