package mapping

import (
	"urlshortener/internal/domain"
	"urlshortener/internal/server/dto"
)

func UrlToUrlWithAliasDto(urlEntity domain.Url) dto.UrlWithAliasDto {
	return dto.UrlWithAliasDto{Url: urlEntity.Url, Alias: urlEntity.Alias}
}
