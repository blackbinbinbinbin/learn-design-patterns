package main

import (
	"errors"
)

func GetGuiFactory(guiType string) (IGui, error) {
	if guiType == "window" {
		return &WindowGui{}, nil
	}
	if guiType == "mac" {
		return &MacGui{}, nil
	}

	return nil, errors.New("not found gui")
}
