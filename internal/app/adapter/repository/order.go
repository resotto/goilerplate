package repository

import (
	"errors"

	"github.com/resotto/goilerplate/internal/app/adapter/postgresql"
	"github.com/resotto/goilerplate/internal/app/adapter/postgresql/model"
	"github.com/resotto/goilerplate/internal/app/domain"
	"github.com/resotto/goilerplate/internal/app/domain/factory"
	"gorm.io/gorm"
)

// Order is the repository of domain.Order
type Order struct{}

// Get gets order
func (o Order) Get() domain.Order {
	db := postgresql.Connection()
	var order model.Order
	// Order has Person/Payment relation and Payment has Card relation which has CardBrand relation.
	result := db.Preload("Person").Preload("Payment.Card.CardBrand").Find(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	orderFactory := factory.Order{}
	return orderFactory.Generate(
		order.Person.ID,
		order.Person.Name,
		order.Person.Weight,
		order.Payment.Card.ID,
		order.Payment.Card.CardBrand.Brand,
		order.ID,
	)
}

// Update updates order
func (o Order) Update(order domain.Order) {
	db := postgresql.Connection()
	card := model.Card{
		ID:    order.Payment.Card.ID,
		Brand: string(order.Payment.Card.Brand),
	}
	payment := model.Payment{
		OrderID: order.ID,
		CardID:  card.ID,
		Card:    card,
	}
	person := model.Person{
		ID:     order.Person.ID,
		Name:   order.Person.Name,
		Weight: order.Person.Weight,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Exec("update persons set name = ?, weight = ? where person_id = ?", person.Name, person.Weight, person.ID).Error
		if err != nil {
			return errors.New("rollback")
		}
		err = tx.Exec("insert into cards values (?, ?)", card.ID, card.Brand).Error
		if err != nil {
			return errors.New("rollback")
		}
		err = tx.Exec("update payments set card_id = ? where order_id = ?", payment.CardID, payment.OrderID).Error
		if err != nil {
			return errors.New("rollback")
		}
		return nil // commit
	})
	if err != nil {
		panic(err)
	}
}
