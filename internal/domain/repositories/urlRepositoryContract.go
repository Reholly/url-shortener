package repositories

import "url-shortener/internal/domain/entities"

type UrlRepositoryContract interface {
	GetByAlias(alias string) (entities.Url, error)
	GetByUrl(url string) (entities.Url, error)
	GetAll() ([]entities.Url, error)
	Add(url entities.Url) error
	Delete(url entities.Url) error
	Update(url entities.Url) error
}
