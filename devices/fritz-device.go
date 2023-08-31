package devices

import (
	"smarthome/interfaces"
	"strings"
)

type FritzDevice struct {
	device *interfaces.FritzDeviceDTO
}

func NewFritzDevice(device *interfaces.FritzDeviceDTO) *FritzDevice {
	return &FritzDevice{device: device}
}

func(d *FritzDevice) GetDevice() *interfaces.Device {
	return &interfaces.Device{
		Name: d.device.Name,
		ID: strings.ReplaceAll(d.device.AIN, " ", ""),
		IsConnected: d.device.Present,
		HasLowBattery: d.device.BatteryLow,
		Type: getType(d.device.Functionbitmask),
		State: interfaces.DeviceState{
			Humidity: getHumidity(d.device.Humidity),
			Temperature: getTemperature(d.device.Temperature),
			TSoll: getTSoll(d.device.HKR),
			Alert: getState(d.device.Alert, d.device.Present),
		},
	}
}

func getState(alert *interfaces.FritzAlertDTO, isConnected bool) *interfaces.AlertState {
	if alert == nil || isConnected == false {
		return nil
	}

	if alert.State == 0 {
		var CLOSE = interfaces.CLOSE
		return &CLOSE
	}

	var OPEN = interfaces.OPEN
	return &OPEN
}

func getTSoll(hkr *interfaces.FritzHKRDTO) *int {
	if hkr != nil && hkr.TSoll >= 16 && hkr.TSoll <= 56 {
		tsoll := hkr.TSoll * 10
		return &tsoll
	}

	return nil
}

func getTemperature(temperature *interfaces.FritzTemperatureDTO) *int {
	if temperature != nil {
		return &temperature.Celsius
	}

	return nil
}

func getHumidity(humidity *interfaces.FritzHumidityDTO) *int {
	if humidity != nil {
		return &humidity.RelHumidity
	}

	return nil
}

func getType(functionbitmask int) interfaces.Type {
    if functionbitmask&(1<<6) != 0 {
		return interfaces.THERMOSTAT
	}

	if functionbitmask&(1<<13) != 0 {
		return interfaces.SENSOR
	}

	if functionbitmask&(1<<5) != 0 && functionbitmask&(1<<8) == 0 {
		return interfaces.BUTTON
	}

	if functionbitmask&(1<<8) != 0 {
		return interfaces.THERMOMETER
	}

	return interfaces.UNKNOWN
}
