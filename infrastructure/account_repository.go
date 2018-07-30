package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/opencrypter/api/domain"
)

type GormAccountRepository struct {
	database *gorm.DB
}

func NewAccountRepository() *GormAccountRepository {
	return &GormAccountRepository{database: openDb()}
}

func (r GormAccountRepository) Save(device *domain.Account) error {
	return r.database.Save(device).Error
}

func (r GormAccountRepository) AccountOfId(id string) (*domain.Account, error) {
	var account domain.Account
	err := r.database.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r GormAccountRepository) AllOfDeviceId(deviceId string) []domain.Account {
	accounts := make([]domain.Account, 0)
	r.database.Where("device_id = ?", deviceId).Find(&accounts)
	return accounts
}
