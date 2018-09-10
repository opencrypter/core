package application

import "github.com/opencrypter/core/domain"

type UpdateDeviceSenderId struct {
	repository domain.DeviceRepository
}

func NewUpdateDeviceSenderId(repository domain.DeviceRepository) UpdateDeviceSenderId {
	return UpdateDeviceSenderId{repository}
}

// Updates the sender id of given device id
func (u UpdateDeviceSenderId) Execute(deviceId string, senderId string) error {
	device, err := u.repository.DeviceOfId(deviceId)
	if err != nil {
		return err
	}
	device.SenderId = &senderId
	err = u.repository.Update(device)
	if err != nil {
		return err
	}
	return nil
}
