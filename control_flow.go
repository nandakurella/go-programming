package main

import "fmt"

func foo() {
	fmt.Println("Iam in foo")
}

func main() {

	n,e := fmt.Println("In main")
	foo()
	fmt.Println(n,e)
	for i := 0; i <= 10; i++ { //loop
		if i%2 == 0 { //conditional
			fmt.Println(i)
		}
	}
}



//control flow
//1. sequence
//2. loop; iterative
//3. conditional
