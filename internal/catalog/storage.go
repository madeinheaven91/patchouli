package catalog

import (
	"context"
	"fmt"

	"github.com/madeinheaven91/patchouli/internal/shared"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB() {
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

func FetchBook(hash string, ctx context.Context) (*Book, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var book Book
	row := conn.QueryRow(ctx, "SELECT * FROM book WHERE hash=$1", hash)
	err = row.Scan(&book.Hash,
		&book.FilePath,
		&book.Title,
		&book.Description,
		&book.Format,
		&book.CategoryID,
		&book.LanguageCode,
		&book.Published)
	return &book, err
}
