package domain

type Account struct {
	ID         string  `gorm:"primary_key;type:uuid";json:"id"`
	DeviceId   *string `gorm:"type:uuid;not null";json:"deviceId"`
	ExchangeId *string `gorm:"type:uuid;not null";json:"exchangeId"`
	Name       *string `gorm:"type:varchar;not null";json:"name"`
	ApiKey     *string `gorm:"not null";json:"apiKey"`
	ApiSecret  *string `gorm:"not null";json:"apiSecret"`
}

func NewAccount(id string, deviceId string, exchangeId string, name string, apiKey string, apiSecret string) *Account {
	return &Account{
		ID:         id,
		DeviceId:   &deviceId,
		ExchangeId: &exchangeId,
		Name:       &name,
		ApiKey:     &apiKey,
		ApiSecret:  &apiSecret,
	}
}

type AccountRepository interface {
	Save(account *Account) error
	AccountOfId(id string) (*Account, error)
	AllOfDeviceId(deviceId string) []Account
}
