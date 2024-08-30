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
			createTask(ptr)
		case 2:
			// изменение название-статус
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
	task.Status = "todo"

	fmt.Print("Введите задачу - ")
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return false
	}

	task.Description = strings.TrimRight(str, "\n")
	tm := time.Now()
	task.Date = tm.Format("02.01.2006 - 15:04:05")

	*ptr = append(*ptr, task)
	fmt.Println("Добавлено")
	return true
}

func deleteTask(ptr *[]Task) bool {
	for _, task := range *ptr {
		fmt.Println(task)
	}

	fmt.Print("Какую задачу удалить? - ")
	indx := 0
	fmt.Scan(&indx)
	*ptr = remove(*ptr, indx-1)

	fmt.Println("Удалено")
	return true
}

func remove(slice []Task, i int) []Task {
	return append(slice[:i], slice[i+1:]...)
}
