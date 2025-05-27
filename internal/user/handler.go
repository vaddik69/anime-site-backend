package user

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Register(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := h.service.Register(req.Email, req.Password); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "registered"})
}

func (h *Handler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	log.Println(token)
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (h *Handler) Profile(c echo.Context) error {
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user_id missing")
	}

	user, err := h.service.GetProfile(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":            user.ID,
		"email":              user.Email,
		"watched_anime_ids":  user.WatchedAnimeIDs,
		"favorite_anime_ids": user.FavoriteAnimeIDs,
	})
}

func (h *Handler) AddWatched(c echo.Context) error {
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user_id missing")
	}

	animeID := c.Param("anime_id")
	if animeID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "anime_id is required")
	}

	if err := h.service.AddWatched(userID, animeID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":   "added to watched",
		"anime_id": animeID,
	})
}

func (h *Handler) AddFavorite(c echo.Context) error {
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user_id missing")
	}

	animeID := c.Param("anime_id")
	if animeID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "anime_id is required")
	}

	if err := h.service.AddFavorite(userID, animeID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":   "added to favorites",
		"anime_id": animeID,
	})
}
func (h *Handler) GetWatchedAnime(c echo.Context) error {
	// Получаем userID из JWT токена
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user_id missing")
	}

	// Получаем список аниме с деталями
	animeList, err := h.service.GetWatchedAnimeDetails(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, animeList)
}
func (h *Handler) GetFavouriteAnime(c echo.Context) error {
	// Получаем userID из JWT токена
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user_id missing")
	}

	// Получаем список аниме с деталями
	animeList, err := h.service.GetFavouriteAnimeDetails(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, animeList)
}
