package main

import (
	"fmt"
	"learn_go/testpackage"
)

func main() {
	var step int = 5
	var name string = "learn_go"
	var left string = "12"
	var right string = "10"
	var isHappy bool = true
	testpackage.MyFunction(step)
	fmt.Println(name)
	fmt.Println(left + right)
	fmt.Println(isHappy)
}
