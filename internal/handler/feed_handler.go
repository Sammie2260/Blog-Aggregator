package handler

import (
	"hellogo/internal/model"
	"hellogo/internal/response"
	"hellogo/internal/service"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FeedHandler struct {
	Service *service.FeedService
}

// ListFeed godoc
// @Summary Get all feeds
// @Description Fetch list of all feeds
// @Tags feeds
// @Produce json
// @Success 200 {array} model.Feed
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /feeds [get]
func (h *FeedHandler) ListFeed(c echo.Context) error {
	feeds, err := h.Service.ListFeed()
	if err != nil {
		log.Println("Error fetching feeds:", err)
		return response.Error(c, "Failed to fetch feeds")

	}
	if len(feeds) == 0 {
		return response.NotFound(c, "No feeds found")

	}
	return response.Success(c, feeds)

}

// GetFeed godoc
// @Summary Get a feed by ID
// @Description Fetch feed details by UUID
// @Tags feeds
// @Produce json
// @Param id path string true "Feed ID"
// @Success 200 {object} model.Feed
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /feeds/{id} [get]
func (h *FeedHandler) GetFeed(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.BadRequest(c, "Invalid UUID")
	}

	feed, err := h.Service.GetFeed(id)
	if err != nil {
		return response.NotFound(c, "Feed not found")
	}
	return response.Success(c, feed)

}

// CreateFeed godoc
// @Summary Create a new feed
// @Description Add a new feed
// @Tags feeds
// @Accept json
// @Produce json
// @Param feed body model.FeedRequest true "Feed data"
// @Success 201 {object} model.Feed
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /feeds [post]
func (h *FeedHandler) CreateFeed(c echo.Context) error {
	input := new(model.FeedRequest)
	if err := c.Bind(input); err != nil {
		return response.BadRequest(c, err.Error())
	}
	if err := c.Validate(input); err != nil {
		return response.BadRequest(c, err.Error())
	}

	uid, err := uuid.Parse(input.UserID)
	if err != nil {
		return response.BadRequest(c, "Invalid UserID")
	}

	newFeed := &model.Feed{
		Name:   input.Name,
		Url:    input.Url,
		UserID: uid,
	}

	created, err := h.Service.CreateFeed(newFeed)
	if err != nil {
		return response.Error(c, "Failed to create feed")
	}

	return response.Success(c, created)
}

//replacefeed ko kaam basically

// UpdateFeed godoc
// @Summary Update a feed by ID
// @Description Update feed partially by UUID
// @Tags feeds
// @Accept json
// @Produce json
// @Param id path string true "Feed ID"
// @Param feed body model.UpdateFeedStruct true "Feed fields to update"
// @Success 200 {object} model.Feed
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /feeds/{id} [patch]
func (h *FeedHandler) UpdateFeed(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.BadRequest(c, "Invalid UUID")
	}

	input := new(model.UpdateFeedStruct)
	if err := c.Bind(input); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}
	if err := c.Validate(input); err != nil {
		return response.BadRequest(c, err.Error())
	}

	feed := &model.Feed{ID: id}
	if input.Name != nil {
		feed.Name = *input.Name
	}
	if input.Url != nil {
		feed.Url = *input.Url
	}
	if input.UserID != nil {
		feed.UserID = *input.UserID
	}

	updated, err := h.Service.UpdateFeed(feed)
	if err != nil {
		return response.Error(c, "Failed to update feed")
	}

	return response.Success(c, updated)
}

// DeleteFeed godoc
// @Summary Delete a feed by ID
// @Description Soft delete a feed or hard delete using query param ?hard=true
// @Tags feeds
// @Produce json
// @Param id path string true "Feed ID"
// @Param hard query string false "Set true for hard delete"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /feeds/{id} [delete]
func (h *FeedHandler) DeleteFeed(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.BadRequest(c, "Invalid UUID")
	}

	if err := h.Service.DeleteFeed(id); err != nil {
		return response.Error(c, "Failed to delete feed")
	}

	return c.NoContent(http.StatusNoContent)
}
