package model

type Package struct {
	CommentComponent
	types    []*Type
	services []*Service
}

func (p *Package) AddService(service *Service) {
	p.services = append(p.services, service)
}

func (p *Package) Services() []*Service {
	return p.services
}
