package post

import (
	"errors"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/thiagosena/gopost/internal"
)

var ErrPostBodyEmpty = errors.New("post body is empty")
var ErrPostBodyExceedsLimit = errors.New("post body exceeds limit")
var ErrPostNotFound = errors.New("post not found")

type Service struct {
	Repository Repository
}

func (service Service) Create(post internal.Post) (internal.Post, error) {
	if post.Body == "" {
		return internal.Post{}, ErrPostBodyEmpty
	}

	if utf8.RuneCountInString(post.Body) > 140 {
		return internal.Post{}, ErrPostBodyExceedsLimit
	}

	return service.Repository.Insert(post)
}

func (service Service) Delete(id uuid.UUID) error {
	return service.Repository.Delete(id)
}

func (service Service) FindOneByID(id uuid.UUID) (internal.Post, error) {
	return service.Repository.FindOneByID(id)
}

func (service Service) FindAll() (internal.Post, error) {
	return service.Repository.FindAll()
}
