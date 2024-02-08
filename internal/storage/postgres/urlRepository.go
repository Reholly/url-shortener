package postgres

import (
	"github.com/jackc/pgconn"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"url-shortener/internal/domain/entities"
	"url-shortener/internal/domain/repositories"
	"url-shortener/internal/storage"
)

const alreadyExistsErrorCode = "23505"

type UrlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) repositories.UrlRepositoryContract {
	repo := &UrlRepository{db: db}
	return repo
}

func (r *UrlRepository) GetAll() ([]entities.Url, error) {
	urlEntities := make([]entities.Url, 0)
	err := r.db.Select(&urlEntities, "select * from url")

	if err != nil {
		return nil, err
	}

	return urlEntities, nil
}

func (r *UrlRepository) GetByAlias(alias string) (entities.Url, error) {
	var url []entities.Url
	err := r.db.Select(&url, "select u.id, u.url, u.alias from url u where u.alias = $1 ", alias)

	if err != nil {
		return entities.Url{}, storage.NotFoundErr
	}

	return url[0], nil
}

func (r *UrlRepository) GetByUrl(url string) (entities.Url, error) {
	var urlEntity entities.Url
	err := r.db.Select(&urlEntity, "select u.id, u.url, u.alias from url u where u.url = $1 limit 1 ", url)

	if err != nil {
		return urlEntity, storage.NotFoundErr
	}

	return urlEntity, nil
}

func (r *UrlRepository) Add(url entities.Url) error {
	row := r.db.QueryRow("insert into url (url, alias) values ($1,$2)", url.Url, url.Alias)

	if row.Err() != nil {
		var pgxError *pgconn.PgError

		if errors.As(row.Err(), &pgxError) {
			if pgxError.Code == alreadyExistsErrorCode {
				return storage.AlreadyExistsErr
			}
		}
		return row.Err()
	}

	return nil
}

func (r *UrlRepository) Delete(url entities.Url) error {
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
func (r *UrlRepository) Update(url entities.Url) error {
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
