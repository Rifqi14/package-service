package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Host                    string
	DbName                  string
	User                    string
	Password                string
	Port                    string
	Location                *time.Location
	DBMaxConnection         int
	DBMAxIdleConnection     int
	DBMaxLifeTimeConnection int
}

func (c Connection) Conn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DbName)

	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	db, _ := DB.DB()

	db.Ping()
	db.SetMaxOpenConns(c.DBMaxConnection)
	db.SetMaxIdleConns(c.DBMAxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(c.DBMaxLifeTimeConnection) * time.Second)

	return DB, err
}
