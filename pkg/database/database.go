package database

import (
	"database/sql"
	"errors"
	"fmt"
	"goex/pkg/config"
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

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTables() error {
	tables := []string{}

	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'").Error
	if err != nil {
		return err
	}

	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	tables := []string{}

	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	DB.Exec("SET foreign_key_checks = 0;")

	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}
