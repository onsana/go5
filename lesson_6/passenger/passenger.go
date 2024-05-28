package passenger

type Passenger struct {
	name string
}

func NewPassenger(name string) *Passenger {
	return &Passenger{name: name}
}

func (p *Passenger) GetName() string {
	return p.name
}
