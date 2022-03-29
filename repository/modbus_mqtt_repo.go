package repository

import (
	models "SQLite_Repo_Pattern/model"
)

type ModbusMqttRepo interface {
	Select() ([]models.ModbusMqtt, error)
	Isert(u models.ModbusMqtt) error
	Update(u models.ModbusMqtt) error
	Delete(u int) error
}
