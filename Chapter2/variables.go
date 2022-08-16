package main

import "fmt"

func main() {
	// Many ways of declaring variables
	var name string
	var age int

	var (
		name1 string = "Miguel"
		age1  int    = 23
	)

	// The type can be inferred
	var (
		name2 = "Miguel"
		age2  = 20
	)

	// Inline multiple initialization
	var name3, age3 = "Miguel", 23

	// Instead of var we could use :=
	name4 := "Miguel"
	age4 := 23

	fmt.Print(name, name1, name2, name3, name4)
	fmt.Print(age, age1, age2, age3, age4)
}
