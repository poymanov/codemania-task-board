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

type GetAllFilter struct {
	ColumnId int
}

type GetAllSort struct {
	SortByPosition string
}

func NewGetAllFilter(boardId int) GetAllFilter {
	return GetAllFilter{
		ColumnId: boardId,
	}
}

func NewGetAllSort(sortByPosition string) GetAllSort {
	return GetAllSort{
		SortByPosition: sortByPosition,
	}
}
