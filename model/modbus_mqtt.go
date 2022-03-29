package model

type ModbusMqtt struct {
	ID     int    `json:"id"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	Clid   string `json:"clid"`
	Id_dev string `json:"id_dev"`
	Topic  string `json:"topic"`
}
