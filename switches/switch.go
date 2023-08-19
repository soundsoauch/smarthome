package switches

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Switch struct {
	HostName    string        `json:"name"`
	Power       *bool         `json:"power,omitempty"`
	IsConnected bool          `json:"isConnected"`
	Energy      *SwitchEnergy `json:"energy,omitempty"`
}

type SwitchEnergy struct {
	Total     float32 `json:"total"`
	Yesterday float32 `json:"yesterday"`
	Today     float32 `json:"today"`
	Power     float32 `json:"power"`
}

type SwitchDto struct {
	Status    StatusDto    `json:"Status"`
	StatusSNS StatusSNSDTO `json:"StatusSNS"`
}

type StatusDto struct {
	Power int `json:"Power"`
}

type StatusSNSDTO struct {
	Energy *EnergyDTO `json:"ENERGY"`
}

type EnergyDTO struct {
	Total     float32 `json:"Total"`
	Yesterday float32 `json:"Yesterday"`
	Today     float32 `json:"Today"`
	Power     float32 `json:"Power"`
}

func New(hostname string) *Switch {
	return &Switch{HostName: hostname}
}

func (s *Switch) GetState(ch chan<- *Switch) {
	resp, err := http.Get("http://" + s.HostName + "/cm?cmnd=Status%200")
	if err != nil {
		s.IsConnected = false
		s.Power = nil
		s.Energy = nil
		ch <- s
		return
	}

	var switchBody *SwitchDto
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

	s.mapResults(switchBody)

	ch <- s
}

func (s *Switch) mapResults(body *SwitchDto) {
	if body.StatusSNS.Energy != nil {
		s.Energy = &SwitchEnergy{
			Total:     body.StatusSNS.Energy.Today,
			Yesterday: body.StatusSNS.Energy.Yesterday,
			Today:     body.StatusSNS.Energy.Today,
			Power:     body.StatusSNS.Energy.Power,
		}
	}
	power := body.Status.Power == 1
	s.Power = &power
	s.IsConnected = true
}
