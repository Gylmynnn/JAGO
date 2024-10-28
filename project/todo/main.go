package main


func main() {
	todos := Todos{}
	store := NewStorage[Todos]("todosdb.json")
	store.Load(&todos)

	command := NewCmdFlags()
	command.Execute(&todos)

	store.Save(todos)

}
