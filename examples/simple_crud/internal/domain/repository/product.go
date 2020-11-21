package repository

import "examples/internal/domain"

type IProduct interface {
	Get(id string) (domain.Product, error)
}