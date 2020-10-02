package model

// CardBrand is the model of card_brands
type CardBrand struct {
	Brand string `gorm:"type:text;primaryKey"`
}
