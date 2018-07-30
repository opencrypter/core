package application

import "github.com/opencrypter/api/domain"

type SecureAccountDto struct {
	Id         string
	ExchangeId *string
	Name       *string
}

type GetAllAccounts struct {
	repository domain.AccountRepository
}

func NewGetAllAccounts(repository domain.AccountRepository) GetAllAccounts {
	return GetAllAccounts{repository}
}

func (g GetAllAccounts) Execute(deviceId string) []SecureAccountDto {
	accounts := g.repository.AllOfDeviceId(deviceId)

	secureAccounts := make([]SecureAccountDto, len(accounts))
	for i := 0; i < len(secureAccounts); i++ {
		secureAccounts[i].Id = accounts[i].ID
		secureAccounts[i].ExchangeId = accounts[i].ExchangeId
		secureAccounts[i].Name = accounts[i].Name
	}
	return secureAccounts
}
