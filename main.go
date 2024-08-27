package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Task struct {
	ID                  int
	Description, Status string
	Date                string
}

func clearScreen() {
	if os.Getenv("OS") == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func exit() {
	return
}

func listing() {
	for num := 1; num != 0; {
		fmt.Println(
			"Листинг задач :\n",
			"0. Назад\n",
			"1. Все задачи\n",
			"2. Ожидающие\n",
			"3. В процессе\n",
			"4. Завершённые",
		)
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			// all
		case 2:
			// todo
		case 3:
			// in progress
		case 4:
			// done
		default:
			return
		}
	}
}

func main() {
	//подгрузка файла
	//создание этого файла, если не найден
	for num := 1; num != 0; {
		fmt.Println(
			"Выберите действие :\n",
			"0. Выход из программы\n",
			"1. Работа с задачами\n",
			"2. Листинг задач",
		)
		fmt.Scan(&num)

		switch num {
		case 0:
			exit()
		case 1:
			editing()
		case 2:
			listing()
		default:
			continue
		}
	}
}

func editing() {
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
