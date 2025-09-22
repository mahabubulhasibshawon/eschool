package middlewares

import (
	"eschool/config"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type Middlewares struct {
	cnf *config.Config
	DB *sqlx.DB
}

func NewMiddlewares(cnf *config.Config, DB *sqlx.DB) *Middlewares {
	return &Middlewares{
		cnf: cnf,
		DB: DB,
	}
}