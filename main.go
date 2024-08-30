package main

import "fmt"

type Task struct {
	ID                  int
	Description, Status string
	Date                string
}

func main() {
	tasks := []Task{
		{1, "заснуть", "-todo-", "30.08.2024 - 18:48:48"},
		{2, "устать", "..in_progress..", "30.08.2024 - 18:48:52"},
		{3, "проснуться", "!!done!!", "30.08.2024 - 18:48:57"},
	}

	for num := 1; num != 0; {
		fmt.Println(
			"Select an action :\n",
			"0. Exit\n",
			"1. Editing tasks\n",
			"2. Task listing",
		)
		fmt.Scan(&num)

		switch num {
		case 0:
			continue
		case 1:
			Editing(&tasks)
		case 2:
			Listing(&tasks)
		default:
			continue
		}
	}
}
