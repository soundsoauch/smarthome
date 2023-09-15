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
}