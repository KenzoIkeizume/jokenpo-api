package datastore

import (
	"jokenpo-api/config"
	"jokenpo-api/domain/model"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	DBMS := "mysql"
	mySQLConfig := &mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
		},
	}
	db, err := gorm.Open(DBMS, mySQLConfig.FormatDSN())

	if err != nil {
		log.Fatalln("database error: ", err)
	}

	migrations(db)

	return db
}

func migrations(db *gorm.DB) {
	db.AutoMigrate(&model.User{})

	// Add table suffix when create tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})

}
