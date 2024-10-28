package repositories

import (
	"context"
	"time"

	"github.com/HemlockPham7/backend/models"
)

type EventRepository struct {
	db any
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	// mock data
	events = append(events, &models.Event{
		ID:        "0092345902384",
		Name:      "Super Event",
		Location:  "Somewhere",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	return nil, nil
}
