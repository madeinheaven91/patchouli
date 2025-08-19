package storage

import (
	"context"

	"github.com/madeinheaven91/patchouli/internal/models"
	"github.com/madeinheaven91/patchouli/internal/service"
	"github.com/madeinheaven91/patchouli/internal/shared"
)

// func FetchRequest(name string, ctx context.Context) (*models.Tag, error) {
// 	conn, err := pool.Acquire(ctx)
// 	if err != nil {
// 		shared.LogError(err)
// 		return nil, err
// 	}
// 	defer conn.Release()
//
// 	var tag models.Tag
// 	row := conn.QueryRow(ctx, "SELECT * FROM tag WHERE name=$1", name)
// 	err = row.Scan(&tag.Name)
// 	return &tag, err
// }
//
// func FetchRequestDocument(ctx context.Context) ([]models.Tag, error) {
// 	conn, err := pool.Acquire(ctx)
// 	if err != nil {
// 		shared.LogError(err)
// 		return nil, err
// 	}
// 	defer conn.Release()
//
// 	tags := make([]models.Tag, 0)
// 	rows, err := conn.Query(ctx, "SELECT * FROM tag")
// 	if err != nil {
// 		shared.LogError(err)
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var tag models.Tag
// 		errR := rows.Scan(&tag.Name)
// 		if errR != nil {
// 			shared.LogError(errR)
// 			continue
// 		}
// 		tags = append(tags, tag)
// 	}
// 	return tags, nil
// }

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
