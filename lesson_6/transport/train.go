package transport

import (
	"fmt"
	"transport/passenger"
)

type Train struct{}

func NewTrain() *Train {
	return &Train{}
}

func (t *Train) AcceptPassengers(p *passenger.Passenger) {
	fmt.Printf("Train: Accepting passenger %s\n", p.GetName())
}

func (t *Train) DropPassengers(p *passenger.Passenger) {
	fmt.Printf("Train: Dropping passenger %s\n", p.GetName())
}
