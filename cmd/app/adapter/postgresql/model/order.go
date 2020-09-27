package model

// Tabler is interface of GORM table name
type Tabler interface {
	TableName() string
}

// Person is the model of persons
// By default, GORM sees its table name as pluralized struct name.
// Given struct name Person, its table name will be interpreted as "people", NOT "persons".
// If you want "persons" as table name, you have to implement Tabler.TableName().
type Person struct {
	ID     string `gorm:"column:person_id;type:uuid"`
	Name   string `gorm:"type:text;not null"`
	Weight int
}

// TableName gets table name of Person
func (Person) TableName() string {
	return "persons"
}

// CardBrand is the model of card_brands
type CardBrand struct {
	Brand string `gorm:"type:text;primaryKey"`
}

// Card is the model of cards
type Card struct {
	ID        string    `gorm:"column:card_id;type:uuid"`
	Brand     string    `gorm:"column:brand;type:text"`
	CardBrand CardBrand `gorm:"foreignKey:Brand;references:Brand"`
}

// Payment is the model of payments
type Payment struct {
	OrderID string `gorm:"type:uuid;primaryKey"`
	CardID  string `gorm:"type:uuid"`
	Card    Card
}

// Order is the model of orders
type Order struct {
	ID       string  `gorm:"column:order_id;type:uuid"`
	PersonID string  `gorm:"type:uuid"`
	Person   Person  // `gorm:"foreignKey:PersonID;references:ID"`
	Payment  Payment // `gorm:"foreignkey:OrderID;references:ID"`
}
