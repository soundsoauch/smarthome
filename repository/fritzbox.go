package repository

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"smarthome/devices"
	"smarthome/interfaces"
)

func GetFritzDevices(sid string) ([]*interfaces.Device, error) {
	resp, err := http.Get("http://fritz.box/webservices/homeautoswitch.lua?sid=" + sid + "&switchcmd=getdevicelistinfos")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fritzDevices *interfaces.FritzDevicesDTO
	err = xml.Unmarshal([]byte(data), &fritzDevices)
	if err != nil {
		return nil, err
	}

	var filteredDevices = FilterFritzDevices(fritzDevices.Devices)
	var result []*interfaces.Device

	for _, filteredDevice := range filteredDevices {
		var fritzDevice *devices.FritzDevice = devices.NewFritzDevice(filteredDevice)
		
		result = append(result, fritzDevice.GetDevice())
	}

	return result, nil
}

func FilterFritzDevices(devices []interfaces.FritzDeviceDTO) []*interfaces.FritzDeviceDTO {
	var (
		filtered        = []*interfaces.FritzDeviceDTO{}
		functionbitmask int
		isAVMMeasure    bool
		isAVMThermostat bool
		isAVMButton		bool
		isHANFUN        bool
	)

	for i := range devices {
		functionbitmask = devices[i].Functionbitmask
		isAVMMeasure = functionbitmask&(1<<8) != 0
		isAVMThermostat = functionbitmask&(1<<6) != 0
		isAVMButton = functionbitmask&(1<<5) != 0
		isHANFUN = functionbitmask&(1<<13) != 0

		if isAVMMeasure || isAVMThermostat || isHANFUN {
			filtered = append(filtered, &devices[i])
		}
		if isAVMButton && !isAVMThermostat {
			filtered = append(filtered, &devices[i])
		}
	}

	return filtered
}