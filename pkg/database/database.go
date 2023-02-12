package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB object
var DB *gorm.DB
var SQLDB *sql.DB

// Connect to the database
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// use gorm.Open connect to the database
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// handle errors
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get the underlying sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
