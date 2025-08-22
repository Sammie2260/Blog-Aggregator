package repository

import (
	"hellogo/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedRepositoryGorm struct {
	DB *gorm.DB
}

func (r *FeedRepositoryGorm) CreateFeedRepo(feed *model.Feed) (*model.Feed, error) {
	if err := r.DB.Create(feed).Error; err != nil {
		return nil, err
	}
	return feed, nil
}

func (r *FeedRepositoryGorm) GetFeedRepo(id uuid.UUID) (*model.Feed, error) {
	var feed model.Feed
	if err := r.DB.First(&feed, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &feed, nil
}


func (r *FeedRepositoryGorm) ListFeedRepo() ([]*model.Feed, error) {
    var feeds []*model.Feed
    if err := r.DB.Find(&feeds).Error; err != nil {
        return nil, err
    }
    return feeds, nil
}



func (r *FeedRepositoryGorm) UpdateFeedRepo(feed *model.Feed) (*model.Feed, error) {
	if err := r.DB.Save(feed).Error; err != nil {
		return nil, err
	}
	return feed, nil
}

func (r *FeedRepositoryGorm) DeleteFeedRepo(id uuid.UUID) error {
	return r.DB.Delete(&model.Feed{}, "id = ?", id).Error
}
