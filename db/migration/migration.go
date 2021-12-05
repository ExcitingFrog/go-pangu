package migration

import (
	"database/sql"
	"fmt"
	"go-pangu/conf"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type VersionStruct struct {
	Version int
	Dirty   bool
}

const versionErr = "file does not exist"

func Migrate(version int) error {
	fmt.Println("migrating")
	db, err := sql.Open("postgres", conf.GetEnv("DATABASE_URL"))
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"DATABASE_NAME", driver)
	if err != nil {
		return err
	}
	for {
		version++
		err = m.Steps(version)
		if err != nil {
			if err.Error() == versionErr {
				break
			}
			return err
		}
	}
	fmt.Println("migrating done")
	return nil
}
