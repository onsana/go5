package transport

import "transport/passenger"

type PublicTransport interface {
	AcceptPassengers(p *passenger.Passenger)
	DropPassengers(p *passenger.Passenger)
}
