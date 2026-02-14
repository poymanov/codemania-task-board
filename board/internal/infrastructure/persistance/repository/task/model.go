package task

import "time"

type Task struct {
	Id int `db:"id"`

	Title string `db:"title"`

	Description string `db:"description"`

	Assignee string `db:"assignee"`

	Position float64 `db:"position"`

	ColumnId int `db:"column_id"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt *time.Time `db:"updated_at"`
}
