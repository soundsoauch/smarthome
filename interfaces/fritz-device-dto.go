package interfaces

type FritzDevicesDTO struct {
	Devices []FritzDeviceDTO `xml:"device"`
}

type FritzDeviceDTO struct {
	Name            string               `xml:"name"`
	Alert           *FritzAlertDTO       `xml:"alert,omitempty"`
	AIN             string               `xml:"identifier,attr"`
	Functionbitmask int                  `xml:"functionbitmask,attr"`
	Present         bool                 `xml:"present"`
	BatteryLow      *bool                `xml:"batterylow"`
	Temperature     *FritzTemperatureDTO `xml:"temperature"`
	Humidity        *FritzHumidityDTO    `xml:"humidity"`
	HKR             *FritzHKRDTO         `xml:"hkr"`
}

type FritzHKRDTO struct {
	TSoll int `xml:"tsoll"`
}

type FritzHumidityDTO struct {
	RelHumidity int `xml:"rel_humidity"`
}

type FritzTemperatureDTO struct {
	Celsius int `xml:"celsius"`
}

type FritzAlertDTO struct {
	State int `xml:"state"`
}
