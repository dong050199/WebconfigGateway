package repoimpl

import (
	"SQLite_Repo_Pattern/model"
	models "SQLite_Repo_Pattern/model"
	repo "SQLite_Repo_Pattern/repository"
	"database/sql"
	"fmt"
)

type ModbusMqttRepoImpl struct {
	Db *sql.DB
}

func NewModbusMqttRepo(db *sql.DB) repo.ModbusMqttRepo {
	return &ModbusMqttRepoImpl{
		Db: db,
	}
}

func (u *ModbusMqttRepoImpl) Select() ([]models.ModbusMqtt, error) {
	ModbusMqtts := make([]model.ModbusMqtt, 0)
	rows, err := u.Db.Query("SELECT * FROM ModbusMqtt")
	if err != nil {
		return ModbusMqtts, err
	}
	for rows.Next() {
		ModbusMqtt := models.ModbusMqtt{}
		err := rows.Scan(&ModbusMqtt.ID, &ModbusMqtt.Ip, &ModbusMqtt.Port, &ModbusMqtt.User, &ModbusMqtt.Pwd, &ModbusMqtt.Clid, &ModbusMqtt.Id_dev, &ModbusMqtt.Topic)
		if err != nil {
			break
		}
		ModbusMqtts = append(ModbusMqtts, ModbusMqtt)
	}
	err = rows.Err()
	if err != nil {
		return ModbusMqtts, err
	}
	return ModbusMqtts, nil
}

func (u *ModbusMqttRepoImpl) Isert(ModbusMqtt models.ModbusMqtt) error {
	insertStatement := `
	INSERT INTO ModbusMqtt (id, ip, port, user, pwd, clid, id_dev, topic)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`
	_, err := u.Db.Exec(insertStatement, ModbusMqtt.ID, ModbusMqtt.Ip, ModbusMqtt.Port, ModbusMqtt.User, ModbusMqtt.Pwd, ModbusMqtt.Clid, ModbusMqtt.Id_dev, ModbusMqtt.Topic)
	if err != nil {
		return err
	}
	fmt.Println("Record added", ModbusMqtt)
	return nil
}

func (u *ModbusMqttRepoImpl) Update(ModbusMqtt models.ModbusMqtt) error {
	stmt, _ := u.Db.Prepare("UPDATE ModbusMqtt set ip = ?, port = ?, user = ?, pwd = ?, clid = ?, id_dev = ?, topic = ? where id = ?")
	_, err := stmt.Exec(ModbusMqtt.Ip, ModbusMqtt.Port, ModbusMqtt.User, ModbusMqtt.Pwd, ModbusMqtt.Clid, ModbusMqtt.Id_dev, ModbusMqtt.Topic, ModbusMqtt.ID)
	if err != nil {
		return err
	}
	fmt.Println("Record updated", ModbusMqtt)
	return nil
}

func (u *ModbusMqttRepoImpl) Delete(ModbusMqtt int) error {
	_, err := u.Db.Exec(fmt.Sprintf("DELETE FROM ModbusMqtt WHERE id = %d", ModbusMqtt))
	if err != nil {
		return err
	}
	return nil
}
