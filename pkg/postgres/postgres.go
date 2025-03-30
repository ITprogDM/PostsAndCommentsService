package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConfigs struct {
	Host, Port, User, Password, Name string
}

func NewPostgresDB(cfg PostgresConfigs) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name))

	if err != nil {
		return nil, err
	}

	return db, nil
}
