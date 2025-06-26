package repository

import (
	redisstrore "github.com/duckcoding00/money-tracker/money-tracker-backend/repository/redisStrore"
	"github.com/duckcoding00/money-tracker/money-tracker-backend/repository/sql"
)

type Repository struct {
	Sql   *sql.Queries
	Redis redisstrore.RedisMethod
}

func NewRepository(sqlQueries *sql.Queries, redisAddr string) *Repository {
	return &Repository{
		Sql:   sqlQueries,
		Redis: redisstrore.NewRedis(redisAddr),
	}
}
