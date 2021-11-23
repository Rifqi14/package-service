package postgresql

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func Migrate(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "../../migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Error migration := ", err.Error())
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
