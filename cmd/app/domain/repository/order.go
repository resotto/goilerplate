package repository

import "github.com/resotto/goilerplate/cmd/app/domain"

// IOrder is interface of order repository
type IOrder interface {
	Get() domain.Order
	Update(domain.Order)
}
