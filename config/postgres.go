package config

import (
	"os"
	"strconv"

	"github.com/ionnotion/statictools-server/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dsn := ""
	opts := gorm.Config{}

	pg, err := gorm.Open(postgres.Open(dsn), &opts)
	if err != nil {
		return nil, err
	}

	if migrate, _ := strconv.ParseBool(os.Getenv("run_migration")); migrate {
		migration.MigratePostgres()
	}

	return pg, nil
}
