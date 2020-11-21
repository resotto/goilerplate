package repository

import "examples/internal/domain"

type Order struct{}

func (o Order) Save(_ domain.Order) error {
	// omitted here for now
	return nil
}
