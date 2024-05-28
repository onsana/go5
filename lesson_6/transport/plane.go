package transport

import (
	"fmt"
	"transport/passenger"
)

type Plane struct{}

func NewPlane() *Plane {
	return &Plane{}
}

func (p *Plane) AcceptPassengers(pas *passenger.Passenger) {
	fmt.Printf("Plane: Accepting passenger %s\n", pas.GetName())
}

func (p *Plane) DropPassengers(pas *passenger.Passenger) {
	fmt.Printf("Plane: Dropping passenger %s\n", pas.GetName())
}
