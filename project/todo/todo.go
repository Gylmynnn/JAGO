package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
	fmt.Println("Successfully add : ", todo.Title)
	todos.print()
}

func (todos *Todos) validateIndex(index int) error {
	if index < 1 || index > len(*todos) {
		err := errors.New("index invalid")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	to := *todos
	if err := to.validateIndex(index); err != nil {
		return err
	}

	*todos = append(to[:index-1], to[index:]...)
	fmt.Println("Successfully Deleted")
	todos.print()
	return nil

}

func (todos *Todos) toggle(index int) error {
	to := (*todos)
	if err := to.validateIndex(index); err != nil {
		return err
	}

	isCompleted := to[index-1].Completed

	if !isCompleted {
		completionTime := time.Now()
		to[index-1].CompletedAt = &completionTime
	}

	to[index-1].Completed = !isCompleted
	title := to[index-1].Title
  fmt.Println("Task :", title, "is : DONE",)
	todos.print()

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	to := *todos
	if err := to.validateIndex(index); err != nil {
		return err
	}

	to[index-1].Title = title
	fmt.Println("Successfully Update :", title)
	todos.print()
	return nil
}

func (todos *Todos) print() {
	fmt.Println("Hallo Gyl let's Progress 🔥")
	fmt.Println("Date 📅 ", time.Now().Format(time.RFC1123))
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("📌 No", "🚀 Title", "🎉 Completed", "⏳ Created At", "⌛ Completed At")

	for index, to := range *todos {
		completed := "⚔  TODO"
		completedAt := ""

		if to.CompletedAt != nil {
			completed = "✨ DONE"
			completedAt = to.CompletedAt.Format(time.RFC1123)
		}
		tbl.AddRow(strconv.Itoa(index+1), to.Title, completed, to.CreatedAt.Format(time.RFC1123), completedAt)
	}

	tbl.Render()

}
