package switchlist

import "smarthome/switches"

type Switchlist struct {
	switches []*switches.Switch
}

func New() *Switchlist {
	return &Switchlist{
		switches: []*switches.Switch{
			switches.New("Deko"),
			switches.New("camera"),
			switches.New("camera-furbo"),
			switches.New("Dimmer"),
			switches.New("BedroomLeft"),
			switches.New("BedroomRight"),
		},
	}
}

func (s *Switchlist) GetSwitchList() ([]*switches.Switch, error) {
	var ch chan *switches.Switch = make(chan *switches.Switch)
	var results []*switches.Switch

	for _, singleSwitch := range s.switches {
		go singleSwitch.GetState(ch)
	}
	for {
		results = append(results, <-ch)

		if len(results) == len(s.switches) {
			break
		}
	}
	return results, nil
}
