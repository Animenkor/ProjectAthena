// Package controllers contains the task list controllers
package controllers

import (
	"bufio"
	"github.com/Animenkor/ProjectAthena/models"
	"github.com/Animenkor/ProjectAthena/views"
	"log"
	"os"
	"strings"
	"time"
)

// Run does the running of the console application
func Run(enablePersistence bool) {
	if enablePersistence {
		models.EnableFilePersistence()
	} else {
		models.DisableFilePersistence()
	}

	models.Initialize()

	views.Clear()
	views.PrintMenu()

	for true {
		executeCommand()
	}
}

func executeCommand() {
	command := AskForInput()
	parseCommand(command)
}

func AskForInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSpace(response)
	return response
}

func parseCommand(input string) {
	switch {
	case input == "1":
		// Add a task
		views.Clear()
		views.PrintTaskInformation()
		response := AskForInput()
		task := CreateTask(response)
		models.AddTask(task)
		views.PrintContinue()
		break
	case input == "2":
		// Remove a task
		views.Clear()
		tasks := models.FindAllTasks()
		views.PrintTaskList(tasks)
		views.PrintRemovingInformation()
		response := AskForInput()
		models.RemoveTask(response)
		views.PrintContinue()
		break
	case input == "3":
		// mark task as completed
		views.Clear()
		tasks := models.FindAllTasks()
		views.PrintTaskList(tasks)
		views.PrintCompleteTask()
		response := AskForInput()
		models.CompleteTask(response)
		views.PrintContinue()
		break
	case input == "4":
		// mark task as not completed
		views.Clear()
		tasks := models.FindAllTasks()
		views.PrintTaskList(tasks)
		views.PrintIncompleteTask()
		response := AskForInput()
		models.IncompleteTask(response)
		views.PrintContinue()
		break

	case input == "5":
		// Edit a task
		views.Clear()
		tasks := models.FindAllTasks()
		views.PrintTaskList(tasks)
		views.PrintEditTaskInformation()
		responseRowIndex := AskForInput()
		models.EditTask(responseRowIndex)
		views.PrintEditFieldTask()
		responseField := AskForInput()
		views.PrintNewFieldTask()
		responseNewInfo := AskForInput()
		models.EditFieldTask(responseRowIndex, responseField, responseNewInfo)
		views.PrintContinue()
		break

	case input == "6":
		// Show all tasks
		views.Clear()
		tasks := models.FindAllTasks()
		views.PrintTaskList(tasks)
		views.PrintContinue()
		break

	case input == "c":
		// Clear view and print menu
		views.Clear()
		views.PrintMenu()
		break
	case input == "q":
		// Terminate application
		views.Clear()
		views.PrintGoodbye()
		time.Sleep(5 * time.Second)
		views.ShutDown()
		break
	}
}

func CreateTask(input string) models.Task {
	inputFormatted := strings.ReplaceAll(input, ", ", ",")
	elements := strings.Split(inputFormatted, ",")

	if len(elements) != 4 {
		log.Fatal("Input not as expected. 4 comma separated values needed!")
	}

	// Parse task
	title := elements[0]
	description := elements[1]
	tag := elements[2]
	var completed bool = models.ToBool(elements[3])

	// Create new task based on parsed values
	task := models.Task{Title: title, Description: description, Tag: tag, Completed: completed}

	return task
}
