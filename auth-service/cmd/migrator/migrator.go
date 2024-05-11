package migrator

import (
	"auth-service/internal/domain/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"log"
)

func Run(connStr, migrDir string) {
	driver := "postgres"

	db, err := sqlx.Open(driver, connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	if err := goose.SetDialect(driver); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, migrDir); err != nil {
		panic(err)
	}

	addRoles := fmt.Sprintf("insert into Roles(role) values ('%s'), ('%s');", model.ArtistRole, model.ListenerRole)

	_, err = db.Exec(addRoles)
}
