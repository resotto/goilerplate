package factory

import (
	"github.com/resotto/goilerplate/internal/app/domain"
	"github.com/resotto/goilerplate/internal/app/domain/valueobject"
)

// Order is the factory of domain.Order
type Order struct{}

// Generate generates domain.Order from primitives
func (of Order) Generate(
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
