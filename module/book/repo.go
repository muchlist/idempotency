package book

import (
	"context"
	"errors"
)

type BookStorer interface {
	GetBook(_ context.Context, isbn string) (BookEntity, error)
	InsertBook(_ context.Context, input BookEntity) error
	UpdateBook(_ context.Context, input BookEntity) error
}

type Repo struct {
	storage map[string]BookEntity
}

func NewRepo() *Repo {
	return &Repo{
		storage: make(map[string]BookEntity),
	}
}

func (r *Repo) GetBook(_ context.Context, isbn string) (BookEntity, error) {
	val, ok := r.storage[isbn]
	if !ok {
		return BookEntity{}, errors.New("isbn not found")
	}
	return val, nil
}

func (r *Repo) InsertBook(_ context.Context, input BookEntity) error {
	if _, ok := r.storage[input.ISBN]; ok {
		return errors.New("duplicate isbn")
	}

	r.storage[input.ISBN] = input
	return nil
}

func (r *Repo) UpdateBook(_ context.Context, input BookEntity) error {
	if _, ok := r.storage[input.ISBN]; !ok {
		return errors.New("isbn not found")
	}

	r.storage[input.ISBN] = input
	return nil
}
