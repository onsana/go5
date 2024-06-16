package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	nums := make(chan int)
	averages := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			num := rand.Intn(100)
			fmt.Printf("Generated number: %d\n", num)
			nums <- num
			time.Sleep(500 * time.Millisecond)
		}
		close(nums)
	}()

	go func() {
		defer wg.Done()
		var sum, count int
		for num := range nums {
			sum += num
			count++
			average := float64(sum) / float64(count)
			averages <- average
		}
		close(averages)
	}()

	go func() {
		defer wg.Done()
		for avg := range averages {
			fmt.Printf("Current average: %.2f\n", avg)
		}
	}()

	wg.Wait()
}
