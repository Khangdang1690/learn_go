package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.Y = 9
	fmt.Println(v.X, v.Y)
}
