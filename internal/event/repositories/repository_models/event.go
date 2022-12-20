package repository_models

import "time"

type CreateEventRepositoryDTO struct {
	TypeTitle    string `db:"type_title"`
	CampaignID   int64  `db:"campaign_id"`
	InsertionID  int64  `db:"insertion_id"`
	UserID       int64  `db:"user_id"`
	CostAmount   uint64 `db:"cost_amount"`
	CostCurrency string `db:"cost_currency"`
}

type EventRepositoryDTO struct {
	ID           int64     `db:"id"`
	TypeTitle    string    `db:"type_title"`
	CampaignID   int64     `db:"campaign_id"`
	InsertionID  int64     `db:"insertion_id"`
	UserID       int64     `db:"user_id"`
	CostAmount   uint64    `db:"cost_amount"`
	CostCurrency string    `db:"cost_currency"`
	CreatedAt    time.Time `db:"created_at"`
}
