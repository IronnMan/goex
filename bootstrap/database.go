package bootstrap

import (
	"errors"
	"fmt"
	"goex/pkg/config"
	"goex/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// SetupDB initialize the database and ORM
func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		// build DSN information
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// initialize sqlite
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// connect to the database and set the log mode of GORM
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// Set the maximum number of open connections
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// set the maximum number of idle connections
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// set an expiration time for each link
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
}
