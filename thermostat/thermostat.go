package thermostat

import (
	"fmt"
)

type Thermostat struct {
	temperature   uint8
	roomName      string
	macAddress    string
	hasLowBattery bool
	timestamp     string
}

func New(roomName string, macAddress string) *Thermostat {
	return &Thermostat{roomName: roomName, macAddress: macAddress}
}

func (t *Thermostat) GetValue() string {
	if t.temperature <= 5.0 {
		return "OFF"
	}

	return fmt.Sprintf("%d°C", t.temperature)
}

func (t *Thermostat) SetValue(value uint8) {
	t.temperature = value
}
