package interfaces

type TasmotaDeviceDto struct {
	Status    TasmotaDeviceStatusDto    `json:"Status"`
	StatusSNS TasmotaDeviceStatusSNSDto `json:"StatusSNS"`
	StatusNET TasmotaDeviceStatusNETDto `json:"StatusNET"`
}

type TasmotaDeviceStatusDto struct {
	Power int `json:"Power"`
	FriendlyName []string `json:"FriendlyName"`
}

type TasmotaDeviceStatusSNSDto struct {
	Energy *TasmotaDeviceEnergyDto `json:"ENERGY"`
}

type TasmotaDeviceStatusNETDto struct {
	HostName string `json:"Hostname"`
}

type TasmotaDeviceEnergyDto struct {
	Total     float32 `json:"Total"`
	Yesterday float32 `json:"Yesterday"`
	Today     float32 `json:"Today"`
	Power     float32 `json:"Power"`
}