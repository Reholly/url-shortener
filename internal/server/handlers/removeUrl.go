package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/storage"
)

type removeUrlRequest struct {
	Url string `json:"url" bson:"url"`
}

type RemoveUrlHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewRemoveUrlHandler(repo storage.UrlRepositoryContract) *RemoveUrlHandler {
	return &RemoveUrlHandler{UrlRepo: repo}
}

func (h *RemoveUrlHandler) RemoveUrl(c echo.Context) error {
	req := new(removeUrlRequest)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, err := h.UrlRepo.GetByUrl(req.Url)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, url)
}
