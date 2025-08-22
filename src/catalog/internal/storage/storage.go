// Package storage
//
// Defines DB operations that are available to use in handlers and services
package storage

import (
	"context"
	"fmt"

	"catalog/internal/shared"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Init() {
	p, err := pgxpool.New(context.Background(),
		fmt.Sprintf("postgres://%s:%s@db:%s/catalog",
			shared.Config.DBUser,
			shared.Config.DBPass,
			shared.Config.DBPort))
	if err != nil {
		panic(err)
	}
	pool = p
}

func Close() {
	pool.Close()
}
