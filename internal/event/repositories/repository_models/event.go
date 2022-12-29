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

type EventSortBy string

func (o EventSortBy) String() string {
	return string(o)
}

type EventSortByList []EventSortBy

func (l EventSortByList) Join() (strList []string) {
	for _, val := range l {
		strList = append(strList, val.String())
	}

	return strList
}

const (
	EventSortByCreatedAt    EventSortBy = "created_at"
	EventSortByTypeTitle    EventSortBy = "type_title"
	EventSortByID           EventSortBy = "id"
	EventSortByCostCurrency EventSortBy = "cost_currency"
)

type EventRepositoryFilter struct {
	TypeTitle    *string         `db:"type_title"`
	CampaignID   *int64          `db:"campaign_id"`
	InsertionID  *int64          `db:"insertion_id"`
	UserID       *int64          `db:"user_id"`
	CostCurrency *string         `db:"cost_currency"`
	SortBy       EventSortByList `db:"sort_by"`
	SortOrder    *SortOrder      `db:"sort_order"`
	PageSize     *uint64         `db:"page_size"`
	PageNumber   *uint64         `db:"page_number"`
}
