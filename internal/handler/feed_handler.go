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
func (h *FeedHandler) ListFeedHandler(c echo.Context) error {
	feeds, err := h.Service.ListFeedService()
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
func (h *FeedHandler) GetFeedHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.BadRequest(c, "Invalid UUID")
	}

	feed, err := h.Service.GetFeedService(id)
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
// @Param feed body model.CreateFeedRequest true "Feed data"
// @Success 201 {object} model.Feed
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /feeds [post]
func (h *FeedHandler) CreateFeedHandler(c echo.Context) error {
    input := new(model.CreateFeedRequest)
    if err := c.Bind(input); err != nil {
        return response.BadRequest(c, err.Error())
    }
    if err := c.Validate(input); err != nil {
        return response.BadRequest(c, err.Error())
    }

    created, err := h.Service.CreateFeedService(input)
    if err != nil {
        return response.Error(c, "Failed to create feed")
    }

    return response.Success(c, created) // created is a DTO, not the DB model
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
func (h *FeedHandler) UpdateFeedHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.BadRequest(c, "Invalid UUID")
	}

	var input model.UpdateFeedStruct
	if err := c.Bind(&input); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}
	if err := c.Validate(&input); err != nil {
		return response.BadRequest(c, err.Error())
	}

	updatedFeed, err := h.Service.UpdateFeedService(id, &input)
	if err != nil {
		if err.Error() == "feed not found" {
			return response.NotFound(c, "Feed not found")
		}
		return response.Error(c, err.Error())
	}

	return response.Success(c, updatedFeed)
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
func (h *FeedHandler) DeleteFeedHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.BadRequest(c, "Invalid UUID")
	}

	if err := h.Service.DeleteFeedService(id); err != nil {
		return response.Error(c, "Failed to delete feed")
	}

	return c.NoContent(http.StatusNoContent)
}
