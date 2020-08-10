// +build integration

package config

import (
	"context"
	"database/sql"
	"log"

	"github.com/selmison/code-micro-videos/models"
)

const (
	AddressServer = "127.0.0.1:3333"
	DBDrive       = "postgres"
	DBName        = "code-micro-videos"
	DBPort        = 5432
	DBUser        = "postgres"
	DBPass        = "postgres"
	DBSSLMode     = "disable"
)

var (
	DBConnStr string
)

func ClearCategoriesTable(dbDriver, dbConnStr string) error {
	db, err := sql.Open(dbDriver, dbConnStr)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := models.Categories().DeleteAll(context.Background(), db); err != nil {
		return err
	}
	return nil
}
