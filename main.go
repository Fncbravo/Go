package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch go crash course"
var fullGolang = "Watch TWN golang Full course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	fmt.Println("##### Welcome to our Todolist App! #####")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)

	// printTasks(taskItems)
	// fmt.Println()

	// taskItems = addTask(taskItems, "Go for a run")
	// taskItems = addTask(taskItems, "Practice coding in go")

	// fmt.Println("Updated list")
	// printTasks(taskItems)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)
}
