package scanner

import (
	"CLI_project/todo"
	"bufio"
	"errors"
	"os"
	"strings"
)

type scanner struct {
	todoLis *todo.List
	events  []Event
}

func NewScaner(list todo.List) *scanner {
	return &scanner{todoLis: &list}
}

func (s *scanner) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPromt()

		ok := scanner.Scan()
		if !ok {
			return
		}
		inputString := scanner.Text()

		result := s.process(inputString)
		if result != nil {
			if errors.Is(result, ErrNeedExit) {
				printExit()
				return
			}
			printResult(result)
		}
		event := NewEvent(result, inputString)
		s.events = append(s.events, event)
	}

}
func (s *scanner) process(inputString string) error {
	fields := strings.Fields(inputString)

	if len(fields) == 0 {
		return ErrEmptiInput
	}

	cmd := fields[0]
	if cmd == "exit" {
		return ErrNeedExit
	}

	if cmd == "add" {
		return s.cmdAdd(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}
	if cmd == "completed" {
		return s.cmdCompleted(fields)
	}
	if cmd == "del" {
		return s.cmdDel(fields)
	}
	if cmd == "help" {
		return s.cmdHelp(fields)
	}
	if cmd == "events" {
		return s.cmdEvents(fields)
	}
	return ErrUnknownCommand
}
func (s *scanner) cmdAdd(fields []string) error {
	if len(fields) < 3 {
		return ErrWrongArgs
	}
	title := fields[1]
	description := ""
	for i := 2; i < len(fields); i++ {
		description += fields[i]
		if i != len(fields)-1 {
			description += " "
		}
	}
	task := todo.NewTask(title, description)
	s.todoLis.AddTask(task)
	printAdd(title)
	return nil
}

func (s *scanner) cmdList(fields []string) error {
	if len(fields) != 1 {
		return ErrWrongArgs
	}
	tasks := s.todoLis.ListTask()
	printTasks(tasks)
	return nil
}
func (s *scanner) cmdCompleted(fields []string) error {
	if len(fields) != 2 {
		return ErrWrongArgs
	}
	title := fields[1]
	completedTaskResult := s.todoLis.CompletedTask(title)
	if completedTaskResult == nil {
		return completedTaskResult
	}
	printComplete(title)
	return nil
}

func (s *scanner) cmdDel(fields []string) error {
	if len(fields) != 2 {
		return ErrWrongArgs
	}
	title := fields[1]
	delTaskResult := s.todoLis.DeleteTask(title)
	if delTaskResult == nil {
		return delTaskResult
	}
	printDel(title)
	return nil
}
func (s *scanner) cmdHelp(fields []string) error {
	if len(fields) != 1 {
		return ErrWrongArgs
	}
	printHelp()
	return nil
}

func (s *scanner) cmdEvents(fields []string) error {
	if len(fields) != 1 {
		return ErrWrongArgs
	}
	printEvents(s.events)
	return nil
}
