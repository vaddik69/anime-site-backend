package comment

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateComment(ctx context.Context, animeID, content string, userID uuid.UUID, parentID *uuid.UUID) (*Comment, error)
	GetComments(ctx context.Context, animeID string, userID uuid.UUID) ([]CommentWithUser, error)
	DeleteComment(ctx context.Context, commentID uuid.UUID, userID uuid.UUID) error
	UpdateComment(ctx context.Context, commentID uuid.UUID, userID uuid.UUID, content string) error
	VoteComment(ctx context.Context, commentID uuid.UUID, userID uuid.UUID, isUpvote bool) error
	RemoveVote(ctx context.Context, commentID uuid.UUID, userID uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateComment(ctx context.Context, animeID, content string, userID uuid.UUID, parentID *uuid.UUID) (*Comment, error) {
	if content == "" {
		return nil, errors.New("comment content cannot be empty")
	}

	if len(content) > 1000 {
		return nil, errors.New("comment is too long")
	}

	comment := &Comment{
		AnimeID:  animeID,
		UserID:   userID,
		Content:  content,
		ParentID: parentID,
	}

	if err := s.repo.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// Добавляем новые методы
func (s *service) VoteComment(ctx context.Context, commentID uuid.UUID, userID uuid.UUID, isUpvote bool) error {

	return s.repo.AddVote(commentID, userID, isUpvote)
}

func (s *service) RemoveVote(ctx context.Context, commentID uuid.UUID, userID uuid.UUID) error {
	return s.repo.RemoveVote(commentID, userID)
}

// Обновляем метод GetComments
func (s *service) GetComments(ctx context.Context, animeID string, userID uuid.UUID) ([]CommentWithUser, error) {
	return s.repo.GetByAnimeID(animeID, userID)
}

func (s *service) DeleteComment(ctx context.Context, commentID uuid.UUID, userID uuid.UUID) error {
	return s.repo.Delete(commentID, userID)
}

func (s *service) UpdateComment(ctx context.Context, commentID uuid.UUID, userID uuid.UUID, content string) error {
	if content == "" {
		return errors.New("comment content cannot be empty")
	}

	return s.repo.Update(&Comment{
		ID:      commentID,
		UserID:  userID,
		Content: content,
	})
}
