package postgres

import (
	"github.com/jackc/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"urlshortener/internal/domain"
	"urlshortener/internal/storage"
)

type UrlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) GetAll() ([]domain.Url, error) {
	urlEntities := make([]domain.Url, 0)
	err := r.db.Select(&urlEntities, "select * from url")

	if err != nil {
		return nil, errors.Wrap(err, "could not select")
	}

	return urlEntities, nil
}

func (r *UrlRepository) GetByAlias(alias string) (domain.Url, error) {
	var url domain.Url
	err := r.db.Select(&url, "select u.id, u.url, u.alias from url u where u.alias = $1 limit 1 ", alias)

	if err != nil {
		return url, storage.NotFoundErr
	}

	return url, nil
}

func (r *UrlRepository) GetByUrl(url string) (domain.Url, error) {
	var urlEntity domain.Url
	err := r.db.Select(&urlEntity, "select u.id, u.url, u.alias from url u where u.url = $1 limit 1 ", url)

	if err != nil {
		return urlEntity, storage.NotFoundErr
	}

	return urlEntity, nil
}

func (r *UrlRepository) Add(url domain.Url) error {
	err := r.db.QueryRow("insert into url (url, alias) values ($1,$2)", url.Url, url.Alias)

	if err != nil {
		var pgxError *pgconn.PgError

		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return storage.AlreadyExistsErr
			}
		}
		return err
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
func (r *UrlRepository) Update(url domain.Url) error {
	alias, err := r.GetByAlias(url.Alias)

	if err != nil {
		return errors.Wrap(err, "could not get by alias")
	}

	_, err = r.db.Exec("update url set url = $1, alias = $2 url where id = $3", alias.Url, alias.Alias, alias.Id)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return storage.AlreadyExistsErr
			}
		}
		return err
	}

	return nil
}
