package models

import "time"

type Event struct {
	ID           int64     `db:"id"`
	TypeID       int64     `db:"type_id"`
	CampaignID   int64     `db:"campaign_id"`
	CostAmount   uint64    `db:"cost_amount"`
	CostCurrency string    `db:"cost_currency"`
	CreatedAt    time.Time `db:"created_at"`
}

type CreateEventInput struct {
	TypeID       int64  `db:"type_id"`
	CampaignID   int64  `db:"campaign_id"`
	CostAmount   uint64 `db:"cost_amount"`
	CostCurrency string `db:"cost_currency"`
}
