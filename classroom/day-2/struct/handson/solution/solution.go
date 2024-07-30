// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	vc := user{
		name:  "Vijay",
		email: "vijayc@mkcl.org",
		age:   42,
	}

	// Display the field values.
	fmt.Println("Name", vc.name)
	fmt.Println("Email", vc.email)
	fmt.Println("Age", vc.age)

	// Declare a variable using an anonymous struct.
	ps := struct {
		name  string
		email string
		age   int
	}{
		name:  "Pranab",
		email: "pranabs@mkcl.org",
		age:   45,
	}

	// Display the field values.
	fmt.Println("Name", ps.name)
	fmt.Println("Email", ps.email)
	fmt.Println("Age", ps.age)
}
