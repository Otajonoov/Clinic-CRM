package db

import (
	"fmt"

	"gitlab.com/clinic-crm/labs/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDBForSuite(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, func() {}
	}
	CleanUpFunc := func() {
		connDB.Close()
	}

	return connDB, CleanUpFunc
}

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDB, err

}
