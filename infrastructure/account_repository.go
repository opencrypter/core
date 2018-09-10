package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/opencrypter/core/domain"
)

type GormAccountRepository struct {
	database *gorm.DB
}

func NewAccountRepository() *GormAccountRepository {
	return &GormAccountRepository{database: Database}
}

func (r GormAccountRepository) Save(account *domain.Account) error {
	err := r.database.Create(account).Error
	if err != nil {
		err = r.database.Save(account).Error
	}
	return err
}

func (r GormAccountRepository) AccountOfId(id string) (*domain.Account, error) {
	var account domain.Account
	err := r.database.Preload("Balances").Preload("Balances.Currency").Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, domain.NewAccountNotFoundError(id)
	}
	return &account, nil
}

func (r GormAccountRepository) AllOfDeviceId(deviceId string) []domain.Account {
	accounts := make([]domain.Account, 0)
	r.database.Where("device_id = ?", deviceId).Find(&accounts)
	return accounts
}
