package repository

import "examples/internal/domain"

type Product struct{}

func (p Product) Get(_ string) (domain.Product, error) {
	// omitted here for now
	return domain.Product{}, nil
}