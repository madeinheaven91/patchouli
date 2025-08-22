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
		&book.AuthorName,
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
			&book.AuthorName,
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

func AddBook(form models.Request, ctx context.Context) (*models.Book, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var book models.Book
	row := conn.QueryRow(ctx,
		`INSERT INTO book (id, filename, title, author_name, description, category, language_code) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING *;`,
		form.ID, form.Filename, form.Title, form.AuthorName, form.Description, form.Category, form.LanguageCode)
	err = row.Scan(&book.ID, &book.Filename, &book.Title, &book.AuthorName, &book.Description, &book.Category, &book.LanguageCode, &book.Published)
	return &book, err
}
