package main

import "fmt"

func main() {
	// explicit declaration and assignment in multiple lines
	var integerVariable int
	integerVariable = 4
	fmt.Println(integerVariable)

	// explicit declaration
	var floatVariable float32 = 3.14
	fmt.Println(floatVariable)

	// implicit declaration
	firstName := "Fred"
	fmt.Println(firstName)

	// implicit declaration
	booleanVariable := true
	fmt.Println(booleanVariable)

	// complex number
	complexVariable := complex(3, 4)
	fmt.Println(complexVariable)

	// multiple declaration in one line
	realValue, ImaginaryValue := real(complexVariable), imag(complexVariable)
	fmt.Println(realValue, ImaginaryValue)
}
