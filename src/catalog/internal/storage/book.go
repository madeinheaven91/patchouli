package storage

import (
	"context"

	"catalog/internal/models"
	"catalog/internal/shared"
)

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
		&book.Filename,
		&book.Title,
		&book.AuthorID,
		&book.Description,
		&book.Category,
		&book.LanguageCode,
		&book.Published)
	return &book, err
}

func FetchAllBooks(ctx context.Context) ([]models.Book, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var book models.Book
	res := make([]models.Book, 0)
	rows, err := conn.Query(ctx, "SELECT * FROM book")
	for rows.Next() {
		err = rows.Scan(&book.ID,
			&book.Filename,
			&book.Title,
			&book.AuthorID,
			&book.Description,
			&book.Category,
			&book.LanguageCode,
			&book.Published)
		if err != nil {
			continue
		}
		res = append(res, book)
	}
	return res, err
}

func AddBook(request models.Request, author models.Author, ctx context.Context) (*models.Book, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var book models.Book
	row := conn.QueryRow(ctx,
		`INSERT INTO book (id, filename, title, author_id, description, category, language_code) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING *;`,
		request.ID, request.Filename, request.Title, author.ID, request.Description, request.Category, request.LanguageCode)
	err = row.Scan(&book.ID, &book.Filename, &book.Title, &book.AuthorID, &book.Description, &book.Category, &book.LanguageCode, &book.Published)
	return &book, err
}
