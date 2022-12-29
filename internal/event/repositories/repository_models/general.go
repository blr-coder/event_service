package repository_models

type SortOrder string

func (o SortOrder) String() string {
	return string(o)
}

const (
	SortOrderASC  SortOrder = "ASC"
	SortOrderDESC SortOrder = "DESC"
)
