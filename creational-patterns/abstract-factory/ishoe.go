package main

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getSize() int
	getLogo() string
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

func (s *Shoe) getLogo() string {
	return s.logo
}
