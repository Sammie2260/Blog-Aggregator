package repository

import (
	"hellogo/internal/model"

	"github.com/google/uuid"
)

type FeedRepository interface {
	CreateFeed(feed *model.Feed) (*model.Feed, error)
	ListFeed() ([]*model.Feed, error)
	GetFeed(id uuid.UUID) (*model.Feed, error)
	UpdateFeed(feed *model.Feed) (*model.Feed, error)
	//ReplaceFeed(feed *model.Feed) (*model.Feed, error)
	DeleteFeed(id uuid.UUID) error
}
