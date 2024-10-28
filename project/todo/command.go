package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Ad string
	Dl int
	Ed string
	Tg int
	Li bool
}

func NewCmdFlags() *CmdFlags {
	command := CmdFlags{}

	flag.StringVar(&command.Ad, "add", "", "Add a new todo")
	flag.StringVar(&command.Ed, "edit", "", "Edit todo by index,tittle. id:title")
	flag.IntVar(&command.Dl, "del", -1, "Delete todo by index")
	flag.IntVar(&command.Tg, "done", -1, "Toggle complete by index true/false")
	flag.BoolVar(&command.Li, "ls", false, "List todos")

	flag.Parse()

	return &command
}

func (command *CmdFlags) Execute(todos *Todos) {
  switch {
  case command.Li:
    todos.print()
  case command.Ad != "":
    todos.add(command.Ad)
  case command.Ed != "":
    parts := strings.SplitN(command.Ed, ":", 2)
  if len(parts) != 2 {
      fmt.Println("invalid format index:title")
      os.Exit(1)
    }
    index, err := strconv.Atoi(parts[0])
    if err != nil {
      fmt.Println("Error invalid index")
      os.Exit(1)
    }
  todos.edit(index, parts[1])
  case command.Tg != -1:
  todos.toggle(command.Tg)
  case command.Dl != -1:
  todos.delete(command.Dl)

  default:
  fmt.Println("Invalid Command")
  }
}
