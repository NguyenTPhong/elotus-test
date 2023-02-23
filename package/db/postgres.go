package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(dsn string, MaxOpenCon, MaxIdleCon, logLevel int) (*gorm.DB, error) {
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
	})
	if err != nil {
		return nil, err
	}
	db, err := connect.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(MaxOpenCon)
	db.SetMaxIdleConns(MaxIdleCon)
	return connect, nil
}

func Close(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		return err
	}
	return database.Close()
}
