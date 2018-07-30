package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/opencrypter/api/domain"
	"regexp"
)

type GormDeviceRepository struct {
	database *gorm.DB
}

func NewDeviceRepository() *GormDeviceRepository {
	return &GormDeviceRepository{database: Database}
}

func (r GormDeviceRepository) Add(device *domain.Device) error {
	err := r.database.Create(device).Error
	if err == nil {
		return nil
	}
	if ok, _ := regexp.Match("(pq: duplicate key).*", []byte(err.Error())); ok {
		return domain.NewDuplicatedDeviceError(device.ID)
	}
	return domain.NewInvalidDeviceError(err.Error())
}

func (r GormDeviceRepository) DeviceOfId(id string) (*domain.Device, error) {
	var device domain.Device
	err := r.database.Where("id = ?", id).First(&device).Error
	if err != nil {
		return nil, domain.NewDeviceNotFoundError(id)
	}
	return &device, nil
}

func (r GormDeviceRepository) Update(device *domain.Device) error {
	err := r.database.Save(device).Error
	if err != nil {
		return err
	}
	return nil
}
