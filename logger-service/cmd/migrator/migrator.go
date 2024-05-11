package migrator

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"log"
)

func MustRun(connStr, migrDir string) {
	driver := "postgres"

	db, err := sqlx.Open(driver, connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	if err := goose.SetDialect(driver); err != nil {
		panic(err.Error())
	}

	if err := goose.Up(db.DB, migrDir); err != nil {
		panic(err.Error())
	}
}
