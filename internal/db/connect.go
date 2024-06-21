package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"productsEsik/internal/config"
)

var DB *sql.DB

func ConDb(cfg *config.Config) error {
	var err error

	dbInfo := cfg.Database

	strCon := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", dbInfo.User, dbInfo.DBName, dbInfo.Password)

	DB, err = sql.Open("postgres", strCon)

	if err != nil {
		return err
	}

	return nil
}
