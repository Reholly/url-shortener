package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortener/internal/domain/entities"
	"url-shortener/internal/domain/repositories"
	"url-shortener/internal/server/dto"
	"url-shortener/internal/utils"
)

const (
	aliasParam = "alias"
)

func AddUrl(repo repositories.UrlRepositoryContract) func(c echo.Context) error {
	return func(c echo.Context) error {
		var req dto.UrlDto
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		alias := utils.GenerateUrlAlias()

		if err := repo.Add(entities.Url{Alias: alias, Url: req.Url}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, nil)
	}
}

func GetUrlByAlias(repo repositories.UrlRepositoryContract) func(c echo.Context) error {
	return func(c echo.Context) error {
		alias := c.QueryParam(aliasParam)

		url, err := repo.GetByAlias(alias)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, dto.UrlToUrlWithAliasDto(url))
	}
}

func RedirectOnUrl(repo repositories.UrlRepositoryContract) func(c echo.Context) error {
	return func(c echo.Context) error {
		alias := c.Param(aliasParam)

		url, err := repo.GetByAlias(alias)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.Redirect(http.StatusPermanentRedirect, url.Url)
	}
}

func RemoveUrl(repo repositories.UrlRepositoryContract) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(dto.UrlDto)

		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		url, err := repo.GetByUrl(req.Url)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		err = repo.Delete(url)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, nil)
	}
}
