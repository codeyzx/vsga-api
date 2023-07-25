package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	dsn := "postgres://qqbdhnat:SP0BJr1Uk_HAoAUxzM7k2eMHlicyU3pj@hansken.db.elephantsql.com/qqbdhnat"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")
}
