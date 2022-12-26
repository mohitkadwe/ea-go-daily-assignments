package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int // This is the variable to be incremented
}

func (numCounter *counter) Add(inpNumber int) {
	numCounter.count += inpNumber
}

// A simple program to trigger 10 goroutines
// & each to add 100 to Counter via iteration one by one
func main() {

	var numCounter counter

	group := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				numCounter.Add(1)
			}
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println(numCounter.count)
}
