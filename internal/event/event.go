package event

import (
	"fmt"
	"sort"
	"time"
)

type Event struct {
	Name       string    `json:"Name"`
	Error      error     `json:"Error"`
	CreateTime time.Time `json:"CreateTime"`
}

func NewEvent(name string, err error) *Event {
	if name == "" {
		fmt.Println("Заголовок не может быть пустым")
		return &Event{}
	}
	return &Event{Name: name, Error: err, CreateTime: time.Now()}
}

func (event *Event) String() string {
	if event.Error == nil {
		return fmt.Sprintf("Name: %s\n[Time: %s]\nerror: %v\n\n",
			event.Name,
			event.CreateTime.Format("2006-01-02 15:04:05"),
			"_")
	}
	return fmt.Sprintf("Name: %s\n[Time: %s]\nerror: %v\n\n",
		event.Name,
		event.CreateTime.Format("2006-01-02 15:04:05"),
		event.Error)

}

// SortedByTime Сортировщик по времени
func SortedByTime(m map[string]*Event) {
	items := make([]*Event, 0, len(m))
	for _, v := range m {
		items = append(items, v)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].CreateTime.Before(items[j].CreateTime)
	})

	for _, it := range items {
		fmt.Println(it)
	}
}
