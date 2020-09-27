package valueobject

// Card is credit card which has serial number and card brand
type Card struct {
	ID    string
	Brand CardBrand
}
