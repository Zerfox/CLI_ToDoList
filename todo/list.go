package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (list *List) AddTask(task Task) error {
	if _, ok := list.tasks[task.Title]; ok {
		return ErrTaskAlreadyExistst
	}
	list.tasks[task.Title] = task
	return nil
}

func (list *List) ListTask() map[string]Task {
	tmp := make(map[string]Task, len(list.tasks))

	for k, v := range list.tasks {
		tmp[k] = v
	}
	return tmp
}

func (list *List) CompletedTask(title string) error {
	task, ok := list.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	task.Complete()
	list.tasks[title] = task
	return nil
}

func (list *List) ListNotCompleteTask(title string) error {
	notCompleteTask := make(map[string]Task)

	for tatle, task := range list.tasks {
		if !task.Completed {
			notCompleteTask[tatle] = task
		}
	}
	return nil
}

func (list *List) DeleteTask(title string) error {
	_, ok := list.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	delete(list.tasks, title)
	return nil
}
