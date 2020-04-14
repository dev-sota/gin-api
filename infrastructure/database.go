package infrastructure

import (
	"github.com/dev-sota/gin-api/entity"
	"github.com/jinzhu/gorm"

	// Use MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBInit is initialize DB
func DBInit() {
	db := gormConnect()

	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	db.AutoMigrate(&entity.User{})
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "gin_api_development"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
