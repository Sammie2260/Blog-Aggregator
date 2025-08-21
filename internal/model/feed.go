package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feed struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"id"`
	Name      string         `gorm:"not null" json:"name" validate:"required,min=1"`
	Url       string         `gorm:"not null" json:"url" validate:"required,url"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id" validate:"required"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt" validate:"datetime=2006-01-02"`
}

// type user struct {
// 	gorm.Model
// 	Name    string
// 	Address string
// }

// yo chai hook create agi call huncha
func (f *Feed) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return
}

type UpdateFeedStruct struct {
	Name   *string    `json:"name" validate:"omitempty,min=1"`
	Url    *string    `json:"url"  validate:"omitempty,url"`
	UserID *uuid.UUID `json:"user_id" validate:"omitempty,uuid"`
}

type FeedRequest struct {
	Name   string `json:"name" validate:"required"`
	Url    string `json:"url" validate:"required,url"`
	UserID string `json:"user_id" validate:"required,uuid"`
} //esma naam ferya chha so thunderclient garda error aauna sakcha publishdate-->createdAt

type FeedResponse struct {
    ID   uuid.UUID `json:"id"`
    Name string    `json:"name"`
    Url  string    `json:"url"`
}