// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

type example11 struct {
	flag    bool
	counter int16
	pi      float32
}
type example22 struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	var ex1 example11
	var ex2 example22
	//TODO: variable assigning
	fmt.Printf("ex1: %v ex2: %v", ex1, ex2)

	//ex1 = ex2 // Compile error type example22 can't be assigned to ex1 of type example11

	ex1 = example11(ex2) // conversion possible for same structure

	ex3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	ex1 = ex3 // unnamed type can bee assigned no compilation error

}
