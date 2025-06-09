package testpackage

import "fmt"

func MyFunction(step int) {
	if step == 1 {
		fmt.Println("step 1")
	} else if step == 2 {
		fmt.Println("step 2")
	} else {
		fmt.Println("step not supported")
	}
	fmt.Println("Hello Go!")
	myPrivateFunction()
}

func myPrivateFunction() {
	fmt.Println("Hello Bro!")
}
