package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/storage"
)

type getUrlByAliasRequest struct {
	Alias string `json:"alias" bson:"alias"`
	Url   string `json:"url" bson:"url"`
}

type GetUrlByAliasHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewGetUrlByAliasHandler(repo storage.UrlRepositoryContract) *GetUrlByAliasHandler {
	return &GetUrlByAliasHandler{UrlRepo: repo}
}

func (h *GetUrlByAliasHandler) GetUrlByAlias(c echo.Context) error {
	req := new(getUrlByAliasRequest)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, err := h.UrlRepo.GetByAlias(req.Alias)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, url)
}
