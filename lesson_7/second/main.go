package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MinMax struct {
	Min int
	Max int
}

func main() {
	nums := make(chan int)
	results := make(chan MinMax)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			num := rand.Intn(100)
			fmt.Printf("Generated number: %d\n", num)
			nums <- num
			time.Sleep(500 * time.Millisecond)
		}
		close(nums)

		result := <-results
		fmt.Printf("Min number: %d, Max number: %d\n", result.Min, result.Max)
	}()

	go func() {
		defer wg.Done()
		var min, max int
		first := true
		for num := range nums {
			if first {
				min, max = num, num
				first = false
			} else {
				if num < min {
					min = num
				}
				if num > max {
					max = num
				}
			}
		}
		results <- MinMax{Min: min, Max: max}
		close(results)
	}()

	wg.Wait()
}
