package interfaces

type Type string

const (
	THERMOSTAT         Type = "THERMOSTAT"
	SENSOR             Type = "SENSOR"
	THERMOMETER        Type = "THERMOMETER"
	BUTTON 			   Type = "BUTTON"
	PLUG               Type = "PLUG"
	UNKNOWN            Type = "UNKNOWN"
)

type AlertState string
const (
	OPEN  AlertState = "OPEN"
	CLOSE AlertState = "CLOSE"
	ON AlertState = "ON"
	OFF AlertState = "OFF"
)

type Device struct {
	Name          string `json:"name"`
	ID            string `json:"id"`
	IsConnected   bool   `json:"isConnected"`
	HasLowBattery *bool  `json:"batterylow,omitempty"`
	Type          Type   `json:"type"`
	State         DeviceState `json:"state"`
}

type DeviceState struct {
	Humidity      *int   `json:"humidity,omitempty"`
	Temperature   *int   `json:"temperature,omitempty"`
	TSoll         *int   `json:"tSoll,omitempty"`
	Alert         *AlertState `json:"alert,omitempty"`
	Energy        *Energy `json:"energy,omitempty"`
}

type Energy struct {
	Total     float32 `json:"Total"`
	Yesterday float32 `json:"Yesterday"`
	Today     float32 `json:"Today"`
	Power     float32 `json:"Power"`
}