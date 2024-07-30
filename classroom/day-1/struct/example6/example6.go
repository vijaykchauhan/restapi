// Sample program to show how to declare and initialize struct types.
package main

import (
	"fmt"
	"unsafe"
)

type example1 struct {
	flag    bool
	counter int16
	pi      float32
}

type example2 struct {
	flag    bool
	counter int16
	flag2   bool
	pi      float32
}
type example3 struct {
	pi      float32
	counter int16
	flag    bool
	flag2   bool
}

func main() {

	var e1 example1
	fmt.Printf("example: %T [%v] %d\n", e1, e1, unsafe.Sizeof(e1))
	fmt.Println("Memory address of e1")
	fmt.Printf("flag: %d\n", &e1.flag)
	fmt.Printf("counter: %d\n", &e1.counter)
	fmt.Printf("pi: %d\n", &e1.pi)

	var e2 example2
	fmt.Printf("example: %T [%v] %d\n", e2, e2, unsafe.Sizeof(e2))

	var e3 example3
	fmt.Printf("example: %T [%v] %d\n", e3, e3, unsafe.Sizeof(e3))

}
