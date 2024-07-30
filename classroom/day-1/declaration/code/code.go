// Sample program to show how to declare variables.
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// Declare variables that are set to their zero value.
	var a int
	var b string  // string is a 2 word 8byte each pointer and length
	var c float64 // 8 bytes and IEEE 754 binary decimal
	var d bool
	sizeInt := unsafe.Sizeof(10)
	fmt.Printf("size a int \t %d \n", sizeInt)

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true
	// sizeInta := unsafe.Sizeof(aa)
	// fmt.Printf("size aa int \t %d \n", sizeInta)
	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	aaa := int32(10)
	aaaa := int64(aaa)
	aaaaa := int32(aaaa)
	fmt.Printf("aaa := int32(10) %T [%v] %d\n", aaa, aaa, unsafe.Sizeof(aaa))
	fmt.Printf("aaa := int32(10) %T [%v] %d\n", aaaa, aaaa, unsafe.Sizeof(aaaa))
	fmt.Printf("aaa := int32(10) %T [%v] %d\n", aaaaa, aaaaa, unsafe.Sizeof(aaaaa))
}

/*
	Zero Values:

	Type Initialized Value
	Boolean false
	Integer 0
	Floating Point 0
	Complex 0i
	String "" (empty string)
	Pointer nil
*/
