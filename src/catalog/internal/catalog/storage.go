package catalog

import (
	"context"
	"fmt"
	"log"

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

func FetchBook(id string, ctx context.Context) (*Book, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer conn.Release()

	var book Book
	row := conn.QueryRow(ctx, "SELECT * FROM book WHERE id=$1", id)
	err = row.Scan(&book.ID,
		&book.FilePath,
		&book.Title,
		&book.Description,
		&book.Format,
		&book.CategoryID,
		&book.LanguageCode,
		&book.Published)
	return &book, err
}

func FetchCategory(id int, ctx context.Context) (*Category, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer conn.Release()

	var category Category
	row := conn.QueryRow(ctx, "SELECT * FROM category WHERE id=$1", id)
	err = row.Scan(&category.ID, &category.Name)
	return &category, err
}

func AddCategory(name string, ctx context.Context) (*Category, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer conn.Release()

	var category Category
	row := conn.QueryRow(ctx,
		`INSERT INTO category (name) VALUES ($1) 
		RETURNING id, name;`,
		name)
	err = row.Scan(&category.ID, &category.Name)
	return &category, err
}

func FetchTag(id int, ctx context.Context) (*Tag, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer conn.Release()

	var tag Tag
	row := conn.QueryRow(ctx, "SELECT * FROM tag WHERE id=$1", id)
	err = row.Scan(&tag.ID, &tag.Name)
	return &tag, err
}

func AddTag(name string, ctx context.Context) (*Tag, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer conn.Release()

	var tag Tag
	row := conn.QueryRow(ctx,
		`INSERT INTO tag (name) VALUES ($1) 
		RETURNING id, name;`,
		name)
	err = row.Scan(&tag.ID, &tag.Name)
	return &tag, err
}

func FetchAuthor(id int, ctx context.Context) (*Author, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer conn.Release()

	var author Author
	row := conn.QueryRow(ctx, "SELECT * FROM author WHERE id=$1", id)
	err = row.Scan(&author.ID, &author.Name, &author.Description, &author.PhotoURL)
	return &author, err
}

func AddAuthor(form AuthorPostForm, ctx context.Context) (*Author, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer conn.Release()

	var author Author
	row := conn.QueryRow(ctx,
		`INSERT INTO author (name, description, photo_url) 
		VALUES ($1, $2, $3) RETURNING *`,
		form.Name, form.Description, form.PhotoURL)
	err = row.Scan(&author.ID, &author.Name, &author.Description, &author.PhotoURL)
	return &author, err
}
