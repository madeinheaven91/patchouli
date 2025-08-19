package storage

import (
	"context"
	"errors"

	"catalog/internal/models"
	"catalog/internal/shared"
)

func FetchTag(name string, ctx context.Context) (*models.Tag, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var tag models.Tag
	row := conn.QueryRow(ctx, "SELECT * FROM tag WHERE name=$1", name)
	err = row.Scan(&tag.Name)
	return &tag, err
}

func FetchAllTags(ctx context.Context) ([]models.Tag, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	tags := make([]models.Tag, 0)
	rows, err := conn.Query(ctx, "SELECT * FROM tag")
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	for rows.Next() {
		var tag models.Tag
		errR := rows.Scan(&tag.Name)
		if errR != nil {
			shared.LogError(errR)
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func AddTag(name string, ctx context.Context) (*models.Tag, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var tag models.Tag
	row := conn.QueryRow(ctx,
		`INSERT INTO tag (name) VALUES ($1) 
		RETURNING name;`,
		name)
	err = row.Scan(&tag.Name)
	return &tag, err
}

func DeleteTag(name string, ctx context.Context) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return err
	}
	defer conn.Release()

	cmd, err := conn.Exec(ctx,
		`DELETE FROM tag WHERE name=$1`,
		name)
	if err != nil {
		shared.LogError(err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		err = errors.New("no row found to delete")
	}
	return err
}

func FetchBookTags(bookID string, ctx context.Context) ([]models.Tag, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	tags := make([]models.Tag, 0)
	rows, err := conn.Query(ctx, "SELECT * FROM tag_to_book WHERE book_id=$1", bookID)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	for rows.Next() {
		var tag models.Tag
		err = rows.Scan(&tag.Name, nil)
		if err != nil {
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func FetchRequestTags(requestID string, ctx context.Context) ([]models.Tag, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	tags := make([]models.Tag, 0)
	rows, err := conn.Query(ctx, "SELECT * FROM tag_to_request WHERE request_id=$1", requestID)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	for rows.Next() {
		var tag models.Tag
		err = rows.Scan(&tag.Name, nil)
		if err != nil {
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func AddTagToBook(name string, bookID string, ctx context.Context) (*models.TagToBook, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var tag models.TagToBook
	row := conn.QueryRow(ctx,
		`INSERT INTO tag_to_book (tag_name, book_id) VALUES ($1, $2) 
		RETURNING *;`,
		name, bookID)
	err = row.Scan(&tag.TagName, &tag.BookID)
	return &tag, err
}

func AddTagToRequest(name string, requestID string, ctx context.Context) (*models.TagToRequest, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return nil, err
	}
	defer conn.Release()

	var tag models.TagToRequest
	row := conn.QueryRow(ctx,
		`INSERT INTO tag_to_request (tag_name, request_id) VALUES ($1, $2) 
		RETURNING *;`,
		name, requestID)
	err = row.Scan(&tag.TagName, &tag.RequestID)
	return &tag, err
}

func DeleteTagToBook(name string, bookID string, ctx context.Context) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return err
	}
	defer conn.Release()

	cmd, err := conn.Exec(ctx,
		`DELETE FROM tag_to_book WHERE name=$1 and id=$2`,
		name, bookID)
	if err != nil {
		shared.LogError(err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		err = errors.New("no row found to delete")
	}
	return err
}

func DeleteTagToRequest(name string, requestID string, ctx context.Context) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return err
	}
	defer conn.Release()

	cmd, err := conn.Exec(ctx,
		`DELETE FROM tag_to_request WHERE name=$1 and id=$2`,
		name, requestID)
	if err != nil {
		shared.LogError(err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		err = errors.New("no row found to delete")
	}
	return err
}
