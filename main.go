package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("##### Welcome to our Todolist App! #####")
	http.ListenAndServe(":8080", nil)

	// var shortGolang = "Watch go crash course"
	// var fullGolang = "Watch TWN golang Full course"
	// var rewardDessert = "Reward myself with a donut"
	// var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	// printTasks(taskItems)
	// fmt.Println()

	// taskItems = addTask(taskItems, "Go for a run")
	// taskItems = addTask(taskItems, "Practice coding in go")

	// fmt.Println("Updated list")
	// printTasks(taskItems)

}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItem []string, newTask string) []string {

	var updatedTaskItems = append(taskItem, newTask)
	return updatedTaskItems
}
