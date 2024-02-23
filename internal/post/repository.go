package post

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/google/uuid"
	"github.com/thiagosena/gopost/internal"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func (r *Repository) Insert(post internal.Post) (internal.Post, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println(r.Conn)

	err := r.Conn.QueryRow(
		ctx,
		"INSERT INTO posts (username, body) VALUES ($1, $2) RETURNING id",
		post.Username,
		post.Body).Scan(&post.ID)

	if err != nil {
		return internal.Post{}, err
	}

	return post, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tag, err := r.Conn.Exec(
		ctx,
		"DELETE FROM posts WHERE id = $1",
		id)

	if tag.RowsAffected() == 0 {
		return ErrPostNotFound
	}

	return err
}

func (r *Repository) FindOneByID(id uuid.UUID) (internal.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var post internal.Post
	err := r.Conn.QueryRow(
		ctx,
		"SELECT username, body, created_at FROM posts WHERE id = $1",
		id).Scan(&post.Username, &post.Body, &post.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return internal.Post{}, ErrPostNotFound
	}

	if err != nil {
		return internal.Post{}, err
	}

	return post, nil
}

func (r *Repository) FindAll() (internal.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var post internal.Post
	err := r.Conn.QueryRow(
		ctx,
		"SELECT id, username, body, created_at FROM posts",
	).Scan(&post.Username, &post.Body, &post.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return internal.Post{}, ErrPostNotFound
	}

	if err != nil {
		return internal.Post{}, err
	}

	return post, nil
}
