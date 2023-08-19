package deviceList

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"smarthome/device"
	"smarthome/interfaces"
)

func GetDeviceList(sid string) ([]*device.Device, error) {
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

	var devices *interfaces.DevicesDTO
	err = xml.Unmarshal([]byte(data), &devices)
	if err != nil {
		return nil, err
	}

	return FilterDevices(devices.Devices), nil
}

func FilterDevices(devices []interfaces.DeviceDTO) []*device.Device {
	var (
		filtered        = []*device.Device{}
		functionbitmask int
		isAVMMeasure    bool
		isAVMThermostat bool
		isHANFUN        bool
	)

	for i := range devices {
		functionbitmask = devices[i].Functionbitmask
		isAVMMeasure = functionbitmask&(1<<8) != 0
		isAVMThermostat = functionbitmask&(1<<6) != 0
		isHANFUN = functionbitmask&(1<<13) != 0

		if isAVMMeasure || isAVMThermostat || isHANFUN {
			filtered = append(filtered, device.New(devices[i]))
		}
	}

	return filtered
}
