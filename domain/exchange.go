package domain

type Exchange struct {
	ID   string
	Name *string `gorm:"type:varchar;not null"`
	Tag  *string `gorm:"type:varchar;not null"`
}

func NewExchange(id string, name string, tag string) *Exchange {
	return &Exchange{ID: id, Name: &name, Tag: &tag}
}

type ExchangeRepository interface {
	Save(exchange *Exchange) error
	All() []Exchange
	ExchangeOfId(id string) (*Exchange, error)
}
