package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/opencrypter/api/domain"
)

type GormExchangeRepository struct {
	database *gorm.DB
}

func NewExchangeRepository() GormExchangeRepository {
	return GormExchangeRepository{database: openDb()}
}

func (r GormExchangeRepository) All() []domain.Exchange {
	exchanges := make([]domain.Exchange, 0)
	r.database.Find(&exchanges)
	return exchanges
}

func (r GormExchangeRepository) Save(exchange *domain.Exchange) error {
	return r.database.Save(exchange).Error
}

func (r GormExchangeRepository) ExchangeOfId(id string) (*domain.Exchange, error) {
	var exchange domain.Exchange
	err := r.database.Where("id = ?", id).First(&exchange).Error
	if err != nil {
		return nil, domain.NewDeviceNotFoundError(id)
	}
	return &exchange, nil
}
