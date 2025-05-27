package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Zipklas/anime-site-backend/internal/comment"
	"github.com/Zipklas/anime-site-backend/internal/kodik"
	"github.com/Zipklas/anime-site-backend/internal/user"
	"github.com/Zipklas/anime-site-backend/pkg/database"

	"github.com/Zipklas/anime-site-backend/internal/shikimori"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()

	// Инициализация базы данных
	db := database.InitPostgres()

	// Инициализация сервиса Shikimori
	shikimoriService := shikimori.NewService()
	shikimoriHandler := shikimori.NewHandler(shikimoriService)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, shikimoriService)
	userHandler := user.NewHandler(userService)
	// Создание нового экземпляра Echo
	e := echo.New()

	e.GET("/kodik.txt", func(c echo.Context) error {
		return c.File("kodik.txt") // файл лежит рядом с main.go
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Роуты для регистрации и логина
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
	e.POST("/api/shikimori/search", shikimoriHandler.SearchAnime)
	e.GET("/api/shikimori/top", shikimoriHandler.GetTopAnime)
	e.GET("/api/shikimori/anime/:id", shikimoriHandler.GetAnimeByID)

	commentRepo := comment.NewRepository(db)
	commentService := comment.NewService(commentRepo)
	commentHandler := comment.NewHandler(commentService)

	// Добавляем роуты
	commentGroup := e.Group("/api/comments")
	commentGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	commentGroup.POST("/:anime_id", commentHandler.CreateComment)
	commentGroup.GET("/:anime_id", commentHandler.GetComments)
	commentGroup.DELETE("/:comment_id", commentHandler.DeleteComment)
	commentGroup.PUT("/:comment_id", commentHandler.UpdateComment)
	// Добавляем после других comment роутов
	commentGroup.PUT("/:comment_id/vote", commentHandler.VoteComment)
	commentGroup.DELETE("/:comment_id/vote", commentHandler.RemoveVote)
	// Добавляем после инициализации других сервисов
	kodikService := kodik.NewService("None")
	kodikHandler := kodik.NewHandler(kodikService)

	// Добавляем роуты для плеера
	e.GET("/api/kodik/search", kodikHandler.SearchVideos)
	e.GET("/api/kodik/videos/:shikimori_id", kodikHandler.GetVideoOptions)

	// Защищенная группа для просмотра
	playerGroup := e.Group("/player")
	playerGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	playerGroup.GET("/:video_id", func(c echo.Context) error {
		// Здесь будет обработчик для самого плеера
		return c.JSON(http.StatusOK, echo.Map{"status": "under construction"})
	})

	// Группа роутов для профиля (с защитой JWT)
	r := e.Group("/profile")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c echo.Context, err error) error {
			log.Printf("Error validating JWT token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		},
	}))

	// Обработчик запроса на получение профиля
	r.GET("", userHandler.Profile)
	r.POST("/watched/:anime_id", userHandler.AddWatched)   // POST /profile/watched/:anime_id
	r.POST("/favorite/:anime_id", userHandler.AddFavorite) // POST /profile/favorite/:anime_id
	r.GET("/watched", userHandler.GetWatchedAnime)         // GET /profile/watched
	r.GET("/favorite", userHandler.GetFavouriteAnime)
	// Запуск сервера
	log.Fatal(e.Start(":8080"))
}
