package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := new(Vertex)
	v.X = 1
	v.Y = 2
	fmt.Println(v)
}
