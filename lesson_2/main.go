package main

import "fmt"

type zooKeeper struct {
	name string
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
	var animals [5]animal
	for i := 0; i < 5; i++ {
		animals[i] = animal{
			id:   uint(i + 1),
			name: "Monkey",
			age:  i + 1,
		}
	}
	var cages [5]cage
	for i := 0; i < 5; i++ {
		cages[i] = cage{
			id:     uint(i + 1),
			status: "Open",
		}
	}

	fmt.Printf("5 monkeys escaped from the zoo: %+v\n", animals)
	fmt.Printf("All cages left open: %+v\n", cages)
	fmt.Printf("5 monkeys started to catch the zookeeper %s\n", zooKeeper.name)

	for i := 0; i < 5; i++ {
		catchMonkey(&animals[i], &cages[i])
	}

}

func catchMonkey(animal *animal, cage *cage) {
	cage.animal = *animal
	cage.status = "Lock"
	fmt.Printf("Monkey number %v was caught and locked in a cage\n", animal.id)
}
