package repoimpl

import (
	"SQLite_Repo_Pattern/model"
	models "SQLite_Repo_Pattern/model"
	repo "SQLite_Repo_Pattern/repository"
	"database/sql"
	"fmt"
)

type ModbusTagRepoImpl struct {
	Db *sql.DB
}

func NewModbusTagRepo(db *sql.DB) repo.ModbusTagRepo {
	return &ModbusTagRepoImpl{
		Db: db,
	}
}

func (u *ModbusTagRepoImpl) Select() ([]models.ModbusTag, error) {
	ModbusTags := make([]model.ModbusTag, 0)
	rows, err := u.Db.Query("SELECT * FROM ModbusTag")
	if err != nil {
		return ModbusTags, err
	}
	for rows.Next() {
		ModbusTag := models.ModbusTag{}
		err := rows.Scan(&ModbusTag.ID, &ModbusTag.Id_dev, &ModbusTag.R_start, &ModbusTag.R_num)
		if err != nil {
			break
		}
		ModbusTags = append(ModbusTags, ModbusTag)
	}
	err = rows.Err()
	if err != nil {
		return ModbusTags, err
	}
	return ModbusTags, nil
}

func (u *ModbusTagRepoImpl) Isert(ModbusTag models.ModbusTag) error {
	insertStatement := `
	INSERT INTO ModbusTag (id, id_dev, r_start, r_num)
	VALUES ($1,$2,$3,$4)
	`
	_, err := u.Db.Exec(insertStatement, ModbusTag.ID, ModbusTag.Id_dev, ModbusTag.R_start, ModbusTag.R_num)
	if err != nil {
		return err
	}
	fmt.Println("Record added", ModbusTag)
	return nil
}

func (u *ModbusTagRepoImpl) Update(ModbusTag models.ModbusTag) error {
	stmt, _ := u.Db.Prepare("UPDATE ModbusTag set id_dev = ?, r_start = ?, r_num = ? where id = ?")
	_, err := stmt.Exec(ModbusTag.Id_dev, ModbusTag.R_start, ModbusTag.R_num, ModbusTag.ID)
	if err != nil {
		return err
	}
	fmt.Println("Record updated", ModbusTag)
	return nil
}

func (u *ModbusTagRepoImpl) Delete(ModbusTag int) error {
	_, err := u.Db.Exec(fmt.Sprintf("DELETE FROM ModbusTag WHERE id = %d", ModbusTag))
	if err != nil {
		return err
	}
	return nil
}
