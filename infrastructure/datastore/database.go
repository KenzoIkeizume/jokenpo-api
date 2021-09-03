package datastore

import (
	config "jokenpo-api/config"
	model "jokenpo-api/domain/model"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
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
		Timeout: config.C.Database.Timeout,
	}
	db, err := gorm.Open(DBMS, mySQLConfig.FormatDSN())

	if err != nil {
		glog.Fatal("Database cannot connect: %+v", err)
	}

	migrations(db)

	return db
}

func migrations(db *gorm.DB) {
	db.AutoMigrate(&model.User{})

	// Add table suffix when create tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})

}
