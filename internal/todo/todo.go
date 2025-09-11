package todo

import (
	"fmt"
	"sort"
	"time"
)

type Todo struct {
	Heading     string    `json:"heading"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreateTime  time.Time `json:"createTime"`
	DoneTime    time.Time `json:"doneTime"`
}

// конструктор структуры

func NewTodo(
	heading string,
	description string,
) *Todo {
	if heading == "" {
		fmt.Println("Заголовок не может быть пустым")
		return &Todo{}
	}
	return &Todo{
		Heading:     heading,
		Description: description,
		CreateTime:  time.Now(),
	}
}
func GetStructToDo(m map[string]*Todo, name string, description string) {
	m[name] = NewTodo(name, description)
}
func (todo *Todo) CompliteTodo() {
	todo.Completed = true
	todo.DoneTime = time.Now()
	fmt.Println("Указанная задача выполнена")
	return
}

func (todo *Todo) String() string {
	status := "Закрыта"
	if todo.Completed != false {
		return fmt.Sprintf("Наименование задачи: %s [time: %s]\nОписание задачи: %s\nСтатус задачи: %s\nВремя выполнения: %s\n",
			todo.Heading,
			todo.CreateTime.Format("2006-01-02 15:04:05"),
			todo.Description,
			status,
			todo.DoneTime.Format("2006-01-02 15:04:05"))
	} else {
		return fmt.Sprintf("Наименование задачи: %s [time: %s]\n Описание задачи: %s\n",
			todo.Heading,
			todo.CreateTime.Format("2006-01-02 15:04:05"),
			todo.Description)
	}
}

func SortedByTime(m map[string]*Todo) {
	items := make([]*Todo, 0, len(m))
	for _, v := range m {
		items = append(items, v)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].CreateTime.Before(items[j].CreateTime) // по возрастанию времени
	})

	for _, it := range items {
		fmt.Println(it)
	}
}
