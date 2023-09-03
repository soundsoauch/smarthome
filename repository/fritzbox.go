package repository

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"smarthome/devices"
	"smarthome/interfaces"
	"strconv"
	"strings"
)

func GetFritzDevices(sid string) ([]*interfaces.Device, error) {
	resp, err := http.Get("http://fritz.box/webservices/homeautoswitch.lua?sid=" + sid + "&switchcmd=getdevicelistinfos")

	if err != nil {
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

func SetFritzState(sid string, id string, deviceState *interfaces.DeviceStateDto) error {
	var tSoll string = getTSoll(deviceState.TSoll)
	var ain string = strings.ReplaceAll(id, " ", "%20")
	resp, err := http.Get("http://fritz.box/webservices/homeautoswitch.lua?sid=" + sid + "&ain=" + ain + "&switchcmd=sethkrtsoll&param=" + tSoll)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func getTSoll(tSoll *int) string {
	if *tSoll >= 80 && *tSoll <= 280 {
		return strconv.Itoa(*tSoll / 10 * 2)
	}

	return "253"; // off
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
		if isAVMButton && !isAVMMeasure {
			filtered = append(filtered, &devices[i])
		}
	}

	return filtered
}