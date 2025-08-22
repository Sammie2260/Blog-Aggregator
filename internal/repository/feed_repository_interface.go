package repository

import (
	"hellogo/internal/model"

	"github.com/google/uuid"
)

type FeedRepository interface {
	CreateFeedRepo(feed *model.Feed) (*model.Feed, error)
	ListFeedRepo() ([]*model.Feed, error)
	GetFeedRepo(id uuid.UUID) (*model.Feed, error)
	UpdateFeedRepo(feed *model.Feed) (*model.Feed, error)
	//ReplaceFeed(feed *model.Feed) (*model.Feed, error)
	DeleteFeedRepo(id uuid.UUID) error
}
