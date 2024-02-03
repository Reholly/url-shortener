package handlers

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
	"urlshortener/internal/domain"
	"urlshortener/internal/storage"
)

const size = 10

type addUrlRequest struct {
	Url string `json:"url" bson:"url"`
}

type AddUrlHandler struct {
	UrlRepo       storage.UrlRepositoryContract
	urlStartsWith string
}

func NewAddUrlHandler(startsWith string, repo storage.UrlRepositoryContract) *AddUrlHandler {
	return &AddUrlHandler{urlStartsWith: startsWith, UrlRepo: repo}
}

func (h *AddUrlHandler) AddUrl(c echo.Context) error {
	req := new(addUrlRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	alias := h.CreateUrlAlias()

	err := h.UrlRepo.Add(domain.Url{Alias: alias, Url: req.Url})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, req)

}

func (h *AddUrlHandler) CreateUrlAlias() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	rndStr := string(b)

	return h.urlStartsWith + rndStr
}
