package ci

import (
	"Library/utils"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
)



func MigrationUp() {
	url, err := utils.ConnectionURLBuilder("migration")
	if err != nil {
		log.Println("Error generating migration url: ", err.Error())
	}
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		log.Fatal("error in creating migration:", err.Error())
	}
	if err := m.Up(); err != nil {
		log.Println("Error in migration ", err.Error())
	}
}
