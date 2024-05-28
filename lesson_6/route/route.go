package route

import "transport/transport"

type Route struct {
	transports []transport.PublicTransport
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) AddTransport(t transport.PublicTransport) {
	r.transports = append(r.transports, t)
}

func (r *Route) GetTransports() []transport.PublicTransport {
	return r.transports
}
