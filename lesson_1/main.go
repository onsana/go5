package main

import (
	"fmt"
	"time"
)

type person struct {
	name     string
	age      int
	from     string
	location string
}

func main() {
	start_time := time.Now()
	ksana := person{name: "Oksana", age: 45, from: "Ukraine", location: "Portugal"}
	ksana.printPerson()
	// Play with numbers
	fmt.Println(multiplication(5, 6))
	var result float32
	result = division(10, 3)
	fmt.Printf("%.2f \n", result)

	// Play with time
	fmt.Println(start_time.Year())
	fmt.Println(start_time.Month())
	fmt.Println(start_time.Day())
	end_time := time.Now().Add(time.Second * 43)
	fmt.Printf("The seconds difference is: %f", end_time.Sub(start_time).Seconds())
}

func multiplication(a int, b int) int {
	return a * b
}

func division(a int, b int) float32 {
	return float32(a) / float32(b)
}

func (p person) printPerson() {
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.from)
	fmt.Println(p.location)
	fmt.Printf("type: %T\n", p)
	hello := "Hello"
	concatenated := hello + ", " + p.name + "!"
	fmt.Println(concatenated)
}
