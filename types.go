package main

import(
	"fmt"
)

var y string

func main(){
	fmt.Println("print the value of y", y, "ending")
	fmt.Println("%T\n", y)

	y = "Ballad"
	fmt.Println(y)
	fmt.Println("%T", y)
}