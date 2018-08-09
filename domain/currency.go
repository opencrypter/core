package domain

type Currency struct {
	ID     string
	Name   *string `gorm:"type:varchar;not null"`
	Symbol *string `gorm:"not null"`
}
