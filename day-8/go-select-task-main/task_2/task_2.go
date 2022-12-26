package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1. Find the issue with below code. Understand the root cause of it.
2. Fix the issue with using select `cancellation` or `timeout`.
*/
func main() {

	newRandStream := func() <-chan int {
		randStream := make(chan int)

		select {
		case <-randStream:
			return randStream
		case <-time.After(1 * time.Millisecond):
			go func() {
				defer fmt.Println("newRandStream closure exited.") // this will never get printed.
				defer close(randStream)
				for {
					randStream <- rand.Int()
				}
			}()
		}
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		select {
		case <-randStream:
			fmt.Printf("%d: %d \n", i, <-randStream)
		case <-time.After(10 * time.Millisecond):
			err := fmt.Errorf("timeout")
			fmt.Println(err)
		}
	}

}
