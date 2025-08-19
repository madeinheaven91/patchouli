package storage

import (
	"context"
	"errors"

	"github.com/madeinheaven91/patchouli/internal/models"
	"github.com/madeinheaven91/patchouli/internal/service"
	"github.com/madeinheaven91/patchouli/internal/shared"
)

func FetchAuthor(id string, ctx context.Context) (*models.Author, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var author models.Author
	row := conn.QueryRow(ctx, "SELECT * FROM author WHERE id=$1", id)
	err = row.Scan(&author.ID, &author.Name, &author.Description, &author.PhotoURL)
	return &author, err
}

func FetchAllAuthors(ctx context.Context) ([]models.Author, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	authors := make([]models.Author, 0)
	rows, err := conn.Query(ctx, "SELECT * FROM author")
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	for rows.Next() {
		var author models.Author
		errR := rows.Scan(&author.ID, &author.Name, &author.Description, &author.PhotoURL)
		if errR != nil {
			shared.LogError(errR)
			continue
		}
		authors = append(authors, author)
	}
	return authors, err
}

func AddAuthor(form service.AuthorPostForm, ctx context.Context) (*models.Author, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var author models.Author
	row := conn.QueryRow(ctx,
		`INSERT INTO author (name, description, photo_url) 
		VALUES ($1, $2, $3) RETURNING *`,
		form.Name, form.Description, form.PhotoURL)
	err = row.Scan(&author.ID, &author.Name, &author.Description, &author.PhotoURL)
	return &author, err
}

func DeleteAuthor(id string, ctx context.Context) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return err
	}
	defer conn.Release()

	cmd, err := conn.Exec(ctx,
		`DELETE FROM author WHERE id=$1`,
		id)
	if err != nil {
		shared.LogError(err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		// FIXME: make error struct and methods and consts
		// also maybe make a generic function for simple db operations like this one
		err = errors.New("no row found to delete")
	}
	return err
}
