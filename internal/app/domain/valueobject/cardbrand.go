package valueobject

import (
	"strings"
)

// CardBrand is credit card brands
type CardBrand string

// VISA is one of CardBrand
const (
	VISA CardBrand = "VISA"
	AMEX CardBrand = "AMEX"
)

// ConvertToCardBrand converts string to CardBrand
func ConvertToCardBrand(s string) CardBrand {
	switch strings.ToUpper(s) {
	case "VISA":
		return VISA
	case "AMEX":
		return AMEX
	default:
		panic("Invalid CardBrand")
	}
}
