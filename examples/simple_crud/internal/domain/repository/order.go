package repository

import "examples/internal/domain"

type IOrder interface {
	Save(order domain.Order) error
}