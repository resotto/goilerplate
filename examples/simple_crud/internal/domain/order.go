package domain

type Order struct {
	ID string
	Customer Customer
	Product Product
}