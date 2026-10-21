package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PsqlDb struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DBName string
}

func NewPsqlDb(user, pass, host, port, dbname string) *PsqlDb {
	return &PsqlDb{
		User:   user,
		Pass:   pass,
		Host:   host,
		Port:   port,
		DBName: dbname,
	}
}

func (db *PsqlDb) Connect() (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db.User, db.Pass, db.Host, db.Port, db.DBName)

	return pgxpool.New(context.Background(), connStr)
}
