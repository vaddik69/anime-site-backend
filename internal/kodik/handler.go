package kodik

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SearchVideos(c echo.Context) error {
	title := c.QueryParam("title")
	if title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "title parameter is required")
	}

	videos, err := h.service.SearchAnime(c.Request().Context(), title)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, videos)
}

func (h *Handler) GetVideoOptions(c echo.Context) error {
	shikimoriID := c.Param("shikimori_id")
	if shikimoriID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "shikimori_id is required")
	}

	videos, err := h.service.GetVideoByShikimoriID(c.Request().Context(), shikimoriID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, videos)
}
