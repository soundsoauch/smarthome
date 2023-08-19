package device

import (
	"smarthome/interfaces"
	"strings"
)

type Type string

const (
	THERMOSTAT         Type = "THERMOSTAT"
	SENSOR             Type = "SENSOR"
	TEMPERATURE_SENSOR Type = "TEMPERATURE_SENSOR"
	UNKNOWN            Type = "UNKNOWN"
)

type State string

const (
	OPEN  State = "OPEN"
	CLOSE State = "CLOSE"
)

type Device struct {
	Name          string `json:"name"`
	AIN           string `json:"ain"`
	IsConnected   bool   `json:"isConnected"`
	HasLowBattery *bool  `json:"batterylow,omitempty"`
	Type          Type   `json:"type"`
	Humidity      *int   `json:"humidity,omitempty"`
	Temperature   *int   `json:"temperature,omitempty"`
	TSoll         *int   `json:"tSoll,omitempty"`
	State         *State `json:"state,omitempty"`
}

func New(deviceDto interfaces.DeviceDTO) *Device {
	return &Device{
		Name:          deviceDto.Name,
		AIN:           strings.ReplaceAll(deviceDto.AIN, " ", ""),
		IsConnected:   deviceDto.Present,
		HasLowBattery: deviceDto.BatteryLow,
		Type:          getType(deviceDto.Functionbitmask),
		Humidity:      getHumidity(deviceDto.Humidity),
		Temperature:   getTemperature(deviceDto.Temperature),
		TSoll:         getTSoll(deviceDto.HKR),
		State:         getState(deviceDto.Alert, deviceDto.Present),
	}
}

func getState(alert *interfaces.AlertDTO, isConnected bool) *State {

	if alert == nil || isConnected == false {
		return nil
	}

	if alert.State == 0 {
		close := CLOSE
		return &close
	}

	open := OPEN
	return &open
}

func getTSoll(hkr *interfaces.HKRDTO) *int {
	if hkr != nil && hkr.TSoll >= 16 && hkr.TSoll <= 56 {
		tsoll := hkr.TSoll * 10
		return &tsoll
	}

	return nil
}

func getTemperature(temperature *interfaces.TemperatureDTO) *int {
	if temperature != nil {
		return &temperature.Celsius
	}

	return nil
}

func getHumidity(humidity *interfaces.HumidityDTO) *int {
	if humidity != nil {
		return &humidity.RelHumidity
	}

	return nil
}

func getType(functionbitmask int) Type {
	if functionbitmask&(1<<6) != 0 {
		return THERMOSTAT
	}

	if functionbitmask&(1<<13) != 0 {
		return SENSOR
	}

	if functionbitmask&(1<<8) != 0 {
		return TEMPERATURE_SENSOR
	}

	return UNKNOWN
}
