package repository_models

import "time"

type CreateEventTypeRepositoryDTO struct {
	Title string `db:"title"`
}

type EventTypeRepositoryDTO struct {
	Title     string     `db:"title"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type OrderBy string

func (o OrderBy) String() string {
	return string(o)
}

type OrderByList []OrderBy

func (l OrderByList) Join() (strList []string) {
	for _, val := range l {
		strList = append(strList, val.String())
	}

	return strList
}

const (
	OrderByTypeCreatedAt OrderBy = "created_at"
	OrderByTypeTitle     OrderBy = "title"
	OrderByTypeID        OrderBy = "id"
)

type EventTypeRepositoryFilter struct {
	Titles     []string    `db:"titles"`
	Search     *string     `db:"search"`
	OrderBy    OrderByList `db:"order_by"`
	SortOrder  *SortOrder  `db:"sort_order"`
	PageSize   *uint64     `db:"page_size"`
	PageNumber *uint64     `db:"page_number"`
}

type UpdateEventTypeRepositoryDTO struct {
	Title    string `db:"title"`
	NewTitle string `db:"new_title"`
}

type DeleteEventTypeRepositoryDTO struct {
	Title string `db:"title"`
}
