// Package storage
//
// Defines DB operations that are available to use in handlers and services
package storage

import (
	"context"
	"fmt"

	"github.com/madeinheaven91/patchouli/internal/models"
	"github.com/madeinheaven91/patchouli/internal/shared"

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

func FetchBook(id string, ctx context.Context) (*models.Book, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var book models.Book
	row := conn.QueryRow(ctx, "SELECT * FROM book WHERE id=$1", id)
	err = row.Scan(&book.ID,
		&book.FilePath,
		&book.Title,
		&book.Description,
		&book.Format,
		&book.Category,
		&book.LanguageCode,
		&book.Published)
	return &book, err
}
