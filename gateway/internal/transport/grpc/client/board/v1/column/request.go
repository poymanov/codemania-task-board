package column

type CreateColumnRequest struct {
	Name string

	BoardId int
}

type DeleteColumnRequest struct {
	Id int
}
