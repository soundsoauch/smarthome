package devices

import (
	"math"
	"smarthome/interfaces"
)

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
		IP: d.device.StatusNET.IPAddress,
		IsConnected: true,
		Type: d.device.Type,
		State: interfaces.DeviceState{
			Alert: GetPowerState(d.device.Status.Power),
			Energy: (*interfaces.Energy)(d.device.StatusSNS.Energy),
			Temperature: GetTasmotaTemperature(d.device.StatusSNS.Temperature),
			Humidity: GetTasmotaHumidity(d.device.StatusSNS.Temperature),
		},
	}
}

func GetTasmotaTemperature(temperature *interfaces.TasmotaDeviceTemperatureDto) *int {
	if temperature == nil {
		return nil
	}
	// round to next .5 and and multiplicate 10
	val := int(math.Round(float64(temperature.Temperature)/0.5) * 5)

	return &val
}

func GetTasmotaHumidity(temperature *interfaces.TasmotaDeviceTemperatureDto) *int {
	if temperature == nil {
		return nil
	}

	val := int(math.Round(float64(temperature.Humidity)))

	return &val
}

func GetPowerState(power int) *interfaces.AlertState {
	var ON = interfaces.ON
	var OFF = interfaces.OFF

	if power == 1 {
		return &ON
	}

	return &OFF
}