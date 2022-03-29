package model

type ModbusTag struct {
	ID      int `json:"id"`
	Id_dev  int `json:"id_dev"`
	R_start int `json:"r_start"`
	R_num   int `json:"r_num"`
}
