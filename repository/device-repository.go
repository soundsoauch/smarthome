package repository

import (
	"errors"
	"smarthome/interfaces"
)

type Repository struct {
	tasmotaDeviceConfigs []TasmotaDeviceConfig
}

type TasmotaDeviceConfig struct {
	hostName string
	deviceType interfaces.Type
}

func NewRepository(config *interfaces.Config) *Repository {
	var tasmotaDeviceConfigs []TasmotaDeviceConfig
	for _, deviceConfig := range config.TasmotaDevices {
		tasmotaDeviceConfigs = append(tasmotaDeviceConfigs, TasmotaDeviceConfig{
			hostName: deviceConfig.Host,
			deviceType: deviceConfig.Type,
		})
	}

	return &Repository{tasmotaDeviceConfigs: tasmotaDeviceConfigs}
}

func(r *Repository) GetDevices(sid string) ([]*interfaces.Device, error) {
	fritzDevices, fritzErr := GetFritzDevices(sid)
	tasmotaDevices, tasmotaErr := GetTasmotaDevices(r.tasmotaDeviceConfigs)
	
	if fritzErr != nil || tasmotaErr  != nil{
		return nil, errors.New(fritzErr.Error() + ", " + tasmotaErr.Error())
	}

	var devices []*interfaces.Device
	for _, fritzDevice := range fritzDevices {
		devices = append(devices, fritzDevice)
	}
	for _, tasmotaDevice := range tasmotaDevices {
		devices = append(devices, tasmotaDevice)
	}

	return devices, nil
}

func(r *Repository) SetDevice(sid string, id string, deviceType string, deviceState *interfaces.DeviceStateDto) error {
	if (deviceType == string(interfaces.PLUG) || deviceType == string(interfaces.CAMERA)) {
		return SetTasmotaState(id, deviceState)
	}

	return SetFritzState(sid, id, deviceState)
}