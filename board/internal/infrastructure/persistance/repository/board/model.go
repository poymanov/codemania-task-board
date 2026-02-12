package board

import "time"

type Board struct {
	Id int `db:"id"`

	Name string `db:"name"`

	Description string `db:"description"`

	OwnerId int `db:"owner_id"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt *time.Time `db:"updated_at"`
}
