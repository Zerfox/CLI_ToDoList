package main

import (
	"CLI_project/scanner"
	"CLI_project/todo"
)

func main() {
	todolist := todo.NewList()

	scan := scanner.NewScaner(*todolist)

	scan.Start()
}
