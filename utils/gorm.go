package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormConn(db, user, pswd, host string, port int) (conn *gorm.DB, err error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", user, pswd, db, host, port)
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CloseGormConn(conn *gorm.DB) error {
	db, err := conn.DB()

	if err != nil {
		return err
	}

	if err = db.Close(); err != nil {
		return err
	}

	return nil
}