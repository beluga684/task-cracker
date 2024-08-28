package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Editing() {
	for num := 1; num != 0; {
		fmt.Println(
			"Работа с задачами :\n",
			"0. Назад\n",
			"1. Новая задача\n",
			"2. Изменить задачу\n",
			"3. Удалить задачу")
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			createTask()
		case 2:
			// изменение название-статус
		case 3:
			// удалить
		default:
			return
		}
	}
}

func createTask() bool {
	var task Task
	task.ID = 1
	task.Status = "todo"

	fmt.Print("Введите задачу - ")
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return false
	}

	task.Description = strings.TrimRight(str, "\n")
	tm := time.Now()
	task.Date = tm.Format("02.01.2006 - 15:04:05")

	fmt.Println("Новая запись :")
	fmt.Println(
		task.ID, "|",
		task.Description, "|",
		task.Status, "|",
		task.Date,
	)
	return true
}
