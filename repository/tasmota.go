package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"smarthome/devices"
	"smarthome/interfaces"
)


func GetState(hostname string, ch chan<- *interfaces.TasmotaDeviceDto) {
	resp, err := http.Get("http://" + hostname + "/cm?cmnd=Status%200")
	if err != nil {
		ch <- nil
		return
	}

	var switchBody *interfaces.TasmotaDeviceDto
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- nil
		return
	}

	err = json.Unmarshal([]byte(body), &switchBody)
	if err != nil {
		ch <- nil
		return
	}

	ch <- switchBody
}


func GetTasmotaDevices(hostnames []string) ([]*interfaces.Device, error) {
	var ch chan *interfaces.TasmotaDeviceDto = make(chan *interfaces.TasmotaDeviceDto)
	var results []*interfaces.Device

	for _, hostname := range hostnames {
		go GetState(hostname, ch)
	}
	for {
		var tasmotaDevice *devices.TasmotaDevice = devices.NewTasmotaDevice(<-ch)

		results = append(results, tasmotaDevice.GetDevice())

		if len(results) == len(hostnames) {
			break
		}
	}

	return results, nil
}