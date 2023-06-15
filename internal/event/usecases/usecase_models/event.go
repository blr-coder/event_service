package usecase_models

import (
	"event_service/internal/event/usecases/usecase_errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Event struct {
	ID           int64     `json:"id"`
	TypeTitle    string    `json:"type_title"`
	CampaignID   int64     `json:"campaign_id"`
	InsertionID  int64     `json:"insertion_id"`
	UserID       int64     `json:"user_id"`
	CostAmount   uint64    `json:"cost_amount"`
	CostCurrency string    `json:"cost_currency"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateEventInput struct {
	TypeTitle   string `json:"type_title"`
	CampaignID  int64  `json:"campaign_id"`
	InsertionID int64  `json:"insertion_id"`
	UserID      int64  `json:"user_id"`
	Cost        *Cost  `json:"cost"`
}

func (i *CreateEventInput) Validate() error {
	err := validation.ValidateStruct(
		i,
		validation.Field(&i.TypeTitle, validation.Required),
		validation.Field(&i.Cost, validation.Required),
	)
	if err != nil {
		// TODO: Is it good solution?
		return usecase_errors.NewValidationErr(err)
	}

	return nil
}

type Cost struct {
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
}

type EventFilter struct {
	TypeTitle    *string       `json:"type_title"`
	CampaignID   *int64        `json:"campaign_id"`
	InsertionID  *int64        `json:"insertion_id"`
	UserID       *int64        `json:"user_id"`
	CostCurrency *string       `json:"cost_currency"`
	SortBy       []EventSortBy `json:"sort_by"`
	SortOrder    *SortOrder    `json:"sort_order"`
	PageSize     *uint64       `json:"page_size"`
	PageNumber   *uint64       `json:"page_number"`
}

type EventSortBy string

const (
	EventSortByCreatedAt    EventSortBy = "created_at"
	EventSortByTypeTitle    EventSortBy = "type_title"
	EventSortByID           EventSortBy = "id"
	EventSortByCostCurrency EventSortBy = "cost_currency"
)

type Events struct {
	Events []*Event
	Count  uint64
}
