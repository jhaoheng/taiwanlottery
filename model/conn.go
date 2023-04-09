package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var IsDebug bool = false

func ConnMySQL() (err error) {
	var loggerConfig logger.Config = logger.Config{
		SlowThreshold:             time.Second * 2, // Slow SQL threshold
		LogLevel:                  logger.Error,    // Log level
		IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,            // Disable color
	}

	if IsDebug {
		loggerConfig.LogLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConfig,
	)

	dsn := "test:test@tcp(127.0.0.1:3311)/taiwanlottery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		err = fmt.Errorf("ERROR: fail connection database %s", err.Error())
		return
	}

	dbConfig, _ := db.DB()
	dbConfig.SetMaxIdleConns(10)
	dbConfig.SetMaxOpenConns(10)
	dbConfig.SetConnMaxLifetime(time.Hour)

	// if err := MigrateSchema(); err != nil {
	// 	panic("failed to initial schema")
	// }
	return nil
}

func MigrateSchema() error {
	err := db.AutoMigrate(&Lottery{}, &Lotto649AllSets{}, &Lotto649Filtered{})
	return err
}
