package database

import (
	"log"
	"os"

	"github.com/Zipklas/anime-site-backend/internal/comment"
	"github.com/Zipklas/anime-site-backend/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	_ = db.AutoMigrate(&user.User{})
	_ = db.AutoMigrate(&comment.Comment{}, &comment.CommentVote{})
	return db
}
