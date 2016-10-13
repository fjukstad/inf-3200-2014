package main

import (
	"fmt"
	"time"
)

//START OMIT
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v \n %s",
		e.When.Local(), e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

//END OMIT
