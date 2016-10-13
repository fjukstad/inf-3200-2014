package main

import (
	"fmt"
	"math/rand"
)

//START OMIT

func generateRandomNumber(ch chan int) {
	ch <- rand.Intn(100)
}

func randomNumbers(num int) chan int {
	ch := make(chan int)
	for i := 0; i < num; i++ {
		go generateRandomNumber(ch)
	}
	return ch

}

func main() {
	num := 10
	ch := randomNumbers(num)
	for i := 0; i < num; i++ {
		number := <-ch
		fmt.Println("Got number", number)
	}
}

//END OMIT
