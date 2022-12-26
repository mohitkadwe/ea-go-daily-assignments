package main

import (
	"fmt"
	"time"
)

// func main() {
// 	main()

// }

func greet() {
	fmt.Println("Hello World")
}

func main() {

	//Write a Go program to run some go routines which do some work like running a loop etc.,

	// go func() {
	// 	for {
	// 		greet()
	// 	}
	// }()

	go greet()
	go greet()

	time.Sleep(1 * time.Second)

}
