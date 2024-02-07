package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/server/dto"
	"urlshortener/internal/storage"
)

type RemoveUrlHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewRemoveUrlHandler(repo storage.UrlRepositoryContract) *RemoveUrlHandler {
	return &RemoveUrlHandler{UrlRepo: repo}
}

func (h *RemoveUrlHandler) RemoveUrl(c echo.Context) error {
	req := new(dto.UrlDto)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, err := h.UrlRepo.GetByUrl(req.Url)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = h.UrlRepo.Delete(url)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
