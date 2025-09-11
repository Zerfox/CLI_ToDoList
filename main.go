package main

import (
	"CLI_project/database"
	"CLI_project/internal/event"
	"CLI_project/internal/todo"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {

	const dir = "/home/katler/Projects/CLI_To-Do_list/database/"
	todoList := make(map[string]*todo.Todo, 10)
	eventMap := make(map[string]*event.Event, 20)

	database.WorkList(dir)

	err := loader(dir, todoList, eventMap)
	if err != nil {
		return
	}

	fmt.Println("Добро пожаловать в CLI_ToDo_List")

	for {
		fmt.Print("\nВведите название команды: ")
		inputUser, err := eventScanIt(eventMap)
		if err != nil {
			fmt.Println(inputUser, err)
			continue
		}

		switch inputUser {

		case "exit": // выход
			resultSaveTodo := database.Save(dir+"dataList.txt", todoList)
			if resultSaveTodo != nil {
				output := "Не удалось выполнить сохранение данных задач"
				eventMap[output] = event.NewEvent(output, resultSaveTodo)
				return
			}
			resultSaveEvent := database.Save(dir+"dataEvent.txt", eventMap)
			if resultSaveEvent != nil {
				output := "Не удалось выполнить сохранение данных событий"
				eventMap[output] = event.NewEvent(output, resultSaveEvent)
				return
			}
			os.Exit(0)

		case "add": // команда добавить
			fmt.Println("Введите заголовок задачи, в формате одного слова: ")
			heading, err := eventScanIt(eventMap)
			if err != nil {
				fmt.Println(inputUser, err)
				break
			}

			fmt.Println("Введите описание этой задачи")
			discription, err := eventScanIt(eventMap)
			if err != nil {
				fmt.Println(inputUser, err)
				break
			}
			todo.GetStructToDo(todoList, heading, discription)

		case "list": // Список всех задач
			fmt.Println("\n[Перечень задач следующий]")
			todo.SortedByTime(todoList)
			break

		case "event": // Список событий
			fmt.Println("Перечень событий")
			event.SortedByTime(eventMap)
			break

		case "done": // Выполнение задачи
			fmt.Println("Введите название задачи которая выполнена: ")
			input, err := eventScanIt(eventMap)
			if err != nil {
				fmt.Println(inputUser, err)
				break
			}
			if _, ok := todoList[input]; ok {
				todoList[input].CompliteTodo()
			} else {
				err = errors.New("ошибка, такой задачи не существует")
				eventMap[input] = event.NewEvent(input, err)
				fmt.Println(inputUser, err)
			}

		case "del": // удаление задачи
			fmt.Println("Введите название задачи для удаления: ")
			input, err := eventScanIt(eventMap)
			if err != nil {
				fmt.Println(inputUser, err)
				break
			}
			if _, ok := todoList[input]; ok {
				deleteTodo(todoList, input)
				fmt.Println("Указанная задача удалена успешно")
			} else {
				err = errors.New("ошибка, такой задачи не существует")
				eventMap[input] = event.NewEvent(input, err)
				fmt.Println(inputUser, err)
			}

		case "help": // вывод списка команд
			fmt.Print("add  -- добавить задачу\n" +
				"del  -- удалить задачу \n" +
				"done -- пометить выполненной задачу \n" +
				"list -- полный список задач\n" +
				"event -- Список событий" +
				"exit -- завершить выполнение программы\n")
		default:
			inputUser = "[Неизвестная команда]"
			err := errors.New("- не корректный ввод команды, help для вызова набора команд")
			eventMap[inputUser] = event.NewEvent(inputUser, err)
			fmt.Println(inputUser, err)
		}

	}

}

// Функция удаления задачи
func deleteTodo(m map[string]*todo.Todo, key string) {
	delete(m, key)
}

// Сканит и отпр. в мапу
func eventScanIt(m map[string]*event.Event) (string, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputUser := scanner.Text()
	if inputUser == "" {
		inputUser = "[Пустое значение]"
		err := errors.New("-ошибка, значение не должно быть пустым")
		m[inputUser] = event.NewEvent(inputUser, err)
		return inputUser, err
	}
	m[inputUser] = event.NewEvent(inputUser, nil)
	return inputUser, nil
}

func loader(dir string, todoList map[string]*todo.Todo, eventMap map[string]*event.Event) error {

	loadResultTodo := database.Load(dir+"dataList.txt", &todoList)
	if loadResultTodo != nil {
		output := "Не удалось выполнить грузку данных задач"
		eventMap[output] = event.NewEvent(output, loadResultTodo)
		fmt.Println(output, loadResultTodo)
	} else {
		fmt.Println("Задачи загружены")

	}

	loadResultEvent := database.Load(dir+"dataEvent.txt", &eventMap)
	if loadResultEvent != nil {
		output := "Не удалось выполнить загрузку данных задач"
		eventMap[output] = event.NewEvent(output, loadResultEvent)
		fmt.Println(output, loadResultEvent)
	} else {
		fmt.Println("События загружены")
	}
	return nil
}
