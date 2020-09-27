package repository

import (
	"github.com/resotto/goilerplate/cmd/app/adapter/postgresql"
	"github.com/resotto/goilerplate/cmd/app/adapter/postgresql/model"
	"github.com/resotto/goilerplate/cmd/app/domain"
	"github.com/resotto/goilerplate/cmd/app/domain/factory"
)

// Order is the repository of domain.Order
type Order struct{}

// Get gets order
func (o Order) Get() domain.Order {
	db := postgresql.Connection()
	var order model.Order
	result := db.First(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	// Order has Person/Payment relation and Payment has Card relation which has CardBrand relation.
	db.Preload("Person").Preload("Payment.Card.CardBrand").Find(&order)

	orderFactory := factory.OrderFactory{}
	return orderFactory.Generate(
		order.Person.ID,
		order.Person.Name,
		order.Person.Weight,
		order.Payment.Card.ID,
		order.Payment.Card.CardBrand.Brand,
		order.ID,
	)
}

// Save saves order
func (o Order) Save(order domain.Order) {
	// TODO
}
