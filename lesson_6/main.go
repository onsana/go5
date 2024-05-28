package main

import (
	"fmt"
	"transport/passenger"
	"transport/route"
	"transport/transport"
)

func main() {
	bus := transport.NewBus()
	train := transport.NewTrain()
	plane := transport.NewPlane()

	r := route.NewRoute()
	r.AddTransport(bus)
	r.AddTransport(train)
	r.AddTransport(plane)

	p := passenger.NewPassenger("John Doe")
	fmt.Println("Starting the journey...")
	for _, vehicle := range r.GetTransports() {
		vehicle.AcceptPassengers(p)
		vehicle.DropPassengers(p)
	}
}
