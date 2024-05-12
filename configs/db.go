package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type connection struct {
	db *sql.DB
}

var db *sql.DB

func InitDB() {

	Password := Password()
	Host := Host()
	Port := Port()
	DbName := DbName()
	User := Username()

	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", User, Password, Host, Port)
	createDB, err := sql.Open("mysql", createDBDsn)
	if err != nil {
		log.Fatalf("Error connecting to MySQL server: %v\n", err)
	}

	_, err = createDB.Exec("CREATE DATABASE IF NOT EXISTS " + DbName)
	if err != nil {
		log.Fatalf("Error creating database: %v\n", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, DbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
}

func NewConnection() *connection {
	return &connection{
		db,
	}
}

func (self *connection) GetDB() *sql.DB {
	return self.db
}
