package scanner

import (
	"CLI_project/todo"
	"fmt"
	"github.com/k0kubun/pp/v3"
)

func printResult(result error) {
	fmt.Println("Результат вызова программы: ", result, "\n")
}

func printPromt() {
	fmt.Println("Введите команду: ")

}
func printExit() {
	fmt.Println("Завершение программы")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' успешно добавлена\n")
}

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Список дел, ", tasks, "\n")

}

func printComplete(title string) {
	fmt.Println("Задача '" + title + "' помечена как выполненная\n")

}
func printDel(title string) {
	fmt.Println("Задача '" + title + "' успешно удалена\n")
}
func printHelp() {
	fmt.Println("Список команд")
	fmt.Println("-Help - output list command\n-Exit - Exit program\n-Add - add task\n-Tasks - output tasks\nComplete - compete task\nDel - delite task")
}

func printEvents(events []Event) {
	pp.Println("Events: ", events, "\n")
}
