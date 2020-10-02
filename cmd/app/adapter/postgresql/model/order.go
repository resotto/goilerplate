package model

// Order is the model of orders
type Order struct {
	ID       string  `gorm:"column:order_id;type:uuid"`
	PersonID string  `gorm:"type:uuid"`
	Person   Person  // `gorm:"foreignKey:PersonID;references:ID"`
	Payment  Payment // `gorm:"foreignkey:OrderID;references:ID"`
}
