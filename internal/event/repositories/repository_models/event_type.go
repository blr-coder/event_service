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
	//var strList []string
	for _, val := range l {
		strList = append(strList, val.String())
	}

	return strList
}

const (
	OrderByTypeCreatedAt OrderBy = "created_at"
	OrderByTypeTitle     OrderBy = "title"
)

type OrderDirection string

func (d OrderDirection) String() string {
	return string(d)
}

const (
	OrderDirectionTypeASC  OrderDirection = "ASC"
	OrderDirectionTypeDESC OrderDirection = "DESC"
)

type EventTypeRepositoryFilter struct {
	Titles         []string        `db:"titles"`
	Search         *string         `db:"search"`
	OrderBy        OrderByList     `db:"order_by"`
	OrderDirection *OrderDirection `db:"order_direction"`
}

type UpdateEventTypeRepositoryDTO struct {
	Title    string `db:"title"`
	NewTitle string `db:"new_title"`
}

type DeleteEventTypeRepositoryDTO struct {
	Title string `db:"title"`
}
