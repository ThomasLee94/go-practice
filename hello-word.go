package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	foo()
	fmt.Println("Print something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("We left main")
}
