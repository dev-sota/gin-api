package database

import (
	"github.com/dev-sota/gin-api/entities"
	"github.com/jinzhu/gorm"

	// Use MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize DB
func Init() {
	gormConnect()
	autoMigration()
}

// Get is called in models
func Get() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// initialize and setup the DB
func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "gin_api_development"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)

	return db
}

// Please add migration
func autoMigration() {
	db.AutoMigrate(&entities.User{})
}
