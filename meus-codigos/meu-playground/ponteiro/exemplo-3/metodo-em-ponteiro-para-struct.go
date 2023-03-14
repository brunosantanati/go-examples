package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// A method on a pointer to a Person struct
func (p *Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old\n", p.Name, p.Age)
}

func main() {
	// Create a new Person struct and initialize it
	person := &Person{
		Name: "John",
		Age:  30,
	}

	// Call the SayHello method on the person pointer
	person.SayHello()
}

/*
Question: Please provide an example using a method on a pointer to a struct, providing also a snipped of code that I can run.

Chat GPT answer: Sure, here's an example of using a method on a pointer to a struct:

package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// A method on a pointer to a Person struct
func (p *Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old\n", p.Name, p.Age)
}

func main() {
    // Create a new Person struct and initialize it
    person := &Person{
        Name: "John",
        Age:  30,
    }

    // Call the SayHello method on the person pointer
    person.SayHello()
}
In this example, we define a Person struct with two fields, Name and Age. We also define a method called SayHello that takes a pointer to a Person struct as a receiver. Inside the method, we use the Name and Age fields of the Person struct to print a greeting message.

In the main function, we create a new Person struct using a pointer and initialize its Name and Age fields. We then call the SayHello method on the person pointer. When we run the program, the output will be:

"Hello, my name is John and I am 30 years old"

Note that in the SayHello method, we use the p pointer to access the Name and Age fields of the Person struct. We don't need to dereference the pointer using the * operator because Go automatically dereferences the pointer for us when we call the method on the p receiver.

###############

Question: Please, explain this part "person := &Person{". Why do we need to use the & symbol?

Chat GPT answer: In the code person := &Person{...}, the & symbol is used to create a pointer to a new Person struct.

The & operator is called the "address of" operator in Go. When applied to a value, it returns the memory address of that value. In this case, the Person{...} syntax creates a new Person struct with the specified field values. The & operator then takes the address of this newly created struct and returns a pointer to it.

By using the & operator, we create a pointer to the Person struct rather than a copy of the struct itself. This is important because if we were to pass the Person struct by value to a function or method, a copy of the struct would be created, and any modifications made to the copy would not affect the original struct.

By using a pointer to the Person struct instead, we can modify the original struct directly and those modifications will be reflected wherever the pointer is used.

In the example code, we use the & operator to create a pointer to a new Person struct, and then we initialize the Name and Age fields of the struct using struct literal syntax. This creates a new Person struct with the specified field values, and returns a pointer to that struct. We then store that pointer in the person variable.
*/
