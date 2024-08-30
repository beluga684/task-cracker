package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Editing(ptr *[]Task) {
	for num := 1; num != 0; {
		fmt.Println(
			"Editing tasks :\n",
			"0. Back\n",
			"1. New task\n",
			"2. Change the task\n",
			"3. Delete the task")
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			createTask(ptr)
		case 2:
			editMod(ptr)
		case 3:
			deleteTask(ptr)
		default:
			return
		}
	}
}

func createTask(ptr *[]Task) bool {
	var task Task
	task.ID = len(*ptr) + 1
	task.Status = "-todo-"

	fmt.Print("Enter the task (0 to cancel) - ")
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return false
	}

	task.Description = strings.TrimRight(str, "\n")
	tm := time.Now()
	task.Date = tm.Format("02.01.2006 - 15:04:05")

	*ptr = append(*ptr, task)
	fmt.Println("Added")
	return true
}

func deleteTask(ptr *[]Task) {
	for _, task := range *ptr {
		fmt.Println(task)
	}

	fmt.Print("Enter the index to delete (0 to discard) - ")
	indx := 0
	fmt.Scan(&indx)
	if indx == 0 {
		fmt.Println("Cancelled")
		return
	} else if indx >= len(*ptr) || indx < 0 {
		fmt.Println("The ID could not be found")
		return
	}
	*ptr = remove(*ptr, indx-1)

	for i := indx - 1; i < len(*ptr); i++ {
		(*ptr)[i].ID -= 1
	}

	fmt.Println("Deleted")
	return
}

func remove(slice []Task, i int) []Task {
	return append(slice[:i], slice[i+1:]...)
}

func editMod(ptr *[]Task) {
	for num := 1; num != 0; {
		fmt.Println(
			"Select :\n",
			"0. Back\n",
			"1. Edit task\n",
			"2. Edit status")
		fmt.Scan(&num)

		if num == 1 || num == 2 {
			editTask(ptr, num)
		}
	}
}

func editTask(ptr *[]Task, mod int) {
	for _, task := range *ptr {
		fmt.Println(task)
	}

	fmt.Print("Enter the ID to edit (0 to discard) - ")
	indx := 0
	fmt.Scan(&indx)

	if indx == 0 {
		fmt.Println("Cancelled")
	} else if indx >= len(*ptr) || indx < 0 {
		fmt.Println("The ID could not be found")
	} else {
		indx -= 1
		if mod == 1 {
			fmt.Print("Enter the task (0 to cancel) - ")
			str, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return
			}
			(*ptr)[indx].Description = strings.TrimRight(str, "\n")
		} else if mod == 2 {
			fmt.Println(
				"Select a status :\n",
				"0. Back\n",
				"1. -todo-\n",
				"2. ..in_progress..\n",
				"3. !!done!!")
			num := -1
			fmt.Scan(&num)

			switch num {
			case 0:
				return
			case 1:
				(*ptr)[indx].Status = "-todo-"
			case 2:
				(*ptr)[indx].Status = "..in_progress.."
			case 3:
				(*ptr)[indx].Status = "!!done!!"
			default:
				return
			}
		}
	}
}

/*
-todo-
..in_progress..
!!done!!
*/
