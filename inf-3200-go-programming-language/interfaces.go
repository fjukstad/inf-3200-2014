package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// START OMIT

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	v := Vertex{3, 4}

	a = v

	fmt.Println(a.Abs())
}

// END OMIT
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
