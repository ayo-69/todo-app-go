package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type Task struct {
	Id      int64  `xorm:"pk autoincr"`
	Name    string `xorm:"varchar(25) notnull"`
	Content string `xorm:"varchar(250) notnull"`
}

func main() {
	fmt.Println("---------------------------")
	fmt.Println("WELCOME TO MY TODO APP")
	fmt.Println("---------------------------")
	fmt.Println()

	//Create and connect to database
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic("Error with the database : ")
	}
	defer engine.Close()

	//Sync the struct with the database table
	err = engine.Sync2(new(Task))
	if err != nil {
		panic("Error syncing struct")
	}

	//Option Handler
	for {
		fmt.Println("---------------------------")
		fmt.Println("Otions")
		fmt.Println("1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Edit a task")
		fmt.Println("4. Remove a task")
		fmt.Println("quit to quit...")
		var choice string
		fmt.Print("Enter an option : ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			AddTask(engine)
		case "2":
			ListTasks(engine)
		case "3":
			EditTask(engine)
		case "4":
			RemoveTask(engine)
		case "quit":
			os.Exit(0)
		default:
			fmt.Println()
			fmt.Println("---------------------------")
			fmt.Println("Invalid option")
			fmt.Println("---------------------------")
			fmt.Println()
		}
	}
}

func AddTask(engine *xorm.Engine) {
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("Adding task")
	fmt.Println("---------------------------")
	fmt.Println()

	//Collect title and content
	var title, content string
	fmt.Print("Enter title : ")
	fmt.Scanln(&title)
	fmt.Print("Enter content : ")
	fmt.Scanln(&content)
	fmt.Println()

	//Insert user
	new_task := Task{Name: title, Content: content}
	_, err := engine.Insert(&new_task)
	if err != nil {
		panic("Couldn't insert task")
	}

	fmt.Println("Title : ", title)
	fmt.Println("Content : ", content)
}

func ListTasks(engine *xorm.Engine) {
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("List all tasks")
	fmt.Println("---------------------------")
	fmt.Println()

	var allTasks []Task
	err := engine.Find(&allTasks)
	if err != nil {
		panic("Couldn't read database")
	}
	for _, task := range allTasks {
		fmt.Printf("ID: %d, Title : %s, Content: %s\n", task.Id, task.Name, task.Content)
	}
}

func EditTask(engine *xorm.Engine) {
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("Edit task ...")
	fmt.Println("---------------------------")
	fmt.Println()

	//Update tasks
	var (
		ID int

		new_name    string
		new_content string

		new_task Task
	)

	fmt.Print("Enter ID : ")
	fmt.Scanln(&ID)
	fmt.Print("Enter new name : ")
	fmt.Scanln(&new_name)
	fmt.Print("Enter new content : ")
	fmt.Scanln(&new_content)

	new_task.Id = int64(ID)
	new_task.Name = new_name
	new_task.Content = new_content

	_, err := engine.ID(new_task.Id).Update(&new_task)
	if err != nil {
		panic("Couldn't update task.")
	}
}
func RemoveTask(engine *xorm.Engine) {
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("Remove task ...")
	fmt.Println("---------------------------")
	fmt.Println()

	var ID int
	fmt.Print("Enter ID : ")
	fmt.Scanln(&ID)
	_, err := engine.ID(ID).Delete(&Task{})
	if err != nil {
		panic("Couldn't remove tasks")
	}
}
