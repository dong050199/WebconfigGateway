package main

import (
	"SQLite_Repo_Pattern/config"
	c "SQLite_Repo_Pattern/controller"
	u "SQLite_Repo_Pattern/controller/user_handler"
	"SQLite_Repo_Pattern/driver"
	"SQLite_Repo_Pattern/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const (
	SQLiteDR = "./data.db"
)

func main() {

	driver.Connect(config.DBPATH)

	r := mux.NewRouter()

	// userCheck := r.PathPrefix("/user").Subrouter()
	// userCheck.Use(middleware.Authentication)
	// userCheck.HandleFunc("/login", u.Login)
	login := r.PathPrefix("/login").Subrouter()
	login.HandleFunc("/user", u.Login)
	login.HandleFunc("/register", u.Register)
	config := r.PathPrefix("/data").Subrouter()
	config.Use(middleware.Authentication)

	config.HandleFunc("/modbus", c.Getmod).Methods("GET")
	//config.HandleFunc("/api/register", u.Register).Methods("POST")
	config.HandleFunc("/modbustag", c.Getmod_tag).Methods("GET")
	config.HandleFunc("/modbusmqtt", c.Getmod_mqtt).Methods("GET")

	config.HandleFunc("/modbus", c.Create_mod).Methods("POST")
	config.HandleFunc("/modbustag", c.Createmod_tag).Methods("POST")
	config.HandleFunc("/modbusmqtt", c.Createmod_mqtt).Methods("POST")

	config.HandleFunc("/modbus/{id}", c.Deletemod).Methods("DELETE")
	config.HandleFunc("/modbustag/{id}", c.Deletemod_tag).Methods("DELETE")
	config.HandleFunc("/modbusmqtt/{id}", c.Deletemod_mqtt).Methods("DELETE")

	config.HandleFunc("/modbus", c.Updatemod).Methods("PUT")
	config.HandleFunc("/modbustag/{id}", c.Updatemod_tag).Methods("PUT")
	config.HandleFunc("/modbusmqtt/{id}", c.Updatemod_mqtt).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
