package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/julioc98/pismo/pkg/env"

	// postgres db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Conn connect on SQL Data Base Gorm
func Conn() (db *gorm.DB) {

	typeDB := env.Get("DB", "postgres")

	pathDB := env.Get("DATABASE_URL", "host=localhost port=5432 user=pismo dbname=pismo password=pismo sslmode=disable")

	db, err := gorm.Open(typeDB, pathDB)
	if err != nil {
		log.Panicf("failed to connect database err: %s", err.Error())
	}

	return db
}
