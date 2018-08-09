package domain

type Currency struct {
	ID     string
	Name   *string `gorm:"type:varchar;not null"`
	Symbol *string `gorm:"not null"`
}

func NewCurrency(id string, name string, symbol string) *Currency {
	return &Currency{ID: id, Name: &name, Symbol: &symbol}
}
