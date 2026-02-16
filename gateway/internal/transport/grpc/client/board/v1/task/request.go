package task

type CreateTaskRequest struct {
	Title string

	Description string

	Assignee string

	ColumnId int
}
