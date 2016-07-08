package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Katuva/go-todo/structs"
	"github.com/olekukonko/tablewriter"
)

func main() {
	fmt.Println("Welcome to go-todo")
	structs.Todo.Parse("todo.json")

	var newTodo structs.STodo
	newTodo.Content = "This is a new todo!"
	structs.Todo.Add(newTodo)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Todo"})
	for i, t := range structs.Todo.Todos {
		table.Append([]string{strconv.Itoa(i + 1), t.Content})
	}
	table.Render()
}
