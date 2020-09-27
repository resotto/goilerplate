package factory

import (
	"github.com/resotto/goilerplate/cmd/app/domain"
	"github.com/resotto/goilerplate/cmd/app/domain/valueobject"
)

// OrderFactory is the factory of domain.Order
type OrderFactory struct{}

// Generate generates domain.Order from primitives
func (of OrderFactory) Generate(
	personID string,
	name string,
	weight int,
	cardID string,
	brand string,
	orderID string,
) domain.Order {
	person := domain.Person{
		ID:     personID,
		Name:   name,
		Weight: weight,
	}
	cardBrand := valueobject.ConvertToCardBrand(brand)
	card := valueobject.Card{
		ID:    cardID,
		Brand: cardBrand,
	}
	payment := valueobject.Payment{
		Card: card,
	}
	return domain.Order{
		ID:      orderID,
		Payment: payment,
		Person:  person,
	}
}
