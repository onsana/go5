package main

import "fmt"

type zooKeeper struct {
	name  string
	cages [5]cage
}

type animal struct {
	id   uint
	name string
	age  int
}

type cage struct {
	id     uint
	status string
	animal animal
}

func main() {
	zooKeeper := zooKeeper{name: "Jonh"}
	for i := 0; i < 5; i++ {
		zooKeeper.cages[i] = cage{
			id:     uint(i + 1),
			status: "Open",
		}
	}

	var animals [5]animal
	for i := 0; i < 5; i++ {
		animals[i] = animal{
			id:   uint(i + 1),
			name: "Monkey",
			age:  i + 1,
		}
	}

	fmt.Printf("5 monkeys escaped from the zoo: %+v\n", animals)
	fmt.Printf("All cages left open \n")
	fmt.Printf("5 monkeys started to catch the zookeeper %s\n", zooKeeper.name)

	zooKeeper.catchAnimal(&animals)
}

func (z zooKeeper) catchAnimal(animals *[5]animal) {
	for i := 0; i < 5; i++ {
		z.cages[i].animal = animals[i]
		z.cages[i].status = "Lock"
		fmt.Printf("Animal number %v was caught and locked in a cage\n", animals[i].id)
	}
}
