package repository

import "examples/internal/domain"

type ICustomer interface {
	Get(id string) (domain.Customer, error)
}