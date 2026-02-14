package task

type NewTask struct {
	Title string

	Description string

	Assignee string

	ColumnId int
}

func NewNewTask(title, description, assignee string, columnId int) NewTask {
	return NewTask{
		Title:       title,
		Description: description,
		Assignee:    assignee,
		ColumnId:    columnId,
	}
}
