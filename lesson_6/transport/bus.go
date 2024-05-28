package transport

import (
	"fmt"
	"transport/passenger"
)

type Bus struct{}

func NewBus() *Bus {
	return &Bus{}
}

func (b *Bus) AcceptPassengers(p *passenger.Passenger) {
	fmt.Printf("Bus: Accepting passenger %s\n", p.GetName())
}

func (b *Bus) DropPassengers(p *passenger.Passenger) {
	fmt.Printf("Bus: Dropping passenger %s\n", p.GetName())
}
