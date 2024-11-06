package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	tasks = make(map[string]*Task)
)

type Task struct {
	Title       string
	Description string
	Completed   bool
}

// Cette fonction permet de cloturer une task
// Comme on la modifie on utilise *Task et pas Task
func (t *Task) Done() {
	t.Completed = true
}

// On parcours notre liste de task et pour chacune d'entre
// On appelle sa fonction d'affichage.
func displayAllTasks() {
	for id, task := range tasks {
		fmt.Printf("\n-> Affichage de la tâche %v\n", id)
		task.Display()
	}
}

// On va créer une nouvelle Task avec une saisie clavier
// ensuite on l'ajoute dans notre liste de Task
func addTask() {
	fmt.Printf("\nSaisissez un titre : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	scan.Scan()
	titre := scan.Text()
	log.Printf("Titre: %v", titre)

	fmt.Printf("\nSaisissez un description : ")
	scan.Scan()
	desc := scan.Text()
	log.Printf("Description: %v", desc)

	t := Task{
		Title:       titre,
		Description: desc,
		Completed:   false,
	}

	// Notre map de task stocke des pointers de task
	// Il faut donc utiliser &t
	tasks[titre] = &t
}

// Cette fonction permet d'afficher une tâche, elle prend aucun argument et affiche sur la sortie standard.
func (t Task) Display() {
	fmt.Printf("Titre: %v\nDescription: %v\nTerminé: %v\n", t.Title, t.Description, t.Completed)
}
