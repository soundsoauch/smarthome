package devices

import "smarthome/interfaces"

type TasmotaDevice struct {
	device    *interfaces.TasmotaDeviceDto
}

func NewTasmotaDevice(device *interfaces.TasmotaDeviceDto) *TasmotaDevice {
	return &TasmotaDevice{device: device}
}

func(d *TasmotaDevice) GetDevice() *interfaces.Device {
	return &interfaces.Device{
		Name: d.device.Status.FriendlyName[0],
		ID: d.device.StatusNET.HostName,
		IsConnected: true,
		Type: interfaces.PLUG,
		State: interfaces.DeviceState{
			Alert: GetPowerState(d.device.Status.Power),
			Energy: (*interfaces.Energy)(d.device.StatusSNS.Energy),
		},
	}
}

func GetPowerState(power int) *interfaces.AlertState {
	var ON = interfaces.ON
	var OFF = interfaces.OFF

	if power == 1 {
		return &ON
	}

	return &OFF
}