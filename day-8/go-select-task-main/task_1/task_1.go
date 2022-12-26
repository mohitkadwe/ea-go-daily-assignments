package main

import (
	"fmt"
	"time"
)

/*
1. Find the issue with below code. Understand the root cause of it.
2. Fix the issue with using select `cancellation` or `timeout`.
*/
func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		select {
		case <-time.After(1 * time.Millisecond):
			go func() {
				defer fmt.Println("doWork exited.")
				defer close(completed)
				for s := range strings {
					// do something interesting
					fmt.Println(s)
				}
			}()

		case <-completed:
			return completed
		}

		return completed
	}

	ch := make(chan string)
	go func() {
		// Close the channel to signal that no more values will be sent.
		defer close(ch)
		ch <- "hello"
		ch <- "world"
	}()

	doWork(ch)
	fmt.Println("Done")
}
