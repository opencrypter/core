package application

import "github.com/opencrypter/api/domain"

type SaveAccount struct {
	repository domain.AccountRepository
}

func NewSaveAccount(repository domain.AccountRepository) SaveAccount {
	return SaveAccount{repository}
}

func (c SaveAccount) Execute(
	id string,
	deviceId *string,
	exchangeId *string,
	name *string,
	apiKey *string,
	apiSecret *string,
) error {
	account := &domain.Account{
		ID:         id,
		DeviceId:   deviceId,
		ExchangeId: exchangeId,
		Name:       name,
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
	}

	return c.repository.Save(account)
}
