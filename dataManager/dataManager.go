package dataManager

import (
	"math/rand"
	"sample-app/thermostat"
	"time"
)

type DataManager struct {
	thermostats map[string]*thermostat.Thermostat
}

func New() DataManager {
	var d DataManager = DataManager{
		thermostats: map[string]*thermostat.Thermostat{
			"livingroom":  thermostat.New("Living room", "00:1A:22:16:9E:37"),
			"kitchen":     thermostat.New("Kitchen", "00:1A:22:16:A4:2A"),
			"floor":       thermostat.New("Floor", "00:1A:22:16:96:2E"),
			"bathroom":    thermostat.New("Bathroom", "00:1A:22:16:8C:00"),
			"bedroom":     thermostat.New("Bathroom", "00:1A:22:16:89:6F"),
			"utilityroom": thermostat.New("Utility room", "00:1A:22:16:89:A0"),
		},
	}

	go d.initSync()

	return d
}

func (d DataManager) GetValueOf(roomName string) string {
	var thermostat, ok = d.thermostats[roomName]

	if ok {
		return thermostat.GetValue()
	}

	return "-"
}

func (d DataManager) initSync() {
	var ticker *time.Ticker = time.NewTicker(2 * time.Second)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for range ticker.C {
		for i := range d.thermostats {
			thermostat := d.thermostats[i]
			// send WOL on iris and fetch state for all then shutdown iris
			thermostat.SetValue(uint8(r1.Intn(20)) + 5)
		}
	}

}
