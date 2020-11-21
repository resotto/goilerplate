package usecase

import (
	"examples/internal/domain"
	"examples/internal/domain/repository"
)

type CreateOrderArgs struct {
	CustomerID         string
	ProductID          string
	CustomerRepository repository.ICustomer
	ProductRepository  repository.IProduct
	OrderRepository    repository.IOrder
}

func CreateOrder(args CreateOrderArgs) domain.Order {
	customer, err := args.CustomerRepository.Get(args.CustomerID)
	if err != nil {
		panic(err)
	}

	product, err := args.ProductRepository.Get(args.ProductID)
	if err != nil {
		panic(err)
	}
	order := domain.Order{
		ID:       "123",
		Customer: customer,
		Product:  product,
	}

	err = args.OrderRepository.Save(order)
	if err != nil {
		panic(err)
	}
	
	return order
}
