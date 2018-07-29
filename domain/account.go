package domain

type Account struct {
	ID         string  `gorm:"primary_key;type:uuid";json:"id"`
	ExchangeId *string `gorm:"type:uuid";json:"exchangeId"`
	Name       *string `gorm:"type:varchar;not null";json:"name"`
	ApiKey     *string `gorm:"not null";json:"apiKey"`
	ApiSecret  *string `gorm:"not null";json:"apiSecret"`
}

type AccountRepository interface {
	Save(account *Account) error
	AccountOfId(id string) (*Account, error)
}
