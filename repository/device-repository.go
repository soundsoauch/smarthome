package repository

import (
	"errors"
	"smarthome/interfaces"
)

type Repository struct {
	tasmotaDeviceHostnames []string
}

func NewRepository() *Repository {
	return &Repository{
		tasmotaDeviceHostnames: []string{
			"Deko", "camera", "camera-furbo", "Dimmer", "BedroomLeft", "BedroomRight", "Kitchen",
		},
	}
}

func(r *Repository) GetDevices(sid string) ([]*interfaces.Device, error) {
	fritzDevices, fritzErr := GetFritzDevices(sid)
	tasmotaDevices, tasmotaErr := GetTasmotaDevices(r.tasmotaDeviceHostnames)
	
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
	if (deviceType == string(interfaces.PLUG)) {
		return SetTasmotaState(id, deviceState)
	} else {
		// delegate fritzbox
	}

	return nil
}