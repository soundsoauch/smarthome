package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"smarthome/devices"
	"smarthome/interfaces"
)


func GetTasmotaState(hostname string, ch chan<- *interfaces.TasmotaDeviceDto) {
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

func SetTasmotaState(hostname string, deviceState *interfaces.DeviceStateDto) error {
	var alert = deviceState.Alert
	
	if alert == nil {
		return errors.New("incorrect State")
	}

	resp, err := http.Get("http://" + hostname + "/cm?cmnd=Power%20" + string(*alert))

	if err != nil {
		return err
	}

	var dto struct {
		Power string `json:"POWER"`
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), &dto)
	if err != nil {
		return err
	}

	return nil
}


func GetTasmotaDevices(hostnames []string) ([]*interfaces.Device, error) {
	var ch chan *interfaces.TasmotaDeviceDto = make(chan *interfaces.TasmotaDeviceDto)
	var results []*interfaces.Device

	for _, hostname := range hostnames {
		go GetTasmotaState(hostname, ch)
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