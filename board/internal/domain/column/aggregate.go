package column

type NewColumn struct {
	Name string

	BoardID int
}

func NewNewColumn(name string, boardID int) NewColumn {
	return NewColumn{
		Name:    name,
		BoardID: boardID,
	}
}
