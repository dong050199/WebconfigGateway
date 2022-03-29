package repository

import (
	models "SQLite_Repo_Pattern/model"
)

type ModbusTagRepo interface {
	Select() ([]models.ModbusTag, error)
	Isert(u models.ModbusTag) error
	Update(u models.ModbusTag) error
	Delete(u int) error
}
