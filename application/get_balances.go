package application

import "github.com/opencrypter/api/domain"

type GetBalances struct {
	repository domain.AccountRepository
}

func NewGetBalances(repository domain.AccountRepository) GetBalances {
	return GetBalances{repository}
}

func (g GetBalances) Execute(accountId string) ([]domain.Balance, error) {
	account, err := g.repository.AccountOfId(accountId)
	if err != nil {
		return nil, err
	}
	return account.Balances, nil
}
