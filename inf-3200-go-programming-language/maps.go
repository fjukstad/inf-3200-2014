package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Tromsø"] = Vertex{69.6828, 18.9428}
	fmt.Println(m["Tromsø"])
}
