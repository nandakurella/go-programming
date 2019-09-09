package main

import "fmt"

//To declare varibales outside function body. Scope is through out the file
var z=35

//if no explicit initialization is done, ZERO value of it type is assigned by default
var a int
var str string
var b bool
var c=5 // declared but not used.Its fine

func main() {
	//short declaration operator
	//declaration and definition. any unused are treated as errors.
	x := 42
	fmt.Println(x)
	//x variable datatype cant be changed
	
	fmt.Println(a) // Print 0
	fmt.Println(b) // Prints false
	fmt.Println(str) // Prints empty string
	
	foo()
}

func foo(){
fmt.Println(z)
}
