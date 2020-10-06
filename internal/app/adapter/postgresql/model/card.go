package model

// Card is the model of cards
type Card struct {
	ID        string    `gorm:"column:card_id;type:uuid"`
	Brand     string    `gorm:"column:brand;type:text"`
	CardBrand CardBrand `gorm:"foreignKey:Brand;references:Brand"`
}
