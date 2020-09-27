package usecase

import (
	"github.com/google/uuid"
	"github.com/resotto/goilerplate/cmd/app/domain"
	"github.com/resotto/goilerplate/cmd/app/domain/repository"
	"github.com/resotto/goilerplate/cmd/app/domain/valueobject"
)

// SaveChangedState saves order whose state has been changed
func SaveChangedState(o repository.IOrder) domain.Order {
	order := o.Get()
	newCardBrand := valueobject.VISA
	if order.Payment.Card.Brand == valueobject.VISA {
		newCardBrand = valueobject.AMEX
	}
	newCard := valueobject.Card{
		ID:    uuid.New().String(),
		Brand: newCardBrand,
	}
	order.Person.Weight++
	order.Payment.Card = newCard
	o.Save(order)
	return order
}
