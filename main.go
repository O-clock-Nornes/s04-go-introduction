package main

func main() {
	t := Task{
		Title:       "Dormir",
		Description: "C'est bien de dormir",
		Completed:   false,
	}
	// Notre map de task stocke des pointers de task
	// Il faut donc utiliser &t
	tasks[t.Title] = &t
	//log.Printf("Notre tâche : %#v", t)
	//t.Display()

	//displayAllTasks()

	//addTask()
	//t.Display()
	t.Done()
	//t.Display()
	displayAllTasks()

}
