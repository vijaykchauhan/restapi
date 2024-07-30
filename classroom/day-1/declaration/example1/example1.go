// Sample program to show how to declare variables.
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)
	fmt.Println(unsafe.Sizeof(b))
}
