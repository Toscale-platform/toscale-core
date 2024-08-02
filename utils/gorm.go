package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGormConn(db, dbschema, user, pswd, host string, port int) (conn *gorm.DB, err error) {

	if len(dbschema) == 0 {
		dbschema = "public"
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", user, pswd, db, host, port)
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", dbschema),
			SingularTable: false,
		},
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