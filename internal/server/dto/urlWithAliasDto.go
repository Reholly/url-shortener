package dto

import "url-shortener/internal/domain/entities"

type UrlWithAliasDto struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

func UrlToUrlWithAliasDto(urlEntity entities.Url) UrlWithAliasDto {
	return UrlWithAliasDto{Url: urlEntity.Url, Alias: urlEntity.Alias}
}
