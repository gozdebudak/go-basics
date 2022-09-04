package main

import "fmt"

// * A variable can be declared and initialized out of the function
var num = 4

// ! num := 4
// ! expected declaration, found num

// ! Remove the funtion name to main before run
func main() {
	// Printing the variable that declared out of function
	fmt.Println(num)

	//declaration with variable type
	var message1 string
	//initialization
	message1 = "Hello Go!"

	fmt.Println(message1)

	//declaration WITH variable type and initialization
	var message2 string = "Hello Go!"

	fmt.Println(message2)

	//declaration WITHOUT variable type and initialization
	// * If you don't know the type of the variable, you don't have to declare its type.
	var message3 = "Hello Go!"

	fmt.Println(message3)

	// declaration of multiple variables at the same time WITHOUT variable type
	// * If the variables don't be initialized, 0 will be passed to variables as default value
	var a, b, c int

	// * Printf does not add a new line so "\n" needs to be added
	fmt.Printf("a: %d b: %d c: %d\n", a, b, c) // a: 0 b: 0 c: 0

	// ! var k, l, m = 3
	// ! cannot initialize 3 variables with 1 values

	// declaration of multiple variables at the same time WITH variable type
	var d, e, f int = 3, 4, 5

	fmt.Printf("a: %d b: %d c: %d\n", d, e, f) // a: 3 b: 4 c: 5

	// declaration of multiple variables which have different data types at the same time
	var x, y, z = 3, true, 4.5

	fmt.Printf("x: %d, y: %t, z: %g\n", x, y, z) // x: 3, y: true, z: 4.5

	// * If you don't want use "var" word when declaring a variable, you can use ":=" to declare and itialize the variable
	u := 45.6
	v, n := "abc", false

	fmt.Println("u:", u)
	fmt.Printf("v: %s, n: %t\n", v, n)
}
