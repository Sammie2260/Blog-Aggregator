package service

import (
	"hellogo/internal/model"
	"hellogo/internal/repository"

	"github.com/google/uuid"
)

type FeedService struct {
	Repo repository.FeedRepository
}

// Get all feeds
func (s *FeedService) ListFeed() ([]*model.FeedResponse, error) {
	feeds, err := s.Repo.ListFeed()
	if err != nil {
		return nil, err
	}

	responses := make([]*model.FeedResponse, len(feeds))
	for i, feed := range feeds {
		responses[i] = &model.FeedResponse{
			ID:   feed.ID,
			Name: feed.Name,
			Url:  feed.Url,
		}
	}

	return responses, nil
}

// Get feed by ID
func (s *FeedService) GetFeed(id uuid.UUID) (*model.FeedResponse, error) {
	feed, err := s.Repo.GetFeed(id)
	if err != nil {
		return nil, err
	}

	response := &model.FeedResponse{
		ID:   feed.ID,
		Name: feed.Name,
		Url:  feed.Url,
	}

	return response, nil
}

// Create new feed
func (s *FeedService) CreateFeed(feed *model.Feed) (*model.Feed, error) {
	feed.ID = uuid.New()
	return s.Repo.CreateFeed(feed)
}

// Update feed
func (s *FeedService) UpdateFeed(feed *model.Feed) (*model.Feed, error) {
	return s.Repo.UpdateFeed(feed)
}

// Delete feed
func (s *FeedService) DeleteFeed(id uuid.UUID) error {
	return s.Repo.DeleteFeed(id)
}
