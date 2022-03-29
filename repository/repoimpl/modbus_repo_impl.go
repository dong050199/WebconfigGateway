package repoimpl

import (
	"SQLite_Repo_Pattern/model"
	models "SQLite_Repo_Pattern/model"
	repo "SQLite_Repo_Pattern/repository"
	"database/sql"
	"fmt"
)

type ModbusRepoImpl struct {
	Db *sql.DB
}

func NewModbusRepo(db *sql.DB) repo.ModbusRepo {
	return &ModbusRepoImpl{
		Db: db,
	}
}

func (u *ModbusRepoImpl) Select() ([]models.Modbus, error) {
	Modbuss := make([]model.Modbus, 0)
	rows, err := u.Db.Query("SELECT * FROM Modbus")
	if err != nil {
		return Modbuss, err
	}
	for rows.Next() {
		Modbus := models.Modbus{}
		err := rows.Scan(&Modbus.ID, &Modbus.Ip, &Modbus.Port)
		if err != nil {
			break
		}
		Modbuss = append(Modbuss, Modbus)
	}
	err = rows.Err()
	if err != nil {
		return Modbuss, err
	}
	return Modbuss, nil
}

func (u *ModbusRepoImpl) Isert(Modbus models.Modbus) error {
	fmt.Println("HELLO")
	insertStatement := `
	INSERT INTO Modbus (id, ip, port)
	VALUES ($1,$2,$3)
	`
	_, err := u.Db.Exec(insertStatement, Modbus.ID, Modbus.Ip, Modbus.Port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Record added", Modbus)
	return nil
}

func (u *ModbusRepoImpl) Update(Modbus models.Modbus) error {
	stmt, _ := u.Db.Prepare("UPDATE Modbus set ip = ?, port = ? where id = ?")
	_, err := stmt.Exec(Modbus.Ip, Modbus.Port, Modbus.ID)
	if err != nil {
		return err
	}
	fmt.Println("Record updated", Modbus)
	return nil
}

func (u *ModbusRepoImpl) Delete(Modbus int) error {
	_, err := u.Db.Exec(fmt.Sprintf("DELETE FROM Modbus WHERE id = %d", Modbus))
	if err != nil {
		return err
	}
	return nil
}
