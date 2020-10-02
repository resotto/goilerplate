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
