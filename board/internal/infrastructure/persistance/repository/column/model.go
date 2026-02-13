package column

import "time"

type Column struct {
	Id int `db:"id"`

	Name string `db:"name"`

	Position float64 `db:"position"`

	BoardId int `db:"board_id"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt *time.Time `db:"updated_at"`
}
