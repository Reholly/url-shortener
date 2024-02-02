package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"urlshortener/internal/domain"
	"urlshortener/internal/storage"
)

type UrlRepositoryContract interface {
	GetByAlias(id int64) ([]domain.Url, error)
	GetAll() ([]domain.Url, error)
	Add(url domain.Url) error
	Delete(url domain.Url) error
}

type UrlRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) GetAll() ([]domain.Url, error) {
	urls := make([]domain.Url, 0)
	err := r.db.Select(&urls, "select * from url")

	if err != nil {
		return nil, errors.Wrap(err, "could not select")
	}

	return urls, nil
}

func (r *UrlRepository) GetByAlias(alias string) (domain.Url, error) {
	var url domain.Url
	err := r.db.Select(&url, "select u.id, u.url, u.alias from url u where u.alias = $1 limit 1 ", alias)

	if err != nil {
		return url, storage.NotFoundErr
	}

	return url, nil
}

func (r *UrlRepository) Add(url domain.Url) error {
	_, err := r.db.Exec("insert into url (url, alias) values ($1,$2)", url.Url, url.Alias)

	if err != nil {
		return errors.Wrap(err, "could not select")
	}

	return nil
}

func (r *UrlRepository) Delete(url domain.Url) error {
	_, err := r.GetByAlias(url.Alias)

	if err != nil {
		return errors.Wrap(err, "could not get by alias")
	}

	_, err = r.db.Exec("delete from url where alias = $1", url.Alias)

	if err != nil {
		return errors.Wrap(err, "could not delete")
	}

	return nil
}
