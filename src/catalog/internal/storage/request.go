package storage

import (
	"context"

	"catalog/internal/models"
	"catalog/internal/service"
	"catalog/internal/shared"
)

func FetchRequest(id string, ctx context.Context) (*models.Request, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var req models.Request
	row := conn.QueryRow(ctx, "SELECT * FROM request WHERE id=$1", id)
	err = row.Scan(&req.ID,
		&req.Filename,
		&req.Title,
		&req.AuthorName,
		&req.Description,
		&req.Category,
		&req.LanguageCode,
		&req.Added)
	return &req, err
}

func FetchAllRequests(ctx context.Context) ([]models.Request, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var req models.Request
	res := make([]models.Request, 0)
	rows, err := conn.Query(ctx, "SELECT * FROM request")
	for rows.Next() {
		err = rows.Scan(&req.ID,
			&req.Filename,
			&req.Title,
			&req.AuthorName,
			&req.Description,
			&req.Category,
			&req.LanguageCode,
			&req.Added)
		if err != nil {
			continue
		}
		res = append(res, req)
	}
	return res, err
}

func AddRequest(form service.RequestPostForm, ctx context.Context) (*models.Request, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var request models.Request
	row := conn.QueryRow(ctx,
		`INSERT INTO request (filename, title, author_name, description, category, language_code) 
		VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING *;`,
		form.Filename, form.Title, form.AuthorName, form.Description, form.Category, form.LanguageCode)
	err = row.Scan(&request.ID, &request.Filename, &request.Title, &request.AuthorName, &request.Description, &request.Category, &request.LanguageCode, &request.Added)
	return &request, err
}
