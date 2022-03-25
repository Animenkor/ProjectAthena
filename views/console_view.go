// Package views contains the task list views
package views

import (
	"fmt"
	"github.com/Animenkor/ProjectAthena/models"
	"os"
	"os/exec"
)

// PrintMenu prints the menu to the console
func PrintMenu() {
	fmt.Println(`
	#************* TO DO List *****************
	#******* CHOOSE YOUR OPTION BELOW *********
	# 1. ADD A TASK
	# 2. REMOVE A TASK
	# 3. COMPLETE A TASK
	# 4. INCOMPLETE A TASK
	# 5. EDIT A TASK
	# 6. SHOW ALL TASKS
	#
	# c. CLEAR VIEW AND PRINT MENU
	# q. TERMINATE TO DO LIST
	`)
}

// Clear clears the console view
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// PrintTaskList prints all the tasks to the console
func PrintTaskList(tasksToPrint []models.Task) {

	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-6s | %-25s | %-50s | %-25s | %-10s\n", "INDEX:", "TITLE: ", "DESCRIPTION: ", "TAG: ", "COMPLETED: ")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------")

	for i, task := range tasksToPrint {
		fmt.Printf("%-6d | %-25s | %-50s | %-25s | %t\n", i+1, task.Title, task.Description, task.Tag, task.Completed)
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------")
}

// PrintContinue prints the continuation information to the console
func PrintContinue() (int, error) {
	return fmt.Println("Press c to continue!")
}

// PrintGoodbye prints a goodbye message to the console
func PrintGoodbye() {
	fmt.Println("Goodbye!")
}

// ShutDown terminates the application
func ShutDown() {
	os.Exit(0)
}

func PrintTaskInformation() {
	fmt.Println("Please enter the new task with the following format: Title, Description, Tag, Completed")
}

func PrintRemovingInformation() {
	fmt.Println("Please enter the the number of the task you want to delete")
}

func PrintCompleteTask() {
	fmt.Println("Please enter the the number of the task you want to mark as completed")
}

func PrintIncompleteTask() {
	fmt.Println("Please enter the the number of the task you want to mark as incompleted")
}

func PrintEditTaskInformation() {
	fmt.Println("Please enter the the number of the task you want to edit")
}

func PrintEditFieldTask() {
	fmt.Println("Do you want to change the title, description or tag")
}

func PrintNewFieldTask() {
	fmt.Println("Please enter the new information")

}
