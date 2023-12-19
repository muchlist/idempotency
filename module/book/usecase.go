package book

import "context"

type BookUsecaseAssumer interface {
	Get(ctx context.Context, isbn string) (BookEntity, error)
	Insert(ctx context.Context, book BookEntity) (BookEntity, error)
	Update(ctx context.Context, book BookEntity) (BookEntity, error)
}

type Usecase struct {
	Repo BookStorer
}

func NewUsecase(repo BookStorer) *Usecase {
	return &Usecase{
		Repo: repo,
	}
}

func (u *Usecase) Get(ctx context.Context, isbn string) (BookEntity, error) {
	book, err := u.Repo.GetBook(ctx, isbn)
	if err != nil {
		return BookEntity{}, err
	}

	return book, nil
}

func (u *Usecase) Insert(ctx context.Context, book BookEntity) (BookEntity, error) {
	if err := u.Repo.InsertBook(ctx, book); err != nil {
		return BookEntity{}, err
	}

	return book, nil
}

func (u *Usecase) Update(ctx context.Context, book BookEntity) (BookEntity, error) {
	if err := u.Repo.UpdateBook(ctx, book); err != nil {
		return BookEntity{}, err
	}

	return book, nil
}
