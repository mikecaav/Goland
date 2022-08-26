package main

import "fmt"

type FunctionType func(text string)

func (functionInstance FunctionType) methodOfInstancedFunction(text string) {
	functionInstance(text)
}

func mainFunction(text string) {
	fmt.Println(text)
}

func main() {
	FunctionType(mainFunction).methodOfInstancedFunction("Wow esto es increíble, puedo hacer cosas raras")
}
