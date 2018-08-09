package domain

type Account struct {
	ID         string  `gorm:"primary_key;type:uuid"`
	DeviceId   *string `gorm:"type:uuid;not null"`
	ExchangeId *string `gorm:"type:uuid;not null"`
	Name       *string `gorm:"type:varchar;not null"`
	ApiKey     *string `gorm:"not null"`
	ApiSecret  *string `gorm:"not null"`
	Balances   []Balance
}

type Balance struct {
	ID         string
	AccountID  *string `gorm:"type:uuid;not null"`
	Currency   *Currency
	CurrencyID *string  `gorm:"type:uuid;not null"`
	Volume     *float64 `gorm:"not null"`
	HasAlert   *bool    `gorm:"has_alert:false;not null"`
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

func NewBalance(id string, accountId string, currencyId string, volume float64, hasAlert bool) *Balance {
	return &Balance{
		ID:         id,
		AccountID:  &accountId,
		CurrencyID: &currencyId,
		Volume:     &volume,
		HasAlert:   &hasAlert,
	}
}

type AccountRepository interface {
	Save(account *Account) error
	AccountOfId(id string) (*Account, error)
	AllOfDeviceId(deviceId string) []Account
}
