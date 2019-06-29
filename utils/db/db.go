package db

import (
	"log"

	"github.com/jinzhu/gorm"
	//import dialect
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

// InitSqlite is
func InitSqlite() {
	db, err := gorm.Open("sqlite3", "./dbFile.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}

// GetDB is
func GetDB() *gorm.DB {
	return db
}
