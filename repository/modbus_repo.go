package repository

import (
	models "SQLite_Repo_Pattern/model"
)

type ModbusRepo interface {
	Select() ([]models.Modbus, error)
	Isert(u models.Modbus) error
	Update(u models.Modbus) error
	Delete(u int) error
}
