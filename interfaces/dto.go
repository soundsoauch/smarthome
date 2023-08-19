package interfaces

type DevicesDTO struct {
	Devices []DeviceDTO `xml:"device"`
}

type DeviceDTO struct {
	Name            string          `xml:"name"`
	Alert           *AlertDTO       `xml:"alert,omitempty"`
	AIN             string          `xml:"identifier,attr"`
	Functionbitmask int             `xml:"functionbitmask,attr"`
	Present         bool            `xml:"present"`
	BatteryLow      *bool           `xml:"batterylow"`
	Temperature     *TemperatureDTO `xml:"temperature"`
	Humidity        *HumidityDTO    `xml:"humidity"`
	HKR             *HKRDTO         `xml:"hkr"`
}

type HKRDTO struct {
	TSoll int `xml:"tsoll"`
}

type HumidityDTO struct {
	RelHumidity int `xml:"rel_humidity"`
}

type TemperatureDTO struct {
	Celsius int `xml:"celsius"`
}

type AlertDTO struct {
	State int `xml:"state"`
}
