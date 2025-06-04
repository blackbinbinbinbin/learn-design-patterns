package main

import (
	"errors"
)

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetISportFactory(name string) (ISportFactory, error) {
	if name == "addidas" {
		return AddidasFactory{}, nil
	}
	if name == "nike" {
		return NikeFactory{}, nil
	}

	return nil, errors.New("not found")
}
