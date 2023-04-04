package model

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnSQLite(path string) (err error) {

	var loggerConfig logger.Config = logger.Config{
		SlowThreshold:             time.Second * 2, // Slow SQL threshold
		LogLevel:                  logger.Error,    // Log level
		IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,           // Disable color
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConfig,
	)

	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	if err := MigrateSchema(); err != nil {
		panic("failed to initial schema")
	}
	return nil
}

func MigrateSchema() error {
	err := db.AutoMigrate(&Lottery{})
	return err
}
