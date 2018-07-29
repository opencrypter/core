package application

import "github.com/opencrypter/api/domain"

type GetAllAccounts struct {
	repository domain.AccountRepository
}

func NewGetAllAccounts(repository domain.AccountRepository) GetAllAccounts {
	return GetAllAccounts{repository}
}

func (g GetAllAccounts) Execute(deviceId string) []domain.Account {
	return g.repository.AllOfDeviceId(deviceId)
}
