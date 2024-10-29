package repositories

import (
	"context"
	"strconv"

	"github.com/HemlockPham7/backend/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	// mock data
	// events = append(events, &models.Event{
	// 	ID:        "0092345902384",
	// 	Name:      "Super Event",
	// 	Location:  "Somewhere",
	// 	Date:      time.Now(),
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })

	// return events, nil
	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	id, _ := strconv.Atoi(eventId) // turn string into integer
	event := &models.Event{}

	res := r.db.Model(event).Where("id = ?", uint(id)).First(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.Create(&event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) UpdateOne(ctx context.Context, eventId string, updateData map[string]interface{}) (*models.Event, error) {
	id, _ := strconv.Atoi(eventId)
	event := &models.Event{}

	updateRes := r.db.Model(event).Where("id = ?", uint(id)).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Model(event).Where("id = ?", uint(id)).First(event)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, eventId string) error {
	res := r.db.Delete(&models.Event{}, eventId)
	return res.Error
}
