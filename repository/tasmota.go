package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"smarthome/devices"
	"smarthome/interfaces"
	"sort"

	"golang.org/x/exp/slices"
)


func GetTasmotaState(tasmotaDeviceConfig TasmotaDeviceConfig, ch chan<- *interfaces.TasmotaDeviceDto) {
	resp, err := http.Get("http://" + tasmotaDeviceConfig.hostName + "/cm?cmnd=Status%200")
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

	switchBody.Type = tasmotaDeviceConfig.deviceType

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


func GetTasmotaDevices(tasmotaDeviceConfigs []TasmotaDeviceConfig) ([]*interfaces.Device, error) {
	var ch chan *interfaces.TasmotaDeviceDto = make(chan *interfaces.TasmotaDeviceDto)
	var results []*interfaces.Device
	var hostNames []string

	for _, tasmotaDeviceConfig := range tasmotaDeviceConfigs {
		hostNames = append(hostNames, tasmotaDeviceConfig.hostName)
		go GetTasmotaState(tasmotaDeviceConfig, ch)
	}
	for {
		var tasmotaDevice *devices.TasmotaDevice = devices.NewTasmotaDevice(<-ch)

		results = append(results, tasmotaDevice.GetDevice())

		if len(results) == len(tasmotaDeviceConfigs) {
			break
		}
	}

	sort.Slice(results, func(i int, j int) bool {
		a := slices.Index(hostNames, results[i].ID)
		b := slices.Index(hostNames, results[j].ID)
		return a < b
	})

	return results, nil
}