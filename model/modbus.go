package model

type Modbus struct {
	ID   int    `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}
