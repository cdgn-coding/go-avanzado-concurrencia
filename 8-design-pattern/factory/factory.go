package main

import "fmt"

func GetComputerFactory(computerType string) (Product, error) {
	if computerType == "laptop" {
		return NewLaptop(), nil
	}

	if computerType == "desktop" {
		return NewDesktop(), nil
	}

	return nil, fmt.Errorf("invalid computer type %s", computerType)
}
