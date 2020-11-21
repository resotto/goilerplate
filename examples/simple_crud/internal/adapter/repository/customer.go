package repository

import "examples/internal/domain"

type Customer struct{}

func (c Customer) Get(_ string) (domain.Customer, error) {
	// omitted here for now
	return domain.Customer{}, nil
}