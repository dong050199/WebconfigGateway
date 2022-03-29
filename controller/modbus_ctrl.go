package controller

import (
	"SQLite_Repo_Pattern/driver"
	models "SQLite_Repo_Pattern/model"
	"SQLite_Repo_Pattern/repository/repoimpl"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const (
	SQLiteDR = "./data_gw.db"
)

//FUNCTION GET DELETE PUT GET MODBUS

func Getmod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusRepo := repoimpl.NewModbusRepo(driver.SQLite.SQL)
	fmt.Println(ModbusRepo.Select())
	data, err := ModbusRepo.Select()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(data)
	//db.SQL.Close()
}

func Deletemod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	params := mux.Vars(r)
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusRepo := repoimpl.NewModbusRepo(driver.SQLite.SQL)
	id, _ := strconv.Atoi(params["id"])
	err = ModbusRepo.Delete(id)
	json.NewEncoder(w).Encode("e")
	//db.SQL.Close()
}

func Create_mod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var modbuss models.Modbus
	_ = json.NewDecoder(r.Body).Decode(&modbuss)

	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusRepo := repoimpl.NewModbusRepo(driver.SQLite.SQL)
	err = ModbusRepo.Isert(modbuss)
	json.NewEncoder(w).Encode(modbuss)
	//db.SQL.Close()
}

func Updatemod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var modbuss models.Modbus
	_ = json.NewDecoder(r.Body).Decode(&modbuss)
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusRepo := repoimpl.NewModbusRepo(driver.SQLite.SQL)
	err = ModbusRepo.Update(modbuss)
	json.NewEncoder(w).Encode(err)
	//db.SQL.Close()
}
