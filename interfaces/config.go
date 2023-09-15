package interfaces

type Config struct {
	Fritzbox struct {
		User string     `yaml:"user"`
		Password string `yaml:"password"`
		Host string     `yaml:"host"`
	} `yaml:"fritzbox"`
	TasmotaDevices []struct {
		Host string `yaml:"host"`
		Type Type   `yaml:"type"`
	} `yaml:"tasmotaDevices"`
	Iris struct {
		Host string      `yaml:"host"`
		User string      `yaml:"user"`
		Mac string       `yaml:"mac"`
		Interface string `yaml:"interface"`
	} `yaml:"iris"`
}