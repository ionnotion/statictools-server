package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ionnotion/statictools-server/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

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
