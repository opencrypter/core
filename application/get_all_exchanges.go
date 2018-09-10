package application

import "github.com/opencrypter/core/domain"

type GetAllExchanges struct {
	repository domain.ExchangeRepository
}

func NewGetAllExchanges(repository domain.ExchangeRepository) GetAllExchanges {
	return GetAllExchanges{repository}
}

func (g GetAllExchanges) Execute() []domain.Exchange {
	return g.repository.All()
}
