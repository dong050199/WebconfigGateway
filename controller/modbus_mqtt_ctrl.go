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

//FUNCTION GET DELETE PUT GET ModbusMqtt

func Getmod_mqtt(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applycation/json")
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusMqttRepo := repoimpl.NewModbusMqttRepo(driver.SQLite.SQL)
	fmt.Println(ModbusMqttRepo.Select())
	data, err := ModbusMqttRepo.Select()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(data)
	//db.SQL.Close()
}

func Deletemod_mqtt(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applycation/json")
	params := mux.Vars(r)
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusMqttRepo := repoimpl.NewModbusMqttRepo(driver.SQLite.SQL)
	id, _ := strconv.Atoi(params["id"])
	err = ModbusMqttRepo.Delete(id)
	json.NewEncoder(w).Encode(err)
}

func Createmod_mqtt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var ModbusMqtts models.ModbusMqtt
	_ = json.NewDecoder(r.Body).Decode(&ModbusMqtts)

	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusMqttRepo := repoimpl.NewModbusMqttRepo(driver.SQLite.SQL)
	err = ModbusMqttRepo.Isert(ModbusMqtts)
	json.NewEncoder(w).Encode(ModbusMqtts)
	//db.SQL.Close()
}

func Updatemod_mqtt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var ModbusMqtts models.ModbusMqtt
	_ = json.NewDecoder(r.Body).Decode(&ModbusMqtts)
	//db := driver.Connect(SQLiteDR)
	err := driver.SQLite.SQL.Ping()
	if err != nil {
		panic(err)
	}
	ModbusMqttRepo := repoimpl.NewModbusMqttRepo(driver.SQLite.SQL)
	err = ModbusMqttRepo.Update(ModbusMqtts)
	json.NewEncoder(w).Encode(err)
	//db.SQL.Close()
}
