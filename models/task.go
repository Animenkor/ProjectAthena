// Package models contains the task list models
package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const FileName = "data.csv"

// Task repository
var tasks []Task

// Task persistence
var filePersistence bool = false

// Task as type
type Task struct {
	Title       string
	Description string
	Tag         string
	Completed   bool
}

// EnableFilePersistence enables the file persistence
func EnableFilePersistence() {
	filePersistence = true
}

// DisableFilePersistence disables the file persistence
func DisableFilePersistence() {
	filePersistence = false
}

// Initialize does the initialization of the repository
func Initialize() {
	if filePersistence {
		tasks, _ = getDataFromFile()
	}
}

func getDataFromFile() ([]Task, error) {
	// open file
	file, err := os.Open(FileName)
	if err != nil {
		return nil, err
	}

	var completedTasks []Task

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)
	for {
		records, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		task := parseTaskData(records)

		// Add task to slice
		completedTasks = append(completedTasks, task)
	}

	// remember to close the file at the end
	file.Close()

	return completedTasks, nil
}

func parseTaskData(rec []string) Task {
	// Parse task
	title := rec[0]
	description := rec[1]
	tag := rec[2]
	completed := ToBool(rec[3])

	// Create new task based on parsed values
	task := Task{Title: title, Description: description, Tag: tag, Completed: completed}
	return task
}

func updateDataInFile() {
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	checkError("Cannot open file", err)
	writer := csv.NewWriter(file)

	for _, task := range tasks {
		taskSerialized := []string{task.Title, task.Description, task.Tag, strconv.FormatBool(task.Completed)}
		err := writer.Write(taskSerialized)
		checkError("Cannot write to file", err)
	}

	writer.Flush()
	file.Close()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// FindAllTasks returns a copy of all tasks
func FindAllTasks() []Task {
	allTasks := make([]Task, len(tasks))
	copy(allTasks, tasks)
	return allTasks
}

// ToInt converts a string to an integer value
func ToInt(info string) int {
	aInt, _ := strconv.Atoi(info)
	return aInt
}

// ToBool converts a string to a boolean value
func ToBool(info string) bool {
	aBool, _ := strconv.ParseBool(info)
	return aBool
}

// AddTask adds a task to the library
func AddTask(task Task) bool {
	if task.Title == "" || task.Description == "" || task.Tag == "" || task.Completed {
		return false
	}

	tasks = append(tasks, task)

	fmt.Println("The Task has been added successfully")

	if filePersistence {
		updateDataInFile()
	}

	return true
}

// CompleteTask marks a task from the list as "completed"

func CompleteTask(id string) {
	idAsInt := ToInt(id) - 1

	for index, incompleteTask := range tasks {

		if index == idAsInt {

			tasks[index].Completed = true
			fmt.Println("The task has been marked as COMPLETED")

		}
		incompleteTask = incompleteTask

	}

	if filePersistence {
		updateDataInFile()
	}
}

func IncompleteTask(id string) {
	idAsInt := ToInt(id) - 1

	for index, completedTask := range tasks {

		if index == idAsInt {

			tasks[index].Completed = false
			fmt.Println("The task has been marked as INCOMPLETED")

		}
		completedTask = completedTask

	}

	if filePersistence {
		updateDataInFile()
	}
}

func RemoveTask(id string) {
	var tempTasks []Task
	idAsInt := ToInt(id) - 1

	for index, currentTask := range tasks {
		if index != idAsInt {
			tempTasks = append(tempTasks, currentTask)
		}

	}

	tasks = tempTasks

	fmt.Println("The Task has been removed successfully")

	if filePersistence {
		updateDataInFile()
	}
}

// EditTask lets you edit an information field of a task
func EditTask(i string) {
	idAsInt := ToInt(i) - 1

	for index, currentTask := range tasks {
		if index == idAsInt {
			fmt.Print("Chosen Task:")
			fmt.Println(currentTask)
		}
	}
}

// EditFieldTask sets a new value to a chosen field
func EditFieldTask(rowIndex, field, newInfo string) {
	idAsInt := ToInt(rowIndex) - 1

	for index := range tasks {
		if index == idAsInt {
			switch {
			case field == "title":
				tasks[index].Title = newInfo
			case field == "description":
				tasks[index].Description = newInfo
			case field == "tag":
				tasks[index].Tag = newInfo
			}
		}
	}
	if filePersistence {
		updateDataInFile()
	}
	fmt.Println("The task has been updated")
}
