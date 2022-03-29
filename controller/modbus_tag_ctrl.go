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

//FUNCTION GET DELETE PUT GET ModbusTag

func Getmod_tag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusTagRepo := repoimpl.NewModbusTagRepo(driver.SQLite.SQL)
	fmt.Println(ModbusTagRepo.Select())
	data, err := ModbusTagRepo.Select()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(data)
	//db.SQL.Close()
}

func Deletemod_tag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	params := mux.Vars(r)
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusTagRepo := repoimpl.NewModbusTagRepo(driver.SQLite.SQL)
	id, _ := strconv.Atoi(params["id"])
	err = ModbusTagRepo.Delete(id)
	json.NewEncoder(w).Encode("e")
	//db.SQL.Close()
}

func Createmod_tag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var modbuss models.ModbusTag
	_ = json.NewDecoder(r.Body).Decode(&modbuss)

	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusTagRepo := repoimpl.NewModbusTagRepo(driver.SQLite.SQL)
	err = ModbusTagRepo.Isert(modbuss)
	json.NewEncoder(w).Encode(modbuss)
	//db.SQL.Close()
}

func Updatemod_tag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var ModbusTags models.ModbusTag
	_ = json.NewDecoder(r.Body).Decode(&ModbusTags)
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusTagRepo := repoimpl.NewModbusTagRepo(driver.SQLite.SQL)
	err = ModbusTagRepo.Update(ModbusTags)
	json.NewEncoder(w).Encode(err)
	//db.SQL.Close()
}
