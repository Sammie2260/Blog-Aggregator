package service

import (
	"errors"
	"fmt"
	"hellogo/internal/model"
	"hellogo/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedService struct {
	Repo repository.FeedRepository
}

func NewFeedService(r repository.FeedRepository) FeedService{
	return FeedService{
		Repo: r,
	}
}

// Get all feeds
func (s *FeedService) ListFeedService() ([]map[string]interface{}, error) {
	feeds, err := s.Repo.ListFeedRepo()
	if err != nil {
		return nil, err
	}

	responses := make([]map[string]interface{}, len(feeds))
	for i, feed := range feeds {
		responses[i] = map[string]interface{}{
			"id":   feed.ID,
			"name": feed.Name,
			"url":  feed.Url,
		}
	}

	return responses, nil
}

// Get feed by ID
func (s *FeedService) GetFeedService(id uuid.UUID) (map[string]interface{}, error) {
	feed, err := s.Repo.GetFeedRepo(id)
	if err != nil {
		return nil, err
	}

	// Shape the response directly in service
	response := map[string]interface{}{
		"id":   feed.ID,
		"name": feed.Name,
		"url":  feed.Url,
	}

	return response, nil
}

// Create new feed
func (s *FeedService) CreateFeedService(req *model.CreateFeedRequest) (*model.Feed, error) {
	uid, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	feed := &model.Feed{
		ID:     uuid.New(),
		Name:   req.Name,
		Url:    req.Url,
		UserID: uid,
	}

	if _, err := s.Repo.CreateFeedRepo(feed); err != nil {
		return nil, err
	}

	// DOmain lai data transfer obj banauni
	return &model.Feed{
		ID:   feed.ID,
		Name: feed.Name,
		Url:  feed.Url,
	}, nil
}

// Update feed
func (s *FeedService) UpdateFeedService(id uuid.UUID, input *model.UpdateFeedStruct) (*model.Feed, error) {
	existingFeed, err := s.Repo.GetFeedRepo(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("feed not found")
		}
		return nil, fmt.Errorf("failed to fetch feed: %w", err)
	}

	if input.Name != nil {
		existingFeed.Name = *input.Name
	}
	if input.Url != nil {
		existingFeed.Url = *input.Url
	}
	if input.UserID != nil {
		existingFeed.UserID = *input.UserID
	}

	updatedFeed, err := s.Repo.UpdateFeedRepo(existingFeed)
	if err != nil {
		return nil, fmt.Errorf("failed to update feed: %w", err)
	}

	return updatedFeed, nil
}

// Delete feed
func (s *FeedService) DeleteFeedService(id uuid.UUID) error {
	return s.Repo.DeleteFeedRepo(id)
}
